package cmd

import (
	"github.com/spf13/cobra"
)

// pricingCmd represents the pricing command
var pricingCmd = &cobra.Command{
	Use:   "pricing",
	Short: "Get spot pricing information",
	Long:  `Get current and historical spot pricing information for different instance types and regions.`,
}

// pricingCurrentCmd gets current pricing
var pricingCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current spot pricing",
	Long:  `Get current spot pricing for all available instance types and regions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement current pricing when API documentation is available
		cmd.Println("Current pricing will be implemented once API documentation is provided.")
	},
}

// pricingHistoryCmd gets pricing history
var pricingHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Get spot pricing history",
	Long:  `Get historical spot pricing data for analysis and trends.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement pricing history when API documentation is available
		cmd.Println("Pricing history will be implemented once API documentation is provided.")
	},
}

func init() {
	rootCmd.AddCommand(pricingCmd)
	pricingCmd.AddCommand(pricingCurrentCmd)
	pricingCmd.AddCommand(pricingHistoryCmd)

	// Add common flags
	pricingCurrentCmd.Flags().StringP("output", "o", "table", "Output format (json, table)")
	pricingCurrentCmd.Flags().StringP("region", "r", "", "Filter by region")
	pricingCurrentCmd.Flags().StringP("instance-type", "t", "", "Filter by instance type")

	pricingHistoryCmd.Flags().StringP("output", "o", "table", "Output format (json, table)")
	pricingHistoryCmd.Flags().StringP("region", "r", "", "Filter by region")
	pricingHistoryCmd.Flags().StringP("instance-type", "t", "", "Filter by instance type")
	pricingHistoryCmd.Flags().StringP("start-date", "", "", "Start date for history (YYYY-MM-DD)")
	pricingHistoryCmd.Flags().StringP("end-date", "", "", "End date for history (YYYY-MM-DD)")
}
