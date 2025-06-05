package serverclasses

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewListCommand creates the serverclasses list command
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all available server classes",
		Long: `List all available Rackspace Spot server classes.

This command retrieves and displays all server classes available across
all regions, including their specifications, pricing, and current availability.`,
		RunE: runList,
	}

	// Add flags for list command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml, wide)")

	return cmd
}

func runList(cmd *cobra.Command, args []string) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	serverClassList, err := client.ListServerClasses(ctx)
	if err != nil {
		return fmt.Errorf("failed to list server classes: %w", err)
	}

	// Read flags directly from command
	outputFormat, _ := cmd.Flags().GetString("output")

	return outputServerClasses(serverClassList, outputFormat)
}
