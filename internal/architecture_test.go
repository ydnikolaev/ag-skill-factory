package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// architecture_test.go enforces "Modern Go" standards via AST analysis.
//
// Rules enforced:
// 1. Forbidden Types: No `interface{}` or `any` (Use Generics).
// 2. Forbidden Logging: No `log` (std) or `fmt.Print*` (Use `log/slog`).
// 3. No Global State: No exported `var` at package level.
// 4. Modern Iteration: Suggest `for range x` instead of `for _ := range x`.
// 5. Context Hygiene: `ctx` must be 1st arg; No `context.Background()` in funcs.
// 6. Package Documentation: Every package must have a doc.go file.

func TestArchitecture(t *testing.T) {
	t.Parallel()

	projectRoot := findProjectRoot(t)

	scanDirs := []string{
		filepath.Join(projectRoot, "cmd"),
		filepath.Join(projectRoot, "internal"),
	}

	walker := newArchWalker(t)

	for _, dir := range scanDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		}
		walkDir(t, walker, dir)
	}

	if len(walker.errors) > 0 {
		t.Errorf("Architecture Violations Found:\n%s", strings.Join(walker.errors, "\n"))
	}
}

// TestPackageDocumentation ensures every package has a doc.go file.
func TestPackageDocumentation(t *testing.T) {
	t.Parallel()

	projectRoot := findProjectRoot(t)

	scanDirs := []string{
		filepath.Join(projectRoot, "cmd"),
		filepath.Join(projectRoot, "internal"),
	}

	missing := findMissingDocFiles(t, scanDirs, projectRoot)

	if len(missing) > 0 {
		t.Errorf("Packages missing doc.go:\n  %s", strings.Join(missing, "\n  "))
	}
}

func findMissingDocFiles(t *testing.T, scanDirs []string, projectRoot string) []string {
	t.Helper()

	var missing []string

	for _, baseDir := range scanDirs {
		if _, err := os.Stat(baseDir); os.IsNotExist(err) {
			continue
		}
		dirMissing := scanDirForMissingDocs(t, baseDir, projectRoot)
		missing = append(missing, dirMissing...)
	}
	return missing
}

func scanDirForMissingDocs(t *testing.T, baseDir, projectRoot string) []string {
	t.Helper()

	var missing []string

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return checkDirForDocGo(path, info, projectRoot, &missing)
	})
	if err != nil {
		t.Fatalf("Walk error in %s: %v", baseDir, err)
	}
	return missing
}

func checkDirForDocGo(path string, info os.FileInfo, projectRoot string, missing *[]string) error {
	if !info.IsDir() {
		return nil
	}

	if shouldSkipDocCheck(info.Name()) {
		return filepath.SkipDir
	}

	if !hasGoFiles(path) {
		return nil
	}

	docPath := filepath.Join(path, "doc.go")
	if _, err := os.Stat(docPath); os.IsNotExist(err) {
		relPath, _ := filepath.Rel(projectRoot, path)
		*missing = append(*missing, relPath)
	}
	return nil
}

func shouldSkipDocCheck(name string) bool {
	if strings.HasPrefix(name, ".") || name == "vendor" {
		return true
	}
	if name == "config" {
		return true
	}
	return false
}

func hasGoFiles(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".go") {
			return true
		}
	}
	return false
}

func findProjectRoot(t *testing.T) string {
	t.Helper()

	root, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	// internal/ -> project_root
	return filepath.Dir(root)
}

func newArchWalker(t *testing.T) *ArchWalker {
	return &ArchWalker{
		t:        t,
		fset:     token.NewFileSet(),
		errors:   make([]string, 0),
		warnings: make([]string, 0),
	}
}

func walkDir(t *testing.T, walker *ArchWalker, dir string) {
	t.Helper()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return walker.processPath(path, info)
	})
	if err != nil {
		t.Fatalf("Walk error in %s: %v", dir, err)
	}
}

// ArchWalker performs AST-based architecture checks.
type ArchWalker struct {
	t        *testing.T
	fset     *token.FileSet
	stack    []ast.Node
	errors   []string
	warnings []string
}

func (w *ArchWalker) processPath(path string, info os.FileInfo) error {
	if info.IsDir() {
		return w.handleDir(info)
	}
	return w.handleFile(path, info)
}

func (w *ArchWalker) handleDir(info os.FileInfo) error {
	if strings.HasPrefix(info.Name(), ".") || info.Name() == "vendor" {
		return filepath.SkipDir
	}
	return nil
}

func (w *ArchWalker) handleFile(path string, info os.FileInfo) error {
	if !strings.HasSuffix(info.Name(), ".go") {
		return nil
	}
	if w.shouldSkipFile(path, info.Name()) {
		return nil
	}
	w.checkFile(path)
	return nil
}

func (w *ArchWalker) shouldSkipFile(path, name string) bool {
	// Skip generated files
	if strings.HasSuffix(name, ".pb.go") || strings.HasSuffix(name, "_mock.go") {
		return true
	}
	// Skip this test file itself
	if strings.Contains(path, "architecture_test.go") {
		return true
	}
	return false
}

func (w *ArchWalker) checkFile(path string) {
	f, err := parser.ParseFile(w.fset, path, nil, parser.ParseComments)
	if err != nil {
		w.t.Logf("Failed to parse %s: %v", path, err)
		return
	}

	comments := w.buildCommentMap(f)
	if w.isFileIgnored(comments) {
		return
	}

	ast.Inspect(f, func(n ast.Node) bool {
		return w.visitNode(n, path, comments, f)
	})
}

func (w *ArchWalker) buildCommentMap(f *ast.File) map[int]string {
	comments := make(map[int]string)
	for _, cg := range f.Comments {
		for _, c := range cg.List {
			line := w.fset.Position(c.Pos()).Line
			comments[line] = c.Text
		}
	}
	return comments
}

func (w *ArchWalker) isFileIgnored(comments map[int]string) bool {
	for _, text := range comments {
		if strings.Contains(text, "arch:ignore-file") {
			return true
		}
	}
	return false
}

func (w *ArchWalker) visitNode(n ast.Node, path string, comments map[int]string, _ *ast.File) bool {
	if n == nil {
		if len(w.stack) > 0 {
			w.stack = w.stack[:len(w.stack)-1]
		}
		return true
	}

	w.stack = append(w.stack, n)

	pos := w.fset.Position(n.Pos())
	if w.isIgnored(pos.Line, comments) {
		return true
	}

	w.checkNode(n, path, pos)
	return true
}

func (w *ArchWalker) isIgnored(line int, comments map[int]string) bool {
	if strings.Contains(comments[line], "arch:ignore") {
		return true
	}
	if strings.Contains(comments[line-1], "arch:ignore") {
		return true
	}
	return false
}

func (w *ArchWalker) report(path string, pos token.Position, rule, msg string) {
	w.errors = append(w.errors, fmt.Sprintf("[%s] %s:%d: %s", rule, filepath.Base(path), pos.Line, msg))
}

func (w *ArchWalker) checkNode(n ast.Node, path string, pos token.Position) {
	switch node := n.(type) {
	case *ast.Field:
		w.checkForbiddenTypes(node, path, pos)
	case *ast.ImportSpec:
		w.checkForbiddenImports(node, path, pos)
	case *ast.CallExpr:
		w.checkForbiddenCalls(node, path, pos)
	case *ast.GenDecl:
		w.checkGlobalState(node, path, pos)
	case *ast.RangeStmt:
		w.checkModernIteration(node, path, pos)
	case *ast.FuncDecl:
		w.checkContextHygiene(node, path, pos)
	}
}

// checkForbiddenTypes flags interface{} and any usage.
func (w *ArchWalker) checkForbiddenTypes(f *ast.Field, path string, pos token.Position) {
	if w.isInSysMethod() {
		return
	}
	if w.isInGenericTypeParams() {
		return
	}

	switch t := f.Type.(type) {
	case *ast.InterfaceType:
		if len(t.Methods.List) == 0 {
			if w.isAllowedAnyPath(path) {
				return
			}
			w.report(path, pos, "NO_ANY", "Use of 'interface{}' is forbidden. Use Generics [T any] or concrete types.")
		}
	case *ast.Ident:
		if t.Name == "any" {
			if w.isAllowedAnyPath(path) {
				return
			}
			w.report(path, pos, "NO_ANY", "Use of 'any' is forbidden. Use Generics [T any] or concrete types.")
		}
	}
}

func (w *ArchWalker) isInSysMethod() bool {
	for i := len(w.stack) - 1; i >= 0; i-- {
		if fd, ok := w.stack[i].(*ast.FuncDecl); ok {
			return fd.Name.Name == "Sys"
		}
	}
	return false
}

func (w *ArchWalker) isInGenericTypeParams() bool {
	if len(w.stack) < 3 {
		return false
	}
	parent := w.stack[len(w.stack)-2]
	grandParent := w.stack[len(w.stack)-3]

	fl, ok := parent.(*ast.FieldList)
	if !ok {
		return false
	}

	if ft, ok := grandParent.(*ast.FuncType); ok {
		return ft.TypeParams == fl
	}
	if ts, ok := grandParent.(*ast.TypeSpec); ok {
		return ts.TypeParams == fl
	}
	return false
}

func (w *ArchWalker) isAllowedAnyPath(path string) bool {
	// Allow in tests
	if strings.HasSuffix(path, "_test.go") {
		return true
	}
	// Allow in cmd/ for CLI tools that need viper
	if strings.Contains(path, "/cmd/") {
		return true
	}
	return false
}

// checkForbiddenImports flags log package.
func (w *ArchWalker) checkForbiddenImports(imp *ast.ImportSpec, path string, pos token.Position) {
	if imp.Path == nil {
		return
	}
	pathVal := strings.Trim(imp.Path.Value, "\"")
	if pathVal == "log" {
		// Allow in cmd/ (entry points)
		if strings.Contains(path, "/cmd/") {
			return
		}
		w.report(path, pos, "NO_LEGACY_LOG", "Import 'log' is forbidden. Use 'log/slog'.")
	}
}

// checkForbiddenCalls flags fmt.Print* and context.Background/TODO.
func (w *ArchWalker) checkForbiddenCalls(call *ast.CallExpr, path string, pos token.Position) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}
	pkgIdent, ok := sel.X.(*ast.Ident)
	if !ok {
		return
	}

	w.checkFmtPrint(pkgIdent.Name, sel.Sel.Name, path, pos)
	w.checkContextBackground(pkgIdent.Name, sel.Sel.Name, path, pos)
}

func (w *ArchWalker) checkFmtPrint(pkg, method, path string, pos token.Position) {
	if pkg != "fmt" || !strings.HasPrefix(method, "Print") {
		return
	}
	// Allow fmt in cmd/ and installer (CLI output)
	if strings.Contains(path, "/cmd/") || strings.Contains(path, "/internal/installer/") {
		return
	}
	w.report(path, pos, "NO_FMT_PRINT", "Avoid 'fmt.Print*'. Use structured logging (log/slog).")
}

func (w *ArchWalker) checkContextBackground(pkg, method, path string, pos token.Position) {
	if pkg != "context" {
		return
	}
	if method != "Background" && method != "TODO" {
		return
	}
	// Allow in tests and cmd/
	if strings.HasSuffix(path, "_test.go") || strings.Contains(path, "/cmd/") {
		return
	}
	w.report(path, pos, "CTX_HYGIENE", "Avoid 'context.Background/TODO' in logic. Pass context from caller.")
}

// checkGlobalState flags exported mutable global variables.
func (w *ArchWalker) checkGlobalState(decl *ast.GenDecl, path string, pos token.Position) {
	if decl.Tok != token.VAR {
		return
	}
	for _, spec := range decl.Specs {
		vSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}
		for _, name := range vSpec.Names {
			if ast.IsExported(name.Name) {
				// Allow error sentinels
				if strings.HasPrefix(name.Name, "Err") {
					continue
				}
				w.report(path, pos, "NO_GLOBALS", fmt.Sprintf("Exported mutable global variable '%s' is forbidden. Use Dependency Injection.", name.Name))
			}
		}
	}
}

// checkModernIteration suggests Go 1.22+ range syntax.
func (w *ArchWalker) checkModernIteration(rs *ast.RangeStmt, path string, pos token.Position) {
	if isBlank(rs.Key) && rs.Value == nil {
		w.report(path, pos, "MODERN_ITER", "Go 1.22+: Use 'for range x' instead of 'for _ := range x'.")
	}
}

func isBlank(expr ast.Expr) bool {
	if expr == nil {
		return false
	}
	id, ok := expr.(*ast.Ident)
	return ok && id.Name == "_"
}

// checkContextHygiene ensures context.Context is the first parameter.
func (w *ArchWalker) checkContextHygiene(fn *ast.FuncDecl, path string, pos token.Position) {
	if fn.Type.Params == nil || len(fn.Type.Params.List) == 0 {
		return
	}

	if w.isFirstParamContext(fn.Type.Params.List[0]) {
		return
	}

	// Check if any other param is context (which is bad)
	w.checkContextNotFirst(fn.Type.Params.List, path, pos)
}

func (w *ArchWalker) isFirstParamContext(field *ast.Field) bool {
	sel, ok := field.Type.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	pkg, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	return pkg.Name == "context" && sel.Sel.Name == "Context"
}

func (w *ArchWalker) checkContextNotFirst(fields []*ast.Field, path string, pos token.Position) {
	for i, field := range fields {
		if i == 0 {
			continue
		}
		sel, ok := field.Type.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		pkg, ok := sel.X.(*ast.Ident)
		if !ok {
			continue
		}
		if pkg.Name == "context" && sel.Sel.Name == "Context" {
			w.report(path, pos, "CTX_HYGIENE", "context.Context must be the first argument.")
			return
		}
	}
}
