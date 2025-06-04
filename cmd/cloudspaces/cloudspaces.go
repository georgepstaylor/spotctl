package cloudspaces

import "github.com/spf13/cobra"

// NewCommand returns the main cloudspaces command with all subcommands
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cloudspaces",
		Short: "Manage Rackspace Spot cloudspaces",
		Long: `Manage and view Rackspace Spot cloudspaces.

This command allows you to list cloudspaces within a specific namespace.
Cloudspaces represent Kubernetes clusters deployed through Rackspace Spot.`,
	}

	// Add all subcommands
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewGetCommand())
	cmd.AddCommand(NewCreateCommand())
	cmd.AddCommand(NewDeleteCommand())

	return cmd
}
