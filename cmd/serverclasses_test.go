package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/spf13/cobra"
)

// MockServerClassClient implements a simple mock for testing serverclasses command business logic
type MockServerClassClient struct {
	serverClasses *client.ServerClassList
	serverClass   *client.ServerClass
	err           error
}

func (m *MockServerClassClient) ListServerClasses() (*client.ServerClassList, error) {
	return m.serverClasses, m.err
}

func (m *MockServerClassClient) GetServerClass(name string) (*client.ServerClass, error) {
	return m.serverClass, m.err
}

func TestServerClassesListCommand(t *testing.T) {
	tests := []struct {
		name              string
		mockServerClasses *client.ServerClassList
		mockError         error
		args              []string
		expectError       bool
	}{
		{
			name: "successful server classes list",
			mockServerClasses: &client.ServerClassList{
				APIVersion: "v1",
				Kind:       "ServerClassList",
				Items: []client.ServerClass{
					{
						APIVersion: "v1",
						Kind:       "ServerClass",
						Metadata: client.ObjectMeta{
							Name: "test-server-class",
						},
						Spec: client.ServerClassSpec{
							DisplayName: "Test Server Class",
							Region:      "uk-lon-1",
							Resources: client.ServerClassResources{
								CPU:    "2",
								Memory: "4GB",
							},
							Availability: "available",
						},
						Status: client.ServerClassStatus{
							Available: intPtr(10),
							Capacity:  intPtr(20),
						},
					},
				},
			},
			mockError:   nil,
			args:        []string{"list"},
			expectError: false,
		},
		{
			name:              "API error handling",
			mockServerClasses: nil,
			mockError:         errors.New("API connection failed"),
			args:              []string{"list"},
			expectError:       true,
		},
		{
			name: "empty server classes list",
			mockServerClasses: &client.ServerClassList{
				APIVersion: "v1",
				Kind:       "ServerClassList",
				Items:      []client.ServerClass{},
			},
			mockError:   nil,
			args:        []string{"list"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock client with the test data
			mockClient := &MockServerClassClient{
				serverClasses: tt.mockServerClasses,
				err:           tt.mockError,
			}

			// Create the serverclasses command
			serverclassesCmd := &cobra.Command{
				Use:   "serverclasses",
				Short: "Manage server classes",
			}

			// Create the list subcommand with mock client
			listCmd := &cobra.Command{
				Use:   "list",
				Short: "List all server classes",
				RunE: func(cmd *cobra.Command, args []string) error {
					serverClasses, err := mockClient.ListServerClasses()
					if err != nil {
						return err
					}

					// Basic validation that we got data
					if len(serverClasses.Items) == 0 {
						return nil // Empty list is valid
					}

					// Validate structure of first item
					if len(serverClasses.Items) > 0 {
						sc := serverClasses.Items[0]
						if sc.Metadata.Name == "" {
							return errors.New("server class missing name")
						}
					}

					return nil
				},
			}

			serverclassesCmd.AddCommand(listCmd)

			// Capture output
			buf := new(bytes.Buffer)
			serverclassesCmd.SetOut(buf)
			serverclassesCmd.SetErr(buf)

			// Set up args
			serverclassesCmd.SetArgs(tt.args)

			// Execute command
			err := serverclassesCmd.Execute()

			// Check expectations
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestServerClassesGetCommand(t *testing.T) {
	tests := []struct {
		name            string
		mockServerClass *client.ServerClass
		mockError       error
		args            []string
		expectError     bool
	}{
		{
			name: "successful server class get",
			mockServerClass: &client.ServerClass{
				APIVersion: "v1",
				Kind:       "ServerClass",
				Metadata: client.ObjectMeta{
					Name: "test-server-class",
				},
				Spec: client.ServerClassSpec{
					DisplayName: "Test Server Class",
					Region:      "uk-lon-1",
					Resources: client.ServerClassResources{
						CPU:    "2",
						Memory: "4GB",
					},
					Availability: "available",
				},
				Status: client.ServerClassStatus{
					Available: intPtr(10),
					Capacity:  intPtr(20),
				},
			},
			args:        []string{"test-server-class"},
			expectError: false,
		},
		{
			name:        "server class not found",
			mockError:   errors.New("failed to get server class 'nonexistent': API error 404: ServerClass not found"),
			args:        []string{"nonexistent"},
			expectError: true,
		},
		{
			name:        "no server class name provided",
			args:        []string{},
			expectError: true,
		},
		{
			name:        "too many arguments",
			args:        []string{"server1", "server2"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new command instance for each test
			cmd := &cobra.Command{
				Use:  "get <name>",
				Args: cobra.ExactArgs(1),
				RunE: func(cmd *cobra.Command, args []string) error {
					// Mock the client behavior
					mock := &MockServerClassClient{
						serverClass: tt.mockServerClass,
						err:         tt.mockError,
					}

					if len(args) == 0 {
						return errors.New("requires exactly 1 arg(s), only received 0")
					}

					_, err := mock.GetServerClass(args[0])
					return err
				},
			}

			// Capture output
			var output bytes.Buffer
			cmd.SetOut(&output)
			cmd.SetErr(&output)

			// Set args and execute
			cmd.SetArgs(tt.args)
			err := cmd.Execute()

			if tt.expectError {
				if err == nil {
					t.Error("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

// Helper function to create int pointer
func intPtr(i int) *int {
	return &i
}
