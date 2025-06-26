package cmd

import (
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
	Long:  `Manage configuration settings for spotctl.`,
}

// configShowCmd shows the current configuration
var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Long:  `Display the current configuration settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := config.GetConfig()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			fmt.Println("\nCurrent configuration from environment and flags:")
		} else {
			fmt.Println("Current configuration:")
		}

		// Show configuration values (mask the refresh token for security)
		refreshToken := viper.GetString("refresh-token")
		if refreshToken != "" {
			maskedToken := refreshToken[:min(8, len(refreshToken))] + "***"
			fmt.Printf("  refresh-token: %s\n", maskedToken)
		} else {
			fmt.Printf("  refresh-token: <not set>\n")
		}

		fmt.Printf("  namespace: %s\n", viper.GetString("namespace"))
		fmt.Printf("  base-url: %s\n", viper.GetString("base-url"))
		fmt.Printf("  debug: %t\n", viper.GetBool("debug"))
		fmt.Printf("  timeout: %d\n", viper.GetInt("timeout"))

		if viper.ConfigFileUsed() != "" {
			fmt.Printf("\nConfig file: %s\n", viper.ConfigFileUsed())
		} else {
			fmt.Printf("\nNo config file found. You can create one at ~/.spot/config.yaml\n")
		}
	},
}

// configSetCmd sets a configuration value
var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Long:  `Set a configuration value and save it to the config file.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		// Validate the key
		validKeys := []string{"refresh-token", "namespace", "base-url", "debug", "timeout"}
		if !contains(validKeys, key) {
			CheckError(fmt.Errorf("invalid configuration key '%s'. Valid keys are: %v", key, validKeys))
		}

		// Load existing config first to preserve other values
		existingCfg, err := config.GetConfig()
		if err != nil {
			// If config doesn't exist or is invalid, start with defaults
			existingCfg = &config.Config{
				RefreshToken: viper.GetString("refresh-token"),
				Namespace:    viper.GetString("namespace"),
				BaseURL:      viper.GetString("base-url"),
				Debug:        viper.GetBool("debug"),
				Timeout:      viper.GetInt("timeout"),
				OutputFormat: viper.GetString("output-format"),
			}
		}

		// Set the new value in viper
		viper.Set(key, value)

		// Create a config object with the updated value
		cfg := &config.Config{
			RefreshToken: existingCfg.RefreshToken,
			Namespace:    existingCfg.Namespace,
			BaseURL:      existingCfg.BaseURL,
			Debug:        existingCfg.Debug,
			Timeout:      existingCfg.Timeout,
			OutputFormat: existingCfg.OutputFormat,
		}

		// Update the specific field that was changed
		switch key {
		case "refresh-token":
			cfg.RefreshToken = value
		case "namespace":
			cfg.Namespace = value
		case "base-url":
			cfg.BaseURL = value
		case "debug":
			cfg.Debug = viper.GetBool("debug")
		case "timeout":
			cfg.Timeout = viper.GetInt("timeout")
		}

		// Save the configuration
		err = config.SaveConfig(cfg)
		CheckError(err)

		fmt.Printf("Configuration saved: %s = %s\n", key, value)
	},
}

// configInitCmd initializes the configuration with prompts
var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration with interactive prompts",
	Long:  `Initialize the configuration file by prompting for required values.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing spotctl configuration...")
		fmt.Println()

		// Prompt for refresh token
		fmt.Print("Enter your Rackspace Spot refresh token: ")
		var refreshToken string
		fmt.Scanln(&refreshToken)

		if refreshToken == "" {
			CheckError(fmt.Errorf("refresh token is required"))
		}

		// Set values
		viper.Set("refresh-token", refreshToken)
		viper.Set("namespace", "") // Default to empty namespace
		viper.Set("base-url", config.Defaults.BaseURL)
		viper.Set("debug", config.Defaults.Debug)
		viper.Set("timeout", config.Defaults.Timeout)

		// Create config object
		cfg := &config.Config{
			RefreshToken: refreshToken,
			Namespace:    "", // Default to empty namespace
			BaseURL:      config.Defaults.BaseURL,
			Debug:        config.Defaults.Debug,
			Timeout:      config.Defaults.Timeout,
			OutputFormat: config.Defaults.OutputFormat,
		}

		// Save the configuration
		err := config.SaveConfig(cfg)
		CheckError(err)

		fmt.Println()
		fmt.Println("Configuration saved successfully!")
		fmt.Println("You can now use spotctl.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configInitCmd) // Add configInitCmd to config command

	// Add output flag to show command
	AddOutputFlag(configShowCmd)
}

// Helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Helper function to get minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
