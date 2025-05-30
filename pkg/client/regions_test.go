package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/georgetaylor/rackspace-spot-cli/pkg/config"
)

// MockTokenManager implements a token manager for testing
type MockTokenManager struct {
	accessToken string
	shouldError bool
	errorMsg    string
}

func (m *MockTokenManager) GetValidAccessToken(ctx context.Context) (string, error) {
	if m.shouldError {
		return "", fmt.Errorf("%s", m.errorMsg)
	}
	return m.accessToken, nil
}

// NewTestClient creates a client with a mock token manager for testing
func NewTestClient(cfg *config.Config, tokenManager TokenManagerInterface) *Client {
	httpClient := &http.Client{
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}

	return &Client{
		httpClient:   httpClient,
		config:       cfg,
		tokenManager: tokenManager,
	}
}

func TestClient_ListRegions(t *testing.T) {
	// Create a mock response
	mockRegionList := &RegionList{
		APIVersion: "v1",
		Kind:       "RegionList",
		Items: []Region{
			{
				APIVersion: "v1",
				Kind:       "Region",
				Metadata: ObjectMeta{
					Name: "us-east-1",
				},
				Spec: RegionSpec{
					Country:     "United States",
					Description: "US East (Virginia)",
					Provider: RegionProvider{
						ProviderType:       "aws",
						ProviderRegionName: "us-east-1",
					},
				},
			},
			{
				APIVersion: "v1",
				Kind:       "Region",
				Metadata: ObjectMeta{
					Name: "eu-west-1",
				},
				Spec: RegionSpec{
					Country:     "Ireland",
					Description: "EU West (Ireland)",
					Provider: RegionProvider{
						ProviderType:       "aws",
						ProviderRegionName: "eu-west-1",
					},
				},
			},
		},
		Metadata: ListMeta{
			ResourceVersion: "1234",
		},
	}

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/regions" {
			t.Errorf("Expected /regions path, got %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") == "" {
			t.Errorf("Expected Authorization header")
		}

		// Return mock response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockRegionList)
	}))
	defer server.Close()

	// Create test config
	cfg := &config.Config{
		BaseURL:      server.URL,
		RefreshToken: "test-refresh-token",
		Timeout:      30,
		Debug:        true,
	}

	// Create client
	client := NewTestClient(cfg, &MockTokenManager{accessToken: "test-access-token"})

	// Test the ListRegions method
	ctx := context.Background()
	regionList, err := client.ListRegions(ctx)
	if err != nil {
		t.Fatalf("ListRegions failed: %v", err)
	}

	// Verify the response
	if regionList.APIVersion != "v1" {
		t.Errorf("Expected APIVersion v1, got %s", regionList.APIVersion)
	}
	if regionList.Kind != "RegionList" {
		t.Errorf("Expected Kind RegionList, got %s", regionList.Kind)
	}
	if len(regionList.Items) != 2 {
		t.Errorf("Expected 2 regions, got %d", len(regionList.Items))
	}

	// Verify first region
	firstRegion := regionList.Items[0]
	if firstRegion.Metadata.Name != "us-east-1" {
		t.Errorf("Expected first region name us-east-1, got %s", firstRegion.Metadata.Name)
	}
	if firstRegion.Spec.Country != "United States" {
		t.Errorf("Expected first region country 'United States', got %s", firstRegion.Spec.Country)
	}
	if firstRegion.Spec.Provider.ProviderType != "aws" {
		t.Errorf("Expected first region provider type 'aws', got %s", firstRegion.Spec.Provider.ProviderType)
	}
}

func TestClient_ListRegions_APIError(t *testing.T) {
	// Create a test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    401,
			"message": "Unauthorized",
			"details": "Invalid access token",
		})
	}))
	defer server.Close()

	// Create test config
	cfg := &config.Config{
		BaseURL:      server.URL,
		RefreshToken: "invalid-refresh-token",
		Timeout:      30,
		Debug:        true,
	}

	// Create client
	client := NewTestClient(cfg, &MockTokenManager{accessToken: "test-access-token"})

	// Test the ListRegions method with error
	ctx := context.Background()
	regionList, err := client.ListRegions(ctx)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if regionList != nil {
		t.Error("Expected nil region list on error")
	}

	// Verify it's an API error
	if apiErr, ok := err.(*APIError); ok {
		if apiErr.Code != 401 {
			t.Errorf("Expected error code 401, got %d", apiErr.Code)
		}
		if apiErr.Message != "Unauthorized" {
			t.Errorf("Expected error message 'Unauthorized', got %s", apiErr.Message)
		}
	} else {
		t.Errorf("Expected APIError, got %T", err)
	}
}
