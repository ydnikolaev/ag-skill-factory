package skills

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show skill inventory with sync status",
	Long: header("LIST") + ` — Show skill inventory

Displays a table of all skills with their installation and sync status.

` + header("COLUMNS") + `
  ` + accent("Skill") + `      Name of the skill
  ` + accent("Installed") + `  Present in .agent/skills/
  ` + accent("Source") + `     Present in factory
  ` + accent("Status") + `     Sync state

` + header("STATUS VALUES") + `
  ` + success("synced") + `        Both local and factory versions exist
  ` + accent("local only") + `    Only in project (custom skill)
  ` + dimmed("not installed") + ` Only in factory (available to install)

` + header("EXAMPLE") + `
  $ skills list
  
  Skill                     Installed  Source     Status
  ──────────────────────────────────────────────────────────
  product-manager           ✓          ✓          synced
  my-custom-skill           ✓          -          local only
  new-factory-skill         -          ✓          not installed`,
	RunE: runList,
}

func runList(_ *cobra.Command, _ []string) error {
	source := viper.GetString("source")
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	localSkillsPath := filepath.Join(cwd, ".agent", "skills")

	// Get source skills
	sourceSkills := getSourceSkills(source)

	// Get local skills
	localSkills := getLocalSkills(localSkillsPath)

	// Collect and sort all skills
	allSkills := collectAllSkills(sourceSkills, localSkills)

	// Print table
	printSkillTable(allSkills, localSkills, sourceSkills)

	return nil
}

func getSourceSkills(source string) map[string]bool {
	skills := make(map[string]bool)
	entries, err := os.ReadDir(source)
	if err != nil {
		return skills
	}

	for _, e := range entries {
		if e.IsDir() && !strings.HasPrefix(e.Name(), "_") {
			skillFile := filepath.Join(source, e.Name(), "SKILL.md")
			if _, err := os.Stat(skillFile); err == nil {
				skills[e.Name()] = true
			}
		}
	}
	return skills
}

func getLocalSkills(path string) map[string]bool {
	skills := make(map[string]bool)
	entries, err := os.ReadDir(path)
	if err != nil {
		return skills
	}

	for _, e := range entries {
		if e.IsDir() {
			skills[e.Name()] = true
		}
	}
	return skills
}

func collectAllSkills(source, local map[string]bool) []string {
	all := make(map[string]bool)
	for k := range source {
		all[k] = true
	}
	for k := range local {
		all[k] = true
	}

	result := make([]string, 0, len(all))
	for k := range all {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}

func printSkillTable(skills []string, local, source map[string]bool) {
	fmt.Println()
	_, _ = color.New(color.Bold).Printf("%-25s %-10s %-10s %-12s\n", "Skill", "Installed", "Source", "Status")
	fmt.Println(strings.Repeat("─", 60))

	for _, skill := range skills {
		installed := dimmed("-")
		if local[skill] {
			installed = color.GreenString("✓")
		}

		inSource := dimmed("-")
		if source[skill] {
			inSource = color.GreenString("✓")
		}

		status := getSkillStatus(local[skill], source[skill])
		fmt.Printf("%-25s %-10s %-10s %-12s\n", skill, installed, inSource, status)
	}
	fmt.Println()
}

func getSkillStatus(isLocal, isSource bool) string {
	switch {
	case isLocal && isSource:
		return color.GreenString("synced")
	case isLocal && !isSource:
		return color.YellowString("local only")
	case !isLocal && isSource:
		return color.CyanString("not installed")
	default:
		return ""
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
