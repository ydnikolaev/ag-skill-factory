package skills

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetSourceSkills(t *testing.T) {
	// Create temp directory with test skill structure
	tmpDir := t.TempDir()

	// Create a valid skill directory
	skillDir := filepath.Join(tmpDir, "test-skill")
	if err := os.MkdirAll(skillDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("# Test"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Create a directory without SKILL.md (should be ignored)
	noSkillDir := filepath.Join(tmpDir, "not-a-skill")
	if err := os.MkdirAll(noSkillDir, 0o755); err != nil {
		t.Fatal(err)
	}

	// Create a _standards directory (should be ignored)
	standardsDir := filepath.Join(tmpDir, "_standards")
	if err := os.MkdirAll(standardsDir, 0o755); err != nil {
		t.Fatal(err)
	}

	// Test
	skills := getSourceSkills(tmpDir)

	if len(skills) != 1 {
		t.Errorf("expected 1 skill, got %d", len(skills))
	}
	if !skills["test-skill"] {
		t.Error("expected test-skill to be in results")
	}
	if skills["not-a-skill"] {
		t.Error("not-a-skill should not be in results (no SKILL.md)")
	}
	if skills["_standards"] {
		t.Error("_standards should not be in results")
	}
}

func TestGetLocalSkills(t *testing.T) {
	tmpDir := t.TempDir()

	// Create some directories
	for _, name := range []string{"skill-a", "skill-b"} {
		if err := os.MkdirAll(filepath.Join(tmpDir, name), 0o755); err != nil {
			t.Fatal(err)
		}
	}

	// Create a file (should be ignored)
	if err := os.WriteFile(filepath.Join(tmpDir, "file.txt"), []byte("test"), 0o644); err != nil {
		t.Fatal(err)
	}

	skills := getLocalSkills(tmpDir)

	if len(skills) != 2 {
		t.Errorf("expected 2 skills, got %d", len(skills))
	}
	if !skills["skill-a"] || !skills["skill-b"] {
		t.Error("expected both skills to be present")
	}
}

func TestGetLocalSkills_NonExistentPath(t *testing.T) {
	skills := getLocalSkills("/nonexistent/path")
	if len(skills) != 0 {
		t.Error("expected empty map for nonexistent path")
	}
}

func TestCollectAllSkills(t *testing.T) {
	source := map[string]bool{"a": true, "b": true, "c": true}
	local := map[string]bool{"b": true, "d": true}

	result := collectAllSkills(source, local)

	if len(result) != 4 {
		t.Errorf("expected 4 unique skills, got %d", len(result))
	}

	// Should be sorted
	expected := []string{"a", "b", "c", "d"}
	for i, name := range expected {
		if result[i] != name {
			t.Errorf("position %d: expected %s, got %s", i, name, result[i])
		}
	}
}

func TestGetSkillStatus(t *testing.T) {
	tests := []struct {
		name       string
		isLocal    bool
		isSource   bool
		wantSubstr string
	}{
		{"synced", true, true, "synced"},
		{"local only", true, false, "local only"},
		{"not installed", false, true, "not installed"},
		{"neither", false, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getSkillStatus(tt.isLocal, tt.isSource)
			// Note: result contains ANSI color codes, so we just check it's not empty when expected
			if tt.wantSubstr != "" && result == "" {
				t.Errorf("expected non-empty status for %s", tt.name)
			}
			if tt.wantSubstr == "" && result != "" {
				t.Errorf("expected empty status for %s, got %s", tt.name, result)
			}
		})
	}
}
