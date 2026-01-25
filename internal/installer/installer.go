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
	Source string
	Target string
	Fs     afero.Fs
}

// InstallResult holds the result of an install operation.
type InstallResult struct {
	SkillCount    int
	WorkflowCount int
	RuleCount     int
	StandardCount int
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

// Install copies the entire blueprint to target .agent/ folder.
func (i *Installer) Install() (*InstallResult, error) {
	result := &InstallResult{}

	color.Cyan("🔧 Installing Antigravity Blueprint...")

	// Create target directory
	if err := i.Fs.MkdirAll(i.Target, 0o755); err != nil {
		return nil, fmt.Errorf("failed to create target: %w", err)
	}

	// Copy each category
	categories := []struct {
		name    string
		counter *int
		isDir   bool
	}{
		{"skills", &result.SkillCount, true},
		{"workflows", &result.WorkflowCount, false},
		{"rules", &result.RuleCount, false},
		{"standards", &result.StandardCount, false},
	}

	for _, cat := range categories {
		srcPath := filepath.Join(i.Source, cat.name)
		dstPath := filepath.Join(i.Target, cat.name)

		if _, err := i.Fs.Stat(srcPath); os.IsNotExist(err) {
			continue
		}

		if err := i.copyDir(srcPath, dstPath); err != nil {
			return nil, fmt.Errorf("failed to copy %s: %w", cat.name, err)
		}

		// Count items
		if cat.isDir {
			*cat.counter = i.countDirs(dstPath)
		} else {
			*cat.counter = i.countFiles(dstPath)
		}

		color.White("   📦 %s: %d", cat.name, *cat.counter)
	}

	return result, nil
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
