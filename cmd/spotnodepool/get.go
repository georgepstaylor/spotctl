package spotnodepool

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewGetCommand returns the spotnodepool get command
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [NAME]",
		Short: "Get a specific spot node pool",
		Long: `Get detailed information about a specific spot node pool by name in the specified namespace.

Examples:
  # Get a specific spot node pool
  spotctl spotnodepool get my-nodepool --namespace org-abc123

  # Get spot node pool with detailed information
  spotctl spotnodepool get my-nodepool --namespace org-abc123 -o wide

  # Get spot node pool with JSON output
  spotctl spotnodepool get my-nodepool --namespace org-abc123 --output json

  # Get spot node pool with YAML output
  spotctl spotnodepool get my-nodepool --namespace org-abc123 --output yaml`,
		Args: cobra.ExactArgs(1),
		RunE: runGet,
	}

	// Add flags for spotnodepool get command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml, wide)")
	cmd.Flags().StringP("namespace", "n", "", "Namespace of the spot node pool (required)")

	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
	spotNodePoolName := args[0] // Get name from positional argument

	// Get flag values
	namespace, _ := cmd.Flags().GetString("namespace")
	outputFormat, _ := cmd.Flags().GetString("output")

	if namespace == "" {
		return fmt.Errorf("namespace is required")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	spotNodePool, err := client.GetSpotNodePool(ctx, namespace, spotNodePoolName)
	if err != nil {
		return fmt.Errorf("failed to get spot node pool: %w", err)
	}

	return outputSpotNodePool(spotNodePool, outputFormat)
}
