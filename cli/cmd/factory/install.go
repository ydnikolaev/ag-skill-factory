package factory

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ydnikolaev/antigravity-factory/internal/installer"
	"github.com/ydnikolaev/antigravity-factory/internal/presets"
)

var presetFlag string
var noInteractive bool

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Copy blueprint to .agent/ in current workspace",
	Long: header("INSTALL") + ` — Copy blueprint to .agent/

Copies the complete agent infrastructure to your project.

` + header("FLAGS") + `
  ` + accent("--preset=<name>") + `   Install specific preset (core, backend, frontend, etc.)
  ` + accent("--no-interactive") + ` Skip interactive selection

` + header("PRESETS") + `
  all       Full blueprint (21 skills)
  core      Pipeline essentials (5 skills)
  backend   Go backend (9 skills)
  frontend  Nuxt/Vue (8 skills)
  fullstack Full stack (12 skills)
  tma       Telegram Mini Apps (8 skills)
  cli       CLI/TUI apps (8 skills)
  minimal   Utilities only (2 skills)

` + header("EXAMPLE") + `
  $ factory install                    # Interactive selection
  $ factory install --preset=backend   # Specific preset`,
	RunE: runInstall,
}

func runInstall(_ *cobra.Command, _ []string) error {
	source := viper.GetString("source")

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	target := filepath.Join(cwd, ".agent")

	// Load presets
	presetsConfig, err := presets.Load(source)
	if err != nil {
		// No presets file - fall back to full install
		presetsConfig = nil
	}

	// Determine preset to use
	selectedPreset := presetFlag
	if selectedPreset == "" && !noInteractive && presetsConfig != nil {
		// Interactive mode
		selectedPreset, err = selectPresetInteractive(presetsConfig)
		if err != nil {
			return err
		}
	}
	if selectedPreset == "" {
		selectedPreset = "all"
	}

	// Always replace existing .agent/
	if _, err := os.Stat(target); err == nil {
		color.Cyan("🔄 Replacing existing .agent/...")
		if err := os.RemoveAll(target); err != nil {
			return fmt.Errorf("failed to remove existing .agent/: %w", err)
		}
	}

	// Create installer
	inst := installer.New(source, target)

	// If preset specified and not "all", filter skills
	if selectedPreset != "all" && presetsConfig != nil {
		allSkills := listSkills(filepath.Join(source, "skills"))
		skillsToInstall, err := presetsConfig.ResolveSkills(selectedPreset, allSkills)
		if err != nil {
			return err
		}
		inst.SetSkillFilter(skillsToInstall)
		color.Cyan("📦 Installing preset: %s (%d skills)", selectedPreset, len(skillsToInstall))
	}

	// Run installation
	result, err := inst.Install()
	if err != nil {
		return fmt.Errorf("installation failed: %w", err)
	}

	color.Green("✅ Installed %d skills, %d workflows, %d rules, %d templates",
		result.SkillCount, result.WorkflowCount, result.RuleCount, result.TemplateCount)
	return nil
}

func selectPresetInteractive(config presets.Config) (string, error) {
	presetList := config.List()
	
	// Sort by name
	sort.Slice(presetList, func(i, j int) bool {
		order := map[string]int{"all": 0, "core": 1, "backend": 2, "frontend": 3, "fullstack": 4, "tma": 5, "cli": 6, "minimal": 7}
		return order[presetList[i].Name] < order[presetList[j].Name]
	})

	options := make([]huh.Option[string], len(presetList))
	for i, p := range presetList {
		options[i] = huh.NewOption(fmt.Sprintf("%-10s — %s", p.Name, p.Description), p.Name)
	}

	var selected string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select preset to install").
				Options(options...).
				Value(&selected),
		),
	)

	if err := form.Run(); err != nil {
		return "", err
	}

	return selected, nil
}

func listSkills(skillsPath string) []string {
	entries, err := os.ReadDir(skillsPath)
	if err != nil {
		return nil
	}
	var skills []string
	for _, e := range entries {
		if e.IsDir() {
			skills = append(skills, e.Name())
		}
	}
	return skills
}

func init() {
	installCmd.Flags().StringVar(&presetFlag, "preset", "", "Preset to install (all, core, backend, frontend, fullstack, tma, cli, minimal)")
	installCmd.Flags().BoolVar(&noInteractive, "no-interactive", false, "Skip interactive selection")
	rootCmd.AddCommand(installCmd)
}
