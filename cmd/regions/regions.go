package regions

import "github.com/spf13/cobra"

// NewCommand returns the main regions command with all subcommands
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "regions",
		Short: "Manage Rackspace Spot regions",
		Long: `Manage and view Rackspace Spot regions.

This command allows you to list and view information about available
Rackspace Spot regions where you can deploy resources.`,
	}

	// Add all subcommands
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewGetCommand())

	return cmd
}
