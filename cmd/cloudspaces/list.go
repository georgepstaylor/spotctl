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

The namespace can be specified via:
- The --namespace/-n flag
- The 'namespace' field in your config file
- The SPOTCTL_NAMESPACE environment variable

Examples:
  # List cloudspaces using namespace from config
  spotctl cloudspaces list

  # List cloudspaces in a specific namespace (overrides config)
  spotctl cloudspaces list --namespace my-namespace

  # List cloudspaces with detailed output
  spotctl cloudspaces list --namespace my-namespace -o wide

  # List cloudspaces with JSON output
  spotctl cloudspaces list --namespace my-namespace --output json`,
		Args: cobra.NoArgs,
		RunE: runList,
	}

	// Add flags for cloudspaces list command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml, wide)")
	cmd.Flags().StringP("namespace", "n", "", "Namespace to list cloudspaces from (overrides config)")

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

	client := client.NewClient(cfg)

	ctx := context.Background()
	cloudSpaceList, err := client.ListCloudSpaces(ctx, namespace)
	if err != nil {
		return fmt.Errorf("failed to list cloudspaces: %w", err)
	}

	outputFormat, _ := cmd.Flags().GetString("output")

	return outputCloudSpaces(cloudSpaceList, outputFormat, namespace)
}
