package cmd

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/georgetaylor/spotctl/pkg/output"
	"github.com/georgetaylor/spotctl/pkg/pager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

var (
	organizationsOutputFormat string
	organizationsShowDetails  bool
	organizationsWideOutput   bool
)

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

	return outputOrganizations(orgList, organizationsOutputFormat, organizationsShowDetails, organizationsWideOutput)
}

func outputOrganizations(orgList *client.OrganizationList, format string, showDetails bool, wideOutput bool) error {
	// Create formatter with options
	options := output.OutputOptions{
		Format:      output.OutputFormat(format),
		ShowDetails: showDetails,
		WideOutput:  wideOutput,
	}

	// Check if pager should be disabled
	noPager := viper.GetBool("no-pager")
	if noPager {
		// Create pager with disabled setting
		pager := pager.NewPager()
		pager.Disable = true
		formatter := output.NewFormatterWithPager(options, pager)

		// Get table configuration for organizations
		tableConfig := getOrganizationsTableConfig()

		// Pass the organizations array directly instead of the full OrganizationList
		return formatter.Output(orgList.Organizations, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Get table configuration for organizations
	tableConfig := getOrganizationsTableConfig()

	// Pass the organizations array directly instead of the full OrganizationList
	return formatter.Output(orgList.Organizations, tableConfig)
}

func getOrganizationsTableConfig() *output.TableConfig {
	return &output.TableConfig{
		Columns: []output.TableColumn{
			{Header: "ID", Field: "id"},
			{Header: "NAME", Field: "name"},
			{Header: "DISPLAY NAME", Field: "display_name"},
		},
		DetailCols: []output.TableColumn{
			{Header: "NAMESPACE", Field: "metadata.namespace"},
		},
		WideCols: []output.TableColumn{
			// Organizations don't have additional wide columns in the current API
		},
	}
}

func init() {
	rootCmd.AddCommand(organizationsCmd)
	organizationsCmd.AddCommand(organizationsListCmd)

	// Add flags for organizations list command
	organizationsListCmd.Flags().StringVarP(&organizationsOutputFormat, "output", "o", "table", "Output format (table, json, yaml)")
	organizationsListCmd.Flags().BoolVar(&organizationsShowDetails, "details", false, "Show detailed organization information")
	organizationsListCmd.Flags().BoolVar(&organizationsWideOutput, "wide", false, "Show additional columns")
}
