package factory

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Version info (set via ldflags during build).
var (
	version = "1.0.0"
	commit  = "dev"
)

// Color shortcuts for consistent styling.
var (
	header  = color.New(color.FgHiCyan, color.Bold).SprintFunc()
	accent  = color.New(color.FgHiYellow).SprintFunc()
	cmd     = color.New(color.FgGreen, color.Bold).SprintFunc()
	dimmed  = color.New(color.Faint).SprintFunc()
	success = color.New(color.FgGreen).SprintFunc()
)

const banner = `
   ______           __                  
  / ____/___ ______/ /_____  _______  __
 / /_  / __ '/ ___/ __/ __ \/ ___/ / / /
/ __/ / /_/ / /__/ /_/ /_/ / /  / /_/ / 
/_/    \__,_/\___/\__/\____/_/   \__, /  
                                /____/   
`

var rootCmd = &cobra.Command{
	Use:   "factory",
	Short: "Antigravity Factory — agent blueprint manager",
	Long: header("FACTORY") + dimmed(" — Antigravity Blueprint Manager") + `

` + dimmed(banner) + `
Manage AI agent infrastructure (skills, workflows, rules, standards).
Copy the blueprint to any project's .agent/ folder.

` + header("COMMANDS") + `
  ` + cmd("factory install") + `    → Copy blueprint to .agent/
  ` + cmd("factory list") + `       → View installed skills/workflows

` + header("BLUEPRINT CONTENTS") + `
  ` + accent("skills/") + `      Agent skills (21 experts)
  ` + accent("workflows/") + `   Automation workflows
  ` + accent("rules/") + `       Team structure (TEAM.md, PIPELINE.md)
  ` + accent("standards/") + `   Protocols (TDD, Git, Tech Debt)

` + header("CONFIGURATION") + `
  Config file: ` + dimmed("~/.config/factory/config.yaml") + `
  
  source: ` + dimmed("Path to blueprint folder") + `

` + header("EXAMPLES") + `
  ` + dimmed("# Install to current project") + `
  $ factory install

  ` + dimmed("# Check installed inventory") + `
  $ factory list`,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.config/factory/config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configPath := home + "/.config/factory"
		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

		// Set defaults — now points to dist/_agent/ (built from src/)
		viper.SetDefault("source", home+"/Developer/antigravity/antigravity-factory/dist/_agent")
	}

	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
}
