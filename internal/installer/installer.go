package installer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"

	"github.com/yuranikolaev/ag-skill-factory/internal/diff"
)

// Installer handles skill installation and synchronization.
type Installer struct {
	Source string
	Target string
}

// InstallResult holds the result of an install operation.
type InstallResult struct {
	SkillCount int
	RuleCount  int
}

// UpdateResult holds the result of an update operation.
type UpdateResult struct {
	UpdatedCount int
}

// New creates a new Installer.
func New(source, target string) *Installer {
	return &Installer{
		Source: source,
		Target: target,
	}
}

// Install copies all skills to target and converts standards to rules.
func (i *Installer) Install() (*InstallResult, error) {
	result := &InstallResult{}

	if err := i.createTargetDirs(); err != nil {
		return nil, err
	}

	if err := i.processSourceEntries(result); err != nil {
		return nil, err
	}

	if err := i.copyMetaFilesToRules(result); err != nil {
		return nil, err
	}

	return result, nil
}

// createTargetDirs creates the target directory structure.
func (i *Installer) createTargetDirs() error {
	dirs := []string{
		filepath.Join(i.Target, "skills"),
		filepath.Join(i.Target, "rules"),
		filepath.Join(i.Target, "workflows"),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("failed to create %s: %w", dir, err)
		}
	}
	return nil
}

// processSourceEntries processes all entries in the source directory.
func (i *Installer) processSourceEntries(result *InstallResult) error {
	entries, err := os.ReadDir(i.Source)
	if err != nil {
		return fmt.Errorf("failed to read source: %w", err)
	}

	for _, entry := range entries {
		if err := i.processEntry(entry, result); err != nil {
			return err
		}
	}
	return nil
}

// processEntry processes a single source entry.
func (i *Installer) processEntry(entry os.DirEntry, result *InstallResult) error {
	if !entry.IsDir() {
		return nil
	}

	name := entry.Name()
	srcPath := filepath.Join(i.Source, name)

	switch name {
	case "_standards":
		count, err := i.convertStandardsToRules(srcPath)
		if err != nil {
			return fmt.Errorf("failed to convert standards: %w", err)
		}
		result.RuleCount = count
	case "references":
		// Skip references folder.
	default:
		if err := i.copySkillIfValid(name, srcPath, result); err != nil {
			return err
		}
	}
	return nil
}

// copySkillIfValid copies a skill if it has SKILL.md.
func (i *Installer) copySkillIfValid(name, srcPath string, result *InstallResult) error {
	skillFile := filepath.Join(srcPath, "SKILL.md")
	if _, err := os.Stat(skillFile); os.IsNotExist(err) {
		return nil
	}

	targetPath := filepath.Join(i.Target, "skills", name)
	if err := i.copyDirWithRewrite(srcPath, targetPath); err != nil {
		return fmt.Errorf("failed to copy skill %s: %w", name, err)
	}
	result.SkillCount++
	return nil
}

// copyDirWithRewrite copies a directory, rewriting paths in .md files.
// Transforms: _standards/X.md → .agent/rules/x.md
func (i *Installer) copyDirWithRewrite(src, dst string) error {
	_ = os.RemoveAll(dst)

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		// Rewrite paths in markdown files
		if strings.HasSuffix(path, ".md") {
			return i.copyFileWithRewrite(path, dstPath)
		}

		return copyFile(path, dstPath)
	})
}

// copyFileWithRewrite copies a file, rewriting _standards paths.
func (i *Installer) copyFileWithRewrite(src, dst string) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// Rewrite _standards references to .agent/rules
	newContent := rewriteStandardsPaths(string(content))

	return os.WriteFile(dst, []byte(newContent), 0o644)
}

// copyMetaFilesToRules copies TEAM.md and PIPELINE.md to rules.
func (i *Installer) copyMetaFilesToRules(result *InstallResult) error {
	files := []string{"TEAM.md", "PIPELINE.md"}
	for _, file := range files {
		src := filepath.Join(i.Source, file)
		if _, err := os.Stat(src); err != nil {
			continue
		}
		dst := filepath.Join(i.Target, "rules", strings.ToLower(file))
		if err := i.convertToRule(src, dst); err != nil {
			color.Yellow("Warning: failed to convert %s: %v", file, err)
		} else {
			result.RuleCount++
		}
	}
	return nil
}

// Update updates skills from source, showing diffs.
func (i *Installer) Update() (*UpdateResult, error) {
	result := &UpdateResult{}

	// Update skills
	localSkillsPath := filepath.Join(i.Target, "skills")
	entries, err := os.ReadDir(localSkillsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read local skills: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		i.updateSingleSkill(entry.Name(), localSkillsPath, result)
	}

	// Update rules/standards
	if err := i.updateRules(); err != nil {
		color.Yellow("Warning: failed to update rules: %v", err)
	}

	return result, nil
}

// updateRules syncs _standards and meta files to rules folder.
func (i *Installer) updateRules() error {
	// Convert _standards to rules
	standardsPath := filepath.Join(i.Source, "_standards")
	if _, err := os.Stat(standardsPath); err == nil {
		if _, err := i.convertStandardsToRules(standardsPath); err != nil {
			return err
		}
	}

	// Copy meta files (TEAM.md, PIPELINE.md)
	files := []string{"TEAM.md", "PIPELINE.md"}
	for _, file := range files {
		src := filepath.Join(i.Source, file)
		if _, err := os.Stat(src); err != nil {
			continue
		}
		dst := filepath.Join(i.Target, "rules", strings.ToLower(file))
		if err := i.convertToRule(src, dst); err != nil {
			color.Yellow("Warning: failed to convert %s: %v", file, err)
		}
	}

	return nil
}

// updateSingleSkill updates a single skill with diff confirmation.
func (i *Installer) updateSingleSkill(name, localSkillsPath string, result *UpdateResult) {
	localPath := filepath.Join(localSkillsPath, name)
	sourcePath := filepath.Join(i.Source, name)

	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return
	}

	changes, err := diff.CompareDirectories(localPath, sourcePath)
	if err != nil {
		color.Yellow("Warning: failed to compare %s: %v", name, err)
		return
	}

	if len(changes) == 0 {
		return
	}

	i.showChangesAndApply(name, changes, sourcePath, localPath, result)
}

// showChangesAndApply displays diff and applies if confirmed.
func (i *Installer) showChangesAndApply(name string, changes []string, src, dst string, result *UpdateResult) {
	color.Cyan("\n=== %s ===", name)
	for _, change := range changes {
		fmt.Println(change)
	}

	if confirm("Apply these changes?") {
		if err := copyDir(src, dst); err != nil {
			color.Red("Failed to update %s: %v", name, err)
		} else {
			result.UpdatedCount++
		}
	}
}

// Backport copies a skill from local back to factory.
func (i *Installer) Backport(skillName string) error {
	localPath := filepath.Join(i.Target, "skills", skillName)
	factoryPath := filepath.Join(i.Source, skillName)

	changes, err := diff.CompareDirectories(factoryPath, localPath)
	if err != nil {
		return fmt.Errorf("failed to compare: %w", err)
	}

	if len(changes) == 0 {
		color.Green("No changes to backport")
		return nil
	}

	color.Cyan("\n=== Changes to backport ===")
	for _, change := range changes {
		fmt.Println(change)
	}

	if !confirm("Backport these changes to factory?") {
		color.Yellow("Cancelled")
		return nil
	}

	return copyDir(localPath, factoryPath)
}

// convertStandardsToRules converts _standards files to rule format.
func (i *Installer) convertStandardsToRules(standardsPath string) (int, error) {
	count := 0

	entries, err := os.ReadDir(standardsPath)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		src := filepath.Join(standardsPath, entry.Name())
		dst := filepath.Join(i.Target, "rules", strings.ToLower(entry.Name()))

		if err := i.convertToRule(src, dst); err != nil {
			return count, err
		}
		count++
	}

	return count, nil
}

// convertToRule converts a markdown file to rule format with YAML frontmatter.
func (i *Installer) convertToRule(src, dst string) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	description := extractFirstHeading(string(content))

	var builder strings.Builder
	builder.WriteString("---\n")
	builder.WriteString(fmt.Sprintf("description: %s\n", description))
	builder.WriteString("---\n\n")
	builder.WriteString(string(content))

	return os.WriteFile(dst, []byte(builder.String()), 0o644)
}

// extractFirstHeading extracts the first markdown heading.
func extractFirstHeading(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return "Rule"
}

// copyDir copies a directory recursively.
func copyDir(src, dst string) error {
	_ = os.RemoveAll(dst)

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		return copyFile(path, dstPath)
	})
}

// copyFile copies a single file.
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() { _ = srcFile.Close() }()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() { _ = dstFile.Close() }()

	_, err = io.Copy(dstFile, srcFile)

	return err
}

// confirm asks for user confirmation.
func confirm(message string) bool {
	fmt.Printf("%s [y/n]: ", message)
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}

// rewriteStandardsPaths transforms _standards/ references to .agent/rules/.
// Example: `_standards/TDD_PROTOCOL.md` → `.agent/rules/tdd_protocol.md`
func rewriteStandardsPaths(content string) string {
	// Pattern: _standards/SOMETHING.md → .agent/rules/something.md
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.Contains(line, "_standards/") {
			lines[i] = rewriteStandardsLine(line)
		}
	}
	return strings.Join(lines, "\n")
}

// rewriteStandardsLine rewrites a single line's _standards references.
func rewriteStandardsLine(line string) string {
	// Find _standards/XXX.md patterns and replace
	// Handle: `_standards/TDD_PROTOCOL.md` → `.agent/rules/tdd_protocol.md`
	result := line

	// Find all occurrences of _standards/something.md
	idx := strings.Index(result, "_standards/")
	for idx != -1 {
		// Find the end of the path (space, `, ), ], or end of line)
		endIdx := idx + len("_standards/")
		for endIdx < len(result) {
			c := result[endIdx]
			if c == ' ' || c == '`' || c == ')' || c == ']' || c == '"' || c == '\'' {
				break
			}
			endIdx++
		}

		oldPath := result[idx:endIdx]
		// Extract filename from _standards/FILENAME.md
		filename := strings.TrimPrefix(oldPath, "_standards/")
		newPath := ".agent/rules/" + strings.ToLower(filename)

		result = result[:idx] + newPath + result[endIdx:]

		// Look for next occurrence
		idx = strings.Index(result[idx+len(newPath):], "_standards/")
		if idx != -1 {
			idx += idx + len(newPath)
		}
	}

	return result
}

