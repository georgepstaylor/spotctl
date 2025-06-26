package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/config"
)

func TestClient_ListOrganizations(t *testing.T) {
	tests := []struct {
		name             string
		serverResponse   string
		serverStatus     int
		expectedError    string
		expectedOrgCount int
	}{
		{
			name:         "successful response",
			serverStatus: http.StatusOK,
			serverResponse: `{
				"start": 0,
				"limit": 10,
				"length": 2,
				"total": 2,
				"organizations": [
					{
						"id": "org-123",
						"name": "test-org",
						"display_name": "Test Organization",
						"metadata": {
							"namespace": "test-namespace"
						}
					},
					{
						"id": "org-456",
						"name": "another-org",
						"display_name": "Another Organization",
						"metadata": {
							"namespace": "another-namespace"
						}
					}
				]
			}`,
			expectedOrgCount: 2,
		},
		{
			name:         "empty response",
			serverStatus: http.StatusOK,
			serverResponse: `{
				"start": 0,
				"limit": 10,
				"length": 0,
				"total": 0,
				"organizations": []
			}`,
			expectedOrgCount: 0,
		},
		{
			name:           "server error",
			serverStatus:   http.StatusInternalServerError,
			serverResponse: `{"code": 500, "message": "Internal server error"}`,
			expectedError:  "API error 500: Internal server error",
		},
		{
			name:           "unauthorized",
			serverStatus:   http.StatusUnauthorized,
			serverResponse: `{"code": 401, "message": "Unauthorized"}`,
			expectedError:  "API error 401: Unauthorized",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify request
				if r.URL.Path != "/auth.ngpc.rxt.io/v1/organizations" {
					t.Errorf("Expected path '/auth.ngpc.rxt.io/v1/organizations', got '%s'", r.URL.Path)
				}
				if r.Method != http.MethodGet {
					t.Errorf("Expected method GET, got %s", r.Method)
				}

				// Check for auth header
				authHeader := r.Header.Get("Authorization")
				if authHeader == "" {
					t.Error("Expected Authorization header")
				}

				w.WriteHeader(tt.serverStatus)
				w.Write([]byte(tt.serverResponse))
			}))
			defer server.Close()

			// Create test config - set base URL to match expected pattern for auth conversion
			cfg := &config.Config{
				RefreshToken: "test-token",
				BaseURL:      server.URL,
				Debug:        false,
				Timeout:      30,
			}

			// Create client
			client := NewClient(cfg)
			client.tokenManager = &MockTokenManager{
				accessToken: "mock-access-token",
			}

			// Test ListOrganizations
			ctx := context.Background()
			result, err := client.ListOrganizations(ctx)

			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error '%s', got nil", tt.expectedError)
					return
				}
				if !contains(err.Error(), tt.expectedError) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
					return
				}
				if result == nil {
					t.Error("Expected non-nil result")
					return
				}
				if len(result.Organizations) != tt.expectedOrgCount {
					t.Errorf("Expected %d organizations, got %d", tt.expectedOrgCount, len(result.Organizations))
				}

				// Verify first organization if present
				if tt.expectedOrgCount > 0 {
					org := result.Organizations[0]
					if org.ID != "org-123" {
						t.Errorf("Expected ID 'org-123', got '%s'", org.ID)
					}
					if org.Name != "test-org" {
						t.Errorf("Expected name 'test-org', got '%s'", org.Name)
					}
					if org.DisplayName != "Test Organization" {
						t.Errorf("Expected display name 'Test Organization', got '%s'", org.DisplayName)
					}
					if org.Metadata.Namespace != "test-namespace" {
						t.Errorf("Expected namespace 'test-namespace', got '%s'", org.Metadata.Namespace)
					}
				}
			}
		})
	}
}

func TestClient_ListOrganizations_InvalidJSON(t *testing.T) {
	// Create mock server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer server.Close()

	// Create test config
	cfg := &config.Config{
		RefreshToken: "test-token",
		BaseURL:      server.URL + "/apis/ngpc.rxt.io/v1",
		Debug:        false,
		Timeout:      30,
	}

	// Create client
	client := NewClient(cfg)
	client.tokenManager = &MockTokenManager{
		accessToken: "mock-access-token",
	}

	// Test ListOrganizations
	ctx := context.Background()
	_, err := client.ListOrganizations(ctx)

	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}
