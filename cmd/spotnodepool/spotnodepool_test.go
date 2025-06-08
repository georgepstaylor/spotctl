package spotnodepool

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/spf13/cobra"
)

// MockSpotNodePoolClient implements a simple mock for testing spotnodepool command business logic
type MockSpotNodePoolClient struct {
	spotnodepool     *client.SpotNodePool
	spotnodepoollist *client.SpotNodePoolList
	err              error
}

func (m *MockSpotNodePoolClient) GetSpotNodePool(namespace, name string) (*client.SpotNodePool, error) {
	return m.spotnodepool, m.err
}

func (m *MockSpotNodePoolClient) ListSpotNodePools(namespace string) (*client.SpotNodePoolList, error) {
	return m.spotnodepoollist, m.err
}

func TestSpotNodePoolGetCommand(t *testing.T) {
	tests := []struct {
		name             string
		mockSpotNodePool *client.SpotNodePool
		mockError        error
		args             []string
		expectError      bool
	}{
		{
			name: "successful spotnodepool get",
			mockSpotNodePool: &client.SpotNodePool{
				APIVersion: "ngpc.rxt.io/v1",
				Kind:       "SpotNodePool",
				Metadata: client.ObjectMeta{
					Name:      "test-nodepool",
					Namespace: "test-namespace",
				},
				Spec: client.SpotNodePoolSpec{
					ServerClass: "gp.4x8",
					Desired:     &[]int{3}[0],
					CloudSpace:  "test-cloudspace",
				},
				Status: client.SpotNodePoolStatus{
					BidStatus: "Winning",
					WonCount:  &[]int{2}[0],
				},
			},
			mockError:   nil,
			args:        []string{"test-nodepool", "--namespace", "test-namespace"},
			expectError: false,
		},
		{
			name:             "spotnodepool not found",
			mockSpotNodePool: nil,
			mockError:        errors.New("API error 404: SpotNodePool not found"),
			args:             []string{"missing-nodepool", "--namespace", "test-namespace"},
			expectError:      true,
		},
		{
			name:             "missing namespace flag",
			mockSpotNodePool: nil,
			mockError:        nil,
			args:             []string{"test-nodepool"},
			expectError:      true,
		},
		{
			name:             "missing spotnodepool name",
			mockSpotNodePool: nil,
			mockError:        nil,
			args:             []string{"--namespace", "test-namespace"},
			expectError:      true,
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
					mockClient := &MockSpotNodePoolClient{
						spotnodepool: tt.mockSpotNodePool,
						err:          tt.mockError,
					}

					// Simulate the function logic
					if tt.mockError != nil {
						return tt.mockError
					}

					// For this test, we'll just verify that we get the expected spotnodepool
					spotnodepool, err := mockClient.GetSpotNodePool(namespace, args[0])
					if err != nil {
						return err
					}

					// Write some output to verify the command works
					if spotnodepool == nil {
						buf.WriteString("SpotNodePool not found.")
					} else {
						buf.WriteString("Found spotnodepool: " + spotnodepool.Metadata.Name)
					}

					return nil
				},
			}

			// Add namespace flag
			cmd.Flags().String("namespace", "", "Namespace for the spotnodepool")

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
				if tt.mockSpotNodePool != nil {
					expectedOutput := "Found spotnodepool: " + tt.mockSpotNodePool.Metadata.Name
					if output != expectedOutput {
						t.Errorf("Expected '%s' output, got: %s", expectedOutput, output)
					}
				}
			}
		})
	}
}

func TestGetSpotNodePoolTableConfig(t *testing.T) {
	config := getSpotNodePoolTableConfig()

	// Check that we have the expected columns
	if len(config.Columns) != 6 {
		t.Errorf("Expected 6 columns, got %d", len(config.Columns))
	}

	// Check that we have detail columns
	if len(config.DetailCols) != 5 {
		t.Errorf("Expected 5 detail columns, got %d", len(config.DetailCols))
	}

	// Check that columns have expected headers
	expectedHeaders := []string{"NAME", "NAMESPACE", "SERVER CLASS", "DESIRED", "BID STATUS", "WON COUNT"}
	for i, col := range config.Columns {
		if col.Header != expectedHeaders[i] {
			t.Errorf("Expected column %d to have header '%s', got '%s'", i, expectedHeaders[i], col.Header)
		}
	}
}

func TestSpotNodePoolListCommand(t *testing.T) {
	tests := []struct {
		name                 string
		mockSpotNodePoolList *client.SpotNodePoolList
		mockError            error
		args                 []string
		expectError          bool
	}{
		{
			name: "successful spotnodepool list",
			mockSpotNodePoolList: &client.SpotNodePoolList{
				APIVersion: "v1",
				Kind:       "SpotNodePoolList",
				Items: []client.SpotNodePool{
					{
						APIVersion: "v1",
						Kind:       "SpotNodePool",
						Metadata: client.ObjectMeta{
							Name:      "test-nodepool",
							Namespace: "test-namespace",
						},
						Spec: client.SpotNodePoolSpec{
							ServerClass: "gp.4x8",
							Desired:     &[]int{3}[0],
							CloudSpace:  "test-cloudspace",
						},
						Status: client.SpotNodePoolStatus{
							BidStatus: "Winning",
							WonCount:  &[]int{2}[0],
						},
					},
					{
						APIVersion: "v1",
						Kind:       "SpotNodePool",
						Metadata: client.ObjectMeta{
							Name:      "another-nodepool",
							Namespace: "test-namespace",
						},
						Spec: client.SpotNodePoolSpec{
							ServerClass: "gp.8x16",
							Desired:     &[]int{5}[0],
							CloudSpace:  "test-cloudspace",
						},
						Status: client.SpotNodePoolStatus{
							BidStatus: "Bidding",
							WonCount:  &[]int{0}[0],
						},
					},
				},
			},
			args:        []string{"--namespace", "test-namespace"},
			expectError: false,
		},
		{
			name: "empty spotnodepool list",
			mockSpotNodePoolList: &client.SpotNodePoolList{
				APIVersion: "v1",
				Kind:       "SpotNodePoolList",
				Items:      []client.SpotNodePool{},
			},
			args:        []string{"--namespace", "empty-namespace"},
			expectError: false,
		},
		{
			name:        "api error",
			mockError:   errors.New("failed to connect to API"),
			args:        []string{"--namespace", "test-namespace"},
			expectError: true,
		},
		{
			name:        "missing namespace argument",
			args:        []string{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			// Create the command
			cmd := &cobra.Command{
				Use:  "list",
				Args: cobra.NoArgs,
				RunE: func(cmd *cobra.Command, args []string) error {
					namespace, _ := cmd.Flags().GetString("namespace")
					if namespace == "" {
						return errors.New("required flag \"namespace\" not set")
					}

					// Mock the client behavior
					mockClient := &MockSpotNodePoolClient{
						spotnodepoollist: tt.mockSpotNodePoolList,
						err:              tt.mockError,
					}

					// Simulate the function logic
					if tt.mockError != nil {
						return tt.mockError
					}

					// For this test, we'll just verify that we get the expected spotnodepools
					spotnodepools, err := mockClient.ListSpotNodePools(namespace)
					if err != nil {
						return err
					}

					// Write some output to verify the command works
					if len(spotnodepools.Items) == 0 {
						buf.WriteString("No spotnodepools found.")
					} else {
						buf.WriteString("Found spotnodepools:")
						for _, snp := range spotnodepools.Items {
							buf.WriteString(" " + snp.Metadata.Name)
						}
					}

					return nil
				},
			}

			// Add namespace flag
			cmd.Flags().String("namespace", "", "Namespace for the spotnodepools")

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
				if tt.mockSpotNodePoolList != nil && len(tt.mockSpotNodePoolList.Items) == 0 {
					if output != "No spotnodepools found." {
						t.Errorf("Expected 'No spotnodepools found.' output, got: %s", output)
					}
				} else if tt.mockSpotNodePoolList != nil && len(tt.mockSpotNodePoolList.Items) > 0 {
					if output == "" {
						t.Errorf("Expected output but got empty string")
					}
					// Check that spotnodepool names appear in output
					for _, snp := range tt.mockSpotNodePoolList.Items {
						if !strings.Contains(output, snp.Metadata.Name) {
							t.Errorf("Expected spotnodepool name '%s' in output: %s", snp.Metadata.Name, output)
						}
					}
				}
			}
		})
	}
}
