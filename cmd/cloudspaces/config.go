package cloudspaces

import (
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/output"
	"github.com/georgetaylor/spotctl/pkg/pager"
	"github.com/spf13/viper"
)

// getCloudSpacesTableConfig returns the table configuration for cloudspaces
// Uses DetailCols to provide additional information when showDetails is true
func getCloudSpacesTableConfig() *output.TableConfig {
	return &output.TableConfig{
		Columns: []output.TableColumn{
			{Header: "NAME", Field: "metadata.name"},
			{Header: "NAMESPACE", Field: "metadata.namespace"},
			{Header: "REGION", Field: "spec.region"},
			{Header: "PHASE", Field: "status.phase", Default: "<none>"},
			{Header: "HEALTH", Field: "status.health", Default: "<none>"},
		},
		DetailCols: []output.TableColumn{
			{Header: "K8S VERSION", Field: "status.currentKubernetesVersion", Default: "<none>"},
			{Header: "CNI", Field: "spec.cni", Default: "<none>"},
			{Header: "DEPLOYMENT TYPE", Field: "spec.deploymentType", Default: "<none>"},
			{Header: "HA CONTROL PLANE", Field: "spec.HAControlPlane", Default: "<none>"},
		},
	}
}

// outputCloudSpaces handles formatting and output of cloudspace lists
func outputCloudSpaces(cloudSpaceList *client.CloudSpaceList, format string, showDetails bool, namespace string) error {
	// Check if no cloudspaces were found
	if len(cloudSpaceList.Items) == 0 {
		if format == "json" {
			fmt.Println("[]")
			return nil
		} else if format == "yaml" {
			fmt.Println("[]")
			return nil
		} else {
			// For table format, show a helpful message
			fmt.Printf("No cloudspaces found in namespace %s\n", namespace)
			return nil
		}
	}

	// Create formatter with options
	options := output.OutputOptions{
		Format:      output.OutputFormat(format),
		ShowDetails: showDetails,
	}

	// Check if pager should be disabled
	noPager := viper.GetBool("no-pager")
	if noPager {
		// Create pager with disabled setting
		pager := pager.NewPager()
		pager.Disable = true
		formatter := output.NewFormatterWithPager(options, pager)

		// Get table configuration for cloudspaces
		tableConfig := getCloudSpacesTableConfig()

		// Pass the cloudspaces array directly instead of the full CloudSpaceList
		return formatter.Output(cloudSpaceList.Items, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Get table configuration for cloudspaces
	tableConfig := getCloudSpacesTableConfig()

	// Pass the cloudspaces array directly instead of the full CloudSpaceList
	return formatter.Output(cloudSpaceList.Items, tableConfig)
}

// outputCreatedCloudSpace handles formatting and output of a newly created cloudspace
func outputCreatedCloudSpace(cloudSpace *client.CloudSpace, format string) error {
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

		// Use standard table config for created cloudspaces
		tableConfig := getCloudSpacesTableConfig()
		return formatter.Output(cloudSpace, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Use standard table config for created cloudspaces
	tableConfig := getCloudSpacesTableConfig()
	return formatter.Output(cloudSpace, tableConfig)
}

// outputCloudSpace handles formatting and output of a single cloudspace
func outputCloudSpace(cloudSpace *client.CloudSpace, format string, showDetails bool) error {
	// Create formatter with options
	options := output.OutputOptions{
		Format:      output.OutputFormat(format),
		ShowDetails: showDetails,
	}

	// Check if pager should be disabled
	noPager := viper.GetBool("no-pager")
	if noPager {
		// Create pager with disabled setting
		pager := pager.NewPager()
		pager.Disable = true
		formatter := output.NewFormatterWithPager(options, pager)

		// Use standard table config, details will be shown based on ShowDetails option
		tableConfig := getCloudSpacesTableConfig()
		return formatter.Output(cloudSpace, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Use standard table config, details will be shown based on ShowDetails option
	tableConfig := getCloudSpacesTableConfig()
	return formatter.Output(cloudSpace, tableConfig)
}
