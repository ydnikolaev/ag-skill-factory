package installer

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/afero"
)

// Test helpers for afero MemMapFs

func newTestInstaller(fs afero.Fs, source, target string) *Installer {
	return NewWithFs(source, target, &MockPrompter{Response: true}, fs)
}

func writeFile(t *testing.T, fs afero.Fs, path, content string) {
	t.Helper()
	dir := filepath.Dir(path)
	if err := fs.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("MkdirAll %s: %v", dir, err)
	}
	if err := afero.WriteFile(fs, path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile %s: %v", path, err)
	}
}

func assertFileExists(t *testing.T, fs afero.Fs, path string) {
	t.Helper()
	if _, err := fs.Stat(path); err != nil {
		t.Errorf("Expected file to exist: %s", path)
	}
}

func assertFileContains(t *testing.T, fs afero.Fs, path, substr string) {
	t.Helper()
	content, err := afero.ReadFile(fs, path)
	if err != nil {
		t.Errorf("Failed to read %s: %v", path, err)
		return
	}
	if !contains(string(content), substr) {
		t.Errorf("File %s should contain %q", path, substr)
	}
}

// === Core Tests ===

func TestNew(t *testing.T) {
	inst := New("/source", "/target")
	if inst.Source != "/source" || inst.Target != "/target" {
		t.Error("New() should set Source and Target")
	}
	if inst.Fs == nil {
		t.Error("New() should set Fs to OsFs")
	}
}

func TestNewWithFs(t *testing.T) {
	fs := afero.NewMemMapFs()
	mock := &MockPrompter{Response: true}
	inst := NewWithFs("/src", "/dst", mock, fs)
	if inst.Fs != fs {
		t.Error("NewWithFs should use provided Fs")
	}
}

// === Install Tests ===

func TestInstall_CreatesDirectoryStructure(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/my-skill/SKILL.md", "# My Skill")
	writeFile(t, fs, "/source/_standards/PROTO.md", "# Protocol")
	writeFile(t, fs, "/source/TEAM.md", "# Team")

	inst := newTestInstaller(fs, "/source", "/target")
	result, err := inst.Install()
	if err != nil {
		t.Fatalf("Install failed: %v", err)
	}
	if result.SkillCount != 1 {
		t.Errorf("Expected 1 skill, got %d", result.SkillCount)
	}
	if result.RuleCount < 1 {
		t.Errorf("Expected at least 1 rule, got %d", result.RuleCount)
	}

	assertFileExists(t, fs, "/target/skills")
	assertFileExists(t, fs, "/target/rules")
	assertFileExists(t, fs, "/target/workflows")
	assertFileExists(t, fs, "/target/skills/my-skill/SKILL.md")
}

func TestInstall_CreateTargetDirsError(t *testing.T) {
	fs := afero.NewReadOnlyFs(afero.NewMemMapFs())
	inst := newTestInstaller(fs, "/source", "/target")

	_, err := inst.Install()
	if err == nil {
		t.Error("Expected error on read-only fs")
	}
}

func TestInstall_ProcessSourceEntriesError(t *testing.T) {
	fs := afero.NewMemMapFs()
	// Don't create /source - will fail to read

	inst := newTestInstaller(fs, "/source", "/target")
	_, err := inst.Install()

	if err == nil {
		t.Error("Expected error for non-existent source")
	}
}

// === ForceRefresh Tests ===

func TestForceRefresh_CopiesAllSkills(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/skill-a/SKILL.md", "# A")
	writeFile(t, fs, "/source/skill-b/SKILL.md", "# B")
	writeFile(t, fs, "/source/_standards/X.md", "# X")
	writeFile(t, fs, "/source/references/Y.md", "# Y")
	writeFile(t, fs, "/source/not-skill/README.md", "# Readme")
	writeFile(t, fs, "/source/file.txt", "file")
	_ = fs.MkdirAll("/target/skills", 0o755)

	inst := newTestInstaller(fs, "/source", "/target")
	result, err := inst.ForceRefresh()
	if err != nil {
		t.Fatalf("ForceRefresh failed: %v", err)
	}
	if result.SkillCount != 2 {
		t.Errorf("Expected 2 skills, got %d", result.SkillCount)
	}
}

func TestForceRefresh_ReadDirError(t *testing.T) {
	fs := afero.NewMemMapFs()
	// No /source created

	inst := newTestInstaller(fs, "/source", "/target")
	_, err := inst.ForceRefresh()

	if err == nil {
		t.Error("Expected error for non-existent source")
	}
}

// === Update Tests ===

func TestUpdate_UpdatesSkillsWithDiff(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/my-skill/SKILL.md", "# New")
	writeFile(t, fs, "/target/skills/my-skill/SKILL.md", "# Old")

	mock := &MockPrompter{Response: true}
	inst := NewWithFs("/source", "/target", mock, fs)

	_, err := inst.Update()
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
}

func TestUpdate_ReadDirError(t *testing.T) {
	fs := afero.NewMemMapFs()
	// No /target/skills created

	inst := newTestInstaller(fs, "/source", "/target")
	_, err := inst.Update()

	if err == nil {
		t.Error("Expected error for non-existent skills dir")
	}
}

func TestUpdate_SkipsNonDirs(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = fs.MkdirAll("/target/skills", 0o755)
	writeFile(t, fs, "/target/skills/not-a-dir.txt", "file")

	inst := newTestInstaller(fs, "/source", "/target")
	result, _ := inst.Update()

	if result.UpdatedCount != 0 {
		t.Errorf("Expected 0 updates, got %d", result.UpdatedCount)
	}
}

// === Backport Tests (use t.TempDir because diff package uses real OS) ===

func TestBackport_NoChanges(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create identical files
	mustWriteFile(t, filepath.Join(sourceDir, "my-skill", "SKILL.md"), []byte("# Same"))
	mustWriteFile(t, filepath.Join(targetDir, "skills", "my-skill", "SKILL.md"), []byte("# Same"))

	mock := &MockPrompter{Response: true}
	inst := NewWithPrompter(sourceDir, targetDir, mock)

	err := inst.Backport("my-skill")
	if err != nil {
		t.Fatalf("Backport failed: %v", err)
	}
	if len(mock.Calls) != 0 {
		t.Error("No confirm should be called for identical content")
	}
}

func TestBackport_WithChanges(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	mustWriteFile(t, filepath.Join(sourceDir, "my-skill", "SKILL.md"), []byte("# Old"))
	mustWriteFile(t, filepath.Join(targetDir, "skills", "my-skill", "SKILL.md"), []byte("# New"))

	mock := &MockPrompter{Response: true}
	inst := NewWithPrompter(sourceDir, targetDir, mock)

	err := inst.Backport("my-skill")
	if err != nil {
		t.Fatalf("Backport failed: %v", err)
	}
	if len(mock.Calls) != 1 {
		t.Errorf("Expected 1 confirm call, got %d", len(mock.Calls))
	}
}

func TestBackport_Cancelled(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	mustWriteFile(t, filepath.Join(sourceDir, "my-skill", "SKILL.md"), []byte("# Old"))
	mustWriteFile(t, filepath.Join(targetDir, "skills", "my-skill", "SKILL.md"), []byte("# New"))

	mock := &MockPrompter{Response: false}
	inst := NewWithPrompter(sourceDir, targetDir, mock)

	_ = inst.Backport("my-skill")

	// Original source should be unchanged
	content, _ := os.ReadFile(filepath.Join(sourceDir, "my-skill", "SKILL.md"))
	if string(content) != "# Old" {
		t.Error("Source should be unchanged on cancel")
	}
}

func TestBackport_CompareError(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// No files - compare will fail
	inst := NewWithPrompter(sourceDir, targetDir, &MockPrompter{})

	err := inst.Backport("nonexistent")
	if err == nil {
		t.Error("Expected error for non-existent skill")
	}
}

// === Operation Tests ===

func TestCreateTargetDirs(t *testing.T) {
	fs := afero.NewMemMapFs()
	inst := newTestInstaller(fs, "/source", "/target")

	err := inst.createTargetDirs()
	if err != nil {
		t.Fatalf("createTargetDirs failed: %v", err)
	}

	assertFileExists(t, fs, "/target/skills")
	assertFileExists(t, fs, "/target/rules")
	assertFileExists(t, fs, "/target/workflows")
}

func TestProcessEntry_Standards(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/_standards/PROTO.md", "# Proto")
	_ = fs.MkdirAll("/target/rules", 0o755)

	inst := newTestInstaller(fs, "/source", "/target")
	entries, _ := afero.ReadDir(fs, "/source")
	result := &InstallResult{}

	for _, entry := range entries {
		_ = inst.processEntry(entry, result)
	}

	if result.RuleCount != 1 {
		t.Errorf("Expected 1 rule, got %d", result.RuleCount)
	}
}

func TestProcessEntry_SkipsReferences(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/references/SKILL.md", "# Ref")
	_ = fs.MkdirAll("/target/skills", 0o755)

	inst := newTestInstaller(fs, "/source", "/target")
	entries, _ := afero.ReadDir(fs, "/source")
	result := &InstallResult{}

	for _, entry := range entries {
		_ = inst.processEntry(entry, result)
	}

	if result.SkillCount != 0 {
		t.Errorf("References should be skipped, got %d skills", result.SkillCount)
	}
}

func TestProcessEntry_SkipsFiles(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/file.txt", "content")

	inst := newTestInstaller(fs, "/source", "/target")
	info, _ := fs.Stat("/source/file.txt")
	result := &InstallResult{}

	err := inst.processEntry(info, result)
	if err != nil {
		t.Fatalf("processEntry failed: %v", err)
	}
	if result.SkillCount != 0 {
		t.Error("Files should be skipped")
	}
}

func TestCopySkillIfValid_NoSkillMd(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/not-skill/README.md", "# Readme")
	_ = fs.MkdirAll("/target/skills", 0o755)

	inst := newTestInstaller(fs, "/source", "/target")
	result := &InstallResult{}

	err := inst.copySkillIfValid("not-skill", "/source/not-skill", result)
	if err != nil {
		t.Fatalf("copySkillIfValid failed: %v", err)
	}
	if result.SkillCount != 0 {
		t.Error("Should skip dirs without SKILL.md")
	}
}

func TestCopyDirWithRewrite_TransformsPaths(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/skill/SKILL.md", "See `_standards/GIT.md` for details.")

	inst := newTestInstaller(fs, "/source", "/target")
	err := inst.copyDirWithRewrite("/source/skill", "/target/skill")
	if err != nil {
		t.Fatalf("copyDirWithRewrite failed: %v", err)
	}

	content, _ := afero.ReadFile(fs, "/target/skill/SKILL.md")
	if contains(string(content), "_standards/") {
		t.Error("Path should be rewritten")
	}
}

func TestCopyDirWithRewrite_WalkError(t *testing.T) {
	fs := afero.NewMemMapFs()
	// Don't create source

	inst := newTestInstaller(fs, "/source", "/target")
	err := inst.copyDirWithRewrite("/source/skill", "/target/skill")

	if err == nil {
		t.Error("Expected error for non-existent source")
	}
}

func TestCopyFileWithRewrite(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/test.md", "Check `_standards/PROTO.md`")

	inst := newTestInstaller(fs, "/source", "/target")
	_ = fs.MkdirAll("/target", 0o755)
	err := inst.copyFileWithRewrite("/source/test.md", "/target/test.md")
	if err != nil {
		t.Fatalf("copyFileWithRewrite failed: %v", err)
	}

	assertFileContains(t, fs, "/target/test.md", ".agent/rules/proto.md")
}

func TestCopyFileWithRewrite_ReadError(t *testing.T) {
	fs := afero.NewMemMapFs()

	inst := newTestInstaller(fs, "/source", "/target")
	err := inst.copyFileWithRewrite("/nonexistent", "/target/out.md")

	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}

// === Converter Tests ===

func TestConvertStandardsToRules(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/_standards/GIT.md", "# Git Protocol")
	writeFile(t, fs, "/source/_standards/TDD.md", "# TDD Protocol")
	writeFile(t, fs, "/source/_standards/README.txt", "not md")
	_ = fs.MkdirAll("/source/_standards/subdir", 0o755)
	_ = fs.MkdirAll("/target/rules", 0o755)

	inst := newTestInstaller(fs, "/source", "/target")
	count, err := inst.convertStandardsToRules("/source/_standards")
	if err != nil {
		t.Fatalf("convertStandardsToRules failed: %v", err)
	}
	if count != 2 {
		t.Errorf("Expected 2 conversions, got %d", count)
	}

	assertFileExists(t, fs, "/target/rules/git.md")
	assertFileContains(t, fs, "/target/rules/git.md", "description: Git Protocol")
}

func TestConvertStandardsToRules_ReadDirError(t *testing.T) {
	fs := afero.NewMemMapFs()

	inst := newTestInstaller(fs, "/source", "/target")
	_, err := inst.convertStandardsToRules("/nonexistent")

	if err == nil {
		t.Error("Expected error for non-existent dir")
	}
}

func TestConvertToRule(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/PROTO.md", "# My Protocol\n\nContent here")
	_ = fs.MkdirAll("/target/rules", 0o755)

	inst := newTestInstaller(fs, "/source", "/target")
	err := inst.convertToRule("/source/PROTO.md", "/target/rules/proto.md")
	if err != nil {
		t.Fatalf("convertToRule failed: %v", err)
	}

	assertFileContains(t, fs, "/target/rules/proto.md", "---")
	assertFileContains(t, fs, "/target/rules/proto.md", "description: My Protocol")
}

func TestConvertToRule_ReadError(t *testing.T) {
	fs := afero.NewMemMapFs()

	inst := newTestInstaller(fs, "/source", "/target")
	err := inst.convertToRule("/nonexistent.md", "/target/out.md")

	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}

func TestExtractFirstHeading(t *testing.T) {
	tests := []struct {
		content  string
		expected string
	}{
		{"# Hello\nWorld", "Hello"},
		{"No heading", "Rule"},
		{"## H2 Only", "Rule"},
		{"", "Rule"},
	}
	for _, tc := range tests {
		got := extractFirstHeading(tc.content)
		if got != tc.expected {
			t.Errorf("extractFirstHeading(%q) = %q, want %q", tc.content, got, tc.expected)
		}
	}
}

func TestCopyMetaFilesToRules(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/TEAM.md", "# Team")
	writeFile(t, fs, "/source/PIPELINE.md", "# Pipeline")
	_ = fs.MkdirAll("/target/rules", 0o755)

	inst := newTestInstaller(fs, "/source", "/target")
	result := &InstallResult{}

	err := inst.copyMetaFilesToRules(result)
	if err != nil {
		t.Fatalf("copyMetaFilesToRules failed: %v", err)
	}

	if result.RuleCount != 2 {
		t.Errorf("Expected 2 rules, got %d", result.RuleCount)
	}

	assertFileExists(t, fs, "/target/rules/team.md")
	assertFileExists(t, fs, "/target/rules/pipeline.md")
}

// === Error Path Tests ===

func TestConvertStandardsToRules_WriteError(t *testing.T) {
	memFs := afero.NewMemMapFs()
	writeFile(t, memFs, "/source/_standards/PROTO.md", "# Proto")
	_ = memFs.MkdirAll("/target/rules", 0o755)

	// Wrap with read-only to prevent writes
	roFs := afero.NewReadOnlyFs(memFs)
	inst := newTestInstaller(roFs, "/source", "/target")

	_, err := inst.convertStandardsToRules("/source/_standards")
	if err == nil {
		t.Error("Expected write error on read-only fs")
	}
}

func TestCopyMetaFilesToRules_WriteError(t *testing.T) {
	memFs := afero.NewMemMapFs()
	writeFile(t, memFs, "/source/TEAM.md", "# Team")
	_ = memFs.MkdirAll("/target/rules", 0o755)

	roFs := afero.NewReadOnlyFs(memFs)
	inst := newTestInstaller(roFs, "/source", "/target")

	result := &InstallResult{}
	// Should not crash, just log warning
	_ = inst.copyMetaFilesToRules(result)

	// RuleCount should be 0 due to write failure
	if result.RuleCount != 0 {
		t.Errorf("Expected 0 rules on write failure, got %d", result.RuleCount)
	}
}

func TestProcessEntry_StandardsConvertError(t *testing.T) {
	memFs := afero.NewMemMapFs()
	writeFile(t, memFs, "/source/_standards/PROTO.md", "# Proto")
	_ = memFs.MkdirAll("/target/skills", 0o755)

	roFs := afero.NewReadOnlyFs(memFs)
	inst := newTestInstaller(roFs, "/source", "/target")

	entries, _ := afero.ReadDir(memFs, "/source")
	result := &InstallResult{}

	err := inst.processEntry(entries[0], result)
	if err == nil {
		t.Error("Expected error from standards conversion")
	}
}

func TestCopySkillIfValid_CopyError(t *testing.T) {
	memFs := afero.NewMemMapFs()
	writeFile(t, memFs, "/source/my-skill/SKILL.md", "# Skill")
	_ = memFs.MkdirAll("/target/skills", 0o755)

	roFs := afero.NewReadOnlyFs(memFs)
	inst := newTestInstaller(roFs, "/source", "/target")

	result := &InstallResult{}
	err := inst.copySkillIfValid("my-skill", "/source/my-skill", result)

	if err == nil {
		t.Error("Expected copy error on read-only fs")
	}
}

func TestCopyFileWithRewrite_WriteError(t *testing.T) {
	memFs := afero.NewMemMapFs()
	writeFile(t, memFs, "/source/test.md", "# Test with _standards/X.md")

	roFs := afero.NewReadOnlyFs(memFs)
	inst := newTestInstaller(roFs, "/source", "/target")

	err := inst.copyFileWithRewrite("/source/test.md", "/target/out.md")
	if err == nil {
		t.Error("Expected write error on read-only fs")
	}
}

func TestConvertToRule_WriteError(t *testing.T) {
	memFs := afero.NewMemMapFs()
	writeFile(t, memFs, "/source/PROTO.md", "# Proto")

	roFs := afero.NewReadOnlyFs(memFs)
	inst := newTestInstaller(roFs, "/source", "/target")

	err := inst.convertToRule("/source/PROTO.md", "/target/proto.md")
	if err == nil {
		t.Error("Expected write error on read-only fs")
	}
}

func TestCopyDirWithRewrite_MkdirError(t *testing.T) {
	memFs := afero.NewMemMapFs()
	writeFile(t, memFs, "/source/skill/SKILL.md", "# Skill")

	roFs := afero.NewReadOnlyFs(memFs)
	inst := newTestInstaller(roFs, "/source", "/target")

	err := inst.copyDirWithRewrite("/source/skill", "/target/skill")
	if err == nil {
		t.Error("Expected mkdir error on read-only fs")
	}
}

func TestForceRefresh_CopyError(t *testing.T) {
	memFs := afero.NewMemMapFs()
	writeFile(t, memFs, "/source/my-skill/SKILL.md", "# Skill")
	_ = memFs.MkdirAll("/target/skills", 0o755)

	roFs := afero.NewReadOnlyFs(memFs)
	inst := newTestInstaller(roFs, "/source", "/target")

	result, _ := inst.ForceRefresh()

	// Should not crash, just log warning and skill count should be 0
	if result.SkillCount != 0 {
		t.Errorf("Expected 0 skills on copy failure, got %d", result.SkillCount)
	}
}

func TestUpdate_UpdateRulesError(t *testing.T) {
	memFs := afero.NewMemMapFs()
	_ = memFs.MkdirAll("/target/skills", 0o755)
	writeFile(t, memFs, "/source/_standards/PROTO.md", "# Proto")

	roFs := afero.NewReadOnlyFs(memFs)
	inst := newTestInstaller(roFs, "/source", "/target")

	// Should not crash, updateRules error just logs warning
	_, _ = inst.Update()
}

// === Utils Tests ===

func TestCopyDir(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/src/file.txt", "content")
	writeFile(t, fs, "/src/sub/nested.txt", "nested")

	inst := newTestInstaller(fs, "/src", "/dst")
	err := inst.copyDir("/src", "/dst")
	if err != nil {
		t.Fatalf("copyDir failed: %v", err)
	}

	assertFileExists(t, fs, "/dst/file.txt")
	assertFileExists(t, fs, "/dst/sub/nested.txt")
}

func TestCopyDir_WalkError(t *testing.T) {
	fs := afero.NewMemMapFs()

	inst := newTestInstaller(fs, "/src", "/dst")
	err := inst.copyDir("/nonexistent", "/dst")

	if err == nil {
		t.Error("Expected error for non-existent source")
	}
}

func TestCopyFile(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/src/file.txt", "hello")

	inst := newTestInstaller(fs, "/src", "/dst")
	_ = fs.MkdirAll("/dst", 0o755)
	err := inst.copyFile("/src/file.txt", "/dst/file.txt")
	if err != nil {
		t.Fatalf("copyFile failed: %v", err)
	}

	content, _ := afero.ReadFile(fs, "/dst/file.txt")
	if string(content) != "hello" {
		t.Error("File content mismatch")
	}
}

func TestCopyFile_OpenError(t *testing.T) {
	fs := afero.NewMemMapFs()

	inst := newTestInstaller(fs, "/src", "/dst")
	err := inst.copyFile("/nonexistent", "/dst/out.txt")

	if err == nil {
		t.Error("Expected error for non-existent source")
	}
}

func TestCopyFile_CreateError(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/src/file.txt", "content")

	// Use read-only fs for destination
	roFs := afero.NewReadOnlyFs(fs)
	inst := NewWithFs("/src", "/dst", &MockPrompter{}, roFs)

	err := inst.copyFile("/src/file.txt", "/dst/out.txt")
	if err == nil {
		t.Error("Expected error on read-only fs")
	}
}

// === UpdateRules Tests ===

func TestUpdateRules(t *testing.T) {
	fs := afero.NewMemMapFs()
	writeFile(t, fs, "/source/_standards/PROTO.md", "# Proto")
	writeFile(t, fs, "/source/TEAM.md", "# Team")
	_ = fs.MkdirAll("/target/rules", 0o755)

	inst := newTestInstaller(fs, "/source", "/target")
	err := inst.updateRules()
	if err != nil {
		t.Fatalf("updateRules failed: %v", err)
	}

	assertFileExists(t, fs, "/target/rules/proto.md")
	assertFileExists(t, fs, "/target/rules/team.md")
}

// === Helper ===

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// === UpdateSingleSkill Tests (uses t.TempDir for diff) ===

func TestUpdateSingleSkill_NoSourceSkill(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	mustWriteFile(t, filepath.Join(targetDir, "skills", "my-skill", "SKILL.md"), []byte("# Local"))

	mock := &MockPrompter{Response: true}
	inst := NewWithPrompter(sourceDir, targetDir, mock)

	result := &UpdateResult{}
	inst.updateSingleSkill("my-skill", filepath.Join(targetDir, "skills"), result)

	if result.UpdatedCount != 0 {
		t.Errorf("Expected 0 updates, got %d", result.UpdatedCount)
	}
}

func TestUpdateSingleSkill_WithChanges(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	mustWriteFile(t, filepath.Join(sourceDir, "my-skill", "SKILL.md"), []byte("# New"))
	mustWriteFile(t, filepath.Join(targetDir, "skills", "my-skill", "SKILL.md"), []byte("# Old"))

	mock := &MockPrompter{Response: true}
	inst := NewWithPrompter(sourceDir, targetDir, mock)

	result := &UpdateResult{}
	inst.updateSingleSkill("my-skill", filepath.Join(targetDir, "skills"), result)

	if result.UpdatedCount != 1 {
		t.Errorf("Expected 1 update, got %d", result.UpdatedCount)
	}
}

func TestUpdateSingleSkill_Declined(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	mustWriteFile(t, filepath.Join(sourceDir, "my-skill", "SKILL.md"), []byte("# New"))
	mustWriteFile(t, filepath.Join(targetDir, "skills", "my-skill", "SKILL.md"), []byte("# Old"))

	mock := &MockPrompter{Response: false}
	inst := NewWithPrompter(sourceDir, targetDir, mock)

	result := &UpdateResult{}
	inst.updateSingleSkill("my-skill", filepath.Join(targetDir, "skills"), result)

	if result.UpdatedCount != 0 {
		t.Errorf("Expected 0 updates on decline, got %d", result.UpdatedCount)
	}
}

func TestUpdateSingleSkill_NoChanges(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Same content - no changes
	mustWriteFile(t, filepath.Join(sourceDir, "my-skill", "SKILL.md"), []byte("# Same"))
	mustWriteFile(t, filepath.Join(targetDir, "skills", "my-skill", "SKILL.md"), []byte("# Same"))

	mock := &MockPrompter{Response: true}
	inst := NewWithPrompter(sourceDir, targetDir, mock)

	result := &UpdateResult{}
	inst.updateSingleSkill("my-skill", filepath.Join(targetDir, "skills"), result)

	if len(mock.Calls) != 0 {
		t.Error("No prompt should be shown for identical content")
	}
}

func TestShowChangesAndApply_Confirmed(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	mustWriteFile(t, filepath.Join(sourceDir, "my-skill", "SKILL.md"), []byte("# New"))
	mustWriteFile(t, filepath.Join(targetDir, "skills", "my-skill", "SKILL.md"), []byte("# Old"))

	mock := &MockPrompter{Response: true}
	inst := NewWithPrompter(sourceDir, targetDir, mock)

	result := &UpdateResult{}
	inst.showChangesAndApply("my-skill",
		[]string{"M SKILL.md"},
		filepath.Join(sourceDir, "my-skill"),
		filepath.Join(targetDir, "skills", "my-skill"),
		result)

	if result.UpdatedCount != 1 {
		t.Errorf("Expected 1 update, got %d", result.UpdatedCount)
	}
}

func TestShowChangesAndApply_Cancelled(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	mustWriteFile(t, filepath.Join(sourceDir, "my-skill", "SKILL.md"), []byte("# New"))
	mustWriteFile(t, filepath.Join(targetDir, "skills", "my-skill", "SKILL.md"), []byte("# Old"))

	mock := &MockPrompter{Response: false}
	inst := NewWithPrompter(sourceDir, targetDir, mock)

	result := &UpdateResult{}
	inst.showChangesAndApply("my-skill",
		[]string{"M SKILL.md"},
		filepath.Join(sourceDir, "my-skill"),
		filepath.Join(targetDir, "skills", "my-skill"),
		result)

	if result.UpdatedCount != 0 {
		t.Errorf("Expected 0 updates on cancel, got %d", result.UpdatedCount)
	}
}

// === Prompter Tests ===

func TestMockPrompter_RecordsCalls(t *testing.T) {
	mock := &MockPrompter{Response: true}

	_ = mock.Confirm("First?")
	_ = mock.Confirm("Second?")

	if len(mock.Calls) != 2 {
		t.Errorf("Expected 2 calls, got %d", len(mock.Calls))
	}
	if mock.Calls[0] != "First?" {
		t.Error("First call not recorded correctly")
	}
}

func TestMockPrompter_ReturnsFalse(t *testing.T) {
	mock := &MockPrompter{Response: false}
	if mock.Confirm("Question?") {
		t.Error("Should return false")
	}
}

func TestStdinPrompter_ConfirmYes(t *testing.T) {
	p := &StdinPrompter{Input: strings.NewReader("y\n")}
	if !p.Confirm("Test?") {
		t.Error("Expected true for 'y'")
	}
}

func TestStdinPrompter_ConfirmYesFull(t *testing.T) {
	p := &StdinPrompter{Input: strings.NewReader("yes\n")}
	if !p.Confirm("Test?") {
		t.Error("Expected true for 'yes'")
	}
}

func TestStdinPrompter_ConfirmNo(t *testing.T) {
	p := &StdinPrompter{Input: strings.NewReader("n\n")}
	if p.Confirm("Test?") {
		t.Error("Expected false for 'n'")
	}
}

func TestStdinPrompter_ConfirmCaseInsensitive(t *testing.T) {
	p := &StdinPrompter{Input: strings.NewReader("YES\n")}
	if !p.Confirm("Test?") {
		t.Error("Expected true for 'YES'")
	}
}

func TestStdinPrompter_ConfirmWithWhitespace(t *testing.T) {
	p := &StdinPrompter{Input: strings.NewReader("  y  \n")}
	if !p.Confirm("Test?") {
		t.Error("Expected true with whitespace")
	}
}

// === Rewriter Tests ===

func TestRewriteStandardsPaths(t *testing.T) {
	input := "See `_standards/GIT.md` for info."
	output := rewriteStandardsPaths(input)

	if contains(output, "_standards/") {
		t.Error("Should transform _standards paths")
	}
	if !contains(output, ".agent/rules/git.md") {
		t.Error("Should contain transformed path")
	}
}

func TestRewriteStandardsPaths_MultipleOccurrences(t *testing.T) {
	input := "Check `_standards/A.md` and `_standards/B.md`."
	output := rewriteStandardsPaths(input)

	if contains(output, "_standards/") {
		t.Error("All occurrences should be transformed")
	}
}

func TestRewriteStandardsPaths_NoChange(t *testing.T) {
	input := "No standards references here."
	if rewriteStandardsPaths(input) != input {
		t.Error("Should not modify content without _standards")
	}
}

// Helper for t.TempDir tests

func mustWriteFile(t *testing.T, path string, content []byte) {
	t.Helper()
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("MkdirAll failed: %v", err)
	}
	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
}
