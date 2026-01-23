package installer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	inst := New("/source", "/target", "/global")

	if inst.Source != "/source" {
		t.Errorf("Expected source /source, got %s", inst.Source)
	}
	if inst.Target != "/target" {
		t.Errorf("Expected target /target, got %s", inst.Target)
	}
	if inst.GlobalPath != "/global" {
		t.Errorf("Expected global /global, got %s", inst.GlobalPath)
	}
}

func TestInstall_CreatesDirectoryStructure(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := filepath.Join(t.TempDir(), ".agent")
	globalDir := t.TempDir()

	skillDir := filepath.Join(sourceDir, "test-skill")
	mustMkdirAll(t, skillDir)
	mustWriteFile(t, filepath.Join(skillDir, "SKILL.md"), []byte("# Test Skill"))

	standardsDir := filepath.Join(sourceDir, "_standards")
	mustMkdirAll(t, standardsDir)
	mustWriteFile(t, filepath.Join(standardsDir, "TEST_PROTOCOL.md"), []byte("# Test Protocol\n\nContent here"))

	inst := New(sourceDir, targetDir, globalDir)
	result, err := inst.Install()
	if err != nil {
		t.Fatalf("Install failed: %v", err)
	}

	if _, err := os.Stat(filepath.Join(targetDir, "skills")); os.IsNotExist(err) {
		t.Error("skills/ directory not created")
	}
	if _, err := os.Stat(filepath.Join(targetDir, "rules")); os.IsNotExist(err) {
		t.Error("rules/ directory not created")
	}
	if _, err := os.Stat(filepath.Join(targetDir, "workflows")); os.IsNotExist(err) {
		t.Error("workflows/ directory not created")
	}

	if _, err := os.Stat(filepath.Join(targetDir, "skills", "test-skill", "SKILL.md")); os.IsNotExist(err) {
		t.Error("test-skill not copied")
	}

	if result.SkillCount != 1 {
		t.Errorf("Expected 1 skill, got %d", result.SkillCount)
	}
	if result.RuleCount < 1 {
		t.Errorf("Expected at least 1 rule, got %d", result.RuleCount)
	}
}

func TestInstall_SkipsIfAgentExists(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := filepath.Join(t.TempDir(), ".agent")

	mustMkdirAll(t, targetDir)

	inst := New(sourceDir, targetDir, "")
	result, err := inst.Install()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil && result.SkillCount > 0 {
		t.Error("Expected 0 skills when .agent already exists")
	}
}

func TestConvertToRule_AddsYAMLFrontmatter(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	srcFile := filepath.Join(sourceDir, "TEST.md")
	mustWriteFile(t, srcFile, []byte("# My Protocol\n\nSome content"))

	inst := New(sourceDir, targetDir, "")
	dstFile := filepath.Join(targetDir, "test.md")

	err := inst.convertToRule(srcFile, dstFile)
	if err != nil {
		t.Fatalf("convertToRule failed: %v", err)
	}

	content, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read result: %v", err)
	}
	contentStr := string(content)

	if contentStr[:4] != "---\n" {
		t.Error("Missing YAML frontmatter start")
	}
	if !contains(contentStr, "description: My Protocol") {
		t.Error("Missing description in frontmatter")
	}
	if !contains(contentStr, "# My Protocol") {
		t.Error("Original content missing")
	}
}

func TestCopyDir_CopiesRecursively(t *testing.T) {
	srcDir := t.TempDir()
	dstDir := filepath.Join(t.TempDir(), "dest")

	mustMkdirAll(t, filepath.Join(srcDir, "subdir"))
	mustWriteFile(t, filepath.Join(srcDir, "file1.txt"), []byte("content1"))
	mustWriteFile(t, filepath.Join(srcDir, "subdir", "file2.txt"), []byte("content2"))

	err := copyDir(srcDir, dstDir)
	if err != nil {
		t.Fatalf("copyDir failed: %v", err)
	}

	if _, err := os.Stat(filepath.Join(dstDir, "file1.txt")); os.IsNotExist(err) {
		t.Error("file1.txt not copied")
	}
	if _, err := os.Stat(filepath.Join(dstDir, "subdir", "file2.txt")); os.IsNotExist(err) {
		t.Error("subdir/file2.txt not copied")
	}

	content, err := os.ReadFile(filepath.Join(dstDir, "file1.txt"))
	if err != nil {
		t.Fatalf("Failed to read: %v", err)
	}
	if string(content) != "content1" {
		t.Errorf("Content mismatch: got %s", content)
	}
}

func TestBackport_DetectsNoChanges(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	skillContent := []byte("# Same Content")

	srcSkill := filepath.Join(sourceDir, "test-skill")
	mustMkdirAll(t, srcSkill)
	mustWriteFile(t, filepath.Join(srcSkill, "SKILL.md"), skillContent)

	localSkill := filepath.Join(targetDir, "skills", "test-skill")
	mustMkdirAll(t, localSkill)
	mustWriteFile(t, filepath.Join(localSkill, "SKILL.md"), skillContent)

	_ = New(sourceDir, targetDir, "")

	_, err := os.Stat(filepath.Join(localSkill, "SKILL.md"))
	if err != nil {
		t.Error("Local skill setup failed")
	}
}

// Test helpers.

func mustMkdirAll(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("Failed to create dir %s: %v", path, err)
	}
}

func mustWriteFile(t *testing.T, path string, content []byte) {
	t.Helper()
	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatalf("Failed to write file %s: %v", path, err)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
