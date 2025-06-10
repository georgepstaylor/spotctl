package spotnodepool

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewEditCommand returns the spotnodepool edit command
func NewEditCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit <spotnodepool-name>",
		Short: "Edit a spot node pool",
		Long: `Edit a spot node pool in the specified namespace using JSON patch operations.

The patch operations should be provided in a JSON file with the following format:
[
  {
    "op": "replace",
    "path": "/spec/desired",
    "value": 5
  },
  {
    "op": "replace",
    "path": "/spec/autoscaling/enabled", 
    "value": true
  },
  {
    "op": "add",
    "path": "/spec/autoscaling/maxNodes",
    "value": 10
  }
]

Supported operations are: add, remove, replace, move, copy, test.

Examples:
  # Edit a spot node pool and show the result in table format
  spotctl spotnodepool edit my-nodepool --namespace org-abc123 --file patch.json

  # Edit with detailed information output
  spotctl spotnodepool edit my-nodepool --namespace org-abc123 --file patch.json -o wide

  # Edit and output the result as JSON
  spotctl spotnodepool edit my-nodepool --namespace org-abc123 --file patch.json --output json

  # Edit and output the result as YAML (skip confirmation)
  spotctl spotnodepool edit my-nodepool --namespace org-abc123 --file patch.json --output yaml --confirm`,
		Args: cobra.ExactArgs(1),
		RunE: runEdit,
	}

	// Add flags for spotnodepool edit command
	cmd.Flags().StringP("namespace", "n", "", "Namespace of the spot node pool (required)")
	cmd.Flags().StringP("file", "f", "", "Path to the JSON file containing patch operations (required)")
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml, wide)")
	cmd.Flags().Bool("confirm", false, "Skip confirmation prompt")

	// Mark flags as required
	cmd.MarkFlagRequired("namespace")
	cmd.MarkFlagRequired("file")

	return cmd
}

func runEdit(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	file, _ := cmd.Flags().GetString("file")
	outputFormat, _ := cmd.Flags().GetString("output")

	// Load the JSON patch operations from the file
	patchOps, err := client.LoadPatchOperations(file)
	if err != nil {
		return fmt.Errorf("failed to load patch operations: %w", err)
	}

	// Display the patch operations that will be applied
	client.DisplayPatchOperations(patchOps)

	// Create a new client
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	apiClient := client.NewClient(cfg)

	// Prompt for confirmation if the --confirm flag is not set
	skipConfirmation, _ := cmd.Flags().GetBool("confirm")
	if !skipConfirmation {
		confirmed, err := client.PromptForConfirmation(fmt.Sprintf("spot node pool '%s'", args[0]))
		if err != nil {
			return err
		}
		if !confirmed {
			fmt.Println("Patch operation cancelled.")
			return nil
		}
	}

	// Apply the patch operations
	updatedSpotNodePool, err := apiClient.EditSpotNodePool(context.Background(), namespace, args[0], patchOps)
	if err != nil {
		return fmt.Errorf("failed to edit spot node pool: %w", err)
	}

	// Output the updated spot node pool using the same formatting as the get command
	return outputSpotNodePool(updatedSpotNodePool, outputFormat)
}
