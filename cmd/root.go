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
	Use:   "rackspace-spot",
	Short: "A CLI tool for interacting with the Rackspace Spot API",
	Long: `rackspace-spot is a command-line interface for managing and interacting 
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rackspace-spot.yaml)")
	rootCmd.PersistentFlags().String("refresh-token", "", "Rackspace Spot refresh token")
	rootCmd.PersistentFlags().String("region", "", "Rackspace region")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug output")

	// Bind flags to viper
	viper.BindPFlag("refresh-token", rootCmd.PersistentFlags().Lookup("refresh-token"))
	viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
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

		// Search config in home directory with name ".rackspace-spot" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".rackspace-spot")
	}

	// Environment variables
	viper.SetEnvPrefix("RACKSPACE_SPOT")
	viper.AutomaticEnv() // read in environment variables that match

	// Explicitly bind environment variables to handle hyphenated keys
	viper.BindEnv("refresh-token", "RACKSPACE_SPOT_REFRESH_TOKEN")
	viper.BindEnv("base-url", "RACKSPACE_SPOT_BASE_URL")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("debug") {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}
