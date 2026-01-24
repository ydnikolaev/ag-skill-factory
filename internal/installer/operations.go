package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"

	"github.com/ydnikolaev/ag-skill-factory/internal/diff"
)

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
		if err := i.copyDirWithRewrite(src, dst); err != nil {
			color.Red("Failed to update %s: %v", name, err)
		} else {
			result.UpdatedCount++
		}
	}
}

// updateRules syncs _standards and meta files to rules folder.
func (i *Installer) updateRules() error {
	color.Cyan("📋 Syncing rules/standards...")

	count := 0

	// Convert _standards to rules
	standardsPath := filepath.Join(i.Source, "_standards")
	if _, err := os.Stat(standardsPath); err == nil {
		n, err := i.convertStandardsToRules(standardsPath)
		if err != nil {
			return err
		}
		count += n
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
		} else {
			count++
		}
	}

	color.Green("✅ Synced %d rules", count)
	return nil
}
