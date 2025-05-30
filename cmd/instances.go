package cmd

import (
	"github.com/spf13/cobra"
)

// instancesCmd represents the instances command
var instancesCmd = &cobra.Command{
	Use:   "instances",
	Short: "Manage spot instances",
	Long:  `Manage Rackspace spot instances - create, list, delete, and monitor instances.`,
}

// instancesListCmd lists spot instances
var instancesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List spot instances",
	Long:  `List all spot instances in your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement instance listing when API documentation is available
		cmd.Println("Instance listing will be implemented once API documentation is provided.")
	},
}

// instancesCreateCmd creates a spot instance
var instancesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a spot instance",
	Long:  `Create a new spot instance with specified configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement instance creation when API documentation is available
		cmd.Println("Instance creation will be implemented once API documentation is provided.")
	},
}

// instancesDeleteCmd deletes a spot instance
var instancesDeleteCmd = &cobra.Command{
	Use:   "delete <instance-id>",
	Short: "Delete a spot instance",
	Long:  `Delete a spot instance by ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement instance deletion when API documentation is available
		cmd.Printf("Instance deletion for %s will be implemented once API documentation is provided.\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(instancesCmd)
	instancesCmd.AddCommand(instancesListCmd)
	instancesCmd.AddCommand(instancesCreateCmd)
	instancesCmd.AddCommand(instancesDeleteCmd)

	// Add common flags
	instancesListCmd.Flags().StringP("output", "o", "table", "Output format (json, table)")
	instancesCreateCmd.Flags().StringP("output", "o", "table", "Output format (json, table)")
}
