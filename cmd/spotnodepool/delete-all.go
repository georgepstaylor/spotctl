package spotnodepool

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewDeleteAllCommand returns the spotnodepool delete-all command
func NewDeleteAllCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-all",
		Short: "Delete all spot node pools in a namespace",
		Long: `Delete all spot node pools in the specified namespace.

This command will first list all spot node pools in the namespace and then
ask for confirmation before proceeding with the deletion.

The namespace can be specified via:
- The --namespace/-n flag
- The 'namespace' field in your config file
- The SPOTCTL_NAMESPACE environment variable

Examples:
  # Delete all spot node pools using namespace from config
  spotctl spotnodepool delete-all

  # Delete all spot node pools with specific namespace (overrides config)
  spotctl spotnodepool delete-all --namespace org-abc123

  # Delete with confirmation
  spotctl spotnodepool delete-all --confirm`,
		Args: cobra.NoArgs,
		RunE: runDeleteAll,
	}

	// Add flags for spotnodepool delete-all command
	cmd.Flags().StringP("namespace", "n", "", "Namespace to delete spot node pools from (overrides config)")
	cmd.Flags().Bool("confirm", false, "Skip confirmation prompt")

	return cmd
}

func runDeleteAll(cmd *cobra.Command, args []string) error {
	namespace, err := getNamespace(cmd)
	if err != nil {
		return err
	}

	confirm, _ := cmd.Flags().GetBool("confirm")

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)
	ctx := context.Background()

	// First, list all spot node pools to show what will be deleted
	fmt.Printf("Listing spot node pools in namespace '%s':\n\n", namespace)

	spotNodePoolList, err := client.ListSpotNodePools(ctx, namespace)
	if err != nil {
		return fmt.Errorf("failed to list spot node pools: %w", err)
	}

	if len(spotNodePoolList.Items) == 0 {
		fmt.Println("No spot node pools found in the namespace.")
		return nil
	}

	// Display the list of spot node pools
	outputFormat, _ := cmd.Flags().GetString("output")
	if outputFormat == "" {
		outputFormat = "table"
	}

	err = outputSpotNodePools(spotNodePoolList, outputFormat, namespace)
	if err != nil {
		return fmt.Errorf("failed to output spot node pools: %w", err)
	}

	fmt.Printf("\nFound %d spot node pool(s) to delete.\n", len(spotNodePoolList.Items))

	// Ask for confirmation unless --confirm flag is used
	if !confirm {
		fmt.Printf("\nAre you sure you want to delete ALL %d spot node pool(s) in namespace '%s'? (y/N): ", len(spotNodePoolList.Items), namespace)
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" && response != "yes" && response != "Yes" {
			fmt.Println("Delete cancelled")
			return nil
		}
	}

	// Proceed with deletion
	deleteResponse, err := client.DeleteAllSpotNodePools(ctx, namespace)
	if err != nil {
		return fmt.Errorf("failed to delete all spot node pools: %w", err)
	}

	// Check if the deletion was successful
	if deleteResponse.Status == "Success" || deleteResponse.Status == "" {
		fmt.Printf("Successfully deleted all %d spot node pool(s) from namespace '%s'\n", len(spotNodePoolList.Items), namespace)
	} else {
		fmt.Printf("Delete operation completed with status: %s\n", deleteResponse.Status)
		if deleteResponse.Message != "" {
			fmt.Printf("Message: %s\n", deleteResponse.Message)
		}
	}
	return nil
}
