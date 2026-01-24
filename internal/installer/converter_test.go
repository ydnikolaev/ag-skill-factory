package installer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConvertStandardsToRules(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create _standards directory with protocols
	standardsDir := filepath.Join(sourceDir, "_standards")
	mustMkdirAll(t, standardsDir)
	mustWriteFile(t, filepath.Join(standardsDir, "TDD_PROTOCOL.md"), []byte("# TDD Protocol\n\nTDD rules here"))
	mustWriteFile(t, filepath.Join(standardsDir, "GIT_PROTOCOL.md"), []byte("# Git Protocol\n\nGit rules here"))
	// Add a non-md file (should be skipped)
	mustWriteFile(t, filepath.Join(standardsDir, "README.txt"), []byte("not markdown"))

	inst := New(sourceDir, targetDir)
	mustMkdirAll(t, filepath.Join(targetDir, "rules"))

	count, err := inst.convertStandardsToRules(standardsDir)
	if err != nil {
		t.Fatalf("convertStandardsToRules failed: %v", err)
	}

	if count != 2 {
		t.Errorf("Expected 2 converted files, got %d", count)
	}

	// Check files were converted with frontmatter
	tddRule := filepath.Join(targetDir, "rules", "tdd_protocol.md")
	if _, err := os.Stat(tddRule); os.IsNotExist(err) {
		t.Error("tdd_protocol.md not created")
	}
	content, _ := os.ReadFile(tddRule)
	if !contains(string(content), "description: TDD Protocol") {
		t.Error("YAML frontmatter not added")
	}
}

func TestCopyMetaFilesToRules(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create TEAM.md and PIPELINE.md
	mustWriteFile(t, filepath.Join(sourceDir, "TEAM.md"), []byte("# Team\n\nTeam content"))
	mustWriteFile(t, filepath.Join(sourceDir, "PIPELINE.md"), []byte("# Pipeline\n\nPipeline content"))

	inst := New(sourceDir, targetDir)
	mustMkdirAll(t, filepath.Join(targetDir, "rules"))

	result := &InstallResult{}
	err := inst.copyMetaFilesToRules(result)
	if err != nil {
		t.Fatalf("copyMetaFilesToRules failed: %v", err)
	}

	if result.RuleCount != 2 {
		t.Errorf("Expected 2 rules, got %d", result.RuleCount)
	}

	// Check lowercase filenames
	teamRule := filepath.Join(targetDir, "rules", "team.md")
	if _, err := os.Stat(teamRule); os.IsNotExist(err) {
		t.Error("team.md not created")
	}
}

func TestForceRefresh(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create source skills
	skill1 := filepath.Join(sourceDir, "skill-one")
	mustMkdirAll(t, skill1)
	mustWriteFile(t, filepath.Join(skill1, "SKILL.md"), []byte("# Skill One"))

	skill2 := filepath.Join(sourceDir, "skill-two")
	mustMkdirAll(t, skill2)
	mustWriteFile(t, filepath.Join(skill2, "SKILL.md"), []byte("# Skill Two"))

	// Create _standards (should be skipped)
	standardsDir := filepath.Join(sourceDir, "_standards")
	mustMkdirAll(t, standardsDir)
	mustWriteFile(t, filepath.Join(standardsDir, "PROTOCOL.md"), []byte("# Protocol"))

	// Create references (should be skipped)
	refsDir := filepath.Join(sourceDir, "references")
	mustMkdirAll(t, refsDir)
	mustWriteFile(t, filepath.Join(refsDir, "SKILL.md"), []byte("# Not a skill"))

	// Create a directory without SKILL.md (should be skipped)
	noSkillDir := filepath.Join(sourceDir, "not-a-skill")
	mustMkdirAll(t, noSkillDir)
	mustWriteFile(t, filepath.Join(noSkillDir, "README.md"), []byte("# Readme"))

	// Create target directories
	mustMkdirAll(t, filepath.Join(targetDir, "skills"))
	mustMkdirAll(t, filepath.Join(targetDir, "rules"))

	inst := New(sourceDir, targetDir)
	result, err := inst.ForceRefresh()
	if err != nil {
		t.Fatalf("ForceRefresh failed: %v", err)
	}

	if result.SkillCount != 2 {
		t.Errorf("Expected 2 skills, got %d", result.SkillCount)
	}

	// Verify skills were copied
	if _, err := os.Stat(filepath.Join(targetDir, "skills", "skill-one", "SKILL.md")); os.IsNotExist(err) {
		t.Error("skill-one not copied")
	}
	if _, err := os.Stat(filepath.Join(targetDir, "skills", "skill-two", "SKILL.md")); os.IsNotExist(err) {
		t.Error("skill-two not copied")
	}

	// Verify _standards and references were NOT copied as skills
	if _, err := os.Stat(filepath.Join(targetDir, "skills", "_standards")); !os.IsNotExist(err) {
		t.Error("_standards should not be copied as skill")
	}
	if _, err := os.Stat(filepath.Join(targetDir, "skills", "references")); !os.IsNotExist(err) {
		t.Error("references should not be copied as skill")
	}
}

func TestUpdate_NoLocalSkills(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create empty local skills directory
	mustMkdirAll(t, filepath.Join(targetDir, "skills"))

	inst := New(sourceDir, targetDir)
	result, err := inst.Update()
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if result.UpdatedCount != 0 {
		t.Errorf("Expected 0 updates with no local skills, got %d", result.UpdatedCount)
	}
}

func TestCopyDir_WithSubdirectories(t *testing.T) {
	srcDir := t.TempDir()
	dstDir := filepath.Join(t.TempDir(), "dest")

	// Create nested structure
	mustMkdirAll(t, filepath.Join(srcDir, "level1", "level2"))
	mustWriteFile(t, filepath.Join(srcDir, "root.txt"), []byte("root"))
	mustWriteFile(t, filepath.Join(srcDir, "level1", "l1.txt"), []byte("level1"))
	mustWriteFile(t, filepath.Join(srcDir, "level1", "level2", "l2.txt"), []byte("level2"))

	err := copyDir(srcDir, dstDir)
	if err != nil {
		t.Fatalf("copyDir failed: %v", err)
	}

	// Verify all files copied
	if _, err := os.Stat(filepath.Join(dstDir, "root.txt")); os.IsNotExist(err) {
		t.Error("root.txt not copied")
	}
	if _, err := os.Stat(filepath.Join(dstDir, "level1", "l1.txt")); os.IsNotExist(err) {
		t.Error("level1/l1.txt not copied")
	}
	if _, err := os.Stat(filepath.Join(dstDir, "level1", "level2", "l2.txt")); os.IsNotExist(err) {
		t.Error("level1/level2/l2.txt not copied")
	}
}

func TestProcessSourceEntries(t *testing.T) {
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create skill
	skillDir := filepath.Join(sourceDir, "my-skill")
	mustMkdirAll(t, skillDir)
	mustWriteFile(t, filepath.Join(skillDir, "SKILL.md"), []byte("# My Skill"))

	inst := New(sourceDir, targetDir)
	mustMkdirAll(t, filepath.Join(targetDir, "skills"))
	mustMkdirAll(t, filepath.Join(targetDir, "rules"))

	result := &InstallResult{}
	err := inst.processSourceEntries(result)
	if err != nil {
		t.Fatalf("processSourceEntries failed: %v", err)
	}

	if result.SkillCount != 1 {
		t.Errorf("Expected 1 skill, got %d", result.SkillCount)
	}
}
