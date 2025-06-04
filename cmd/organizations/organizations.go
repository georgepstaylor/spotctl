package organizations

import "github.com/spf13/cobra"

// NewCommand returns the main organizations command with all subcommands
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "organizations",
		Short: "Manage Rackspace Spot organizations",
		Long: `Manage and view Rackspace Spot organizations.

This command allows you to list organizations that you have access to
in the Rackspace Spot platform.`,
	}

	// Add all subcommands
	cmd.AddCommand(NewListCommand())

	return cmd
}
