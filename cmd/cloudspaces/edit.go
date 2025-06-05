package cloudspaces

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// PatchOperation represents a JSON patch operation
type PatchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

// loadPatchOperations loads JSON patch operations from a file
func loadPatchOperations(filePath string) ([]PatchOperation, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var patchOps []PatchOperation
	if err := json.Unmarshal(data, &patchOps); err != nil {
		return nil, fmt.Errorf("failed to parse JSON patch operations: %w", err)
	}

	return patchOps, nil
}

// displayPatchOperations shows the patch operations that will be applied
func displayPatchOperations(patchOps []PatchOperation) {
	fmt.Printf("Applying %d patch operation(s):\n", len(patchOps))
	for i, op := range patchOps {
		fmt.Printf("  %d. %s %s", i+1, op.Op, op.Path)
		if op.Value != nil {
			// Try to format the value nicely
			switch v := op.Value.(type) {
			case string:
				fmt.Printf(" = %q", v)
			case bool:
				fmt.Printf(" = %t", v)
			case float64:
				// Check if it's actually an integer
				if v == float64(int(v)) {
					fmt.Printf(" = %d", int(v))
				} else {
					fmt.Printf(" = %g", v)
				}
			default:
				// For complex values, show as JSON
				valueJson, _ := json.Marshal(v)
				fmt.Printf(" = %s", valueJson)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// promptForConfirmation asks the user to confirm the patch operation
func promptForConfirmation(cloudspaceName string) (bool, error) {
	fmt.Printf("\nDo you want to apply these patches to cloudspace '%s'? (y/N): ", cloudspaceName)
	var response string
	fmt.Scanln(&response)
	return response == "y" || response == "Y" || response == "yes" || response == "Yes", nil
}

// NewEditCommand returns the cloudspaces edit command
func NewEditCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit <cloudspace-name>",
		Short: "Edit a cloudspace",
		Long: `Edit a cloudspace in the specified namespace using JSON patch operations.

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
  # Edit a cloudspace and show the result in table format
  spotctl cloudspaces edit my-cloudspace --namespace org-abc123 --file patch.json

  # Edit with detailed information output
  spotctl cloudspaces edit my-cloudspace --namespace org-abc123 --file patch.json -o wide

  # Edit and output the result as JSON
  spotctl cloudspaces edit my-cloudspace --namespace org-abc123 --file patch.json --output json

  # Edit and output the result as YAML (skip confirmation)
  spotctl cloudspaces edit my-cloudspace --namespace org-abc123 --file patch.json --output yaml --confirm`,
		Args: cobra.ExactArgs(1),
		RunE: runEdit,
	}

	// Add flags for cloudspaces edit command
	cmd.Flags().StringP("namespace", "n", "", "Namespace of the cloudspace (required)")
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
	patchOps, err := loadPatchOperations(file)
	if err != nil {
		return fmt.Errorf("failed to load patch operations: %w", err)
	}

	// Display the patch operations that will be applied
	displayPatchOperations(patchOps)

	// Create a new client
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	client := client.NewClient(cfg)

	// Prompt for confirmation if the --confirm flag is not set
	skipConfirmation, _ := cmd.Flags().GetBool("confirm")
	if !skipConfirmation {
		confirmed, err := promptForConfirmation(args[0])
		if err != nil {
			return err
		}
		if !confirmed {
			fmt.Println("Patch operation cancelled.")
			return nil
		}
	}

	// Apply the patch operations
	updatedCloudSpace, err := client.EditCloudSpace(context.Background(), namespace, args[0], patchOps)
	if err != nil {
		return fmt.Errorf("failed to edit cloudspace: %w", err)
	}

	// Output the updated cloudspace using the same formatting as the get command
	return outputCloudSpace(updatedCloudSpace, outputFormat)
}
