package cloudspaces

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewDeleteCommand returns the cloudspaces delete command
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [NAME]",
		Short: "Delete a cloudspace",
		Long: `Delete a cloudspace by name in the specified namespace.

The namespace can be specified via:
- The --namespace/-n flag
- The 'namespace' field in your config file
- The SPOTCTL_NAMESPACE environment variable

Examples:
  # Delete a cloudspace using namespace from config
  spotctl cloudspaces delete my-cloudspace

  # Delete a cloudspace with specific namespace (overrides config)
  spotctl cloudspaces delete my-cloudspace --namespace org-abc123

  # Delete with confirmation
  spotctl cloudspaces delete my-cloudspace --confirm`,
		Args: cobra.ExactArgs(1),
		RunE: runDelete,
	}

	// Add flags for cloudspaces delete command
	cmd.Flags().StringP("namespace", "n", "", "Namespace of the cloudspace (overrides config)")
	cmd.Flags().Bool("confirm", false, "Skip confirmation prompt")

	return cmd
}

func runDelete(cmd *cobra.Command, args []string) error {
	cloudspaceName := args[0] // Get name from positional argument

	namespace, err := getNamespace(cmd)
	if err != nil {
		return err
	}

	confirm, _ := cmd.Flags().GetBool("confirm")

	// Ask for confirmation unless --confirm flag is used
	if !confirm {
		fmt.Printf("Are you sure you want to delete cloudspace '%s' in namespace '%s'? (y/N): ", cloudspaceName, namespace)
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" && response != "yes" && response != "Yes" {
			fmt.Println("Delete cancelled")
			return nil
		}
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	apiClient := client.NewClient(cfg)

	ctx := context.Background()
	deleteResponse, err := apiClient.DeleteCloudSpace(ctx, namespace, cloudspaceName)
	if err != nil {
		return fmt.Errorf("failed to delete cloudspace: %w", err)
	}

	// Check if the deletion was successful
	if deleteResponse.Status == "Success" || deleteResponse.Status == "" {
		fmt.Printf("Cloudspace '%s' deleted successfully from namespace '%s'\n", cloudspaceName, namespace)
	} else {
		fmt.Printf("Delete operation completed with status: %s\n", deleteResponse.Status)
		if deleteResponse.Message != "" {
			fmt.Printf("Message: %s\n", deleteResponse.Message)
		}
	}
	return nil
}
