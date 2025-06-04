package serverclasses

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewGetCommand creates the serverclasses get command
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <name>",
		Short: "Get a specific server class by name",
		Long: `Get detailed information about a specific Rackspace Spot server class.

This command retrieves and displays information for a single server class
by its name, including specifications, pricing, and current availability.

Examples:
  spotctl serverclasses get standard-2
  spotctl serverclasses get compute-optimized-4 -o json`,
		Args: cobra.ExactArgs(1),
		RunE: runGet,
	}

	// Add flags for get command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	cmd.Flags().Bool("details", false, "Show detailed server class information")
	cmd.Flags().Bool("wide", false, "Show additional columns including availability and pricing")

	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	name := args[0]
	serverClass, err := client.GetServerClass(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to get server class '%s': %w", name, err)
	}

	// Read flags directly from command
	outputFormat, _ := cmd.Flags().GetString("output")
	showDetails, _ := cmd.Flags().GetBool("details")
	wideOutput, _ := cmd.Flags().GetBool("wide")

	return outputServerClass(serverClass, outputFormat, showDetails, wideOutput)
}
