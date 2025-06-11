package cloudspaces

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewEditCommand returns the cloudspaces edit command
func NewEditCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit <cloudspace-name>",
		Short: "Edit a cloudspace",
		Long: `Edit a cloudspace in the specified namespace using JSON patch operations.

The namespace can be specified via:
- The --namespace/-n flag
- The 'namespace' field in your config file
- The SPOTCTL_NAMESPACE environment variable

The patch operations should be provided in a JSON file with the following format:
[
  {
    "op": "replace",
    "path": "/spec/someSpec1",
    "value": "anotherValue"
  },
  {
    "op": "add",
    "path": "/spec/someSpec2", 
    "value": "newValue"
  }
]

Supported operations are: add, remove, replace, move, copy, test.

Examples:
  # Edit a cloudspace using namespace from config
  spotctl cloudspaces edit my-cloudspace --file patch.json

  # Edit with specific namespace (overrides config)
  spotctl cloudspaces edit my-cloudspace --namespace org-abc123 --file patch.json

  # Edit with detailed information output
  spotctl cloudspaces edit my-cloudspace --file patch.json -o wide

  # Edit and output the result as JSON
  spotctl cloudspaces edit my-cloudspace --file patch.json --output json

  # Edit and output the result as YAML (skip confirmation)
  spotctl cloudspaces edit my-cloudspace --file patch.json --output yaml --confirm`,
		Args: cobra.ExactArgs(1),
		RunE: runEdit,
	}

	// Add flags for cloudspaces edit command
	cmd.Flags().StringP("namespace", "n", "", "Namespace of the cloudspace (overrides config)")
	cmd.Flags().StringP("file", "f", "", "Path to the JSON file containing patch operations (required)")
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml, wide)")
	cmd.Flags().Bool("confirm", false, "Skip confirmation prompt")

	// Mark only file as required (namespace comes from config/flag/env)
	cmd.MarkFlagRequired("file")

	return cmd
}

func runEdit(cmd *cobra.Command, args []string) error {
	namespace, err := getNamespace(cmd)
	if err != nil {
		return err
	}

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
		confirmed, err := client.PromptForConfirmation(fmt.Sprintf("cloudspace '%s'", args[0]))
		if err != nil {
			return err
		}
		if !confirmed {
			fmt.Println("Patch operation cancelled.")
			return nil
		}
	}

	// Apply the patch operations
	updatedCloudSpace, err := apiClient.EditCloudSpace(context.Background(), namespace, args[0], patchOps)
	if err != nil {
		return fmt.Errorf("failed to edit cloudspace: %w", err)
	}

	// Output the updated cloudspace using the same formatting as the get command
	return outputCloudSpace(updatedCloudSpace, outputFormat)
}
