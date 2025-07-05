package cmd

import (
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version information for spotctl.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("spotctl \n")
		fmt.Printf("--------------------------------\n")
		fmt.Printf("version: %s\n", version.Version)
		fmt.Printf("commit hash: %s\n", version.Commit)
		fmt.Printf("build date: %s\n", version.Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
