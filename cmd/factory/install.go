package factory

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ydnikolaev/antigravity-factory/internal/installer"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Copy blueprint to .agent/ in current workspace",
	Long: header("INSTALL") + ` — Copy blueprint to .agent/

Copies the complete agent infrastructure to your project:

` + header("CREATES") + `
  .agent/
  ├── ` + accent("skills/") + `     21 expert skills
  ├── ` + accent("workflows/") + `  Automation workflows
  ├── ` + accent("rules/") + `      TEAM.md, PIPELINE.md
  └── ` + accent("standards/") + `  TDD, Git, Tech Debt protocols

` + header("NOTES") + `
  • Blueprint is ` + cmd("copied") + ` as-is (no transformations)
  • Existing .agent/ will be ` + cmd("replaced") + `

` + header("EXAMPLE") + `
  $ cd my-project
  $ factory install
  ✅ Installed 21 skills, 2 workflows, 2 rules, 5 standards`,
	RunE: runInstall,
}

func runInstall(_ *cobra.Command, _ []string) error {
	source := viper.GetString("source")

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	target := filepath.Join(cwd, ".agent")

	// Always replace existing .agent/ (this is a blueprint copy)
	if _, err := os.Stat(target); err == nil {
		color.Cyan("🔄 Replacing existing .agent/...")
		if err := os.RemoveAll(target); err != nil {
			return fmt.Errorf("failed to remove existing .agent/: %w", err)
		}
	}

	// Create installer
	inst := installer.New(source, target)

	// Run installation
	result, err := inst.Install()
	if err != nil {
		return fmt.Errorf("installation failed: %w", err)
	}

	color.Green("✅ Installed %d skills, %d workflows, %d rules, %d standards",
		result.SkillCount, result.WorkflowCount, result.RuleCount, result.StandardCount)
	return nil
}

func init() {
	rootCmd.AddCommand(installCmd)
}
