package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/georgetaylor/spotctl/pkg/errors"
	"github.com/spf13/cobra"
)

// AddOutputFlag adds a common --output flag to commands
func AddOutputFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("output", "o", "table", "Output format (json, table, yaml)")
}

// GetOutputFormat gets the output format from command flags
func GetOutputFormat(cmd *cobra.Command) string {
	output, _ := cmd.Flags().GetString("output")
	return output
}

// ConfirmAction prompts the user for confirmation
func ConfirmAction(message string) bool {
	fmt.Printf("%s (y/N): ", message)
	var response string
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y" || strings.ToLower(response) == "yes"
}

// CheckError checks for an error and exits if one exists
func CheckError(err error) {
	if err != nil {
		if appErr, ok := err.(*errors.Error); ok {
			fmt.Fprintf(os.Stderr, "Error: %s\n", appErr.Message)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		os.Exit(1)
	}
}
