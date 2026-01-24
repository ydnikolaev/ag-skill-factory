package skills

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ydnikolaev/ag-skill-factory/internal/installer"
)

var forceUpdate bool

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Pull latest skill changes from factory",
	Long: header("UPDATE") + ` — Pull latest from factory

Synchronizes your local skills with the factory source.

` + header("PROCESS") + `
  1. Compares each local skill with factory version
  2. Shows ` + accent("colored diff") + ` for changed files
  3. Asks for ` + cmd("confirmation") + ` before applying
  4. Updates only confirmed skills

` + header("FLAGS") + `
  --force    Skip diff comparison, refresh all skills

` + header("NOTES") + `
  • Safe operation — prompts before each change
  • Preserves local-only skills (not in factory)
  • Also updates rules/standards

` + header("EXAMPLE") + `
  $ skills update
  === my-skill ===
  ` + dimmed("- old line") + `
  ` + success("+ new line") + `
  Apply changes? [y/n]: y
  ✅ Updated 3 skills

  $ skills update --force
  ✅ Force refreshed 15 skills`,
	RunE: runUpdate,
}

func runUpdate(_ *cobra.Command, _ []string) error {
	source := viper.GetString("source")

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	target := filepath.Join(cwd, ".agent")

	// Check if .agent exists
	if _, err := os.Stat(target); os.IsNotExist(err) {
		color.Red("❌ .agent/ not found. Run 'skills install' first.")
		return nil
	}

	// Create installer
	inst := installer.New(source, target)

	if forceUpdate {
		// Force refresh all skills
		result, err := inst.ForceRefresh()
		if err != nil {
			return fmt.Errorf("force refresh failed: %w", err)
		}
		color.Green("✅ Force refreshed %d skills", result.SkillCount)
	} else {
		// Normal update with diff
		result, err := inst.Update()
		if err != nil {
			return fmt.Errorf("update failed: %w", err)
		}

		if result.UpdatedCount == 0 {
			color.Green("✅ Everything up to date")
		} else {
			color.Green("✅ Updated %d skills", result.UpdatedCount)
		}
	}
	return nil
}

func init() {
	updateCmd.Flags().BoolVarP(&forceUpdate, "force", "f", false, "Force refresh all skills (skip diff)")
	rootCmd.AddCommand(updateCmd)
}
