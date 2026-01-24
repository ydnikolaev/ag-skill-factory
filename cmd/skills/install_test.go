package skills

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestRunInstall_AlreadyExists(t *testing.T) {
	// Create temp directory as workspace
	tmpDir := t.TempDir()

	// Create .agent directory to simulate existing installation
	agentDir := filepath.Join(tmpDir, ".agent")
	if err := os.MkdirAll(agentDir, 0o755); err != nil {
		t.Fatal(err)
	}

	// Change to temp directory
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(oldWd) }()

	// Set up viper
	viper.Set("source", t.TempDir())

	// Run install - should return nil (no error) but not do anything
	err = runInstall(nil, nil)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestRunInstall_FreshInstall(t *testing.T) {
	// Create temp directories
	workspaceDir := t.TempDir()
	sourceDir := t.TempDir()

	// Create a test skill in source
	skillDir := filepath.Join(sourceDir, "test-skill")
	if err := os.MkdirAll(skillDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("# Test Skill"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Change to workspace
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(workspaceDir); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(oldWd) }()

	// Set up viper
	viper.Set("source", sourceDir)

	// Run install
	err = runInstall(nil, nil)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	// Verify .agent was created
	agentDir := filepath.Join(workspaceDir, ".agent")
	if _, err := os.Stat(agentDir); os.IsNotExist(err) {
		t.Error(".agent directory was not created")
	}

	// Verify skill was copied
	installedSkill := filepath.Join(agentDir, "skills", "test-skill", "SKILL.md")
	if _, err := os.Stat(installedSkill); os.IsNotExist(err) {
		t.Error("skill was not copied to .agent/skills/")
	}
}

func TestRunInstall_InvalidSource(t *testing.T) {
	workspaceDir := t.TempDir()

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(workspaceDir); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(oldWd) }()

	// Set invalid source
	viper.Set("source", "/nonexistent/path/to/source")

	err = runInstall(nil, nil)
	if err == nil {
		t.Error("expected error for invalid source")
	}
}
