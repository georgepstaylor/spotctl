package regions

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewListCommand creates the regions list command
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all available regions",
		Long: `List all available Rackspace Spot regions.

This command retrieves and displays all regions where Rackspace Spot
services are available, including their details such as provider
information and location.`,
		RunE: runList,
	}

	// Add flags for list command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	cmd.Flags().Bool("details", false, "Show additional details")

	return cmd
}

func runList(cmd *cobra.Command, args []string) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if cfg.RefreshToken == "" {
		return fmt.Errorf("refresh token not configured. Run 'rackspace-spot-cli config init' to set up authentication")
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	regionList, err := client.ListRegions(ctx)
	if err != nil {
		return fmt.Errorf("failed to list regions: %w", err)
	}

	// Get flag values
	outputFormat, _ := cmd.Flags().GetString("output")
	showDetails, _ := cmd.Flags().GetBool("details")

	return outputRegions(regionList, outputFormat, showDetails)
}
