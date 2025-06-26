package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// Generic API operation patterns to reduce duplication in client methods

// ListOptions contains parameters for list operations
type ListOptions struct {
	Namespace  string
	APIVersion APIVersion
}

// GetOptions contains parameters for get operations
type GetOptions struct {
	Namespace  string
	Name       string
	APIVersion APIVersion
}

// CreateOptions contains parameters for create operations
type CreateOptions struct {
	Namespace  string
	APIVersion APIVersion
}

// EditOptions contains parameters for edit operations
type EditOptions struct {
	Namespace  string
	Name       string
	APIVersion APIVersion
}

// DeleteOptions contains parameters for delete operations
type DeleteOptions struct {
	Namespace    string
	Name         string
	APIVersion   APIVersion
	ResourceType string // Used for success message generation
}

// genericList performs a generic list operation
func genericList[T any](c *Client, ctx context.Context, endpoint string, opts ListOptions) (*T, error) {
	// Use default API version if not specified
	apiVersion := opts.APIVersion
	if apiVersion == "" {
		apiVersion = APIVersionDefault
	}

	resp, err := c.MakeRequest(ctx, http.MethodGet, endpoint, nil, apiVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// genericGet performs a generic get operation
func genericGet[T any](c *Client, ctx context.Context, endpoint string, opts GetOptions) (*T, error) {
	// Use default API version if not specified
	apiVersion := opts.APIVersion
	if apiVersion == "" {
		apiVersion = APIVersionDefault
	}

	resp, err := c.MakeRequest(ctx, http.MethodGet, endpoint, nil, apiVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// genericCreate performs a create operation
func genericCreate[T any](c *Client, ctx context.Context, endpoint string, body interface{}, opts CreateOptions) (*T, error) {
	// Use default API version if not specified
	apiVersion := opts.APIVersion
	if apiVersion == "" {
		apiVersion = APIVersionDefault
	}

	resp, err := c.MakeRequest(ctx, http.MethodPost, endpoint, body, apiVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// genericEdit performs an edit operation using JSON patch
func genericEdit[T any](c *Client, ctx context.Context, endpoint string, patchOps []PatchOperation, opts EditOptions) (*T, error) {
	// Use default API version if not specified
	apiVersion := opts.APIVersion
	if apiVersion == "" {
		apiVersion = APIVersionDefault
	}

	resp, err := c.MakeRequest(ctx, http.MethodPatch, endpoint, patchOps, apiVersion, "application/json-patch+json")
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// genericDelete performs a delete operation
func genericDelete[T any](c *Client, ctx context.Context, endpoint string, opts DeleteOptions) (*T, error) {
	// Use default API version if not specified
	apiVersion := opts.APIVersion
	if apiVersion == "" {
		apiVersion = APIVersionDefault
	}

	resp, err := c.MakeRequest(ctx, http.MethodDelete, endpoint, nil, apiVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.HandleAPIError(resp); err != nil {
		return nil, err
	}

	// For successful deletes, HTTP 200 or 202 means success according to API docs
	if resp.StatusCode == 200 || resp.StatusCode == 202 {
		// Try to parse the response, but if it fails, return a default success response
		var result T
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			// Create a default DeleteResponse for cases where we can't parse the response
			// but got a successful status code
			if opts.ResourceType != "" {
				// This is a bit of a hack, but since we're working with generics,
				// we'll assume T is *DeleteResponse for delete operations
				defaultResponse := &DeleteResponse{
					Status:  "Success",
					Message: fmt.Sprintf("%s deleted successfully", opts.ResourceType),
				}
				// Type assertion to convert to T (assuming T is *DeleteResponse)
				return any(defaultResponse).(*T), nil
			}
		}
		return &result, nil
	}

	// Try to parse the actual response for error details
	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		// If we can't parse the response but got a successful status code, create default response
		if opts.ResourceType != "" {
			defaultResponse := &DeleteResponse{
				Status:  "Success",
				Message: fmt.Sprintf("%s deleted successfully", opts.ResourceType),
			}
			return any(defaultResponse).(*T), nil
		}
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// Validation helpers
func validateNamespace(namespace string) error {
	if namespace == "" {
		return fmt.Errorf("namespace is required")
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

func validateCreateInput(input interface{}) error {
	if input == nil {
		return fmt.Errorf("input object is required")
	}

	// Handle typed nil pointers (e.g., *CloudSpace = nil)
	// Use reflection to check if the underlying value is nil
	if reflect.ValueOf(input).Kind() == reflect.Ptr && reflect.ValueOf(input).IsNil() {
		return fmt.Errorf("input object is required")
	}

	return nil
}

func validatePatchOperations(patchOps []PatchOperation) error {
	if patchOps == nil {
		return fmt.Errorf("patch operations are required")
	}
	return nil
}
