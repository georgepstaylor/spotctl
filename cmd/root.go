package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spotctl",
	Short: "A CLI tool for interacting with the Rackspace Spot API",
	Long: `spotctl is a command-line interface for managing and interacting 
with Rackspace Spot resources through their public API.

This tool allows you to manage spot instances, monitor pricing, 
and perform various operations on your Rackspace Spot infrastructure.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.spot/config.yaml)")
	rootCmd.PersistentFlags().String("refresh-token", "", "Rackspace Spot refresh token")
	rootCmd.PersistentFlags().String("region", "", "Rackspace region")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug output")
	rootCmd.PersistentFlags().Bool("no-pager", false, "Disable pager for long output")

	// Bind flags to viper
	viper.BindPFlag("refresh-token", rootCmd.PersistentFlags().Lookup("refresh-token"))
	viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("no-pager", rootCmd.PersistentFlags().Lookup("no-pager"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in .spot directory with name "config" (without extension).
		viper.AddConfigPath(home + "/.spot")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	// Environment variables
	viper.SetEnvPrefix("SPOTCTL")
	viper.AutomaticEnv() // read in environment variables that match

	// Explicitly bind environment variables to handle hyphenated keys
	viper.BindEnv("refresh-token", "SPOTCTL_REFRESH_TOKEN")
	viper.BindEnv("base-url", "SPOTCTL_BASE_URL")
	viper.BindEnv("no-pager", "SPOTCTL_NO_PAGER")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("debug") {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}
