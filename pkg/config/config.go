package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	RefreshToken string `mapstructure:"refresh-token"`
	Region       string `mapstructure:"region"`
	BaseURL      string `mapstructure:"base-url"`
	Debug        bool   `mapstructure:"debug"`
	Timeout      int    `mapstructure:"timeout"`
}

// GetConfig returns the current configuration
func GetConfig() (*Config, error) {
	var cfg Config

	// Set defaults
	viper.SetDefault("base-url", "https://spot.rackspace.com/apis/ngpc.rxt.io/v1")
	viper.SetDefault("timeout", 30)
	viper.SetDefault("region", "uk-lon-1")

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate required fields
	if cfg.RefreshToken == "" {
		return nil, fmt.Errorf("refresh token is required. Set it via --refresh-token flag, config file, or SPOTCTL_REFRESH_TOKEN environment variable")
	}

	return &cfg, nil
}

// ValidateConfig checks if the configuration is valid
func ValidateConfig(cfg *Config) error {
	if cfg.RefreshToken == "" {
		return fmt.Errorf("refresh token is required")
	}
	if cfg.Region == "" {
		return fmt.Errorf("region is required")
	}
	if cfg.BaseURL == "" {
		return fmt.Errorf("base URL is required")
	}
	return nil
}

// SaveConfig saves the current configuration to file
func SaveConfig(cfg *Config) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	viper.Set("refresh-token", cfg.RefreshToken)
	viper.Set("region", cfg.Region)
	viper.Set("base-url", cfg.BaseURL)
	viper.Set("debug", cfg.Debug)
	viper.Set("timeout", cfg.Timeout)

	configDir := fmt.Sprintf("%s/.spot", home)
	configPath := fmt.Sprintf("%s/config.yaml", configDir)

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	return viper.WriteConfigAs(configPath)
}
