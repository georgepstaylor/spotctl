package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// FormatOutput formats data according to the specified output format
func FormatOutput(data interface{}, format string) error {
	switch strings.ToLower(format) {
	case "json":
		return outputJSON(data)
	case "table":
		return outputTable(data)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
}

// outputJSON outputs data in JSON format
func outputJSON(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// outputTable outputs data in table format (placeholder for now)
func outputTable(data interface{}) error {
	// For now, just output JSON. Later we can implement proper table formatting
	return outputJSON(data)
}

// FormatTime formats a time string in a user-friendly way
func FormatTime(timeStr string) string {
	if timeStr == "" {
		return "N/A"
	}

	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return timeStr
	}

	return t.Format("2006-01-02 15:04:05 MST")
}

// CheckError checks for an error and exits if one exists
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// AddOutputFlag adds a common --output flag to commands
func AddOutputFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("output", "o", "table", "Output format (json, table)")
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
