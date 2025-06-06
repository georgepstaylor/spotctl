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

		fmt.Printf("  region: %s\n", viper.GetString("region"))
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
		validKeys := []string{"refresh-token", "region", "base-url", "debug", "timeout"}
		if !contains(validKeys, key) {
			CheckError(fmt.Errorf("invalid configuration key '%s'. Valid keys are: %v", key, validKeys))
		}

		// Set the value in viper
		viper.Set(key, value)

		// Create a config object to save
		cfg := &config.Config{
			RefreshToken: viper.GetString("refresh-token"),
			Region:       viper.GetString("region"),
			BaseURL:      viper.GetString("base-url"),
			Debug:        viper.GetBool("debug"),
			Timeout:      viper.GetInt("timeout"),
		}

		// Save the configuration
		err := config.SaveConfig(cfg)
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

		// Prompt for region (required)
		fmt.Print("Enter your default region (e.g., uk-lon-1, aus-syd-1): ")
		var region string
		fmt.Scanln(&region)

		if region == "" {
			CheckError(fmt.Errorf("region is required"))
		}

		// Set values
		viper.Set("refresh-token", refreshToken)
		viper.Set("region", region)
		viper.Set("base-url", config.Defaults.BaseURL)
		viper.Set("debug", config.Defaults.Debug)
		viper.Set("timeout", config.Defaults.Timeout)

		// Create config object
		cfg := &config.Config{
			RefreshToken: refreshToken,
			Region:       region,
			BaseURL:      config.Defaults.BaseURL,
			Debug:        config.Defaults.Debug,
			Timeout:      config.Defaults.Timeout,
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
