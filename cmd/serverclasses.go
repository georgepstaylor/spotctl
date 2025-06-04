package cmd

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// serverclassesCmd represents the serverclasses command
var serverclassesCmd = &cobra.Command{
	Use:   "serverclasses",
	Short: "Manage Rackspace Spot server classes",
	Long: `Manage and view Rackspace Spot server classes.

This command allows you to list and view information about available
Rackspace Spot server classes including pricing, resources, and availability.`,
}

// serverclassesListCmd represents the serverclasses list command
var serverclassesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available server classes",
	Long: `List all available Rackspace Spot server classes.

This command retrieves and displays all server classes available across
all regions, including their specifications, pricing, and current availability.`,
	RunE: runServerClassesList,
}

// serverclassesGetCmd represents the serverclasses get command
var serverclassesGetCmd = &cobra.Command{
	Use:   "get <name>",
	Short: "Get a specific server class by name",
	Long: `Get detailed information about a specific Rackspace Spot server class.

This command retrieves and displays information for a single server class
by its name, including specifications, pricing, and current availability.

Examples:
  spotctl serverclasses get standard-2
  spotctl serverclasses get compute-optimized-4 -o json`,
	Args: cobra.ExactArgs(1),
	RunE: runServerClassesGet,
}

func runServerClassesList(cmd *cobra.Command, args []string) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	serverClassList, err := client.ListServerClasses(ctx)
	if err != nil {
		return fmt.Errorf("failed to list server classes: %w", err)
	}

	// Read flags directly from command
	outputFormat, _ := cmd.Flags().GetString("output")
	showDetails, _ := cmd.Flags().GetBool("details")
	wideOutput, _ := cmd.Flags().GetBool("wide")

	return outputServerClasses(serverClassList, outputFormat, showDetails, wideOutput)
}

func runServerClassesGet(cmd *cobra.Command, args []string) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	name := args[0]
	serverClass, err := client.GetServerClass(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to get server class '%s': %w", name, err)
	}

	// Read flags directly from command
	outputFormat, _ := cmd.Flags().GetString("output")
	showDetails, _ := cmd.Flags().GetBool("details")
	wideOutput, _ := cmd.Flags().GetBool("wide")

	return outputServerClass(serverClass, outputFormat, showDetails, wideOutput)
}

func init() {
	rootCmd.AddCommand(serverclassesCmd)
	serverclassesCmd.AddCommand(serverclassesListCmd)
	serverclassesCmd.AddCommand(serverclassesGetCmd)

	// Add flags for serverclasses list command
	serverclassesListCmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	serverclassesListCmd.Flags().Bool("details", false, "Show detailed server class information")
	serverclassesListCmd.Flags().Bool("wide", false, "Show additional columns including availability and pricing")

	// Add flags for serverclasses get command
	serverclassesGetCmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	serverclassesGetCmd.Flags().Bool("details", false, "Show detailed server class information")
	serverclassesGetCmd.Flags().Bool("wide", false, "Show additional columns including availability and pricing")
}
