package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/config"
)

func TestListCloudSpaces(t *testing.T) {
	tests := []struct {
		name              string
		namespace         string
		mockResponse      string
		mockStatus        int
		expectError       bool
		expectedErrorMsg  string
		expectedCount     int
		expectedFirstName string
	}{
		{
			name:      "successful cloudspaces list",
			namespace: "test-namespace",
			mockResponse: `{
				"apiVersion": "v1",
				"kind": "CloudSpaceList",
				"items": [
					{
						"apiVersion": "v1",
						"kind": "CloudSpace",
						"metadata": {
							"name": "test-cloudspace",
							"namespace": "test-namespace",
							"creationTimestamp": "2023-01-01T00:00:00Z"
						},
						"spec": {
							"region": "uk-lon-1",
							"cloud": "default",
							"kubernetesVersion": "1.25.0",
							"HAControlPlane": true,
							"cni": "calico",
							"deploymentType": "standard"
						},
						"status": {
							"phase": "Running",
							"health": "Healthy",
							"currentKubernetesVersion": "1.25.0"
						}
					}
				],
				"metadata": {
					"resourceVersion": "12345"
				}
			}`,
			mockStatus:        200,
			expectError:       false,
			expectedCount:     1,
			expectedFirstName: "test-cloudspace",
		},
		{
			name:      "empty cloudspaces list",
			namespace: "empty-namespace",
			mockResponse: `{
				"apiVersion": "v1",
				"kind": "CloudSpaceList",
				"items": [],
				"metadata": {
					"resourceVersion": "12345"
				}
			}`,
			mockStatus:    200,
			expectError:   false,
			expectedCount: 0,
		},
		{
			name:             "missing namespace",
			namespace:        "",
			expectError:      true,
			expectedErrorMsg: "namespace is required",
		},
		{
			name:             "404 error",
			namespace:        "test-namespace",
			mockResponse:     `{"error": "namespace not found"}`,
			mockStatus:       404,
			expectError:      true,
			expectedErrorMsg: "API error 404",
		},
		{
			name:             "unauthorized error",
			namespace:        "test-namespace",
			mockResponse:     `{"error": "unauthorized"}`,
			mockStatus:       401,
			expectError:      true,
			expectedErrorMsg: "API error 401",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip server setup for namespace validation test
			if tt.namespace == "" {
				client := &Client{}
				_, err := client.ListCloudSpaces(context.Background(), tt.namespace)
				if !tt.expectError {
					t.Errorf("expected error but got none")
					return
				}
				if err == nil || err.Error() != tt.expectedErrorMsg {
					t.Errorf("expected error message %q, got %q", tt.expectedErrorMsg, err.Error())
				}
				return
			}

			// Create mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/namespaces/" + tt.namespace + "/cloudspaces"
				if r.URL.Path != expectedPath {
					t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
				}

				if r.Method != "GET" {
					t.Errorf("expected GET method, got %s", r.Method)
				}

				w.WriteHeader(tt.mockStatus)
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			// Create client with mock server
			cfg := &config.Config{
				RefreshToken: "test-token",
				BaseURL:      server.URL,
				Debug:        false,
				Timeout:      30,
			}

			// Create client with mock token manager
			client := NewClient(cfg)
			client.tokenManager = &MockTokenManager{
				accessToken: "mock-access-token",
			}

			// Call the method
			cloudSpaceList, err := client.ListCloudSpaces(context.Background(), tt.namespace)

			// Check error expectation
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}
				if tt.expectedErrorMsg != "" && !strings.Contains(err.Error(), tt.expectedErrorMsg) {
					t.Errorf("expected error message to contain %q, got %q", tt.expectedErrorMsg, err.Error())
				}
				return
			}

			// Check no error expected
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Check result
			if cloudSpaceList == nil {
				t.Errorf("expected cloudSpaceList but got nil")
				return
			}

			if len(cloudSpaceList.Items) != tt.expectedCount {
				t.Errorf("expected %d cloudspaces, got %d", tt.expectedCount, len(cloudSpaceList.Items))
				return
			}

			if tt.expectedCount > 0 && cloudSpaceList.Items[0].Metadata.Name != tt.expectedFirstName {
				t.Errorf("expected first cloudspace name %q, got %q", tt.expectedFirstName, cloudSpaceList.Items[0].Metadata.Name)
			}

			// Additional checks for successful cases
			if tt.expectedCount > 0 {
				cloudspace := cloudSpaceList.Items[0]
				if cloudspace.Spec.Region != "uk-lon-1" {
					t.Errorf("expected region %q, got %q", "uk-lon-1", cloudspace.Spec.Region)
				}
				if cloudspace.Status.Phase != "Running" {
					t.Errorf("expected phase %q, got %q", "Running", cloudspace.Status.Phase)
				}
				if cloudspace.Status.Health != "Healthy" {
					t.Errorf("expected health %q, got %q", "Healthy", cloudspace.Status.Health)
				}
			}
		})
	}
}
