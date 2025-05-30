package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/georgetaylor/rackspace-spot-cli/pkg/config"
)

// Client represents the Rackspace Spot API client
type Client struct {
	httpClient   *http.Client
	config       *config.Config
	tokenManager TokenManagerInterface
}

// APIError represents an error response from the API
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("API error %d: %s (%s)", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("API error %d: %s", e.Code, e.Message)
}

// NewClient creates a new Rackspace Spot API client
func NewClient(cfg *config.Config) *Client {
	httpClient := &http.Client{
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}

	tokenManager := NewTokenManager(cfg.RefreshToken, httpClient, cfg.Debug)

	return &Client{
		httpClient:   httpClient,
		config:       cfg,
		tokenManager: tokenManager,
	}
}

// makeRequest performs an HTTP request to the API
func (c *Client) makeRequest(ctx context.Context, method, endpoint string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.config.BaseURL, endpoint)

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Get valid access token
	accessToken, err := c.tokenManager.GetValidAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("User-Agent", "rackspace-spot-cli/1.0.0")

	if c.config.Debug {
		fmt.Printf("Making %s request to %s\n", method, url)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, endpoint string) (*http.Response, error) {
	return c.makeRequest(ctx, http.MethodGet, endpoint, nil)
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return c.makeRequest(ctx, http.MethodPost, endpoint, body)
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return c.makeRequest(ctx, http.MethodPut, endpoint, body)
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, endpoint string) (*http.Response, error) {
	return c.makeRequest(ctx, http.MethodDelete, endpoint, nil)
}

// HandleAPIError checks if the response indicates an API error and returns an APIError
func (c *Client) HandleAPIError(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HTTP %d: failed to read error response", resp.StatusCode)
	}

	var apiErr APIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	if apiErr.Code == 0 {
		apiErr.Code = resp.StatusCode
	}

	return &apiErr
}

// ListRegions retrieves all available regions
func (c *Client) ListRegions(ctx context.Context) (*RegionList, error) {
	resp, err := c.Get(ctx, "/regions")
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var regionList RegionList
	if err := json.NewDecoder(resp.Body).Decode(&regionList); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &regionList, nil
}

// GetRegion retrieves a specific region by name
func (c *Client) GetRegion(ctx context.Context, name string) (*Region, error) {
	if name == "" {
		return nil, fmt.Errorf("region name is required")
	}

	resp, err := c.Get(ctx, fmt.Sprintf("/regions/%s", name))
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var region Region
	if err := json.NewDecoder(resp.Body).Decode(&region); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &region, nil
}
