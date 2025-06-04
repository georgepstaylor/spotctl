package cmd

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// organizationsCmd represents the organizations command
var organizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "Manage Rackspace Spot organizations",
	Long: `Manage and view Rackspace Spot organizations.

This command allows you to list organizations that you have access to
in the Rackspace Spot platform.`,
}

// organizationsListCmd represents the organizations list command
var organizationsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available organizations",
	Long: `List all Rackspace Spot organizations that you have access to.

This command retrieves and displays all organizations available to
your authenticated account, including organization details and metadata.`,
	RunE: runOrganizationsList,
}

func runOrganizationsList(cmd *cobra.Command, args []string) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	orgList, err := client.ListOrganizations(ctx)
	if err != nil {
		return fmt.Errorf("failed to list organizations: %w", err)
	}

	// Get flag values
	outputFormat, _ := cmd.Flags().GetString("output")
	showDetails, _ := cmd.Flags().GetBool("details")
	wideOutput, _ := cmd.Flags().GetBool("wide")

	return outputOrganizations(orgList, outputFormat, showDetails, wideOutput)
}

func init() {
	rootCmd.AddCommand(organizationsCmd)
	organizationsCmd.AddCommand(organizationsListCmd)

	// Add flags for organizations list command
	organizationsListCmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	organizationsListCmd.Flags().Bool("details", false, "Show detailed organization information")
	organizationsListCmd.Flags().Bool("wide", false, "Show additional columns")
}
