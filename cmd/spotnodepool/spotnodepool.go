package spotnodepool

import "github.com/spf13/cobra"

// NewCommand returns the main spotnodepool command with all subcommands
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "spotnodepool",
		Short: "Manage Rackspace Spot node pools",
		Long: `Manage and view Rackspace Spot node pools.

This command allows you to view spot node pools within a specific namespace.
Spot node pools represent groups of worker nodes deployed through Rackspace Spot.`,
	}

	// Add all subcommands
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewGetCommand())

	return cmd
}
