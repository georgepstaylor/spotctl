package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/georgetaylor/spotctl/pkg/config"
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
	req.Header.Set("User-Agent", "spotctl/0.0.1")

	if c.config.Debug {
		fmt.Printf("Making %s request to %s\n", method, url)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

// makeAuthRequest is similar to makeRequest but uses the auth API subdomain
func (c *Client) makeAuthRequest(ctx context.Context, method, endpoint string, body interface{}) (*http.Response, error) {
	// Convert the base URL to use auth subdomain
	// Change from "https://spot.rackspace.com/apis/ngpc.rxt.io/v1" to "https://spot.rackspace.com/apis/auth.ngpc.rxt.io/v1"
	authBaseURL := fmt.Sprintf("%s", c.config.BaseURL)
	if strings.Contains(authBaseURL, "/apis/ngpc.rxt.io/") {
		authBaseURL = strings.Replace(authBaseURL, "/apis/ngpc.rxt.io/", "/apis/auth.ngpc.rxt.io/", 1)
	}

	url := fmt.Sprintf("%s%s", authBaseURL, endpoint)

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
	req.Header.Set("User-Agent", "spotctl/0.0.1")

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

// GetAuth performs a GET request to the auth API (different subdomain)
func (c *Client) GetAuth(ctx context.Context, endpoint string) (*http.Response, error) {
	return c.makeAuthRequest(ctx, http.MethodGet, endpoint, nil)
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

// ListServerClasses retrieves all available server classes
func (c *Client) ListServerClasses(ctx context.Context) (*ServerClassList, error) {
	resp, err := c.Get(ctx, "/serverclasses")
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var serverClassList ServerClassList
	if err := json.NewDecoder(resp.Body).Decode(&serverClassList); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &serverClassList, nil
}

// GetServerClass retrieves a specific server class by name
func (c *Client) GetServerClass(ctx context.Context, name string) (*ServerClass, error) {
	if name == "" {
		return nil, fmt.Errorf("server class name is required")
	}

	resp, err := c.Get(ctx, fmt.Sprintf("/serverclasses/%s", name))
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var serverClass ServerClass
	if err := json.NewDecoder(resp.Body).Decode(&serverClass); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &serverClass, nil
}

// ListOrganizations retrieves all organizations for the authenticated user
func (c *Client) ListOrganizations(ctx context.Context) (*OrganizationList, error) {
	// Organizations API uses a different subdomain
	resp, err := c.GetAuth(ctx, "/organizations")
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var orgList OrganizationList
	if err := json.NewDecoder(resp.Body).Decode(&orgList); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &orgList, nil
}

// ListCloudSpaces retrieves all cloudspaces for a given namespace
func (c *Client) ListCloudSpaces(ctx context.Context, namespace string) (*CloudSpaceList, error) {
	if namespace == "" {
		return nil, fmt.Errorf("namespace is required")
	}

	endpoint := fmt.Sprintf("/namespaces/%s/cloudspaces", namespace)
	resp, err := c.Get(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var cloudSpaceList CloudSpaceList
	if err := json.NewDecoder(resp.Body).Decode(&cloudSpaceList); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &cloudSpaceList, nil
}

// CreateCloudSpace creates a new cloudspace in the specified namespace
func (c *Client) CreateCloudSpace(ctx context.Context, namespace string, cloudSpace *CloudSpace) (*CloudSpace, error) {
	if namespace == "" {
		return nil, fmt.Errorf("namespace is required")
	}
	if cloudSpace == nil {
		return nil, fmt.Errorf("cloudspace configuration is required")
	}

	endpoint := fmt.Sprintf("/namespaces/%s/cloudspaces", namespace)
	resp, err := c.Post(ctx, endpoint, cloudSpace)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var createdCloudSpace CloudSpace
	if err := json.NewDecoder(resp.Body).Decode(&createdCloudSpace); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &createdCloudSpace, nil
}

// DeleteCloudSpace deletes a cloudspace by name in the specified namespace
func (c *Client) DeleteCloudSpace(ctx context.Context, namespace, name string) (*DeleteResponse, error) {
	if namespace == "" {
		return nil, fmt.Errorf("namespace is required")
	}
	if name == "" {
		return nil, fmt.Errorf("cloudspace name is required")
	}

	endpoint := fmt.Sprintf("/namespaces/%s/cloudspaces/%s", namespace, name)
	resp, err := c.Delete(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	// For successful deletes, HTTP 200 or 202 means success according to API docs
	// We'll return a simple success response instead of trying to parse the complex Kubernetes Status object
	if resp.StatusCode == 200 || resp.StatusCode == 202 {
		return &DeleteResponse{
			Status:  "Success",
			Message: "CloudSpace deleted successfully",
		}, nil
	}

	// If we get here, try to parse the actual response for error details
	var deleteResponse DeleteResponse
	if err := json.NewDecoder(resp.Body).Decode(&deleteResponse); err != nil {
		// If we can't parse the response but got a successful status code, assume success
		return &DeleteResponse{
			Status:  "Success",
			Message: "CloudSpace deleted successfully",
		}, nil
	}

	return &deleteResponse, nil
}

// GetCloudSpace retrieves a specific cloudspace by name in the specified namespace
func (c *Client) GetCloudSpace(ctx context.Context, namespace, name string) (*CloudSpace, error) {
	if namespace == "" {
		return nil, fmt.Errorf("namespace is required")
	}
	if name == "" {
		return nil, fmt.Errorf("cloudspace name is required")
	}

	endpoint := fmt.Sprintf("/namespaces/%s/cloudspaces/%s", namespace, name)
	resp, err := c.Get(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var cloudSpace CloudSpace
	if err := json.NewDecoder(resp.Body).Decode(&cloudSpace); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &cloudSpace, nil
}
