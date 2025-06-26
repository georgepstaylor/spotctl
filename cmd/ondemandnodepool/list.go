package ondemandnodepools

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewListCommand returns the ondemandnodepool list command
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List on demand node pools in a namespace",
		Long: `List all on demand node pools in the specified namespace.

The namespace can be specified via:
- The --namespace/-n flag
- The 'namespace' field in your config file
- The SPOTCTL_NAMESPACE environment variable

Examples:
  # List on demand node pools using namespace from config
  spotctl ondemandnodepool list

  # List on demand node pools with specific namespace (overrides config)
  spotctl ondemandnodepool list --namespace org-abc123

  # List with detailed output
  spotctl ondemandnodepool list --output wide

  # List with JSON output
  spotctl ondemandnodepool list --output json`,
		Args: cobra.NoArgs,
		RunE: runList,
	}

	// Add flags for ondemandnodepool list command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml, wide)")
	cmd.Flags().StringP("namespace", "n", "", "Namespace to list on demand node pools from (overrides config)")

	return cmd
}

func runList(cmd *cobra.Command, args []string) error {
	namespace, err := getNamespace(cmd)
	if err != nil {
		return err
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	apiClient := client.NewClient(cfg)

	ctx := context.Background()
	onDemandNodePoolList, err := apiClient.ListOnDemandNodePools(ctx, namespace)
	if err != nil {
		return fmt.Errorf("failed to list on demand node pools: %w", err)
	}

	outputFormat, _ := cmd.Flags().GetString("output")

	return outputOnDemandNodePools(onDemandNodePoolList, outputFormat, namespace)
}
