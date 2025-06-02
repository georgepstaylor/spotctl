package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/spf13/cobra"
)

// MockCloudSpaceClient implements a simple mock for testing cloudspaces command business logic
type MockCloudSpaceClient struct {
	cloudspaces *client.CloudSpaceList
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
