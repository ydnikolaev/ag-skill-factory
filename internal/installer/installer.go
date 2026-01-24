// Package installer handles skill installation and synchronization.
package installer

import (
	"fmt"
	"os"
	"path/filepath"

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

// ForceRefresh copies all skills with path rewriting, skipping diff check.
func (i *Installer) ForceRefresh() (*InstallResult, error) {
	result := &InstallResult{}

	color.Cyan("🔄 Force refreshing all skills...")

	// Read all skills from source
	entries, err := os.ReadDir(i.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to read source: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		if name == "_standards" || name == "references" {
			continue
		}

		srcPath := filepath.Join(i.Source, name)
		skillFile := filepath.Join(srcPath, "SKILL.md")
		if _, err := os.Stat(skillFile); os.IsNotExist(err) {
			continue
		}

		targetPath := filepath.Join(i.Target, "skills", name)
		if err := i.copyDirWithRewrite(srcPath, targetPath); err != nil {
			color.Yellow("Warning: failed to copy %s: %v", name, err)
			continue
		}
		result.SkillCount++
	}

	// Also refresh rules
	if err := i.updateRules(); err != nil {
		color.Yellow("Warning: failed to update rules: %v", err)
	}

	return result, nil
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
