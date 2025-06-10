package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/georgetaylor/spotctl/pkg/errors"
)

// APIVersion represents the different API versions available
type APIVersion string

// API version constants
const (
	// APIVersionDefault is the default API version for most operations
	APIVersionDefault APIVersion = "ngpc.rxt.io/v1"
	// APIVersionAuth is the API version for authentication-related operations
	APIVersionAuth APIVersion = "auth.ngpc.rxt.io/v1"
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

// requestOptions contains options for making HTTP requests
type requestOptions struct {
	method      string
	endpoint    string
	body        interface{}
	apiVersion  APIVersion
	contentType string
}

// prepareRequest prepares an HTTP request with the given options
func (c *Client) prepareRequest(ctx context.Context, opts requestOptions) (*http.Request, error) {
	// Construct the full URL by combining base URL, API version, and endpoint
	url := fmt.Sprintf("%s/%s%s", c.config.BaseURL, opts.apiVersion.String(), opts.endpoint)

	var reqBody io.Reader
	if opts.body != nil {
		jsonBody, err := json.Marshal(opts.body)
		if err != nil {
			return nil, errors.NewInternalError(fmt.Sprintf("failed to marshal request body for %s %s", opts.method, url), err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, opts.method, url, reqBody)
	if err != nil {
		return nil, errors.NewInternalError(fmt.Sprintf("failed to create request for %s %s", opts.method, url), err)
	}

	accessToken, err := c.tokenManager.GetValidAccessToken(ctx)
	if err != nil {
		return nil, errors.NewAPIError(http.StatusUnauthorized, fmt.Sprintf("failed to get access token for %s %s", opts.method, url), err)
	}

	contentType := "application/json"
	if opts.contentType != "" {
		contentType = opts.contentType
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("User-Agent", "spotctl/0.0.1")

	if c.config.Debug {
		fmt.Printf("Making %s request to %s (API: %s)\n", opts.method, url, opts.apiVersion)
	}

	return req, nil
}

// doRequest executes an HTTP request and handles the response
func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewAPIError(0, fmt.Sprintf("request failed for %s %s", req.Method, req.URL.String()), err)
	}

	// For successful responses, return the response without closing the body
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return resp, nil
	}

	// For error responses, read and parse the error body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewAPIError(resp.StatusCode, fmt.Sprintf("failed to read error response for %s %s", req.Method, req.URL.String()), err)
	}

	var apiErr struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details,omitempty"`
	}

	if err := json.Unmarshal(body, &apiErr); err != nil {
		return nil, errors.NewAPIError(resp.StatusCode, fmt.Sprintf("invalid error response format for %s %s: %s", req.Method, req.URL.String(), string(body)), nil)
	}

	if apiErr.Code == 0 {
		apiErr.Code = resp.StatusCode
	}

	return nil, &APIError{
		Code:    apiErr.Code,
		Message: apiErr.Message,
		Details: apiErr.Details,
	}
}

// MakeRequest performs an HTTP request to the API
func (c *Client) MakeRequest(ctx context.Context, method, endpoint string, body interface{}, apiVersion APIVersion, contentType ...string) (*http.Response, error) {
	opts := requestOptions{
		method:      method,
		endpoint:    endpoint,
		body:        body,
		apiVersion:  apiVersion,
		contentType: "",
	}

	if len(contentType) > 0 {
		opts.contentType = contentType[0]
	}

	req, err := c.prepareRequest(ctx, opts)
	if err != nil {
		return nil, err
	}

	return c.doRequest(req)
}

// Convenience methods for common HTTP methods
func (c *Client) Get(ctx context.Context, endpoint string, apiVersion ...APIVersion) (*http.Response, error) {
	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	return c.MakeRequest(ctx, http.MethodGet, endpoint, nil, version)
}

func (c *Client) Post(ctx context.Context, endpoint string, body interface{}, apiVersion ...APIVersion) (*http.Response, error) {
	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	return c.MakeRequest(ctx, http.MethodPost, endpoint, body, version)
}

func (c *Client) Put(ctx context.Context, endpoint string, body interface{}, apiVersion ...APIVersion) (*http.Response, error) {
	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	return c.MakeRequest(ctx, http.MethodPut, endpoint, body, version)
}

func (c *Client) Delete(ctx context.Context, endpoint string, apiVersion ...APIVersion) (*http.Response, error) {
	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	return c.MakeRequest(ctx, http.MethodDelete, endpoint, nil, version)
}

func (c *Client) Patch(ctx context.Context, endpoint string, body interface{}, apiVersion ...APIVersion) (*http.Response, error) {
	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	return c.MakeRequest(ctx, http.MethodPatch, endpoint, body, version, "application/json")
}

// ListRegions retrieves a list of available regions
func (c *Client) ListRegions(ctx context.Context, apiVersion ...APIVersion) (*RegionList, error) {
	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	return genericList[RegionList](c, ctx, "/regions", ListOptions{APIVersion: version})
}

// GetRegion retrieves a specific region by name
func (c *Client) GetRegion(ctx context.Context, name string, apiVersion ...APIVersion) (*Region, error) {
	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("region name cannot be empty")
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	return genericGet[Region](c, ctx, fmt.Sprintf("/regions/%s", name), GetOptions{Name: name, APIVersion: version})
}

// ListServerClasses retrieves all available server classes
func (c *Client) ListServerClasses(ctx context.Context, apiVersion ...APIVersion) (*ServerClassList, error) {
	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	return genericList[ServerClassList](c, ctx, "/serverclasses", ListOptions{APIVersion: version})
}

// GetServerClass retrieves a specific server class by name
func (c *Client) GetServerClass(ctx context.Context, name string, apiVersion ...APIVersion) (*ServerClass, error) {
	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("server class name is required")
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	return genericGet[ServerClass](c, ctx, fmt.Sprintf("/serverclasses/%s", name), GetOptions{Name: name, APIVersion: version})
}

// ListOrganizations retrieves all organizations for the authenticated user
func (c *Client) ListOrganizations(ctx context.Context, apiVersion ...APIVersion) (*OrganizationList, error) {
	// Organizations API defaults to auth version but can be overridden
	version := APIVersionAuth
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	return genericList[OrganizationList](c, ctx, "/organizations", ListOptions{APIVersion: version})
}

// ListCloudSpaces retrieves all cloudspaces for a given namespace
func (c *Client) ListCloudSpaces(ctx context.Context, namespace string, apiVersion ...APIVersion) (*CloudSpaceList, error) {
	if err := validateNamespace(namespace); err != nil {
		return nil, err
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	endpoint := fmt.Sprintf("/namespaces/%s/cloudspaces", namespace)
	return genericList[CloudSpaceList](c, ctx, endpoint, ListOptions{Namespace: namespace, APIVersion: version})
}

// CreateCloudSpace creates a new cloudspace in the specified namespace
func (c *Client) CreateCloudSpace(ctx context.Context, namespace string, cloudSpace *CloudSpace, apiVersion ...APIVersion) (*CloudSpace, error) {
	if err := validateNamespace(namespace); err != nil {
		return nil, err
	}
	if err := validateCreateInput(cloudSpace); err != nil {
		return nil, fmt.Errorf("cloudspace configuration is required")
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	endpoint := fmt.Sprintf("/namespaces/%s/cloudspaces", namespace)
	return genericCreate[CloudSpace](c, ctx, endpoint, cloudSpace, CreateOptions{Namespace: namespace, APIVersion: version})
}

// DeleteCloudSpace deletes a cloudspace by name in the specified namespace
func (c *Client) DeleteCloudSpace(ctx context.Context, namespace, name string, apiVersion ...APIVersion) (*DeleteResponse, error) {
	if err := validateNamespace(namespace); err != nil {
		return nil, err
	}
	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("cloudspace name is required")
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	endpoint := fmt.Sprintf("/namespaces/%s/cloudspaces/%s", namespace, name)
	return genericDelete[DeleteResponse](c, ctx, endpoint, DeleteOptions{
		Namespace:    namespace,
		Name:         name,
		ResourceType: "CloudSpace",
		APIVersion:   version,
	})
}

// GetCloudSpace retrieves a specific cloudspace by name in the specified namespace
func (c *Client) GetCloudSpace(ctx context.Context, namespace, name string, apiVersion ...APIVersion) (*CloudSpace, error) {
	if err := validateNamespace(namespace); err != nil {
		return nil, err
	}
	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("cloudspace name is required")
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	endpoint := fmt.Sprintf("/namespaces/%s/cloudspaces/%s", namespace, name)
	return genericGet[CloudSpace](c, ctx, endpoint, GetOptions{Namespace: namespace, Name: name, APIVersion: version})
}

// EditCloudSpace edits a cloudspace using JSON patch operations
func (c *Client) EditCloudSpace(ctx context.Context, namespace, name string, patchOps []PatchOperation, apiVersion ...APIVersion) (*CloudSpace, error) {
	if err := validateNamespace(namespace); err != nil {
		return nil, err
	}
	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("cloudspace name is required")
	}
	if err := validatePatchOperations(patchOps); err != nil {
		return nil, err
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	endpoint := fmt.Sprintf("/namespaces/%s/cloudspaces/%s", namespace, name)
	return genericEdit[CloudSpace](c, ctx, endpoint, patchOps, EditOptions{Namespace: namespace, Name: name, APIVersion: version})
}

// ListSpotNodePools retrieves all spot node pools for a given namespace
func (c *Client) ListSpotNodePools(ctx context.Context, namespace string, apiVersion ...APIVersion) (*SpotNodePoolList, error) {
	if err := validateNamespace(namespace); err != nil {
		return nil, err
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	endpoint := fmt.Sprintf("/namespaces/%s/spotnodepools", namespace)
	return genericList[SpotNodePoolList](c, ctx, endpoint, ListOptions{Namespace: namespace, APIVersion: version})
}

// CreateSpotNodePool creates a new spot node pool in the specified namespace
func (c *Client) CreateSpotNodePool(ctx context.Context, namespace string, spotNodePool *SpotNodePool, apiVersion ...APIVersion) (*SpotNodePool, error) {
	if err := validateNamespace(namespace); err != nil {
		return nil, err
	}
	if err := validateCreateInput(spotNodePool); err != nil {
		return nil, fmt.Errorf("spot node pool configuration is required")
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	endpoint := fmt.Sprintf("/namespaces/%s/spotnodepools", namespace)
	return genericCreate[SpotNodePool](c, ctx, endpoint, spotNodePool, CreateOptions{Namespace: namespace, APIVersion: version})
}

// EditSpotNodePool edits a spot node pool using JSON patch operations
func (c *Client) EditSpotNodePool(ctx context.Context, namespace, name string, patchOps []PatchOperation, apiVersion ...APIVersion) (*SpotNodePool, error) {
	if err := validateNamespace(namespace); err != nil {
		return nil, err
	}
	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("spot node pool name is required")
	}
	if err := validatePatchOperations(patchOps); err != nil {
		return nil, err
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	endpoint := fmt.Sprintf("/namespaces/%s/spotnodepools/%s", namespace, name)
	return genericEdit[SpotNodePool](c, ctx, endpoint, patchOps, EditOptions{Namespace: namespace, Name: name, APIVersion: version})
}

// GetSpotNodePool retrieves a specific spot node pool by name in the specified namespace
func (c *Client) GetSpotNodePool(ctx context.Context, namespace, name string, apiVersion ...APIVersion) (*SpotNodePool, error) {
	if err := validateNamespace(namespace); err != nil {
		return nil, err
	}
	if err := validateName(name); err != nil {
		return nil, fmt.Errorf("spot node pool name is required")
	}

	version := APIVersionDefault
	if len(apiVersion) > 0 {
		version = apiVersion[0]
	}
	endpoint := fmt.Sprintf("/namespaces/%s/spotnodepools/%s", namespace, name)
	return genericGet[SpotNodePool](c, ctx, endpoint, GetOptions{Namespace: namespace, Name: name, APIVersion: version})
}

// HandleAPIError processes API error responses and returns appropriate error types
func (c *Client) HandleAPIError(resp *http.Response) error {
	if resp == nil {
		return errors.NewInternalError("received nil response", nil)
	}

	// For successful responses (2xx), return nil
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.NewAPIError(resp.StatusCode, "failed to read error response body", err)
	}

	var apiErr struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details string `json:"details,omitempty"`
	}

	if err := json.Unmarshal(body, &apiErr); err != nil {
		return errors.NewAPIError(resp.StatusCode, fmt.Sprintf("invalid error response format: %s", string(body)), nil)
	}

	if apiErr.Code == 0 {
		apiErr.Code = resp.StatusCode
	}

	// Add request context to error message
	errorMsg := fmt.Sprintf("%s %s failed: %s", resp.Request.Method, resp.Request.URL.String(), apiErr.Message)
	if apiErr.Details != "" {
		errorMsg = fmt.Sprintf("%s (Details: %s)", errorMsg, apiErr.Details)
	}

	return &APIError{
		Code:    apiErr.Code,
		Message: errorMsg,
		Details: apiErr.Details,
	}
}
