# Static Analysis Command Reference

Quick reference for code analysis commands.

## Go Projects

### LOC (Lines of Code)

```bash
# Find largest files
find . -name "*.go" ! -name "*_test.go" -exec wc -l {} \; | sort -rn | head -20

# Count total LOC
find . -name "*.go" ! -name "*_test.go" -exec cat {} \; | wc -l

# Files over 300 lines (flag threshold)
find . -name "*.go" ! -name "*_test.go" -exec sh -c 'wc -l "$1" | awk "\$1 > 300 {print}"' _ {} \;

# Files over 500 lines (god file threshold)
find . -name "*.go" ! -name "*_test.go" -exec sh -c 'wc -l "$1" | awk "\$1 > 500 {print}"' _ {} \;
```

### Missing Tests

```bash
# Files without corresponding _test.go
find . -name "*.go" ! -name "*_test.go" | while read f; do
  test_file="${f%.go}_test.go"
  [ ! -f "$test_file" ] && echo "No test: $f"
done

# Count tested vs untested
echo "Tested: $(find . -name "*_test.go" | wc -l)"
echo "Source: $(find . -name "*.go" ! -name "*_test.go" | wc -l)"
```

### Complexity Analysis

```bash
# Using golangci-lint (if configured)
golangci-lint run --out-format=json 2>/dev/null | jq '.Issues[] | select(.FromLinter == "gocyclo")'

# Using gocyclo directly
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
gocyclo -over 10 .

# Using goconst (for magic values)
go install github.com/jgautheron/goconst/cmd/goconst@latest
goconst ./...
```

### Import Analysis

```bash
# Circular dependencies
go install golang.org/x/tools/cmd/guru@latest
guru -scope ./... imports ./...

# Package dependencies (basic)
go list -m -json all | jq '.Path'

# Check specific import patterns (domain → infra leak)
grep -r "import.*internal/infra" internal/domain/ 2>/dev/null && echo "VIOLATION: domain imports infra"
```

### Dead Code Detection

```bash
# Using deadcode
go install golang.org/x/tools/cmd/deadcode@latest
deadcode -test ./...

# Using staticcheck
staticcheck -checks U1000 ./...
```

## Nuxt/Vue Projects

### LOC

```bash
# Largest Vue files
find . -name "*.vue" -exec wc -l {} \; | sort -rn | head -20

# Largest TypeScript files
find . -name "*.ts" ! -path "./node_modules/*" -exec wc -l {} \; | sort -rn | head -20
```

### Missing Tests

```bash
# Components without tests
find ./components -name "*.vue" | while read f; do
  basename=$(basename "$f" .vue)
  find ./tests -name "*${basename}*" | grep -q . || echo "No test: $f"
done
```

### Complexity (ESLint)

```bash
# Check with ESLint complexity rule
npx eslint --rule 'complexity: ["error", 10]' ./

# List files with most issues
npx eslint --format compact ./ | cut -d: -f1 | sort | uniq -c | sort -rn | head -20
```

## Universal Commands

### Directory Structure Analysis

```bash
# Files per directory (hotspots)
find . -type f -name "*.go" | xargs dirname | sort | uniq -c | sort -rn | head -20

# Deep nesting (smell)
find . -type d -depth +5
```

### Git-Based Analysis

```bash
# Most changed files (hotspots)
git log --pretty=format: --name-only | sort | uniq -c | sort -rn | head -20

# Files with most authors (ownership unclear)
for f in $(find . -name "*.go" | head -50); do
  authors=$(git log --pretty=format:"%an" "$f" 2>/dev/null | sort -u | wc -l)
  echo "$authors $f"
done | sort -rn | head -20
```

## Output Interpretation

| Metric | Flag | Action |
|--------|------|--------|
| LOC > 300 | ⚠️ Warning | Consider splitting |
| LOC > 500 | 🔴 Critical | Must split (god file) |
| Cyclomatic > 10 | ⚠️ Warning | Refactor logic |
| Cyclomatic > 15 | 🔴 Critical | Must simplify |
| No test file | ⚠️ Warning | Add tests |
| Domain → Infra import | 🔴 Critical | Architecture violation |
| Circular dep | 🔴 Critical | Restructure packages |
