package serverclasses

import "github.com/spf13/cobra"

// NewCommand returns the main serverclasses command with all subcommands
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serverclasses",
		Short: "Manage Rackspace Spot server classes",
		Long: `Manage and view Rackspace Spot server classes.

This command allows you to list and view information about available
server classes in Rackspace Spot regions.`,
	}

	// Add all subcommands
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewGetCommand())

	return cmd
}
