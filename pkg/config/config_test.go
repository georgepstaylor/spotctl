package config

import (
	"testing"
)

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				RefreshToken: "test-token",
				Region:       "us-east-1",
				BaseURL:      "https://api.test.com",
			},
			wantErr: false,
		},
		{
			name: "missing refresh token",
			config: &Config{
				Region:  "us-east-1",
				BaseURL: "https://api.test.com",
			},
			wantErr: true,
		},
		{
			name: "missing region",
			config: &Config{
				RefreshToken: "test-token",
				BaseURL:      "https://api.test.com",
			},
			wantErr: true,
		},
		{
			name: "missing base URL",
			config: &Config{
				RefreshToken: "test-token",
				Region:       "us-east-1",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateConfig(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
