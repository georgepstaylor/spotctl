package cmd

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// regionsCmd represents the regions command
var regionsCmd = &cobra.Command{
	Use:   "regions",
	Short: "Manage Rackspace Spot regions",
	Long: `Manage and view Rackspace Spot regions.

This command allows you to list and view information about available
Rackspace Spot regions where you can deploy resources.`,
}

// regionsListCmd represents the regions list command
var regionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available regions",
	Long: `List all available Rackspace Spot regions.

This command retrieves and displays all regions where Rackspace Spot
services are available, including their details such as provider
information and location.`,
	RunE: runRegionsList,
}

// regionsGetCmd represents the regions get command
var regionsGetCmd = &cobra.Command{
	Use:   "get <name>",
	Short: "Get a specific region by name",
	Long: `Get detailed information about a specific Rackspace Spot region.

This command retrieves and displays information for a single region
by its name, including provider details, location, and metadata.

Examples:
  rackspace-spot regions get uk-lon-1
  rackspace-spot regions get us-central-dfw-1 -o json`,
	Args: cobra.ExactArgs(1),
	RunE: runRegionsGet,
}

func init() {
	// Add regions command to root
	rootCmd.AddCommand(regionsCmd)

	// Add list subcommand to regions
	regionsCmd.AddCommand(regionsListCmd)

	// Add get subcommand to regions
	regionsCmd.AddCommand(regionsGetCmd)

	// Add flags for regions list command
	regionsListCmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	regionsListCmd.Flags().Bool("details", false, "Show detailed region information")
	regionsListCmd.Flags().Bool("wide", false, "Show additional columns including metadata")

	// Add flags for regions get command
	regionsGetCmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	regionsGetCmd.Flags().Bool("details", false, "Show detailed region information")
	regionsGetCmd.Flags().Bool("wide", false, "Show additional columns including metadata")
}

func runRegionsList(cmd *cobra.Command, args []string) error {
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
	wideOutput, _ := cmd.Flags().GetBool("wide")

	return outputRegions(regionList, outputFormat, showDetails, wideOutput)
}

func runRegionsGet(cmd *cobra.Command, args []string) error {
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
