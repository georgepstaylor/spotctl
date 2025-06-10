package client

import (
	"encoding/json"
	"fmt"
	"os"
)

// PatchOperation represents a JSON patch operation
type PatchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

// LoadPatchOperations loads JSON patch operations from a file
func LoadPatchOperations(filePath string) ([]PatchOperation, error) {
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

// DisplayPatchOperations shows the patch operations that will be applied
func DisplayPatchOperations(patchOps []PatchOperation) {
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

// PromptForConfirmation asks the user to confirm the patch operation
func PromptForConfirmation(resourceName string) (bool, error) {
	fmt.Printf("\nDo you want to apply these patches to '%s'? (y/N): ", resourceName)
	var response string
	fmt.Scanln(&response)
	return response == "y" || response == "Y" || response == "yes" || response == "Yes", nil
}
