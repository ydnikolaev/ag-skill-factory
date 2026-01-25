package factory

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print the version, commit, and build information for the factory CLI.`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("factory %s (commit: %s)\n", version, commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
