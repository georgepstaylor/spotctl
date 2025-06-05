package cloudspaces

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/spf13/cobra"
)

// MockCloudSpaceClient implements a simple mock for testing cloudspaces command business logic
type MockCloudSpaceClient struct {
	cloudspaces *client.CloudSpaceList
	cloudspace  *client.CloudSpace
	err         error
}

func (m *MockCloudSpaceClient) ListCloudSpaces(namespace string) (*client.CloudSpaceList, error) {
	return m.cloudspaces, m.err
}

func (m *MockCloudSpaceClient) CreateCloudSpace(namespace string, cloudSpace *client.CloudSpace) (*client.CloudSpace, error) {
	if m.err != nil {
		return nil, m.err
	}
	// Return a created cloudspace with some status information
	created := *cloudSpace
	created.Status = client.CloudSpaceStatus{
		Phase: "Creating",
	}
	return &created, nil
}

func (m *MockCloudSpaceClient) GetCloudSpace(namespace, name string) (*client.CloudSpace, error) {
	return m.cloudspace, m.err
}

func (m *MockCloudSpaceClient) EditCloudSpace(ctx context.Context, namespace, name string, patchOps interface{}) (*client.CloudSpace, error) {
	if m.err != nil {
		return nil, m.err
	}
	// Return the mock cloudspace (simulating successful patch)
	return m.cloudspace, nil
}

func TestCloudspacesListCommand(t *testing.T) {
	tests := []struct {
		name            string
		mockCloudSpaces *client.CloudSpaceList
		mockError       error
		args            []string
		expectError     bool
	}{
		{
			name: "successful cloudspaces list",
			mockCloudSpaces: &client.CloudSpaceList{
				APIVersion: "v1",
				Kind:       "CloudSpaceList",
				Items: []client.CloudSpace{
					{
						APIVersion: "v1",
						Kind:       "CloudSpace",
						Metadata: client.ObjectMeta{
							Name:      "test-cloudspace",
							Namespace: "test-namespace",
						},
						Spec: client.CloudSpaceSpec{
							Region: "uk-lon-1",
							Cloud:  "default",
							CNI:    "calico",
						},
						Status: client.CloudSpaceStatus{
							Phase:  "Running",
							Health: "Healthy",
						},
					},
					{
						APIVersion: "v1",
						Kind:       "CloudSpace",
						Metadata: client.ObjectMeta{
							Name:      "another-cloudspace",
							Namespace: "test-namespace",
						},
						Spec: client.CloudSpaceSpec{
							Region: "uk-lon-1",
							Cloud:  "default",
							CNI:    "flannel",
						},
						Status: client.CloudSpaceStatus{
							Phase:  "Provisioning",
							Health: "Unknown",
						},
					},
				},
			},
			args:        []string{"test-namespace"},
			expectError: false,
		},
		{
			name: "empty cloudspaces list",
			mockCloudSpaces: &client.CloudSpaceList{
				APIVersion: "v1",
				Kind:       "CloudSpaceList",
				Items:      []client.CloudSpace{},
			},
			args:        []string{"empty-namespace"},
			expectError: false,
		},
		{
			name:        "api error",
			mockError:   errors.New("failed to connect to API"),
			args:        []string{"test-namespace"},
			expectError: true,
		},
		{
			name:        "missing namespace argument",
			args:        []string{},
			expectError: true,
		},
		{
			name:        "too many arguments",
			args:        []string{"namespace1", "namespace2"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			// Create the command
			cmd := &cobra.Command{
				Use:  "list [NAMESPACE]",
				Args: cobra.ExactArgs(1),
				RunE: func(cmd *cobra.Command, args []string) error {
					// Mock the client behavior
					mockClient := &MockCloudSpaceClient{
						cloudspaces: tt.mockCloudSpaces,
						err:         tt.mockError,
					}

					// Simulate the function logic
					if tt.mockError != nil {
						return tt.mockError
					}

					// For this test, we'll just verify that we get the expected cloudspaces
					cloudspaces, err := mockClient.ListCloudSpaces(args[0])
					if err != nil {
						return err
					}

					// Write some output to verify the command works
					if len(cloudspaces.Items) == 0 {
						buf.WriteString("No cloudspaces found.")
					} else {
						buf.WriteString("Found cloudspaces:")
						for _, cs := range cloudspaces.Items {
							buf.WriteString(" " + cs.Metadata.Name)
						}
					}

					return nil
				},
			}

			// Set output to our buffer
			cmd.SetOut(&buf)
			cmd.SetErr(&buf)

			// Set the arguments
			cmd.SetArgs(tt.args)

			// Execute the command
			err := cmd.Execute()

			// Check error expectation
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				// Check output for successful cases
				output := buf.String()
				if tt.mockCloudSpaces != nil && len(tt.mockCloudSpaces.Items) == 0 {
					if output != "No cloudspaces found." {
						t.Errorf("Expected 'No cloudspaces found.' output, got: %s", output)
					}
				} else if tt.mockCloudSpaces != nil && len(tt.mockCloudSpaces.Items) > 0 {
					if output == "" {
						t.Errorf("Expected output but got empty string")
					}
					// Check that cloudspace names appear in output
					for _, cs := range tt.mockCloudSpaces.Items {
						if !bytes.Contains(buf.Bytes(), []byte(cs.Metadata.Name)) {
							t.Errorf("Expected cloudspace name '%s' in output: %s", cs.Metadata.Name, output)
						}
					}
				}
			}
		})
	}
}

func TestGetCloudSpacesTableConfig(t *testing.T) {
	config := getCloudSpacesTableConfig()

	if config == nil {
		t.Fatal("Expected table config but got nil")
	}

	// Check that basic columns are present
	expectedColumns := []string{"NAME", "NAMESPACE", "REGION", "PHASE", "HEALTH"}
	if len(config.Columns) != len(expectedColumns) {
		t.Errorf("Expected %d columns, got %d", len(expectedColumns), len(config.Columns))
	}

	for i, expected := range expectedColumns {
		if i >= len(config.Columns) {
			t.Errorf("Missing expected column: %s", expected)
			continue
		}
		if config.Columns[i].Header != expected {
			t.Errorf("Expected column header %s, got %s", expected, config.Columns[i].Header)
		}
	}

	// Check that detail columns are present
	if len(config.DetailCols) == 0 {
		t.Error("Expected detail columns but got none")
	}

	expectedDetailColumns := []string{"K8S VERSION", "CNI", "DEPLOYMENT TYPE", "HA CONTROL PLANE"}
	for i, expected := range expectedDetailColumns {
		if i >= len(config.DetailCols) {
			t.Errorf("Missing expected detail column: %s", expected)
			continue
		}
		if config.DetailCols[i].Header != expected {
			t.Errorf("Expected detail column header %s, got %s", expected, config.DetailCols[i].Header)
		}
	}
}

func TestCloudspacesCreateCommand(t *testing.T) {
	tests := []struct {
		name        string
		mockError   error
		args        []string
		flags       map[string]string
		expectError bool
	}{
		{
			name:      "successful cloudspace create",
			mockError: nil,
			args:      []string{"test-namespace"},
			flags: map[string]string{
				"name":               "test-cloudspace",
				"region":             "uk-lon-1",
				"kubernetes-version": "1.31.1",
			},
			expectError: false,
		},
		{
			name:      "missing namespace argument",
			mockError: nil,
			args:      []string{},
			flags: map[string]string{
				"name":   "test-cloudspace",
				"region": "uk-lon-1",
			},
			expectError: true,
		},
		{
			name:      "API error",
			mockError: errors.New("API error"),
			args:      []string{"test-namespace"},
			flags: map[string]string{
				"name":   "test-cloudspace",
				"region": "uk-lon-1",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock client with the test data
			mockClient := &MockCloudSpaceClient{
				err: tt.mockError,
			}

			// Create fresh commands for each test to avoid flag persistence
			cloudspacesCmd := &cobra.Command{
				Use:   "cloudspaces",
				Short: "Manage Rackspace Spot cloudspaces",
			}

			createCmd := &cobra.Command{
				Use:  "create [NAMESPACE]",
				Args: cobra.ExactArgs(1),
				RunE: func(cmd *cobra.Command, args []string) error {
					namespace := args[0]

					// Get flag values
					name, _ := cmd.Flags().GetString("name")
					region, _ := cmd.Flags().GetString("region")
					kubernetesVersion, _ := cmd.Flags().GetString("kubernetes-version")
					webhook, _ := cmd.Flags().GetString("webhook")
					haControlPlane, _ := cmd.Flags().GetBool("ha-control-plane")
					cloud, _ := cmd.Flags().GetString("cloud")
					cni, _ := cmd.Flags().GetString("cni")

					// Validate required fields
					if name == "" {
						return errors.New("cloudspace name is required")
					}
					if region == "" {
						return errors.New("region is required")
					}

					// Create the CloudSpace object
					cloudSpace := &client.CloudSpace{
						APIVersion: "ngpc.rxt.io/v1",
						Kind:       "CloudSpace",
						Metadata: client.ObjectMeta{
							Name:      name,
							Namespace: namespace,
						},
						Spec: client.CloudSpaceSpec{
							Region:            region,
							KubernetesVersion: kubernetesVersion,
							Webhook:           webhook,
							HAControlPlane:    haControlPlane,
							Cloud:             cloud,
							CNI:               cni,
						},
					}

					_, err := mockClient.CreateCloudSpace(namespace, cloudSpace)
					return err
				},
			}

			// Add flags
			createCmd.Flags().String("name", "", "Name of the cloudspace")
			createCmd.Flags().String("region", "", "Region for the cloudspace")
			createCmd.Flags().String("kubernetes-version", "1.31.1", "Kubernetes version")
			createCmd.Flags().String("webhook", "", "Webhook URL")
			createCmd.Flags().Bool("ha-control-plane", false, "Enable HA control plane")
			createCmd.Flags().String("cloud", "", "Cloud provider")
			createCmd.Flags().String("cni", "", "CNI plugin")

			cloudspacesCmd.AddCommand(createCmd)

			// Set flag values
			for flag, value := range tt.flags {
				createCmd.Flags().Set(flag, value)
			}

			// Create a buffer to capture output
			var buf bytes.Buffer
			cloudspacesCmd.SetOut(&buf)
			cloudspacesCmd.SetErr(&buf)

			// Set args and execute
			cloudspacesCmd.SetArgs(append([]string{"create"}, tt.args...))
			err := cloudspacesCmd.Execute()

			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestCloudspacesGetCommand(t *testing.T) {
	tests := []struct {
		name           string
		mockCloudSpace *client.CloudSpace
		mockError      error
		args           []string
		expectError    bool
	}{
		{
			name: "successful cloudspace get",
			mockCloudSpace: &client.CloudSpace{
				APIVersion: "ngpc.rxt.io/v1",
				Kind:       "CloudSpace",
				Metadata: client.ObjectMeta{
					Name:      "test-cloudspace",
					Namespace: "test-namespace",
				},
				Spec: client.CloudSpaceSpec{
					Region:            "uk-lon-1",
					Cloud:             "default",
					KubernetesVersion: "1.31.1",
					HAControlPlane:    true,
					CNI:               "calico",
				},
				Status: client.CloudSpaceStatus{
					Phase:  "Running",
					Health: "Healthy",
				},
			},
			mockError:   nil,
			args:        []string{"test-cloudspace", "--namespace", "test-namespace"},
			expectError: false,
		},
		{
			name:           "cloudspace not found",
			mockCloudSpace: nil,
			mockError:      errors.New("API error 404: CloudSpace not found"),
			args:           []string{"missing-cloudspace", "--namespace", "test-namespace"},
			expectError:    true,
		},
		{
			name:           "missing namespace flag",
			mockCloudSpace: nil,
			mockError:      nil,
			args:           []string{"test-cloudspace"},
			expectError:    true,
		},
		{
			name:           "missing cloudspace name",
			mockCloudSpace: nil,
			mockError:      nil,
			args:           []string{"--namespace", "test-namespace"},
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			// Create the command
			cmd := &cobra.Command{
				Use:  "get [NAME]",
				Args: cobra.ExactArgs(1),
				RunE: func(cmd *cobra.Command, args []string) error {
					// Check if namespace flag is provided
					namespace, _ := cmd.Flags().GetString("namespace")
					if namespace == "" {
						return errors.New("required flag \"namespace\" not set")
					}

					// Mock the client behavior
					mockClient := &MockCloudSpaceClient{
						cloudspace: tt.mockCloudSpace,
						err:        tt.mockError,
					}

					// Simulate the function logic
					if tt.mockError != nil {
						return tt.mockError
					}

					// For this test, we'll just verify that we get the expected cloudspace
					cloudspace, err := mockClient.GetCloudSpace(namespace, args[0])
					if err != nil {
						return err
					}

					// Write some output to verify the command works
					if cloudspace == nil {
						buf.WriteString("CloudSpace not found.")
					} else {
						buf.WriteString("Found cloudspace: " + cloudspace.Metadata.Name)
					}

					return nil
				},
			}

			// Add namespace flag
			cmd.Flags().String("namespace", "", "Namespace for the cloudspace")

			// Set output to our buffer
			cmd.SetOut(&buf)
			cmd.SetErr(&buf)

			// Set the arguments
			cmd.SetArgs(tt.args)

			// Execute the command
			err := cmd.Execute()

			// Check error expectation
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				// Check output for successful cases
				output := buf.String()
				if tt.mockCloudSpace != nil {
					expectedOutput := "Found cloudspace: " + tt.mockCloudSpace.Metadata.Name
					if output != expectedOutput {
						t.Errorf("Expected '%s' output, got: %s", expectedOutput, output)
					}
				}
			}
		})
	}
}

func TestCloudspacesEditCommand(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "edit-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a valid patch file
	validPatchFile := filepath.Join(tmpDir, "valid-patch.json")
	validPatch := `[
		{
			"op": "replace",
			"path": "/spec/webhook",
			"value": "https://example.com/webhook"
		}
	]`
	if err := os.WriteFile(validPatchFile, []byte(validPatch), 0644); err != nil {
		t.Fatalf("Failed to create valid patch file: %v", err)
	}

	// Create an invalid patch file
	invalidPatchFile := filepath.Join(tmpDir, "invalid-patch.json")
	invalidPatch := `invalid json`
	if err := os.WriteFile(invalidPatchFile, []byte(invalidPatch), 0644); err != nil {
		t.Fatalf("Failed to create invalid patch file: %v", err)
	}

	// Create an empty patch file
	emptyPatchFile := filepath.Join(tmpDir, "empty-patch.json")
	emptyPatch := `[]`
	if err := os.WriteFile(emptyPatchFile, []byte(emptyPatch), 0644); err != nil {
		t.Fatalf("Failed to create empty patch file: %v", err)
	}

	// Mock cloudspace that would be returned after edit
	mockUpdatedCloudSpace := &client.CloudSpace{
		APIVersion: "ngpc.rxt.io/v1",
		Kind:       "CloudSpace",
		Metadata: client.ObjectMeta{
			Name:      "test-cloudspace",
			Namespace: "test-namespace",
			Labels: map[string]string{
				"environment": "test",
			},
		},
		Spec: client.CloudSpaceSpec{
			Region:            "uk-lon-1",
			KubernetesVersion: "1.31.1",
			Cloud:             "default",
			CNI:               "calico",
		},
		Status: client.CloudSpaceStatus{
			Phase:  "Running",
			Health: "Healthy",
		},
	}

	tests := []struct {
		name        string
		args        []string
		flags       map[string]string
		mockError   error
		expectError bool
		expectOut   []string // strings that should appear in output
	}{
		{
			name: "successful edit with valid patch",
			args: []string{"test-cloudspace"},
			flags: map[string]string{
				"namespace": "test-namespace",
				"file":      validPatchFile,
				"confirm":   "true",
				"output":    "table",
			},
			mockError:   nil,
			expectError: false,
			expectOut:   []string{"Applying 1 patch operation(s)", "replace /spec/webhook"},
		},
		{
			name: "successful edit with json output",
			args: []string{"test-cloudspace"},
			flags: map[string]string{
				"namespace": "test-namespace",
				"file":      validPatchFile,
				"confirm":   "true",
				"output":    "json",
			},
			mockError:   nil,
			expectError: false,
			expectOut:   []string{"Applying 1 patch operation(s)", "\"name\":\"test-cloudspace\""},
		},
		{
			name: "successful edit with empty patch",
			args: []string{"test-cloudspace"},
			flags: map[string]string{
				"namespace": "test-namespace",
				"file":      emptyPatchFile,
				"confirm":   "true",
			},
			mockError:   nil,
			expectError: false,
			expectOut:   []string{"Applying 0 patch operation(s)"},
		},
		{
			name: "missing namespace flag",
			args: []string{"test-cloudspace"},
			flags: map[string]string{
				"file": validPatchFile,
			},
			expectError: true,
		},
		{
			name: "missing file flag",
			args: []string{"test-cloudspace"},
			flags: map[string]string{
				"namespace": "test-namespace",
			},
			expectError: true,
		},
		{
			name: "missing cloudspace name",
			args: []string{},
			flags: map[string]string{
				"namespace": "test-namespace",
				"file":      validPatchFile,
			},
			expectError: true,
		},
		{
			name: "nonexistent patch file",
			args: []string{"test-cloudspace"},
			flags: map[string]string{
				"namespace": "test-namespace",
				"file":      "/nonexistent/file.json",
				"confirm":   "true",
			},
			expectError: true,
		},
		{
			name: "invalid patch file",
			args: []string{"test-cloudspace"},
			flags: map[string]string{
				"namespace": "test-namespace",
				"file":      invalidPatchFile,
				"confirm":   "true",
			},
			expectError: true,
		},
		{
			name: "API error during edit",
			args: []string{"test-cloudspace"},
			flags: map[string]string{
				"namespace": "test-namespace",
				"file":      validPatchFile,
				"confirm":   "true",
			},
			mockError:   errors.New("API error: cloudspace not found"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := &MockCloudSpaceClient{
				cloudspace: mockUpdatedCloudSpace,
				err:        tt.mockError,
			}

			// Capture output
			var buf bytes.Buffer
			var stdoutBuf bytes.Buffer

			// Create the edit command
			editCmd := NewEditCommand()
			editCmd.SetOut(&buf)
			editCmd.SetErr(&buf)

			// Override the runEdit function to use our mock client
			editCmd.RunE = func(cmd *cobra.Command, args []string) error {
				namespace, _ := cmd.Flags().GetString("namespace")
				file, _ := cmd.Flags().GetString("file")
				outputFormat, _ := cmd.Flags().GetString("output")

				// Load the JSON patch operations from the file
				patchOps, err := loadPatchOperations(file)
				if err != nil {
					return err
				}

				// Capture displayPatchOperations output
				old := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w

				// Display the patch operations that will be applied
				displayPatchOperations(patchOps)

				// Restore stdout and capture the output
				w.Close()
				os.Stdout = old
				output, _ := io.ReadAll(r)
				stdoutBuf.Write(output)

				// Skip confirmation since we set confirm=true in flags or this is a test
				skipConfirmation, _ := cmd.Flags().GetBool("confirm")
				if !skipConfirmation {
					// In tests, we always confirm to avoid hanging
					buf.WriteString("Confirmation skipped in test\n")
				}

				// Apply the patch operations using mock client
				_, err = mockClient.EditCloudSpace(context.Background(), namespace, args[0], patchOps)
				if err != nil {
					return err
				}

				// Output the updated cloudspace - simplified for testing
				if outputFormat == "json" {
					buf.WriteString(`{"apiVersion":"ngpc.rxt.io/v1","kind":"CloudSpace","metadata":{"name":"test-cloudspace"}}`)
				} else {
					buf.WriteString("CloudSpace updated successfully")
				}

				return nil
			}

			// Set flags
			for flag, value := range tt.flags {
				if err := editCmd.Flags().Set(flag, value); err != nil {
					t.Fatalf("Failed to set flag %s=%s: %v", flag, value, err)
				}
			}

			// Set args and execute
			editCmd.SetArgs(tt.args)
			err := editCmd.Execute()

			// Check error expectation
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				// Check expected output strings in both buffers
				combinedOutput := stdoutBuf.String() + buf.String()
				for _, expected := range tt.expectOut {
					if !strings.Contains(combinedOutput, expected) {
						t.Errorf("Expected output to contain %q, but got: %s", expected, combinedOutput)
					}
				}
			}
		})
	}
}

func TestLoadPatchOperations(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "patch-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name        string
		fileContent string
		expectError bool
		expectCount int
	}{
		{
			name: "valid patch operations",
			fileContent: `[
				{
					"op": "replace",
					"path": "/spec/region",
					"value": "uk-lon-1"
				},
				{
					"op": "add",
					"path": "/metadata/labels/env",
					"value": "test"
				}
			]`,
			expectError: false,
			expectCount: 2,
		},
		{
			name:        "empty array",
			fileContent: `[]`,
			expectError: false,
			expectCount: 0,
		},
		{
			name:        "invalid json",
			fileContent: `invalid json`,
			expectError: true,
		},
		{
			name:        "valid json but wrong structure",
			fileContent: `{"not": "an array"}`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test file
			testFile := filepath.Join(tmpDir, tt.name+".json")
			if err := os.WriteFile(testFile, []byte(tt.fileContent), 0644); err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			// Test loadPatchOperations
			ops, err := loadPatchOperations(testFile)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if len(ops) != tt.expectCount {
					t.Errorf("Expected %d operations, got %d", tt.expectCount, len(ops))
				}
			}
		})
	}

	// Test nonexistent file
	t.Run("nonexistent file", func(t *testing.T) {
		_, err := loadPatchOperations("/nonexistent/file.json")
		if err == nil {
			t.Errorf("Expected error for nonexistent file but got none")
		}
	})
}

func TestDisplayPatchOperations(t *testing.T) {
	tests := []struct {
		name       string
		operations []PatchOperation
		expectOut  []string
	}{
		{
			name: "multiple operations with different types",
			operations: []PatchOperation{
				{Op: "replace", Path: "/spec/region", Value: "uk-lon-1"},
				{Op: "add", Path: "/metadata/labels/env", Value: "test"},
				{Op: "remove", Path: "/spec/oldField"},
				{Op: "replace", Path: "/spec/count", Value: float64(42)},
				{Op: "replace", Path: "/spec/enabled", Value: true},
			},
			expectOut: []string{
				"Applying 5 patch operation(s)",
				"1. replace /spec/region = \"uk-lon-1\"",
				"2. add /metadata/labels/env = \"test\"",
				"3. remove /spec/oldField",
				"4. replace /spec/count = 42",
				"5. replace /spec/enabled = true",
			},
		},
		{
			name:       "empty operations",
			operations: []PatchOperation{},
			expectOut:  []string{"Applying 0 patch operation(s)"},
		},
		{
			name: "operation with complex value",
			operations: []PatchOperation{
				{Op: "add", Path: "/spec/config", Value: map[string]interface{}{"key": "value", "nested": map[string]interface{}{"inner": "data"}}},
			},
			expectOut: []string{
				"Applying 1 patch operation(s)",
				"1. add /spec/config = {",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture output
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call function
			displayPatchOperations(tt.operations)

			// Restore stdout and read output
			w.Close()
			os.Stdout = old

			output, _ := io.ReadAll(r)
			outputStr := string(output)

			// Check expected strings
			for _, expected := range tt.expectOut {
				if !strings.Contains(outputStr, expected) {
					t.Errorf("Expected output to contain %q, but got: %s", expected, outputStr)
				}
			}
		})
	}
}
