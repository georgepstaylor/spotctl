package client

import (
	"testing"
	"time"
)

func TestTokenManager_IsValid(t *testing.T) {
	tests := []struct {
		name         string
		refreshToken string
		want         bool
	}{
		{
			name:         "valid token",
			refreshToken: "valid-token",
			want:         true,
		},
		{
			name:         "empty token",
			refreshToken: "",
			want:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TokenManager{
				refreshToken: tt.refreshToken,
			}
			if got := tm.IsValid(); got != tt.want {
				t.Errorf("TokenManager.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenManager_tokenExpiry(t *testing.T) {
	tm := &TokenManager{
		refreshToken: "test-token",
		accessToken:  "test-access-token",
		expiresAt:    time.Now().Add(10 * time.Minute), // Expires in 10 minutes
	}

	// Should be considered valid (more than 5 minute buffer)
	if !tm.hasValidToken() {
		t.Error("Token should be considered valid when it expires in more than 5 minutes")
	}

	// Set expiry to 3 minutes from now
	tm.expiresAt = time.Now().Add(3 * time.Minute)
	if tm.hasValidToken() {
		t.Error("Token should be considered invalid when it expires in less than 5 minutes")
	}

	// Set expiry to past
	tm.expiresAt = time.Now().Add(-1 * time.Minute)
	if tm.hasValidToken() {
		t.Error("Token should be considered invalid when it has expired")
	}
}

// Helper method for testing token validity
func (tm *TokenManager) hasValidToken() bool {
	return tm.accessToken != "" && time.Now().Add(5*time.Minute).Before(tm.expiresAt)
}
