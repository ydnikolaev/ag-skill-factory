// Package installer handles skill installation and synchronization.
package installer

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/afero"

	"github.com/ydnikolaev/ag-skill-factory/internal/diff"
)

// Installer handles skill installation and synchronization.
type Installer struct {
	Source   string
	Target   string
	Prompter Prompter
	Fs       afero.Fs
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

// New creates a new Installer with default StdinPrompter and OsFs.
func New(source, target string) *Installer {
	return &Installer{
		Source:   source,
		Target:   target,
		Prompter: &StdinPrompter{},
		Fs:       afero.NewOsFs(),
	}
}

// NewWithPrompter creates an Installer with a custom Prompter (for testing).
func NewWithPrompter(source, target string, prompter Prompter) *Installer {
	return &Installer{
		Source:   source,
		Target:   target,
		Prompter: prompter,
		Fs:       afero.NewOsFs(),
	}
}

// NewWithFs creates an Installer with custom Prompter and Fs (for testing).
func NewWithFs(source, target string, prompter Prompter, fs afero.Fs) *Installer {
	return &Installer{
		Source:   source,
		Target:   target,
		Prompter: prompter,
		Fs:       fs,
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
	entries, err := afero.ReadDir(i.Fs, i.Source)
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
		if _, err := i.Fs.Stat(skillFile); err != nil {
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
	entries, err := afero.ReadDir(i.Fs, localSkillsPath)
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

	if !i.Prompter.Confirm("Backport these changes to factory?") {
		color.Yellow("Cancelled")
		return nil
	}

	return i.copyDir(localPath, factoryPath)
}
