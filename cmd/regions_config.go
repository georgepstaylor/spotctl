package cmd

import (
	"github.com/georgetaylor/spotctl/pkg/output"
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
