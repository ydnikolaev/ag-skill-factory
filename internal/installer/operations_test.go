package installer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile_Basic(t *testing.T) {
	srcDir := t.TempDir()
	dstDir := t.TempDir()

	srcFile := filepath.Join(srcDir, "test.txt")
	dstFile := filepath.Join(dstDir, "test.txt")
	content := []byte("Hello, World!")

	mustWriteFile(t, srcFile, content)

	err := copyFile(srcFile, dstFile)
	if err != nil {
		t.Fatalf("copyFile failed: %v", err)
	}

	readContent, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read copied file: %v", err)
	}

	if string(readContent) != string(content) {
		t.Errorf("Content mismatch: expected %q, got %q", content, readContent)
	}
}

func TestCopyFile_NonExistent(t *testing.T) {
	dstDir := t.TempDir()
	err := copyFile("/nonexistent/file.txt", filepath.Join(dstDir, "out.txt"))
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}

func TestCopyFile_InvalidDestination(t *testing.T) {
	srcDir := t.TempDir()
	srcFile := filepath.Join(srcDir, "test.txt")
	mustWriteFile(t, srcFile, []byte("content"))

	// Try to copy to invalid path
	err := copyFile(srcFile, "/nonexistent/dir/file.txt")
	if err == nil {
		t.Error("Expected error for invalid destination")
	}
}

func TestExtractFirstHeading_Basic(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "standard heading",
			content:  "# My Protocol\n\nContent here",
			expected: "My Protocol",
		},
		{
			name:     "heading with trailing content",
			content:  "# TDD Protocol\n## Section\nMore content",
			expected: "TDD Protocol",
		},
		{
			name:     "no heading",
			content:  "Just plain text\nNo heading here",
			expected: "Rule",
		},
		{
			name:     "h2 only",
			content:  "## Not H1\n### Also not",
			expected: "Rule",
		},
		{
			name:     "empty content",
			content:  "",
			expected: "Rule",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractFirstHeading(tt.content)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestInstaller_CreateTargetDirs(t *testing.T) {
	targetDir := t.TempDir()
	inst := New("/tmp/source", targetDir)

	err := inst.createTargetDirs()
	if err != nil {
		t.Fatalf("createTargetDirs failed: %v", err)
	}

	dirs := []string{"skills", "rules", "workflows"}
	for _, dir := range dirs {
		path := filepath.Join(targetDir, dir)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Directory %s not created", dir)
		}
	}
}

func TestInstaller_CopySkillIfValid_WithSKILLmd(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create a valid skill with SKILL.md
	skillDir := filepath.Join(sourceDir, "test-skill")
	mustMkdirAll(t, skillDir)
	mustWriteFile(t, filepath.Join(skillDir, "SKILL.md"), []byte("# Test Skill"))

	inst := New(sourceDir, targetDir)
	mustMkdirAll(t, filepath.Join(targetDir, "skills"))

	result := &InstallResult{}
	err := inst.copySkillIfValid("test-skill", skillDir, result)
	if err != nil {
		t.Fatalf("copySkillIfValid failed: %v", err)
	}

	if result.SkillCount != 1 {
		t.Errorf("Expected 1 skill, got %d", result.SkillCount)
	}

	copiedSkill := filepath.Join(targetDir, "skills", "test-skill", "SKILL.md")
	if _, err := os.Stat(copiedSkill); os.IsNotExist(err) {
		t.Error("Skill was not copied")
	}
}

func TestInstaller_CopySkillIfValid_WithoutSKILLmd(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create a directory without SKILL.md
	skillDir := filepath.Join(sourceDir, "not-a-skill")
	mustMkdirAll(t, skillDir)
	mustWriteFile(t, filepath.Join(skillDir, "readme.md"), []byte("Not a skill"))

	inst := New(sourceDir, targetDir)
	mustMkdirAll(t, filepath.Join(targetDir, "skills"))

	result := &InstallResult{}
	err := inst.copySkillIfValid("not-a-skill", skillDir, result)
	if err != nil {
		t.Fatalf("copySkillIfValid failed: %v", err)
	}

	if result.SkillCount != 0 {
		t.Errorf("Expected 0 skills (no SKILL.md), got %d", result.SkillCount)
	}
}

func TestInstaller_ProcessEntry_SkipsNonDirectory(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create a file (not directory)
	mustWriteFile(t, filepath.Join(sourceDir, "file.txt"), []byte("content"))

	inst := New(sourceDir, targetDir)
	if err := inst.createTargetDirs(); err != nil {
		t.Fatalf("createTargetDirs failed: %v", err)
	}

	entries, _ := os.ReadDir(sourceDir)
	result := &InstallResult{}

	for _, entry := range entries {
		err := inst.processEntry(entry, result)
		if err != nil {
			t.Fatalf("processEntry failed: %v", err)
		}
	}

	// No skills should be counted for a file
	if result.SkillCount != 0 {
		t.Errorf("Expected 0 skills for file entry, got %d", result.SkillCount)
	}
}

func TestInstaller_ProcessEntry_SkipsReferences(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create references directory (should be skipped)
	refsDir := filepath.Join(sourceDir, "references")
	mustMkdirAll(t, refsDir)
	mustWriteFile(t, filepath.Join(refsDir, "SKILL.md"), []byte("# Fake"))

	inst := New(sourceDir, targetDir)
	if err := inst.createTargetDirs(); err != nil {
		t.Fatalf("createTargetDirs failed: %v", err)
	}

	entries, _ := os.ReadDir(sourceDir)
	result := &InstallResult{}

	for _, entry := range entries {
		err := inst.processEntry(entry, result)
		if err != nil {
			t.Fatalf("processEntry failed: %v", err)
		}
	}

	if result.SkillCount != 0 {
		t.Errorf("Expected 0 skills (references skipped), got %d", result.SkillCount)
	}
}

func TestInstaller_CopyDirWithRewrite_TransformsPaths(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create a skill with _standards references
	skillDir := filepath.Join(sourceDir, "test-skill")
	mustMkdirAll(t, skillDir)
	content := "# Skill\n\nSee `_standards/TDD_PROTOCOL.md` for TDD rules."
	mustWriteFile(t, filepath.Join(skillDir, "SKILL.md"), []byte(content))

	inst := New(sourceDir, targetDir)
	dstDir := filepath.Join(targetDir, "test-skill")

	err := inst.copyDirWithRewrite(skillDir, dstDir)
	if err != nil {
		t.Fatalf("copyDirWithRewrite failed: %v", err)
	}

	copiedContent, err := os.ReadFile(filepath.Join(dstDir, "SKILL.md"))
	if err != nil {
		t.Fatalf("Failed to read copied file: %v", err)
	}

	if contains(string(copiedContent), "_standards/") {
		t.Error("_standards/ paths were not rewritten")
	}
	if !contains(string(copiedContent), ".agent/rules/tdd_protocol.md") {
		t.Error("Path not transformed to .agent/rules/")
	}
}

func TestInstaller_CopyFileWithRewrite(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	srcFile := filepath.Join(sourceDir, "test.md")
	content := "Check `_standards/GIT_PROTOCOL.md` for git rules."
	mustWriteFile(t, srcFile, []byte(content))

	inst := New(sourceDir, targetDir)
	dstFile := filepath.Join(targetDir, "test.md")

	err := inst.copyFileWithRewrite(srcFile, dstFile)
	if err != nil {
		t.Fatalf("copyFileWithRewrite failed: %v", err)
	}

	result, _ := os.ReadFile(dstFile)
	if contains(string(result), "_standards/") {
		t.Error("Path not rewritten")
	}
	if !contains(string(result), ".agent/rules/git_protocol.md") {
		t.Error("Expected transformed path")
	}
}
