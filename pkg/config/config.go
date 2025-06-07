package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/georgetaylor/spotctl/pkg/errors"
	"github.com/spf13/viper"
)

// Default configuration values
var Defaults = struct {
	BaseURL      string
	Timeout      int
	Debug        bool
	OutputFormat string
}{
	BaseURL:      "https://spot.rackspace.com/apis",
	Timeout:      30,
	Debug:        false,
	OutputFormat: "table",
}

// Config represents the application configuration
type Config struct {
	RefreshToken string `mapstructure:"refresh-token"`
	Region       string `mapstructure:"region"`
	BaseURL      string `mapstructure:"base-url"`
	Debug        bool   `mapstructure:"debug"`
	Timeout      int    `mapstructure:"timeout"`
	OutputFormat string `mapstructure:"output-format"`
}

// ValidateConfig validates the configuration
func ValidateConfig(cfg *Config) error {
	if cfg.RefreshToken == "" {
		return errors.NewValidationError(
			"refresh token is required. Set it via --refresh-token flag, config file, or SPOTCTL_REFRESH_TOKEN environment variable",
			nil,
		)
	}

	if cfg.Region == "" {
		return errors.NewValidationError(
			"region is required. Set it via --region flag or in your config file",
			nil,
		)
	}

	if cfg.BaseURL == "" {
		return errors.NewValidationError(
			"base URL is required. Set it via --base-url flag or in your config file",
			nil,
		)
	}

	return nil
}

// GetConfig returns the current configuration
func GetConfig() (*Config, error) {
	var cfg Config

	// Set defaults
	viper.SetDefault("base-url", Defaults.BaseURL)
	viper.SetDefault("timeout", Defaults.Timeout)
	viper.SetDefault("debug", Defaults.Debug)
	viper.SetDefault("output-format", Defaults.OutputFormat)

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.NewConfigError("failed to unmarshal config", err)
	}

	// Validate required fields with descriptive messages
	if err := ValidateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// SaveConfig saves the current configuration to file
func SaveConfig(cfg *Config) error {
	// Validate before saving with descriptive messages
	if err := ValidateConfig(cfg); err != nil {
		return err
	}

	// Set values in viper
	viper.Set("refresh-token", cfg.RefreshToken)
	viper.Set("region", cfg.Region)
	viper.Set("base-url", cfg.BaseURL)
	viper.Set("debug", cfg.Debug)
	viper.Set("timeout", cfg.Timeout)
	viper.Set("output-format", cfg.OutputFormat)

	// Get config path
	configPath := os.Getenv("SPOTCTL_CONFIG")
	if configPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return errors.NewConfigError("failed to get user home directory", err)
		}
		configPath = filepath.Join(home, ".config", "spotctl", "config.yaml")
	}

	// Ensure directory exists
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.NewConfigError(fmt.Sprintf("failed to create config directory at %s", dir), err)
	}

	// Write config file
	if err := viper.WriteConfigAs(configPath); err != nil {
		return errors.NewConfigError(fmt.Sprintf("failed to write config file to %s", configPath), err)
	}

	return nil
}

// InitConfig initializes the configuration
func InitConfig() error {
	// Set config file name and type
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Set config file locations
	viper.AddConfigPath("$HOME/.config/spotctl")
	viper.AddConfigPath(".")

	// Read environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("SPOTCTL")

	// Bind environment variables
	viper.BindEnv("refresh-token", "SPOTCTL_REFRESH_TOKEN")
	viper.BindEnv("region", "SPOTCTL_REGION")
	viper.BindEnv("base-url", "SPOTCTL_BASE_URL")
	viper.BindEnv("debug", "SPOTCTL_DEBUG")
	viper.BindEnv("timeout", "SPOTCTL_TIMEOUT")
	viper.BindEnv("output-format", "SPOTCTL_OUTPUT_FORMAT")

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return errors.NewConfigError("failed to read config file", err)
		}
		// Config file not found is not an error, we'll use defaults
	}

	return nil
}
