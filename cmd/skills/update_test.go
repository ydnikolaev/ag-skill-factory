package skills

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestRunUpdate_NoAgentDir(t *testing.T) {
	// Create temp workspace without .agent
	workspaceDir := t.TempDir()

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(workspaceDir); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(oldWd) }()

	viper.Set("source", t.TempDir())

	// Should not error, but print message
	err = runUpdate(nil, nil)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestRunUpdate_WithAgentDir(t *testing.T) {
	workspaceDir := t.TempDir()
	sourceDir := t.TempDir()

	// Create .agent directory
	agentDir := filepath.Join(workspaceDir, ".agent")
	skillsDir := filepath.Join(agentDir, "skills")
	if err := os.MkdirAll(skillsDir, 0o755); err != nil {
		t.Fatal(err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(workspaceDir); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(oldWd) }()

	viper.Set("source", sourceDir)

	// Should succeed with no skills to update
	err = runUpdate(nil, nil)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestRunUpdate_ForceMode(t *testing.T) {
	workspaceDir := t.TempDir()
	sourceDir := t.TempDir()

	// Create .agent directory
	agentDir := filepath.Join(workspaceDir, ".agent")
	skillsDir := filepath.Join(agentDir, "skills")
	if err := os.MkdirAll(skillsDir, 0o755); err != nil {
		t.Fatal(err)
	}

	// Create a skill in source
	skillDir := filepath.Join(sourceDir, "test-skill")
	if err := os.MkdirAll(skillDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("# Test"), 0o644); err != nil {
		t.Fatal(err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(workspaceDir); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(oldWd) }()

	viper.Set("source", sourceDir)

	// Set force mode
	oldForce := forceUpdate
	forceUpdate = true
	defer func() { forceUpdate = oldForce }()

	err = runUpdate(nil, nil)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	// Verify skill was force-copied
	installedSkill := filepath.Join(skillsDir, "test-skill", "SKILL.md")
	if _, err := os.Stat(installedSkill); os.IsNotExist(err) {
		t.Error("skill was not force-refreshed")
	}
}
