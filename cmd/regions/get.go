package regions

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewGetCommand creates the regions get command
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <n>",
		Short: "Get a specific region by name",
		Long: `Get detailed information about a specific Rackspace Spot region.

This command retrieves and displays information for a single region
by its name, including provider details, location, and metadata.

Examples:
  rackspace-spot regions get uk-lon-1
  rackspace-spot regions get us-central-dfw-1 -o json`,
		Args: cobra.ExactArgs(1),
		RunE: runGet,
	}

	// Add flags for get command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	cmd.Flags().Bool("details", false, "Show additional details")
	cmd.Flags().BoolP("wide", "w", false, "Show wide output with additional columns")

	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
	regionName := args[0]

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if cfg.RefreshToken == "" {
		return fmt.Errorf("refresh token not configured. Run 'rackspace-spot-cli config init' to set up authentication")
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	region, err := client.GetRegion(ctx, regionName)
	if err != nil {
		return fmt.Errorf("failed to get region '%s': %w", regionName, err)
	}

	// Get flag values
	outputFormat, _ := cmd.Flags().GetString("output")
	showDetails, _ := cmd.Flags().GetBool("details")
	wideOutput, _ := cmd.Flags().GetBool("wide")

	return outputRegion(region, outputFormat, showDetails, wideOutput)
}
