package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/config"
)

func TestClient_ListServerClasses(t *testing.T) {
	tests := []struct {
		name           string
		serverResponse string
		serverStatus   int
		expectedError  string
		expectedItems  int
	}{
		{
			name:         "successful response",
			serverStatus: http.StatusOK,
			serverResponse: `{
				"apiVersion": "v1",
				"kind": "ServerClassList",
				"items": [
					{
						"apiVersion": "v1",
						"kind": "ServerClass",
						"metadata": {
							"name": "test-server-class"
						},
						"spec": {
							"displayName": "Test Server Class",
							"region": "uk-lon-1",
							"resources": {
								"cpu": "2",
								"memory": "4GB"
							},
							"availability": "available"
						},
						"status": {
							"available": 10,
							"capacity": 20
						}
					}
				]
			}`,
			expectedItems: 1,
		},
		{
			name:         "empty response",
			serverStatus: http.StatusOK,
			serverResponse: `{
				"apiVersion": "v1",
				"kind": "ServerClassList",
				"items": []
			}`,
			expectedItems: 0,
		},
		{
			name:           "server error",
			serverStatus:   http.StatusInternalServerError,
			serverResponse: `{"code": 500, "message": "Internal server error"}`,
			expectedError:  "API error 500: Internal server error",
		},
		{
			name:           "not found",
			serverStatus:   http.StatusNotFound,
			serverResponse: `{"code": 404, "message": "Not found"}`,
			expectedError:  "API error 404: Not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify request
				if r.URL.Path != "/ngpc.rxt.io/v1/serverclasses" {
					t.Errorf("Expected path '/ngpc.rxt.io/v1/serverclasses', got '%s'", r.URL.Path)
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

			// Create test config
			cfg := &config.Config{
				RefreshToken: "test-token",
				BaseURL:      server.URL,
				Region:       "uk-lon-1",
				Debug:        false,
				Timeout:      30,
			}

			// Create client with mock token manager
			client := NewClient(cfg)
			client.tokenManager = &MockTokenManager{
				accessToken: "mock-access-token",
			}

			// Test ListServerClasses
			ctx := context.Background()
			result, err := client.ListServerClasses(ctx)

			// Check error expectations
			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error '%s', got nil", tt.expectedError)
					return
				}
				if !contains(err.Error(), tt.expectedError) {
					t.Errorf("Expected error to contain '%s', got '%s'", tt.expectedError, err.Error())
				}
				return
			}

			// Check success expectations
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Error("Expected result, got nil")
				return
			}

			if len(result.Items) != tt.expectedItems {
				t.Errorf("Expected %d items, got %d", tt.expectedItems, len(result.Items))
			}

			// Verify the structure of returned data for successful cases
			if tt.expectedItems > 0 {
				serverClass := result.Items[0]
				if serverClass.Metadata.Name != "test-server-class" {
					t.Errorf("Expected name 'test-server-class', got '%s'", serverClass.Metadata.Name)
				}
				if serverClass.Spec.DisplayName != "Test Server Class" {
					t.Errorf("Expected display name 'Test Server Class', got '%s'", serverClass.Spec.DisplayName)
				}
				if serverClass.Spec.Region != "uk-lon-1" {
					t.Errorf("Expected region 'uk-lon-1', got '%s'", serverClass.Spec.Region)
				}
				if serverClass.Status.Available == nil || *serverClass.Status.Available != 10 {
					t.Errorf("Expected available count 10, got %v", serverClass.Status.Available)
				}
			}
		})
	}
}

func TestClient_ListServerClasses_InvalidJSON(t *testing.T) {
	// Create mock server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer server.Close()

	// Create test config
	cfg := &config.Config{
		RefreshToken: "test-token",
		BaseURL:      server.URL,
		Region:       "uk-lon-1",
		Debug:        false,
		Timeout:      30,
	}

	// Create client with mock token manager
	client := NewClient(cfg)
	client.tokenManager = &MockTokenManager{
		accessToken: "mock-access-token",
	}

	// Test ListServerClasses
	ctx := context.Background()
	_, err := client.ListServerClasses(ctx)

	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}

	if !contains(err.Error(), "failed to decode response") {
		t.Errorf("Expected decode error, got: %v", err)
	}
}

func TestClient_GetServerClass(t *testing.T) {
	tests := []struct {
		name            string
		serverClassName string
		serverResponse  string
		serverStatus    int
		expectedError   string
		expectedName    string
	}{
		{
			name:            "successful response",
			serverClassName: "test-server-class",
			serverStatus:    http.StatusOK,
			serverResponse: `{
				"apiVersion": "v1",
				"kind": "ServerClass",
				"metadata": {
					"name": "test-server-class"
				},
				"spec": {
					"displayName": "Test Server Class",
					"region": "uk-lon-1",
					"resources": {
						"cpu": "2",
						"memory": "4GB"
					},
					"availability": "available"
				},
				"status": {
					"available": 10,
					"capacity": 20
				}
			}`,
			expectedName: "test-server-class",
		},
		{
			name:            "not found",
			serverClassName: "nonexistent",
			serverStatus:    http.StatusNotFound,
			serverResponse:  `{"code": 404, "message": "ServerClass not found"}`,
			expectedError:   "API error 404: ServerClass not found",
		},
		{
			name:            "server error",
			serverClassName: "test-server-class",
			serverStatus:    http.StatusInternalServerError,
			serverResponse:  `{"code": 500, "message": "Internal server error"}`,
			expectedError:   "API error 500: Internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify request
				expectedPath := "/ngpc.rxt.io/v1/serverclasses/" + tt.serverClassName
				if r.URL.Path != expectedPath {
					t.Errorf("Expected path '%s', got '%s'", expectedPath, r.URL.Path)
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

			// Create test config
			cfg := &config.Config{
				RefreshToken: "test-token",
				BaseURL:      server.URL,
				Region:       "uk-lon-1",
				Debug:        false,
				Timeout:      30,
			}

			// Create client
			client := NewClient(cfg)
			client.tokenManager = &MockTokenManager{
				accessToken: "mock-access-token",
			}

			// Test GetServerClass
			ctx := context.Background()
			result, err := client.GetServerClass(ctx, tt.serverClassName)

			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error '%s', got nil", tt.expectedError)
					return
				}
				if !contains(err.Error(), tt.expectedError) {
					t.Errorf("Expected error to contain '%s', got '%s'", tt.expectedError, err.Error())
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
				if result.Metadata.Name != tt.expectedName {
					t.Errorf("Expected name '%s', got '%s'", tt.expectedName, result.Metadata.Name)
				}
			}
		})
	}
}

func TestClient_GetServerClass_InvalidJSON(t *testing.T) {
	// Create mock server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer server.Close()

	// Create test config
	cfg := &config.Config{
		RefreshToken: "test-token",
		BaseURL:      server.URL,
		Region:       "uk-lon-1",
		Debug:        false,
		Timeout:      30,
	}

	// Create client
	client := NewClient(cfg)
	client.tokenManager = &MockTokenManager{
		accessToken: "mock-access-token",
	}

	// Test GetServerClass
	ctx := context.Background()
	_, err := client.GetServerClass(ctx, "test-name")

	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}

	if !contains(err.Error(), "failed to decode response") {
		t.Errorf("Expected decode error, got: %v", err)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || (len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || indexContains(s, substr) >= 0)))
}

func indexContains(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
