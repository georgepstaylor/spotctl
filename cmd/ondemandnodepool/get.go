package ondemandnodepools

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewGetCommand returns the ondemandnodepool get command
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [NAME]",
		Short: "Get a specific on demand node pool",
		Long: `Get details of a specific on demand node pool by name in the specified namespace.

The namespace can be specified via:
- The --namespace/-n flag
- The 'namespace' field in your config file
- The SPOTCTL_NAMESPACE environment variable

Examples:
  # Get an on demand node pool using namespace from config
  spotctl ondemandnodepool get my-pool

  # Get an on demand node pool with specific namespace (overrides config)
  spotctl ondemandnodepool get my-pool --namespace org-abc123

  # Get with JSON output
  spotctl ondemandnodepool get my-pool --output json`,
		Args: cobra.ExactArgs(1),
		RunE: runGet,
	}

	// Add flags for ondemandnodepool get command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml, wide)")
	cmd.Flags().StringP("namespace", "n", "", "Namespace of the on demand node pool (overrides config)")

	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
	onDemandNodePoolName := args[0] // Get name from positional argument

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
	onDemandNodePool, err := apiClient.GetOnDemandNodePool(ctx, namespace, onDemandNodePoolName)
	if err != nil {
		return fmt.Errorf("failed to get on demand node pool: %w", err)
	}

	outputFormat, _ := cmd.Flags().GetString("output")

	return outputOnDemandNodePool(onDemandNodePool, outputFormat)
}
