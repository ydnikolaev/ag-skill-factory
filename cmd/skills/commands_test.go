package skills

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestRootCommand_Help(t *testing.T) {
	viper.Reset()

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Root command failed: %v", err)
	}

	output := buf.String()

	commands := []string{"install", "update", "backport", "list"}
	for _, cmd := range commands {
		if !containsStr(output, cmd) {
			t.Errorf("Help missing '%s' command", cmd)
		}
	}
}

func TestInstallCommand_FailsWithoutAgent(t *testing.T) {
	tempDir := t.TempDir()
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Logf("Failed to restore directory: %v", err)
		}
	})

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	viper.Reset()
	sourceDir := t.TempDir()
	viper.Set("source", sourceDir)
	viper.Set("global_path", t.TempDir())

	skillDir := filepath.Join(sourceDir, "mock-skill")
	if err := os.MkdirAll(skillDir, 0o755); err != nil {
		t.Fatalf("Failed to create skill dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("# Mock"), 0o644); err != nil {
		t.Fatalf("Failed to write SKILL.md: %v", err)
	}

	buf := new(bytes.Buffer)
	installCmd.SetOut(buf)
	installCmd.SetErr(buf)

	err = installCmd.RunE(installCmd, []string{})
	if err != nil {
		t.Fatalf("Install command failed: %v", err)
	}

	if _, err := os.Stat(filepath.Join(tempDir, ".agent", "skills")); os.IsNotExist(err) {
		t.Error(".agent/skills not created")
	}
}

func TestListCommand_ShowsSkills(t *testing.T) {
	tempDir := t.TempDir()
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Logf("Failed to restore directory: %v", err)
		}
	})

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	sourceDir := t.TempDir()
	createTestSkills(t, sourceDir)

	viper.Reset()
	viper.Set("source", sourceDir)

	buf := new(bytes.Buffer)
	listCmd.SetOut(buf)
	listCmd.SetErr(buf)

	err = listCmd.RunE(listCmd, []string{})
	if err != nil {
		t.Fatalf("List command failed: %v", err)
	}
}

func createTestSkills(t *testing.T, sourceDir string) {
	t.Helper()
	skills := []string{"skill-one", "skill-two"}
	for _, s := range skills {
		skillPath := filepath.Join(sourceDir, s)
		if err := os.MkdirAll(skillPath, 0o755); err != nil {
			t.Fatalf("Failed to create %s: %v", s, err)
		}
		if err := os.WriteFile(filepath.Join(skillPath, "SKILL.md"), []byte("# "+s), 0o644); err != nil {
			t.Fatalf("Failed to write SKILL.md for %s: %v", s, err)
		}
	}
}

func TestBackportCommand_RequiresArg(t *testing.T) {
	tests := []struct {
		args    []string
		wantErr bool
	}{
		{[]string{}, true},
		{[]string{"skill-name"}, false},
		{[]string{"arg1", "arg2"}, true},
	}

	for _, tt := range tests {
		err := backportCmd.Args(backportCmd, tt.args)
		if (err != nil) != tt.wantErr {
			t.Errorf("Args(%v) error = %v, wantErr %v", tt.args, err, tt.wantErr)
		}
	}
}

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
