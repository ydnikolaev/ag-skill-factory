package skills

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/yuranikolaev/ag-skill-factory/internal/installer"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Bootstrap .agent/ structure in current workspace",
	Long: header("INSTALL") + ` — Bootstrap .agent/ structure

Creates the complete skill infrastructure in your project:

` + header("CREATES") + `
  .agent/
  ├── ` + accent("skills/") + `     Full copies of all skills from factory
  ├── ` + accent("rules/") + `      Standards converted to rule format
  └── ` + accent("workflows/") + `  Empty, ready for custom workflows

` + header("NOTES") + `
  • Skills are ` + cmd("copied") + `, not symlinked (agents need local files)
  • Rules get YAML frontmatter for workspace-local injection

` + header("EXAMPLE") + `
  $ cd my-project
  $ skills install
  ✅ Installed 12 skills, 5 rules`,
	RunE: runInstall,
}

func runInstall(_ *cobra.Command, _ []string) error {
	source := viper.GetString("source")

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	target := filepath.Join(cwd, ".agent")

	// Check if .agent already exists
	if _, err := os.Stat(target); err == nil {
		color.Yellow("⚠️  .agent/ already exists. Use 'skills update' to refresh.")
		return nil
	}

	// Create installer
	inst := installer.New(source, target)

	// Run installation
	result, err := inst.Install()
	if err != nil {
		return fmt.Errorf("installation failed: %w", err)
	}

	color.Green("✅ Installed %d skills, %d rules", result.SkillCount, result.RuleCount)
	return nil
}

func init() {
	rootCmd.AddCommand(installCmd)
}
