package spotnodepool

import (
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/output"
	"github.com/georgetaylor/spotctl/pkg/pager"
	"github.com/spf13/viper"
)

// getSpotNodePoolTableConfig returns the table configuration for spot node pools
// Uses DetailCols to provide additional information when showDetails is true
func getSpotNodePoolTableConfig() *output.TableConfig {
	return &output.TableConfig{
		Columns: []output.TableColumn{
			{Header: "NAME", Field: "metadata.name"},
			{Header: "NAMESPACE", Field: "metadata.namespace"},
			{Header: "SERVER CLASS", Field: "spec.serverClass", Default: "<none>"},
			{Header: "DESIRED", Field: "spec.desired", Default: "<none>"},
			{Header: "BID STATUS", Field: "status.bidStatus", Default: "<none>"},
			{Header: "WON COUNT", Field: "status.wonCount", Default: "<none>"},
		},
		DetailCols: []output.TableColumn{
			{Header: "CLOUD SPACE", Field: "spec.cloudSpace", Default: "<none>"},
			{Header: "BID PRICE", Field: "spec.bidPrice", Default: "<none>"},
			{Header: "AUTOSCALING", Field: "spec.autoscaling.enabled", Default: "<none>"},
			{Header: "MIN NODES", Field: "spec.autoscaling.minNodes", Default: "<none>"},
			{Header: "MAX NODES", Field: "spec.autoscaling.maxNodes", Default: "<none>"},
		},
	}
}

// outputSpotNodePool handles formatting and output of a single spot node pool
func outputSpotNodePool(spotNodePool *client.SpotNodePool, format string) error {
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

		// Use standard table config
		tableConfig := getSpotNodePoolTableConfig()
		return formatter.Output(spotNodePool, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Use standard table config
	tableConfig := getSpotNodePoolTableConfig()
	return formatter.Output(spotNodePool, tableConfig)
}

// outputSpotNodePools handles formatting and output of spot node pool lists
func outputSpotNodePools(spotNodePoolList *client.SpotNodePoolList, format string, namespace string) error {
	// Check if no spot node pools were found
	if len(spotNodePoolList.Items) == 0 {
		if format == "json" {
			fmt.Println("[]")
			return nil
		} else if format == "yaml" {
			fmt.Println("[]")
			return nil
		} else {
			// For table format, show a helpful message
			fmt.Printf("No spot node pools found in namespace %s\n", namespace)
			return nil
		}
	}

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

		// Get table configuration for spot node pools
		tableConfig := getSpotNodePoolTableConfig()

		// Pass the spot node pools array directly instead of the full SpotNodePoolList
		return formatter.Output(spotNodePoolList.Items, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Get table configuration for spot node pools
	tableConfig := getSpotNodePoolTableConfig()

	// Pass the spot node pools array directly instead of the full SpotNodePoolList
	return formatter.Output(spotNodePoolList.Items, tableConfig)
}
