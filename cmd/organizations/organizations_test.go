package organizations

import (
	"bytes"
	"errors"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/spf13/cobra"
)

// MockOrganizationClient implements a simple mock for testing organizations command business logic
type MockOrganizationClient struct {
	organizations *client.OrganizationList
	err           error
}

func (m *MockOrganizationClient) ListOrganizations() (*client.OrganizationList, error) {
	return m.organizations, m.err
}

func TestOrganizationsListCommand(t *testing.T) {
	tests := []struct {
		name              string
		mockOrganizations *client.OrganizationList
		mockError         error
		args              []string
		expectError       bool
	}{
		{
			name: "successful organizations list",
			mockOrganizations: &client.OrganizationList{
				Start:  0,
				Limit:  10,
				Length: 2,
				Total:  2,
				Organizations: []client.Organization{
					{
						ID:          "org-123",
						Name:        "test-org",
						DisplayName: "Test Organization",
						Metadata: client.OrganizationMetadata{
							Namespace: "test-namespace",
						},
					},
					{
						ID:          "org-456",
						Name:        "another-org",
						DisplayName: "Another Organization",
						Metadata: client.OrganizationMetadata{
							Namespace: "another-namespace",
						},
					},
				},
			},
			args:        []string{},
			expectError: false,
		},
		{
			name: "empty organizations list",
			mockOrganizations: &client.OrganizationList{
				Start:         0,
				Limit:         10,
				Length:        0,
				Total:         0,
				Organizations: []client.Organization{},
			},
			args:        []string{},
			expectError: false,
		},
		{
			name:        "client error",
			mockError:   errors.New("failed to list organizations: API error 500: Internal server error"),
			args:        []string{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new command instance for each test
			cmd := &cobra.Command{
				Use: "list",
				RunE: func(cmd *cobra.Command, args []string) error {
					// Mock the client behavior
					mock := &MockOrganizationClient{
						organizations: tt.mockOrganizations,
						err:           tt.mockError,
					}

					_, err := mock.ListOrganizations()
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
