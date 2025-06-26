package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/georgetaylor/spotctl/pkg/config"
)

func TestListSpotNodePools(t *testing.T) {
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
			name:      "successful spotnodepools list",
			namespace: "test-namespace",
			mockResponse: `{
				"apiVersion": "v1",
				"kind": "SpotNodePoolList",
				"items": [
					{
						"apiVersion": "v1",
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
							"desired": 3,
							"serverClass": "gp.vs1.large-lon"
						},
						"status": {
							"bidStatus": "Winning",
							"wonCount": 2
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
			expectedFirstName: "test-spotnodepool",
		},
		{
			name:      "empty spotnodepools list",
			namespace: "empty-namespace",
			mockResponse: `{
				"apiVersion": "v1",
				"kind": "SpotNodePoolList",
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
				_, err := client.ListSpotNodePools(context.Background(), tt.namespace)
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
				expectedPath := "/ngpc.rxt.io/v1/namespaces/" + tt.namespace + "/spotnodepools"
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
			spotNodePoolList, err := client.ListSpotNodePools(context.Background(), tt.namespace)

			// Check error expectation
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

			// Check success cases
			if spotNodePoolList == nil {
				t.Errorf("expected spotnodepool list but got nil")
				return
			}

			if len(spotNodePoolList.Items) != tt.expectedCount {
				t.Errorf("expected %d items, got %d", tt.expectedCount, len(spotNodePoolList.Items))
				return
			}

			if tt.expectedCount > 0 && spotNodePoolList.Items[0].Metadata.Name != tt.expectedFirstName {
				t.Errorf("expected first item name %q, got %q", tt.expectedFirstName, spotNodePoolList.Items[0].Metadata.Name)
			}
		})
	}
}

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
					"serverClass": "gp.vs1.large-lon"
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

			if spotNodePool.Spec.ServerClass != "gp.vs1.large-lon" {
				t.Errorf("expected server class %q, got %q", "gp.vs1.large-lon", spotNodePool.Spec.ServerClass)
			}

			if spotNodePool.Status.BidStatus != "Winning" {
				t.Errorf("expected bid status %q, got %q", "Winning", spotNodePool.Status.BidStatus)
			}
		})
	}
}

func TestCreateSpotNodePool(t *testing.T) {
	tests := []struct {
		name             string
		namespace        string
		spotNodePool     *SpotNodePool
		mockResponse     string
		mockStatus       int
		expectError      bool
		expectedErrorMsg string
	}{
		{
			name:      "successful spotnodepool create",
			namespace: "test-namespace",
			spotNodePool: &SpotNodePool{
				APIVersion: "ngpc.rxt.io/v1",
				Kind:       "SpotNodePool",
				Metadata: ObjectMeta{
					Name:      "test-spotnodepool",
					Namespace: "test-namespace",
				},
				Spec: SpotNodePoolSpec{
					ServerClass: "gp.vs1.large-lon",
					CloudSpace:  "test-cloudspace",
					Desired:     &[]int{3}[0],
					BidPrice:    "0.50",
				},
			},
			mockResponse: `{
				"apiVersion": "ngpc.rxt.io/v1",
				"kind": "SpotNodePool",
				"metadata": {
					"name": "test-spotnodepool",
					"namespace": "test-namespace",
					"creationTimestamp": "2024-01-01T00:00:00Z"
				},
				"spec": {
					"serverClass": "gp.vs1.large-lon",
					"cloudSpace": "test-cloudspace",
					"desired": 3,
					"bidPrice": "0.50"
				},
				"status": {
					"bidStatus": "Pending"
				}
			}`,
			mockStatus:  201,
			expectError: false,
		},
		{
			name:      "empty namespace error",
			namespace: "",
			spotNodePool: &SpotNodePool{
				APIVersion: "ngpc.rxt.io/v1",
				Kind:       "SpotNodePool",
				Metadata: ObjectMeta{
					Name: "test-spotnodepool",
				},
			},
			expectError:      true,
			expectedErrorMsg: "namespace is required",
		},
		{
			name:             "nil spotnodepool error",
			namespace:        "test-namespace",
			spotNodePool:     nil,
			expectError:      true,
			expectedErrorMsg: "spot node pool configuration is required",
		},
		{
			name:      "API error response",
			namespace: "test-namespace",
			spotNodePool: &SpotNodePool{
				APIVersion: "ngpc.rxt.io/v1",
				Kind:       "SpotNodePool",
				Metadata: ObjectMeta{
					Name:      "test-spotnodepool",
					Namespace: "test-namespace",
				},
			},
			mockResponse: `{"code": 400, "message": "Invalid request"}`,
			mockStatus:   400,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip creating server for error test cases that don't need HTTP
			if tt.expectedErrorMsg != "" && (tt.namespace == "" || tt.spotNodePool == nil) {
				cfg := &config.Config{
					BaseURL:      "https://spot.rackspace.com",
					RefreshToken: "test-token",
					Timeout:      30,
				}

				// Use mock token manager for non-HTTP tests too
				mockTokenManager := &MockTokenManager{
					accessToken: "test-access-token",
					shouldError: false,
				}

				client := NewTestClient(cfg, mockTokenManager)
				_, err := client.CreateSpotNodePool(context.Background(), tt.namespace, tt.spotNodePool)

				if !tt.expectError {
					t.Errorf("expected no error but got: %v", err)
					return
				}

				if err == nil {
					t.Errorf("expected error but got none")
					return
				}

				if !strings.Contains(err.Error(), tt.expectedErrorMsg) {
					t.Errorf("expected error message to contain %q, got %q", tt.expectedErrorMsg, err.Error())
				}
				return
			}

			// Create mock server for HTTP test cases
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Handle authentication endpoint
				if r.URL.Path == "/auth/token" {
					w.WriteHeader(200)
					w.Write([]byte(`{"access_token": "test-token", "expires_in": 3600}`))
					return
				}

				// Verify the request
				expectedPath := "/ngpc.rxt.io/v1/namespaces/" + tt.namespace + "/spotnodepools"
				if r.URL.Path != expectedPath {
					t.Errorf("expected path %q, got %q", expectedPath, r.URL.Path)
				}

				if r.Method != "POST" {
					t.Errorf("expected POST method, got %q", r.Method)
				}

				w.WriteHeader(tt.mockStatus)
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			cfg := &config.Config{
				BaseURL:      server.URL,
				RefreshToken: "test-token",
				Timeout:      30,
			}

			// Use the mock token manager
			mockTokenManager := &MockTokenManager{
				accessToken: "test-access-token",
				shouldError: false,
			}

			client := NewTestClient(cfg, mockTokenManager)
			createdSpotNodePool, err := client.CreateSpotNodePool(context.Background(), tt.namespace, tt.spotNodePool)

			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
				return
			}

			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
				return
			}

			if !tt.expectError {
				if createdSpotNodePool == nil {
					t.Errorf("expected spotnodepool but got nil")
					return
				}

				if createdSpotNodePool.Metadata.Name != "test-spotnodepool" {
					t.Errorf("expected name %q, got %q", "test-spotnodepool", createdSpotNodePool.Metadata.Name)
				}

				if createdSpotNodePool.Metadata.Namespace != "test-namespace" {
					t.Errorf("expected namespace %q, got %q", "test-namespace", createdSpotNodePool.Metadata.Namespace)
				}

				if createdSpotNodePool.Spec.ServerClass != "gp.vs1.large-lon" {
					t.Errorf("expected server class %q, got %q", "gp.vs1.large-lon", createdSpotNodePool.Spec.ServerClass)
				}

				if createdSpotNodePool.Status.BidStatus != "Pending" {
					t.Errorf("expected bid status %q, got %q", "Pending", createdSpotNodePool.Status.BidStatus)
				}
			}
		})
	}
}

func TestDeleteSpotNodePool(t *testing.T) {
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
			name:             "successful delete spotnodepool",
			namespace:        "test-namespace",
			spotNodePoolName: "test-spotnodepool",
			mockResponse: `{
				"apiVersion": "v1",
				"kind": "Status",
				"status": "Success",
				"message": "Spot node pool deleted successfully"
			}`,
			mockStatus:  200,
			expectError: false,
		},
		{
			name:             "missing namespace",
			namespace:        "",
			spotNodePoolName: "test-spotnodepool",
			expectError:      true,
			expectedErrorMsg: "namespace is required",
		},
		{
			name:             "missing spotnodepool name",
			namespace:        "test-namespace",
			spotNodePoolName: "",
			expectError:      true,
			expectedErrorMsg: "spot node pool name is required",
		},
		{
			name:             "404 error",
			namespace:        "test-namespace",
			spotNodePoolName: "test-spotnodepool",
			mockResponse:     `{"error": "spot node pool not found"}`,
			mockStatus:       404,
			expectError:      true,
			expectedErrorMsg: "API error 404",
		},
		{
			name:             "unauthorized error",
			namespace:        "test-namespace",
			spotNodePoolName: "test-spotnodepool",
			mockResponse:     `{"error": "unauthorized"}`,
			mockStatus:       401,
			expectError:      true,
			expectedErrorMsg: "API error 401",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip server setup for validation tests
			if tt.namespace == "" || tt.spotNodePoolName == "" {
				client := &Client{}
				_, err := client.DeleteSpotNodePool(context.Background(), tt.namespace, tt.spotNodePoolName)
				if !tt.expectError {
					t.Errorf("expected error but got none")
					return
				}
				if err == nil || !strings.Contains(err.Error(), tt.expectedErrorMsg) {
					t.Errorf("expected error message to contain %q, got %q", tt.expectedErrorMsg, err.Error())
				}
				return
			}

			// Create mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/ngpc.rxt.io/v1/namespaces/" + tt.namespace + "/spotnodepools/" + tt.spotNodePoolName
				if r.URL.Path != expectedPath {
					t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
				}

				if r.Method != "DELETE" {
					t.Errorf("expected DELETE method, got %s", r.Method)
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
			deleteResponse, err := client.DeleteSpotNodePool(context.Background(), tt.namespace, tt.spotNodePoolName)

			// Check error expectation
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

			// Check success cases
			if deleteResponse == nil {
				t.Errorf("expected delete response but got nil")
				return
			}

			if deleteResponse.Status != "Success" {
				t.Errorf("expected status Success, got %s", deleteResponse.Status)
			}
		})
	}
}

func TestGetOnDemandNodePool(t *testing.T) {
	tests := []struct {
		name                 string
		namespace            string
		onDemandNodePoolName string
		mockResponse         string
		mockStatus           int
		expectError          bool
		expectedErrorMsg     string
	}{
		{
			name:                 "successful ondemandnodepool get",
			namespace:            "test-namespace",
			onDemandNodePoolName: "test-ondemandnodepool",
			mockResponse: `{
				"apiVersion": "v1",
				"kind": "OnDemandNodePool", 
				"metadata": {
					"name": "test-ondemandnodepool",
					"namespace": "test-namespace"
				},
				"spec": {
					"serverClass": "gp.vs1.large-lon",
					"desired": 3,
					"cloudSpace": "test-cloudspace"
				},
				"status": {
					"reservedCount": 3,
					"reservedStatus": "Active"
				}
			}`,
			mockStatus:  200,
			expectError: false,
		},
		{
			name:                 "ondemandnodepool not found",
			namespace:            "test-namespace",
			onDemandNodePoolName: "nonexistent-ondemandnodepool",
			mockResponse:         `{"error": "ondemandnodepool not found"}`,
			mockStatus:           404,
			expectError:          true,
			expectedErrorMsg:     "API error 404",
		},
		{
			name:                 "missing namespace",
			namespace:            "",
			onDemandNodePoolName: "test-ondemandnodepool",
			expectError:          true,
			expectedErrorMsg:     "namespace is required",
		},
		{
			name:                 "missing ondemandnodepool name",
			namespace:            "test-namespace",
			onDemandNodePoolName: "",
			expectError:          true,
			expectedErrorMsg:     "on demand node pool name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip server setup for validation tests
			if tt.namespace == "" || tt.onDemandNodePoolName == "" {
				client := &Client{}
				_, err := client.GetOnDemandNodePool(context.Background(), tt.namespace, tt.onDemandNodePoolName)
				if !tt.expectError {
					t.Errorf("expected error but got none")
					return
				}
				if err == nil || !strings.Contains(err.Error(), tt.expectedErrorMsg) {
					t.Errorf("expected error message to contain %q, got %q", tt.expectedErrorMsg, err.Error())
				}
				return
			}

			// Create mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/ngpc.rxt.io/v1/namespaces/" + tt.namespace + "/ondemandnodepools/" + tt.onDemandNodePoolName
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
			onDemandNodePool, err := client.GetOnDemandNodePool(context.Background(), tt.namespace, tt.onDemandNodePoolName)

			// Check error expectation
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

			// Check success cases
			if onDemandNodePool == nil {
				t.Errorf("expected ondemandnodepool but got nil")
				return
			}

			if onDemandNodePool.Metadata.Name != "test-ondemandnodepool" {
				t.Errorf("expected name %q, got %q", "test-ondemandnodepool", onDemandNodePool.Metadata.Name)
			}

			if onDemandNodePool.Metadata.Namespace != "test-namespace" {
				t.Errorf("expected namespace %q, got %q", "test-namespace", onDemandNodePool.Metadata.Namespace)
			}

			if onDemandNodePool.Spec.ServerClass != "gp.vs1.large-lon" {
				t.Errorf("expected server class %q, got %q", "gp.vs1.large-lon", onDemandNodePool.Spec.ServerClass)
			}

			if onDemandNodePool.Status.ReservedStatus != "Active" {
				t.Errorf("expected reserved status %q, got %q", "Active", onDemandNodePool.Status.ReservedStatus)
			}
		})
	}
}

func TestDeleteAllSpotNodePools(t *testing.T) {
	tests := []struct {
		name             string
		namespace        string
		mockResponse     string
		mockStatus       int
		expectError      bool
		expectedErrorMsg string
	}{
		{
			name:      "successful delete all spotnodepools",
			namespace: "test-namespace",
			mockResponse: `{
				"apiVersion": "v1",
				"kind": "Status",
				"status": "Success",
				"message": "All spot node pools deleted successfully"
			}`,
			mockStatus:  200,
			expectError: false,
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
				_, err := client.DeleteAllSpotNodePools(context.Background(), tt.namespace)
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
				expectedPath := "/ngpc.rxt.io/v1/namespaces/" + tt.namespace + "/spotnodepools"
				if r.URL.Path != expectedPath {
					t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
				}

				if r.Method != "DELETE" {
					t.Errorf("expected DELETE method, got %s", r.Method)
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
			deleteResponse, err := client.DeleteAllSpotNodePools(context.Background(), tt.namespace)

			// Check error expectation
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

			// Check success cases
			if deleteResponse == nil {
				t.Errorf("expected delete response but got nil")
				return
			}

			if deleteResponse.Status != "Success" {
				t.Errorf("expected status Success, got %s", deleteResponse.Status)
			}
		})
	}
}
