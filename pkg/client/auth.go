package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TokenManagerInterface defines the interface for token management
type TokenManagerInterface interface {
	GetValidAccessToken(ctx context.Context) (string, error)
}

const (
	// OAuth constants for Rackspace Spot
	OAuthURL  = "https://login.spot.rackspace.com/oauth/token"
	ClientID  = "mwG3lUMV8KyeMqHe4fJ5Bb3nM1vBvRNa"
	GrantType = "refresh_token"
)

// TokenResponse represents the OAuth token response
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// TokenManager handles OAuth token management
type TokenManager struct {
	refreshToken string
	accessToken  string
	expiresAt    time.Time
	httpClient   *http.Client
	mutex        sync.RWMutex
	debug        bool
}

// NewTokenManager creates a new token manager
func NewTokenManager(refreshToken string, httpClient *http.Client, debug bool) *TokenManager {
	return &TokenManager{
		refreshToken: refreshToken,
		httpClient:   httpClient,
		debug:        debug,
	}
}

// GetValidAccessToken returns a valid access token, refreshing if necessary
func (tm *TokenManager) GetValidAccessToken(ctx context.Context) (string, error) {
	tm.mutex.RLock()
	// Check if we have a valid token (with 5 minute buffer)
	if tm.accessToken != "" && time.Now().Add(5*time.Minute).Before(tm.expiresAt) {
		token := tm.accessToken
		tm.mutex.RUnlock()
		return token, nil
	}
	tm.mutex.RUnlock()

	// Need to refresh the token
	return tm.refreshAccessToken(ctx)
}

// refreshAccessToken gets a new access token using the refresh token
func (tm *TokenManager) refreshAccessToken(ctx context.Context) (string, error) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	// Double-check in case another goroutine already refreshed
	if tm.accessToken != "" && time.Now().Add(5*time.Minute).Before(tm.expiresAt) {
		return tm.accessToken, nil
	}

	if tm.debug {
		fmt.Println("Refreshing OAuth access token...")
	}

	// Prepare form data
	data := url.Values{}
	data.Set("grant_type", GrantType)
	data.Set("client_id", ClientID)
	data.Set("refresh_token", tm.refreshToken)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, OAuthURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "rackspace-spot-cli/1.0.0")

	resp, err := tm.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token request failed with status %d", resp.StatusCode)
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	// Update token information - use ID token for API authentication
	tm.accessToken = tokenResp.IDToken
	tm.expiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

	if tm.debug {
		fmt.Printf("Access token refreshed, expires at: %s\n", tm.expiresAt.Format(time.RFC3339))
	}

	return tm.accessToken, nil
}

// IsValid checks if the refresh token can be used
func (tm *TokenManager) IsValid() bool {
	return tm.refreshToken != ""
}
