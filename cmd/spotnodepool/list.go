package spotnodepool

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewListCommand returns the spotnodepool list command
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List spot node pools in a namespace",
		Long: `List all spot node pools in the specified namespace.

Examples:
  # List spot node pools in a specific namespace
  spotctl spotnodepool list --namespace my-namespace

  # List spot node pools with detailed output
  spotctl spotnodepool list --namespace my-namespace -o wide

  # List spot node pools with JSON output
  spotctl spotnodepool list --namespace my-namespace --output json`,
		Args: cobra.NoArgs,
		RunE: runList,
	}

	// Add flags for spotnodepool list command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml, wide)")
	cmd.Flags().StringP("namespace", "n", "", "Namespace to list spot node pools from (required)")

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
	spotNodePoolList, err := client.ListSpotNodePools(ctx, namespace)
	if err != nil {
		return fmt.Errorf("failed to list spot node pools: %w", err)
	}

	outputFormat, _ := cmd.Flags().GetString("output")

	return outputSpotNodePools(spotNodePoolList, outputFormat, namespace)
}
