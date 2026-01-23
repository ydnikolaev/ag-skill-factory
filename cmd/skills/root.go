package skills

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
	version = "0.1.0"
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
   _____ __ __ _ ____    
  / ___// //_/(_) / /____
  \__ \/ ,<  / / / / ___/
 ___/ / /| |/ / / (__  ) 
/____/_/ |_/_/_/_/____/  
`

var rootCmd = &cobra.Command{
	Use:   "skills",
	Short: "Manage Antigravity skills in workspaces",
	Long: header("SKILLS") + dimmed(" — Antigravity Skill Manager") + `

` + dimmed(banner) + `
A CLI tool for managing AI agent skills across workspaces.
Enables ` + accent("install") + `, ` + accent("update") + `, ` + accent("backport") + `, and ` + accent("list") + ` operations.

` + header("WORKFLOW") + `
  1. ` + cmd("skills install") + `    → Bootstrap .agent/ in new project
  2. ` + cmd("skills list") + `       → View skill status
  3. ` + cmd("skills update") + `     → Pull latest from factory
  4. ` + cmd("skills backport") + `   → Push changes back to factory

` + header("CONFIGURATION") + `
  Config file: ` + dimmed("~/.config/ag-skills/config.yaml") + `
  
  source:      ` + dimmed("Path to skill factory (squads/)") + `
  global_path: ` + dimmed("Path for global skill copies") + `

` + header("EXAMPLES") + `
  ` + dimmed("# Install skills to current project") + `
  $ skills install

  ` + dimmed("# Check for updates") + `
  $ skills update

  ` + dimmed("# Push local changes back to factory") + `
  $ skills backport my-skill`,
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.config/ag-skills/config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configPath := home + "/.config/ag-skills"
		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

		// Set defaults
		viper.SetDefault("source", home+"/Developer/antigravity/ag-skill-factory/squads")
		viper.SetDefault("global_path", home+"/.gemini/antigravity/global_skills")
	}

	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
}
