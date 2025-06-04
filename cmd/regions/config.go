package regions

import (
	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/output"
	"github.com/georgetaylor/spotctl/pkg/pager"
	"github.com/spf13/viper"
)

// getRegionsTableConfig returns the table configuration for regions
func getRegionsTableConfig() *output.TableConfig {
	return &output.TableConfig{
		Columns: []output.TableColumn{
			{Header: "NAME", Field: "metadata.name", Default: "N/A"},
			{Header: "COUNTRY", Field: "spec.country", Default: "N/A"},
			{Header: "PROVIDER", Field: "spec.provider.providerType", Default: "N/A"},
		},
		DetailCols: []output.TableColumn{
			{Header: "PROVIDER REGION", Field: "spec.provider.providerRegionName", Default: "N/A"},
			{Header: "DESCRIPTION", Field: "spec.description", Default: "N/A", Width: 50},
		},
		WideCols: []output.TableColumn{
			{Header: "API VERSION", Field: "apiVersion", Default: "N/A"},
			{Header: "KIND", Field: "kind", Default: "N/A"},
			{Header: "UID", Field: "metadata.uid", Default: "N/A"},
		},
	}
}

// outputRegions handles formatting and output of region lists
func outputRegions(regionList *client.RegionList, format string, showDetails bool, wideOutput bool) error {
	// Create formatter with options
	options := output.OutputOptions{
		Format:      output.OutputFormat(format),
		ShowDetails: showDetails,
		WideOutput:  wideOutput,
	}

	// Check if pager should be disabled
	noPager := viper.GetBool("no-pager")
	if noPager {
		// Create pager with disabled setting
		pager := pager.NewPager()
		pager.Disable = true
		formatter := output.NewFormatterWithPager(options, pager)

		// Get table configuration for regions
		tableConfig := getRegionsTableConfig()

		// Output using the shared formatter
		return formatter.Output(regionList, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Get table configuration for regions
	tableConfig := getRegionsTableConfig()

	// Output using the shared formatter
	return formatter.Output(regionList, tableConfig)
}

// outputRegion handles formatting and output of a single region
func outputRegion(region *client.Region, format string, showDetails bool, wideOutput bool) error {
	// Create formatter with options
	options := output.OutputOptions{
		Format:      output.OutputFormat(format),
		ShowDetails: showDetails,
		WideOutput:  wideOutput,
	}

	// Check if pager should be disabled
	noPager := viper.GetBool("no-pager")
	if noPager {
		// Create pager with disabled setting
		pager := pager.NewPager()
		pager.Disable = true
		formatter := output.NewFormatterWithPager(options, pager)

		// Get table configuration for regions
		tableConfig := getRegionsTableConfig()

		// Output using the shared formatter
		return formatter.Output(region, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Get table configuration for regions (same config works for single region)
	tableConfig := getRegionsTableConfig()

	// Output using the shared formatter
	return formatter.Output(region, tableConfig)
}
