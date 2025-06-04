package organizations

import (
	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/output"
	"github.com/georgetaylor/spotctl/pkg/pager"
	"github.com/spf13/viper"
)

// getOrganizationsTableConfig returns the table configuration for organizations
func getOrganizationsTableConfig() *output.TableConfig {
	return &output.TableConfig{
		Columns: []output.TableColumn{
			{Header: "ID", Field: "id"},
			{Header: "NAME", Field: "name"},
			{Header: "DISPLAY NAME", Field: "display_name"},
		},
		DetailCols: []output.TableColumn{
			{Header: "NAMESPACE", Field: "metadata.namespace"},
		},
	}
}

// outputOrganizations handles formatting and output of organization lists
func outputOrganizations(orgList *client.OrganizationList, format string, showDetails bool) error {
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

		// Get table configuration for organizations
		tableConfig := getOrganizationsTableConfig()

		// Pass the organizations array directly instead of the full OrganizationList
		return formatter.Output(orgList.Organizations, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Get table configuration for organizations
	tableConfig := getOrganizationsTableConfig()

	// Pass the organizations array directly instead of the full OrganizationList
	return formatter.Output(orgList.Organizations, tableConfig)
}
