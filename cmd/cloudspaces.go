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

// cloudspacesCmd represents the cloudspaces command
var cloudspacesCmd = &cobra.Command{
	Use:   "cloudspaces",
	Short: "Manage Rackspace Spot cloudspaces",
	Long: `Manage and view Rackspace Spot cloudspaces.

This command allows you to list cloudspaces within a specific namespace.
Cloudspaces represent Kubernetes clusters deployed through Rackspace Spot.`,
}

// cloudspacesListCmd represents the cloudspaces list command
var cloudspacesListCmd = &cobra.Command{
	Use:   "list [NAMESPACE]",
	Short: "List cloudspaces in a namespace",
	Long: `List all cloudspaces in the specified namespace.

Examples:
  # List cloudspaces in a specific namespace
  spotctl cloudspaces list my-namespace

  # List cloudspaces with wide output
  spotctl cloudspaces list my-namespace --wide

  # List cloudspaces with JSON output
  spotctl cloudspaces list my-namespace --output json`,
	Args: cobra.ExactArgs(1),
	RunE: runCloudspacesList,
}

var (
	cloudspacesOutputFormat string
	cloudspacesShowDetails  bool
	cloudspacesWideOutput   bool
)

func runCloudspacesList(cmd *cobra.Command, args []string) error {
	namespace := args[0]

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	cloudSpaceList, err := client.ListCloudSpaces(ctx, namespace)
	if err != nil {
		return fmt.Errorf("failed to list cloudspaces: %w", err)
	}

	return outputCloudSpaces(cloudSpaceList, cloudspacesOutputFormat, cloudspacesShowDetails, cloudspacesWideOutput)
}

func outputCloudSpaces(cloudSpaceList *client.CloudSpaceList, format string, showDetails bool, wideOutput bool) error {
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

		// Get table configuration for cloudspaces
		tableConfig := getCloudSpacesTableConfig()

		// Pass the cloudspaces array directly instead of the full CloudSpaceList
		return formatter.Output(cloudSpaceList.Items, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Get table configuration for cloudspaces
	tableConfig := getCloudSpacesTableConfig()

	// Pass the cloudspaces array directly instead of the full CloudSpaceList
	return formatter.Output(cloudSpaceList.Items, tableConfig)
}

func getCloudSpacesTableConfig() *output.TableConfig {
	return &output.TableConfig{
		Columns: []output.TableColumn{
			{Header: "NAME", Field: "metadata.name"},
			{Header: "NAMESPACE", Field: "metadata.namespace"},
			{Header: "REGION", Field: "spec.region"},
			{Header: "PHASE", Field: "status.phase"},
			{Header: "HEALTH", Field: "status.health"},
		},
		DetailCols: []output.TableColumn{
			{Header: "K8S VERSION", Field: "status.currentKubernetesVersion"},
			{Header: "CNI", Field: "spec.cni"},
			{Header: "DEPLOYMENT TYPE", Field: "spec.deploymentType"},
			{Header: "HA CONTROL PLANE", Field: "spec.HAControlPlane"},
		},
	}
}

func init() {
	rootCmd.AddCommand(cloudspacesCmd)
	cloudspacesCmd.AddCommand(cloudspacesListCmd)

	// Add flags for cloudspaces list command
	cloudspacesListCmd.Flags().StringVarP(&cloudspacesOutputFormat, "output", "o", "table", "Output format (table, json, yaml)")
	cloudspacesListCmd.Flags().BoolVarP(&cloudspacesShowDetails, "details", "d", false, "Show detailed output")
	cloudspacesListCmd.Flags().BoolVarP(&cloudspacesWideOutput, "wide", "w", false, "Show additional columns")
}
