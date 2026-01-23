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

var backportCmd = &cobra.Command{
	Use:   "backport <skill-name>",
	Short: "Push local skill changes back to factory",
	Long: header("BACKPORT") + ` — Push changes to factory

Copies skill modifications from your project back to the factory.
Essential for evolving skills based on real-world usage.

` + header("PROCESS") + `
  1. Compares local skill with factory version
  2. Shows ` + accent("colored diff") + ` of your changes
  3. Asks for ` + cmd("confirmation") + ` before copying
  4. Overwrites factory version with local

` + header("USE CASES") + `
  • Fixed a bug in a skill
  • Added new examples or resources
  • Improved SKILL.md documentation
  • Refined checklist or templates

` + header("EXAMPLE") + `
  $ skills backport product-manager
  === Changes to backport ===
  ` + success("+ Added new interview questions") + `
  Backport to factory? [y/n]: y
  ✅ Backported 'product-manager' to factory`,
	Args: cobra.ExactArgs(1),
	RunE: runBackport,
}

func runBackport(_ *cobra.Command, args []string) error {
	skillName := args[0]
	source := viper.GetString("source")

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	localSkill := filepath.Join(cwd, ".agent", "skills", skillName)

	// Check if local skill exists
	if _, err := os.Stat(localSkill); os.IsNotExist(err) {
		color.Red("❌ Skill '%s' not found in .agent/skills/", skillName)
		return nil
	}

	// Create installer
	inst := installer.New(source, filepath.Join(cwd, ".agent"), "")

	// Run backport with diff
	err = inst.Backport(skillName)
	if err != nil {
		return fmt.Errorf("backport failed: %w", err)
	}

	color.Green("✅ Backported '%s' to factory", skillName)
	return nil
}

func init() {
	rootCmd.AddCommand(backportCmd)
}
