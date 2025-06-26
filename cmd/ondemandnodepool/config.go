package ondemandnodepools

import (
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/georgetaylor/spotctl/pkg/output"
	"github.com/georgetaylor/spotctl/pkg/pager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getSpotNodePoolTableConfig returns the table configuration for spot node pools
// Uses DetailCols to provide additional information when showDetails is true
func getOnDemandNodePoolTableConfig() *output.TableConfig {
	return &output.TableConfig{
		Columns: []output.TableColumn{
			{Header: "NAME", Field: "metadata.name"},
			{Header: "NAMESPACE", Field: "metadata.namespace"},
			{Header: "SERVER CLASS", Field: "spec.serverClass", Default: "<none>"},
			{Header: "DESIRED", Field: "spec.desired", Default: "<none>"},
		},
		DetailCols: []output.TableColumn{
			{Header: "CLOUD SPACE", Field: "spec.cloudSpace", Default: "<none>"},
		},
	}
}

// outputOnDemandNodePool handles formatting and output of a single on demand node pool
func outputOnDemandNodePool(onDemandNodePool *client.OnDemandNodePool, format string) error {
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
		tableConfig := getOnDemandNodePoolTableConfig()
		return formatter.Output(onDemandNodePool, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Use standard table config
	tableConfig := getOnDemandNodePoolTableConfig()
	return formatter.Output(onDemandNodePool, tableConfig)
}

// outputOnDemandNodePools handles formatting and output of on demand node pool lists
func outputOnDemandNodePools(onDemandNodePoolList *client.OnDemandNodePoolList, format string, namespace string) error {
	// Check if no on demand node pools were found
	if len(onDemandNodePoolList.Items) == 0 {
		if format == "json" {
			fmt.Println("[]")
			return nil
		} else if format == "yaml" {
			fmt.Println("[]")
			return nil
		} else {
			// For table format, show a helpful message
			fmt.Printf("No on demand node pools found in namespace %s\n", namespace)
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
		tableConfig := getOnDemandNodePoolTableConfig()

		// Pass the on demand node pools array directly instead of the full OnDemandNodePoolList
		return formatter.Output(onDemandNodePoolList.Items, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Get table configuration for spot node pools
	tableConfig := getOnDemandNodePoolTableConfig()

	// Pass the on demand node pools array directly instead of the full OnDemandNodePoolList
	return formatter.Output(onDemandNodePoolList.Items, tableConfig)
}

// outputCreatedOnDemandNodePool handles formatting and output of a newly created on demand node pool
func outputCreatedOnDemandNodePool(onDemandNodePool *client.OnDemandNodePool, format string) error {
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

		// Use standard table config for created on demand node pools
		tableConfig := getOnDemandNodePoolTableConfig()
		return formatter.Output(onDemandNodePool, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// Use standard table config for created on demand node pools
	tableConfig := getOnDemandNodePoolTableConfig()
	return formatter.Output(onDemandNodePool, tableConfig)
}

// getNamespace resolves the namespace to use, with flag taking precedence over config
func getNamespace(cmd *cobra.Command) (string, error) {
	// Check if namespace was provided via flag
	namespace, _ := cmd.Flags().GetString("namespace")
	if namespace != "" {
		return namespace, nil
	}

	// Fall back to config namespace
	cfg, err := config.GetConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	if cfg.Namespace != "" {
		return cfg.Namespace, nil
	}

	// No namespace configured
	return "", fmt.Errorf("namespace is required: set it via --namespace flag, config file, or SPOTCTL_NAMESPACE environment variable")
}
