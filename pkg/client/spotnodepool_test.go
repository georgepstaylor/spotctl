package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/config"
)

func TestGetSpotNodePool(t *testing.T) {
	tests := []struct {
		name             string
		namespace        string
		spotNodePoolName string
		mockResponse     string
		mockStatus       int
		expectError      bool
		expectedErrorMsg string
	}{
		{
			name:             "successful spotnodepool get",
			namespace:        "test-namespace",
			spotNodePoolName: "test-spotnodepool",
			mockResponse: `{
				"apiVersion": "ngpc.rxt.io/v1",
				"kind": "SpotNodePool",
				"metadata": {
					"name": "test-spotnodepool",
					"namespace": "test-namespace",
					"creationTimestamp": "2023-01-01T00:00:00Z"
				},
				"spec": {
					"autoscaling": {
						"enabled": true,
						"maxNodes": 10,
						"minNodes": 1
					},
					"bidPrice": "0.50",
					"cloudSpace": "test-cloudspace",
					"customAnnotations": {},
					"customLabels": {},
					"customTaints": [],
					"desired": 3,
					"serverClass": "gp.4x8"
				},
				"status": {
					"bidStatus": "Winning",
					"customMetadataStatus": {
						"annotations": [],
						"labels": [],
						"taints": []
					},
					"wonCount": 2
				}
			}`,
			mockStatus:  200,
			expectError: false,
		},
		{
			name:             "spotnodepool not found",
			namespace:        "test-namespace",
			spotNodePoolName: "missing-spotnodepool",
			mockResponse:     `{"code": 404, "message": "SpotNodePool not found"}`,
			mockStatus:       404,
			expectError:      true,
			expectedErrorMsg: "API error 404: SpotNodePool not found",
		},
		{
			name:             "missing namespace",
			namespace:        "",
			spotNodePoolName: "test-spotnodepool",
			mockResponse:     "",
			mockStatus:       0,
			expectError:      true,
			expectedErrorMsg: "namespace is required",
		},
		{
			name:             "missing spotnodepool name",
			namespace:        "test-namespace",
			spotNodePoolName: "",
			mockResponse:     "",
			mockStatus:       0,
			expectError:      true,
			expectedErrorMsg: "spot node pool name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip server setup for validation tests
			if tt.namespace == "" || tt.spotNodePoolName == "" {
				client := &Client{}
				_, err := client.GetSpotNodePool(context.Background(), tt.namespace, tt.spotNodePoolName)
				if !tt.expectError {
					t.Errorf("expected error but got none")
					return
				}
				if err == nil || !strings.Contains(err.Error(), tt.expectedErrorMsg) {
					t.Errorf("expected error message to contain %q, got %q", tt.expectedErrorMsg, err.Error())
				}
				return
			}

			// Setup mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/ngpc.rxt.io/v1/namespaces/" + tt.namespace + "/spotnodepools/" + tt.spotNodePoolName
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

			// Create test client with mock token manager
			cfg := &config.Config{
				RefreshToken: "test-token",
				BaseURL:      server.URL,
				Debug:        false,
				Timeout:      30,
			}

			mockTokenManager := &MockTokenManager{
				accessToken: "mock-access-token",
			}
			client := NewTestClient(cfg, mockTokenManager)

			// Test the method
			ctx := context.Background()
			spotNodePool, err := client.GetSpotNodePool(ctx, tt.namespace, tt.spotNodePoolName)

			// Check error expectations
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
				return
			}

			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
				return
			}

			if tt.expectError {
				if !strings.Contains(err.Error(), tt.expectedErrorMsg) {
					t.Errorf("expected error message to contain %q, got %q", tt.expectedErrorMsg, err.Error())
				}
				return
			}

			// Check response for successful cases
			if spotNodePool == nil {
				t.Errorf("expected spotnodepool but got nil")
				return
			}

			if spotNodePool.Metadata.Name != "test-spotnodepool" {
				t.Errorf("expected name %q, got %q", "test-spotnodepool", spotNodePool.Metadata.Name)
			}

			if spotNodePool.Metadata.Namespace != "test-namespace" {
				t.Errorf("expected namespace %q, got %q", "test-namespace", spotNodePool.Metadata.Namespace)
			}

			if spotNodePool.Spec.ServerClass != "gp.4x8" {
				t.Errorf("expected server class %q, got %q", "gp.4x8", spotNodePool.Spec.ServerClass)
			}

			if spotNodePool.Status.BidStatus != "Winning" {
				t.Errorf("expected bid status %q, got %q", "Winning", spotNodePool.Status.BidStatus)
			}
		})
	}
}
