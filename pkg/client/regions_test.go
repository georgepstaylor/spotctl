package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"errors"

	"github.com/georgetaylor/spotctl/pkg/config"
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
					Name: "uk-lon-1",
				},
				Spec: RegionSpec{
					Country:     "United Kingdom",
					Description: "London",
					Provider: RegionProvider{
						ProviderType:       "ospc",
						ProviderRegionName: "uk-lon-1",
					},
				},
			},
			{
				APIVersion: "v1",
				Kind:       "Region",
				Metadata: ObjectMeta{
					Name: "us-central-dfw-2",
				},
				Spec: RegionSpec{
					Country:     "United States",
					Description: "Dallas Fort Worth",
					Provider: RegionProvider{
						ProviderType:       "ospc",
						ProviderRegionName: "us-central-dfw-2",
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
		if r.URL.Path != "/apis/ngpc.rxt.io/v1/regions" {
			t.Errorf("Expected /apis/ngpc.rxt.io/v1/regions path, got %s", r.URL.Path)
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
		RefreshToken: "test-token",
		BaseURL:      server.URL + "/apis/ngpc.rxt.io/v1",
		Region:       "uk-lon-1",
		Debug:        false,
		Timeout:      30,
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
	if firstRegion.Metadata.Name != "uk-lon-1" {
		t.Errorf("Expected first region name uk-lon-1, got %s", firstRegion.Metadata.Name)
	}
	if firstRegion.Spec.Country != "United Kingdom" {
		t.Errorf("Expected first region country 'United Kingdom', got %s", firstRegion.Spec.Country)
	}
	if firstRegion.Spec.Provider.ProviderType != "ospc" {
		t.Errorf("Expected first region provider type 'ospc', got %s", firstRegion.Spec.Provider.ProviderType)
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
		RefreshToken: "test-token",
		BaseURL:      server.URL + "/apis/ngpc.rxt.io/v1",
		Region:       "uk-lon-1",
		Debug:        false,
		Timeout:      30,
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
	var apiErr *APIError
	if errors.As(err, &apiErr) {
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
