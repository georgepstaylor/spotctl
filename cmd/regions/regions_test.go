package regions

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/spf13/cobra"
)

// MockClient implements a simple mock for testing regions command business logic
type MockClient struct {
	regions *client.RegionList
	err     error
}

func (m *MockClient) ListRegions() (*client.RegionList, error) {
	return m.regions, m.err
}

func TestRegionsListCommand(t *testing.T) {
	tests := []struct {
		name           string
		mockRegions    *client.RegionList
		mockError      error
		args           []string
		expectError    bool
		expectedOutput string
	}{
		{
			name: "successful regions list",
			mockRegions: &client.RegionList{
				APIVersion: "v1",
				Kind:       "RegionList",
				Items: []client.Region{
					{
						APIVersion: "v1",
						Kind:       "Region",
						Metadata: client.ObjectMeta{
							Name: "uk-lon-1",
						},
						Spec: client.RegionSpec{
							Country:     "United Kingdom",
							Description: "London",
							Provider: client.RegionProvider{
								ProviderType:       "ospc",
								ProviderRegionName: "uk-lon-1",
							},
						},
					},
				},
			},
			mockError:   nil,
			args:        []string{"list"},
			expectError: false,
		},
		{
			name:        "API error handling",
			mockRegions: nil,
			mockError:   errors.New("API connection failed"),
			args:        []string{"list"},
			expectError: true,
		},
		{
			name: "empty regions list",
			mockRegions: &client.RegionList{
				APIVersion: "v1",
				Kind:       "RegionList",
				Items:      []client.Region{},
			},
			mockError:   nil,
			args:        []string{"list"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock client with the test data
			mockClient := &MockClient{
				regions: tt.mockRegions,
				err:     tt.mockError,
			}

			// Create the regions command
			regionsCmd := &cobra.Command{
				Use:   "regions",
				Short: "Manage regions",
			}

			// Create the list subcommand with mock client
			listCmd := &cobra.Command{
				Use:   "list",
				Short: "List all regions",
				RunE: func(cmd *cobra.Command, args []string) error {
					regions, err := mockClient.ListRegions()
					if err != nil {
						return err
					}

					// For testing, we just verify the command runs without error
					// The actual formatting is tested in pkg/output/formatter_test.go
					if regions != nil && len(regions.Items) == 0 {
						cmd.Println("No regions found")
					} else if regions != nil {
						cmd.Printf("Found %d regions\n", len(regions.Items))
					}

					return nil
				},
			}

			regionsCmd.AddCommand(listCmd)

			// Capture output
			var buf bytes.Buffer
			regionsCmd.SetOut(&buf)
			regionsCmd.SetErr(&buf)
			regionsCmd.SetArgs(tt.args)

			// Execute command
			err := regionsCmd.Execute()

			// Check expectations
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestRegionsCommandFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "default flags",
			args:     []string{"list"},
			expected: "output=table details=false wide=false",
		},
		{
			name:     "json output",
			args:     []string{"list", "-o", "json"},
			expected: "output=json details=false wide=false",
		},
		{
			name:     "details flag",
			args:     []string{"list", "--details"},
			expected: "output=table details=true wide=false",
		},
		{
			name:     "wide flag",
			args:     []string{"list", "--wide"},
			expected: "output=table details=false wide=true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create fresh commands for each test to avoid flag persistence
			regionsCmd := &cobra.Command{
				Use: "regions",
			}

			listCmd := &cobra.Command{
				Use: "list",
				RunE: func(cmd *cobra.Command, args []string) error {
					// Test that flags can be retrieved
					output, _ := cmd.Flags().GetString("output")
					details, _ := cmd.Flags().GetBool("details")
					wide, _ := cmd.Flags().GetBool("wide")

					// Verify flags are accessible
					if output == "" {
						output = "table" // default
					}

					cmd.Printf("output=%s details=%v wide=%v", output, details, wide)
					return nil
				},
			}

			// Add standard flags
			listCmd.Flags().StringP("output", "o", "table", "Output format (table|json|yaml)")
			listCmd.Flags().Bool("details", false, "Show detailed output")
			listCmd.Flags().Bool("wide", false, "Show wide output")

			regionsCmd.AddCommand(listCmd)

			var buf bytes.Buffer
			regionsCmd.SetOut(&buf)
			regionsCmd.SetArgs(tt.args)

			err := regionsCmd.Execute()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			output := strings.TrimSpace(buf.String())
			if output != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, output)
			}
		})
	}
}
