package ondemandnodepools

import "github.com/spf13/cobra"

// NewCommand returns the main ondemandnodepool command with all subcommands
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ondemandnodepool",
		Short: "Manage Rackspace On Demand node pools",
		Long: `Manage and view Rackspace On Demand node pools.

This command allows you to view on demand node pools within a specific namespace.
On demand node pools represent groups of worker nodes deployed through Rackspace On Demand.`,
	}

	// Add implemented subcommands
	cmd.AddCommand(NewGetCommand())
	// TODO: Add other commands as they are implemented
	// cmd.AddCommand(NewListCommand())
	// cmd.AddCommand(NewCreateCommand())
	// cmd.AddCommand(NewEditCommand())
	// cmd.AddCommand(NewDeleteCommand())
	// cmd.AddCommand(NewDeleteAllCommand())

	return cmd
}
