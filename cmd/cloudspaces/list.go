package cloudspaces

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewListCommand returns the cloudspaces list command
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List cloudspaces in a namespace",
		Long: `List all cloudspaces in the specified namespace.

Examples:
  # List cloudspaces in a specific namespace
  spotctl cloudspaces list --namespace my-namespace

  # List cloudspaces with wide output
  spotctl cloudspaces list --namespace my-namespace --wide

  # List cloudspaces with JSON output
  spotctl cloudspaces list --namespace my-namespace --output json`,
		Args: cobra.NoArgs,
		RunE: runList,
	}

	// Add flags for cloudspaces list command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	cmd.Flags().Bool("details", false, "Show detailed cloudspace information")
	cmd.Flags().Bool("wide", false, "Show additional columns")
	cmd.Flags().StringP("namespace", "n", "", "Namespace to list cloudspaces from (required)")

	return cmd
}

func runList(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	if namespace == "" {
		return fmt.Errorf("namespace is required")
	}

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

	outputFormat, _ := cmd.Flags().GetString("output")
	showDetails, _ := cmd.Flags().GetBool("details")
	wideOutput, _ := cmd.Flags().GetBool("wide")

	return outputCloudSpaces(cloudSpaceList, outputFormat, showDetails, wideOutput, namespace)
}
