package serverclasses

import (
	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/output"
	"github.com/georgetaylor/spotctl/pkg/pager"
	"github.com/spf13/viper"
)

// getServerClassesTableConfig returns the table configuration for server classes
func getServerClassesTableConfig() *output.TableConfig {
	return &output.TableConfig{
		Columns: []output.TableColumn{
			{Header: "NAME", Field: "metadata.name"},
			{Header: "DISPLAY NAME", Field: "spec.displayName"},
			{Header: "REGION", Field: "spec.region"},
			{Header: "CPU", Field: "spec.resources.cpu"},
			{Header: "MEMORY", Field: "spec.resources.memory"},
			{Header: "AVAILABILITY", Field: "spec.availability"},
		},
		DetailCols: []output.TableColumn{
			{Header: "CATEGORY", Field: "spec.category"},
			{Header: "FLAVOR TYPE", Field: "spec.flavorType"},
			{Header: "PROVIDER TYPE", Field: "spec.provider.providerType"},
			{Header: "ON-DEMAND COST", Field: "spec.onDemandPricing.cost"},
			{Header: "SPOT PRICE", Field: "status.spotPricing.marketPricePerHour", Default: "N/A"},
			{Header: "HAMMER PRICE", Field: "status.spotPricing.hammerPricePerHour", Default: "N/A"},
		},
	}
}

// outputServerClasses handles formatting and output of server class lists
func outputServerClasses(serverClassList *client.ServerClassList, format string) error {
	// Create formatter with options
	options := output.OutputOptions{
		Format: output.OutputFormat(format),
	}

	// Check if pager should be disabled
	noPager := viper.GetBool("no-pager")
	if noPager {
		// Create pager with disabled setting
		pager := pager.NewPager()
		pager.Disable = true
		formatter := output.NewFormatterWithPager(options, pager)

		// Get table configuration for server classes
		tableConfig := getServerClassesTableConfig()

		// Output using the shared formatter
		return formatter.Output(serverClassList, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Get table configuration for server classes
	tableConfig := getServerClassesTableConfig()

	// Output using the shared formatter
	return formatter.Output(serverClassList, tableConfig)
}

// outputServerClass handles formatting and output of a single server class
func outputServerClass(serverClass *client.ServerClass, format string) error {
	// Create formatter with options
	options := output.OutputOptions{
		Format: output.OutputFormat(format),
	}

	// Check if pager should be disabled
	noPager := viper.GetBool("no-pager")
	if noPager {
		// Create pager with disabled setting
		pager := pager.NewPager()
		pager.Disable = true
		formatter := output.NewFormatterWithPager(options, pager)

		// Get table configuration for server classes
		tableConfig := getServerClassesTableConfig()

		// Output using the shared formatter
		return formatter.Output(serverClass, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Get table configuration for server classes
	tableConfig := getServerClassesTableConfig()

	// Output using the shared formatter
	return formatter.Output(serverClass, tableConfig)
}
