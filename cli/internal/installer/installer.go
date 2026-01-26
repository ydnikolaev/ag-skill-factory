// Package installer handles blueprint installation to project .agent/ folder.
package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/afero"
)

// Installer handles blueprint installation.
type Installer struct {
	Source      string
	Target      string
	Fs          afero.Fs
	SkillFilter map[string]bool // nil = all skills
}

// InstallResult holds the result of an install operation.
type InstallResult struct {
	SkillCount    int
	WorkflowCount int
	RuleCount     int
	StandardCount int
	TemplateCount int
}

// New creates a new Installer with OsFs.
func New(source, target string) *Installer {
	return &Installer{
		Source: source,
		Target: target,
		Fs:     afero.NewOsFs(),
	}
}

// NewWithFs creates an Installer with custom Fs (for testing).
func NewWithFs(source, target string, fs afero.Fs) *Installer {
	return &Installer{
		Source: source,
		Target: target,
		Fs:     fs,
	}
}

// SetSkillFilter sets which skills to install (nil = all).
func (i *Installer) SetSkillFilter(skills []string) {
	if len(skills) == 0 {
		i.SkillFilter = nil
		return
	}
	i.SkillFilter = make(map[string]bool)
	for _, s := range skills {
		i.SkillFilter[s] = true
	}
}

// Install copies the entire blueprint to target .agent/ folder.
func (i *Installer) Install() (*InstallResult, error) {
	result := &InstallResult{}

	color.Cyan("🔧 Installing Antigravity Blueprint...")

	// Create target directory
	if err := i.Fs.MkdirAll(i.Target, 0o755); err != nil {
		return nil, fmt.Errorf("failed to create target: %w", err)
	}

	// Copy skills (with filtering)
	skillsSrc := filepath.Join(i.Source, "skills")
	skillsDst := filepath.Join(i.Target, "skills")
	if _, err := i.Fs.Stat(skillsSrc); err == nil {
		if err := i.copySkills(skillsSrc, skillsDst); err != nil {
			return nil, fmt.Errorf("failed to copy skills: %w", err)
		}
		result.SkillCount = i.countDirs(skillsDst)
		color.White("   📦 skills: %d", result.SkillCount)
	}

	// Copy other categories (no filtering)
	others := []struct {
		name    string
		counter *int
	}{
		{"workflows", &result.WorkflowCount},
		{"rules", &result.RuleCount},
		{"standards", &result.StandardCount},
	}

	for _, cat := range others {
		srcPath := filepath.Join(i.Source, cat.name)
		dstPath := filepath.Join(i.Target, cat.name)

		if _, err := i.Fs.Stat(srcPath); os.IsNotExist(err) {
			continue
		}

		if err := i.copyDir(srcPath, dstPath); err != nil {
			return nil, fmt.Errorf("failed to copy %s: %w", cat.name, err)
		}

		*cat.counter = i.countFiles(dstPath)
		color.White("   📦 %s: %d", cat.name, *cat.counter)
	}

	// Copy project/docs/templates from dist/project/ (sibling of _agent/)
	// Source is dist/_agent/, so dist/project/ is at ../project/
	distDir := filepath.Dir(i.Source) // dist/_agent/ -> dist/
	templatesSrc := filepath.Join(distDir, "project", "docs", "templates")
	
	// Target project is sibling of .agent/
	projectDir := filepath.Dir(i.Target) // .agent/ -> cwd
	templatesDst := filepath.Join(projectDir, "project", "docs", "templates")
	
	if _, err := i.Fs.Stat(templatesSrc); err == nil {
		if err := i.copyDir(templatesSrc, templatesDst); err != nil {
			return nil, fmt.Errorf("failed to copy templates: %w", err)
		}
		result.TemplateCount = i.countFiles(templatesDst)
		color.White("   📄 templates: %d", result.TemplateCount)
	}

	return result, nil
}

// copySkills copies skills respecting the filter.
func (i *Installer) copySkills(src, dst string) error {
	entries, err := afero.ReadDir(i.Fs, src)
	if err != nil {
		return err
	}

	// Remove existing skills directory (replace mode)
	_ = i.Fs.RemoveAll(dst)

	if err := i.Fs.MkdirAll(dst, 0o755); err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		skillName := entry.Name()

		// Apply filter
		if i.SkillFilter != nil && !i.SkillFilter[skillName] {
			continue
		}

		srcSkill := filepath.Join(src, skillName)
		dstSkill := filepath.Join(dst, skillName)
		if err := i.copyDir(srcSkill, dstSkill); err != nil {
			return err
		}
	}

	return nil
}

// copyDir recursively copies a directory.
func (i *Installer) copyDir(src, dst string) error {
	// Remove existing destination
	_ = i.Fs.RemoveAll(dst)

	return afero.Walk(i.Fs, src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return i.Fs.MkdirAll(dstPath, info.Mode())
		}

		return i.copyFile(path, dstPath)
	})
}

// copyFile copies a single file.
func (i *Installer) copyFile(src, dst string) error {
	content, err := afero.ReadFile(i.Fs, src)
	if err != nil {
		return err
	}
	return afero.WriteFile(i.Fs, dst, content, 0o644)
}

// countDirs counts subdirectories in a path.
func (i *Installer) countDirs(path string) int {
	count := 0
	entries, err := afero.ReadDir(i.Fs, path)
	if err != nil {
		return 0
	}
	for _, e := range entries {
		if e.IsDir() && !strings.HasPrefix(e.Name(), ".") {
			count++
		}
	}
	return count
}

// countFiles counts .md files in a path.
func (i *Installer) countFiles(path string) int {
	count := 0
	entries, err := afero.ReadDir(i.Fs, path)
	if err != nil {
		return 0
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
			count++
		}
	}
	return count
}
