package factory

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
	Short: "Show installed blueprint inventory",
	Long: header("LIST") + ` — Show blueprint inventory

Displays all installed components from the blueprint.

` + header("CATEGORIES") + `
  ` + accent("Skills") + `     Expert agents (21 total)
  ` + accent("Workflows") + `  Automation scripts
  ` + accent("Rules") + `      Team structure
  ` + accent("Standards") + `  Development protocols

` + header("EXAMPLE") + `
  $ factory list
  
  📦 Blueprint Inventory
  
  Skills (21)
  ──────────────────────────────────────
  backend-go-expert    frontend-nuxt
  ...

  Workflows (2)
  ──────────────────────────────────────
  doc-cleanup    refactor`,
	RunE: runList,
}

func runList(_ *cobra.Command, _ []string) error {
	source := viper.GetString("source")
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	agentPath := filepath.Join(cwd, ".agent")

	// Check if .agent exists
	if _, err := os.Stat(agentPath); os.IsNotExist(err) {
		color.Yellow("⚠️  No .agent/ found. Run 'factory install' first.")
		fmt.Println()
		printSourceInventory(source)
		return nil
	}

	// Print installed inventory
	printInstalledInventory(agentPath)
	return nil
}

func printSourceInventory(source string) {
	fmt.Println(header("📦 Available Blueprint"))
	fmt.Println()

	// Skills
	skills := listDirItems(filepath.Join(source, "skills"))
	printCategory("Skills", skills)

	// Workflows
	workflows := listFiles(filepath.Join(source, "workflows"))
	printCategory("Workflows", workflows)

	// Rules
	rules := listFiles(filepath.Join(source, "rules"))
	printCategory("Rules", rules)

	// Standards
	standards := listFiles(filepath.Join(source, "standards"))
	printCategory("Standards", standards)
}

func printInstalledInventory(agentPath string) {
	fmt.Println(header("📦 Installed Blueprint"))
	fmt.Println()

	// Skills
	skills := listDirItems(filepath.Join(agentPath, "skills"))
	printCategory("Skills", skills)

	// Workflows
	workflows := listFiles(filepath.Join(agentPath, "workflows"))
	printCategory("Workflows", workflows)

	// Rules
	rules := listFiles(filepath.Join(agentPath, "rules"))
	printCategory("Rules", rules)

	// Standards
	standards := listFiles(filepath.Join(agentPath, "standards"))
	printCategory("Standards", standards)
}

func listDirItems(path string) []string {
	var items []string
	entries, err := os.ReadDir(path)
	if err != nil {
		return items
	}

	for _, e := range entries {
		if e.IsDir() && !strings.HasPrefix(e.Name(), ".") {
			items = append(items, e.Name())
		}
	}
	sort.Strings(items)
	return items
}

func listFiles(path string) []string {
	var items []string
	entries, err := os.ReadDir(path)
	if err != nil {
		return items
	}

	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
			name := strings.TrimSuffix(e.Name(), ".md")
			items = append(items, name)
		}
	}
	sort.Strings(items)
	return items
}

func printCategory(name string, items []string) {
	if len(items) == 0 {
		return
	}

	_, _ = color.New(color.FgCyan, color.Bold).Printf("%s (%d)\n", name, len(items))
	fmt.Println(strings.Repeat("─", 40))

	// Print in columns (3 per row)
	for i := 0; i < len(items); i += 3 {
		row := make([]string, 0, 3)
		for j := i; j < i+3 && j < len(items); j++ {
			row = append(row, items[j])
		}
		fmt.Printf("  %-20s %-20s %-20s\n", padItems(row)...)
	}
	fmt.Println()
}

func padItems(items []string) []interface{} {
	result := make([]interface{}, 3)
	for i := 0; i < 3; i++ {
		if i < len(items) {
			result[i] = items[i]
		} else {
			result[i] = ""
		}
	}
	return result
}

func init() {
	rootCmd.AddCommand(listCmd)
}
