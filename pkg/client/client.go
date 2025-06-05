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

// APIVersion represents the different API versions available
type APIVersion string

// API version constants
const (
	// APIVersionDefault is the default API version for most operations
	APIVersionDefault APIVersion = "ngpc.rxt.io"
	// APIVersionAuth is the API version for authentication-related operations
	APIVersionAuth APIVersion = "auth.ngpc.rxt.io"
	// Future API versions can be added here, for example:
	// APIVersionV2     APIVersion = "v2.ngpc.rxt.io"
	// APIVersionBeta   APIVersion = "beta.ngpc.rxt.io"
)

// String returns the string representation of the API version
func (v APIVersion) String() string {
	return string(v)
}

// NewAPIVersion creates a new APIVersion from a string
// This can be used for custom or future API versions
func NewAPIVersion(version string) APIVersion {
	return APIVersion(version)
}

// IsValid checks if the API version is one of the known versions
func (v APIVersion) IsValid() bool {
	switch v {
	case APIVersionDefault, APIVersionAuth:
		return true
	default:
		return false
	}
}

// GetAllAPIVersions returns a slice of all known API versions
func GetAllAPIVersions() []APIVersion {
	return []APIVersion{
		APIVersionDefault,
		APIVersionAuth,
	}
}

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

// MakeRequest performs an HTTP request to the API with dynamic apiVersion support
// This is the primary method for making requests with full control over method, endpoint, body, and API version
// The contentType parameter is optional - if not provided, defaults to "application/json"
//
// For convenience, use the wrapper methods:
// - Get, Post, Put, Delete (for default API version)
// - GetAuth, PostAuth, PutAuth, DeleteAuth (for auth API version)
//
// Example usage:
//
//	resp, err := client.MakeRequest(ctx, http.MethodGet, "/custom-endpoint", nil, APIVersionAuth)
//	resp, err := client.MakeRequest(ctx, http.MethodPatch, "/endpoint", data, APIVersionDefault, "application/json-patch+json")
func (c *Client) MakeRequest(ctx context.Context, method, endpoint string, body interface{}, apiVersion APIVersion, contentType ...string) (*http.Response, error) {
	// Build the URL based on the apiVersion
	baseURL := c.config.BaseURL
	if apiVersion != APIVersionDefault {
		// Replace the default API version with the specified one
		baseURL = strings.Replace(baseURL, "/apis/ngpc.rxt.io/", fmt.Sprintf("/apis/%s/", apiVersion.String()), 1)
	}

	url := fmt.Sprintf("%s%s", baseURL, endpoint)

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

	// Set content type - default to application/json if not specified
	ct := "application/json"
	if len(contentType) > 0 && contentType[0] != "" {
		ct = contentType[0]
	}

	// Set headers
	req.Header.Set("Content-Type", ct)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("User-Agent", "spotctl/0.0.1")

	if c.config.Debug {
		if ct != "application/json" {
			fmt.Printf("Making %s request to %s (API: %s, Content-Type: %s)\n", method, url, apiVersion, ct)
		} else {
			fmt.Printf("Making %s request to %s (API: %s)\n", method, url, apiVersion)
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, endpoint string) (*http.Response, error) {
	return c.MakeRequest(ctx, http.MethodGet, endpoint, nil, APIVersionDefault)
}

// GetAuth performs a GET request to the auth API (different subdomain)
func (c *Client) GetAuth(ctx context.Context, endpoint string) (*http.Response, error) {
	return c.MakeRequest(ctx, http.MethodGet, endpoint, nil, APIVersionAuth)
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return c.MakeRequest(ctx, http.MethodPost, endpoint, body, APIVersionDefault)
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return c.MakeRequest(ctx, http.MethodPut, endpoint, body, APIVersionDefault)
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, endpoint string) (*http.Response, error) {
	return c.MakeRequest(ctx, http.MethodDelete, endpoint, nil, APIVersionDefault)
}

// Patch performs a PATCH request
func (c *Client) Patch(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return c.MakeRequest(ctx, http.MethodPatch, endpoint, body, APIVersionDefault)
}

// PatchWithContentType performs a PATCH request with a specific content type
func (c *Client) PatchWithContentType(ctx context.Context, endpoint string, body interface{}, contentType string) (*http.Response, error) {
	return c.MakeRequest(ctx, http.MethodPatch, endpoint, body, APIVersionDefault, contentType)
}

// PostAuth performs a POST request to the auth API
func (c *Client) PostAuth(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return c.MakeRequest(ctx, http.MethodPost, endpoint, body, APIVersionAuth)
}

// PutAuth performs a PUT request to the auth API
func (c *Client) PutAuth(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return c.MakeRequest(ctx, http.MethodPut, endpoint, body, APIVersionAuth)
}

// DeleteAuth performs a DELETE request to the auth API
func (c *Client) DeleteAuth(ctx context.Context, endpoint string) (*http.Response, error) {
	return c.MakeRequest(ctx, http.MethodDelete, endpoint, nil, APIVersionAuth)
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

// EditCloudSpace applies JSON patch operations to update a cloudspace
func (c *Client) EditCloudSpace(ctx context.Context, namespace, name string, patchOps interface{}) (*CloudSpace, error) {
	if namespace == "" {
		return nil, fmt.Errorf("namespace is required")
	}
	if name == "" {
		return nil, fmt.Errorf("cloudspace name is required")
	}

	endpoint := fmt.Sprintf("/namespaces/%s/cloudspaces/%s", namespace, name)
	resp, err := c.PatchWithContentType(ctx, endpoint, patchOps, "application/json-patch+json")
	if err != nil {
		return nil, fmt.Errorf("failed to make patch request: %w", err)
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
