package installer

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestNew(t *testing.T) {
	inst := New("/source", "/target")
	if inst.Source != "/source" {
		t.Errorf("expected source /source, got %s", inst.Source)
	}
	if inst.Target != "/target" {
		t.Errorf("expected target /target, got %s", inst.Target)
	}
	if inst.Fs == nil {
		t.Error("expected Fs to be set")
	}
}

func TestNewWithFs(t *testing.T) {
	fs := afero.NewMemMapFs()
	inst := NewWithFs("/source", "/target", fs)
	if inst.Fs != fs {
		t.Error("expected custom Fs")
	}
}

func TestInstall_Success(t *testing.T) {
	fs := afero.NewMemMapFs()

	// Setup source blueprint
	_ = fs.MkdirAll("/source/skills/test-skill", 0o755)
	_ = afero.WriteFile(fs, "/source/skills/test-skill/SKILL.md", []byte("# Test"), 0o644)

	_ = fs.MkdirAll("/source/workflows", 0o755)
	_ = afero.WriteFile(fs, "/source/workflows/test.md", []byte("# Workflow"), 0o644)

	_ = fs.MkdirAll("/source/rules", 0o755)
	_ = afero.WriteFile(fs, "/source/rules/TEAM.md", []byte("# Team"), 0o644)

	_ = fs.MkdirAll("/source/standards", 0o755)
	_ = afero.WriteFile(fs, "/source/standards/TDD.md", []byte("# TDD"), 0o644)

	inst := NewWithFs("/source", "/target", fs)
	result, err := inst.Install()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.SkillCount != 1 {
		t.Errorf("expected 1 skill, got %d", result.SkillCount)
	}
	if result.WorkflowCount != 1 {
		t.Errorf("expected 1 workflow, got %d", result.WorkflowCount)
	}
	if result.RuleCount != 1 {
		t.Errorf("expected 1 rule, got %d", result.RuleCount)
	}
	if result.StandardCount != 1 {
		t.Errorf("expected 1 standard, got %d", result.StandardCount)
	}

	// Verify files were copied
	if _, err := fs.Stat("/target/skills/test-skill/SKILL.md"); err != nil {
		t.Error("expected skill file to exist")
	}
	if _, err := fs.Stat("/target/workflows/test.md"); err != nil {
		t.Error("expected workflow file to exist")
	}
	if _, err := fs.Stat("/target/rules/TEAM.md"); err != nil {
		t.Error("expected rule file to exist")
	}
	if _, err := fs.Stat("/target/standards/TDD.md"); err != nil {
		t.Error("expected standard file to exist")
	}
}

func TestInstall_ReplacesExisting(t *testing.T) {
	fs := afero.NewMemMapFs()

	// Setup existing target
	_ = fs.MkdirAll("/target/skills/old-skill", 0o755)
	_ = afero.WriteFile(fs, "/target/skills/old-skill/SKILL.md", []byte("old"), 0o644)

	// Setup source
	_ = fs.MkdirAll("/source/skills/new-skill", 0o755)
	_ = afero.WriteFile(fs, "/source/skills/new-skill/SKILL.md", []byte("new"), 0o644)

	inst := NewWithFs("/source", "/target", fs)
	_, err := inst.Install()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Old skill should be gone
	if _, err := fs.Stat("/target/skills/old-skill"); err == nil {
		t.Error("expected old-skill to be removed")
	}

	// New skill should exist
	if _, err := fs.Stat("/target/skills/new-skill/SKILL.md"); err != nil {
		t.Error("expected new-skill to exist")
	}
}

func TestInstall_EmptyCategory(t *testing.T) {
	fs := afero.NewMemMapFs()

	// Setup source with only skills
	_ = fs.MkdirAll("/source/skills/test-skill", 0o755)
	_ = afero.WriteFile(fs, "/source/skills/test-skill/SKILL.md", []byte("# Test"), 0o644)

	inst := NewWithFs("/source", "/target", fs)
	result, err := inst.Install()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.SkillCount != 1 {
		t.Errorf("expected 1 skill, got %d", result.SkillCount)
	}
	if result.WorkflowCount != 0 {
		t.Errorf("expected 0 workflows, got %d", result.WorkflowCount)
	}
}

func TestCopyDir(t *testing.T) {
	fs := afero.NewMemMapFs()

	// Setup source
	_ = fs.MkdirAll("/src/subdir", 0o755)
	_ = afero.WriteFile(fs, "/src/file.txt", []byte("content"), 0o644)
	_ = afero.WriteFile(fs, "/src/subdir/nested.txt", []byte("nested"), 0o644)

	inst := NewWithFs("/source", "/target", fs)
	err := inst.copyDir("/src", "/dst")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check files exist
	if _, err := fs.Stat("/dst/file.txt"); err != nil {
		t.Error("expected file.txt to exist")
	}
	if _, err := fs.Stat("/dst/subdir/nested.txt"); err != nil {
		t.Error("expected subdir/nested.txt to exist")
	}

	// Check content
	content, _ := afero.ReadFile(fs, "/dst/file.txt")
	if string(content) != "content" {
		t.Errorf("expected 'content', got '%s'", content)
	}
}

func TestCountDirs(t *testing.T) {
	fs := afero.NewMemMapFs()

	_ = fs.MkdirAll("/test/dir1", 0o755)
	_ = fs.MkdirAll("/test/dir2", 0o755)
	_ = fs.MkdirAll("/test/.hidden", 0o755)
	_ = afero.WriteFile(fs, "/test/file.txt", []byte("x"), 0o644)

	inst := NewWithFs("/source", "/target", fs)
	count := inst.countDirs("/test")

	if count != 2 {
		t.Errorf("expected 2 dirs, got %d", count)
	}
}

func TestCountFiles(t *testing.T) {
	fs := afero.NewMemMapFs()

	_ = fs.MkdirAll("/test", 0o755)
	_ = afero.WriteFile(fs, "/test/file1.md", []byte("x"), 0o644)
	_ = afero.WriteFile(fs, "/test/file2.md", []byte("x"), 0o644)
	_ = afero.WriteFile(fs, "/test/file.txt", []byte("x"), 0o644) // Not .md

	inst := NewWithFs("/source", "/target", fs)
	count := inst.countFiles("/test")

	if count != 2 {
		t.Errorf("expected 2 md files, got %d", count)
	}
}

func TestInstall_MultipleSkills(t *testing.T) {
	fs := afero.NewMemMapFs()

	skills := []string{"skill1", "skill2", "skill3"}
	for _, s := range skills {
		_ = fs.MkdirAll(filepath.Join("/source/skills", s), 0o755)
		_ = afero.WriteFile(fs, filepath.Join("/source/skills", s, "SKILL.md"), []byte("# "+s), 0o644)
	}

	inst := NewWithFs("/source", "/target", fs)
	result, err := inst.Install()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.SkillCount != 3 {
		t.Errorf("expected 3 skills, got %d", result.SkillCount)
	}
}
