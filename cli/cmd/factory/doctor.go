package factory

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ydnikolaev/antigravity-factory/internal/doctor"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run diagnostics on the blueprint",
	Long: header("DOCTOR") + ` — Diagnose blueprint health

Checks for common issues:
  • Broken ` + accent("See examples/...") + ` links
  • Unknown ` + accent("@skill-name") + ` references
  • Hardcoded paths ` + accent("/Users/...") + `
  • Missing required sections

` + header("EXIT CODES") + `
  0 = All checks passed
  1 = Errors found (CI should fail)

` + header("EXAMPLE") + `
  $ factory doctor
  ✅ 21 skills checked, 0 errors, 2 warnings`,
	RunE: runDoctor,
}

var exitOnError bool

func runDoctor(_ *cobra.Command, _ []string) error {
	source := viper.GetString("source")

	fmt.Println(header("🔍 Running diagnostics..."))
	fmt.Println()

	result, err := doctor.Check(source)
	if err != nil {
		return fmt.Errorf("doctor failed: %w", err)
	}

	// Print errors
	if len(result.Errors) > 0 {
		color.Red("❌ Errors (%d):", len(result.Errors))
		for _, e := range result.Errors {
			fmt.Printf("   • %s\n", e)
		}
		fmt.Println()
	}

	// Print warnings
	if len(result.Warnings) > 0 {
		color.Yellow("⚠️  Warnings (%d):", len(result.Warnings))
		for _, w := range result.Warnings {
			fmt.Printf("   • %s\n", w)
		}
		fmt.Println()
	}

	// Summary
	if len(result.Errors) == 0 && len(result.Warnings) == 0 {
		color.Green("✅ All checks passed!")
	} else if len(result.Errors) == 0 {
		color.Green("✅ No errors (with %d warnings)", len(result.Warnings))
	} else {
		color.Red("❌ %d errors found", len(result.Errors))
		if exitOnError {
			os.Exit(1)
		}
	}

	return nil
}

func init() {
	doctorCmd.Flags().BoolVar(&exitOnError, "exit-on-error", false, "Exit with code 1 if errors found")
	rootCmd.AddCommand(doctorCmd)
}
