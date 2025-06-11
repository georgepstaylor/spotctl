package cloudspaces

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewGetCommand returns the cloudspaces get command
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [NAME]",
		Short: "Get a specific cloudspace",
		Long: `Get detailed information about a specific cloudspace by name in the specified namespace.

The namespace can be specified via:
- The --namespace/-n flag
- The 'namespace' field in your config file
- The SPOTCTL_NAMESPACE environment variable

Examples:
  # Get a cloudspace using namespace from config
  spotctl cloudspaces get my-cloudspace

  # Get cloudspace with specific namespace (overrides config)
  spotctl cloudspaces get my-cloudspace --namespace org-abc123

  # Get cloudspace with detailed information
  spotctl cloudspaces get my-cloudspace -o wide

  # Get cloudspace with JSON output
  spotctl cloudspaces get my-cloudspace --output json

  # Get cloudspace with YAML output
  spotctl cloudspaces get my-cloudspace --output yaml`,
		Args: cobra.ExactArgs(1),
		RunE: runGet,
	}

	// Add flags for cloudspaces get command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml, wide)")
	cmd.Flags().StringP("namespace", "n", "", "Namespace of the cloudspace (overrides config)")

	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
	cloudspaceName := args[0] // Get name from positional argument

	namespace, err := getNamespace(cmd)
	if err != nil {
		return err
	}

	outputFormat, _ := cmd.Flags().GetString("output")

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	cloudSpace, err := client.GetCloudSpace(ctx, namespace, cloudspaceName)
	if err != nil {
		return fmt.Errorf("failed to get cloudspace: %w", err)
	}

	return outputCloudSpace(cloudSpace, outputFormat)
}
