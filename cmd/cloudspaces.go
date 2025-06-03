package cmd

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/georgetaylor/spotctl/pkg/output"
	"github.com/georgetaylor/spotctl/pkg/pager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloudspacesCmd represents the cloudspaces command
var cloudspacesCmd = &cobra.Command{
	Use:   "cloudspaces",
	Short: "Manage Rackspace Spot cloudspaces",
	Long: `Manage and view Rackspace Spot cloudspaces.

This command allows you to list cloudspaces within a specific namespace.
Cloudspaces represent Kubernetes clusters deployed through Rackspace Spot.`,
}

// cloudspacesListCmd represents the cloudspaces list command
var cloudspacesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List cloudspaces in a namespace",
	Long: `List all cloudspaces in the specified namespace.

Examples:
  # List cloudspaces in a specific namespace
  spotctl cloudspaces list --namespace my-namespace

  # List cloudspaces with wide output
  spotctl cloudspaces list --namespace my-namespace --wide

  # List cloudspaces with JSON output
  spotctl cloudspaces list --namespace my-namespace --output json`,
	Args: cobra.NoArgs,
	RunE: runCloudspacesList,
}

// cloudspacesCreateCmd represents the cloudspaces create command
var cloudspacesCreateCmd = &cobra.Command{
	Use:   "create [NAME]",
	Short: "Create a new cloudspace",
	Long: `Create a new cloudspace in the specified namespace.

Examples:
  # Create a cloudspace with required parameters
  spotctl cloudspaces create my-cloudspace --namespace org-abc123 --region uk-lon-1 --kubernetes-version 1.31.1

  # Create a cloudspace with webhook
  spotctl cloudspaces create my-cloudspace --namespace org-abc123 --region uk-lon-1 --webhook https://hooks.slack.com/services/...

  # Create a cloudspace with HA control plane
  spotctl cloudspaces create my-cloudspace --namespace org-abc123 --region uk-lon-1 --ha-control-plane`,
	Args: cobra.ExactArgs(1),
	RunE: runCloudspacesCreate,
}

var (
	cloudspacesOutputFormat string
	cloudspacesShowDetails  bool
	cloudspacesWideOutput   bool
	cloudspacesNamespace    string // For list command
	// Flags for cloudspaces create command
	cloudspaceNamespace         string
	cloudspaceRegion            string
	cloudspaceKubernetesVersion string
	cloudspaceWebhook           string
	cloudspaceHAControlPlane    bool
	cloudspaceCNI               string
)

func runCloudspacesList(cmd *cobra.Command, args []string) error {
	if cloudspacesNamespace == "" {
		return fmt.Errorf("namespace is required")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	cloudSpaceList, err := client.ListCloudSpaces(ctx, cloudspacesNamespace)
	if err != nil {
		return fmt.Errorf("failed to list cloudspaces: %w", err)
	}

	return outputCloudSpaces(cloudSpaceList, cloudspacesOutputFormat, cloudspacesShowDetails, cloudspacesWideOutput)
}

func outputCloudSpaces(cloudSpaceList *client.CloudSpaceList, format string, showDetails bool, wideOutput bool) error {
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
			fmt.Printf("No cloudspaces found in namespace %s\n", cloudspacesNamespace)
			return nil
		}
	}

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

func runCloudspacesCreate(cmd *cobra.Command, args []string) error {
	cloudspaceName := args[0] // Get name from positional argument

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	apiClient := client.NewClient(cfg)

	// Create the CloudSpace object
	cloudSpace := &client.CloudSpace{
		APIVersion: "ngpc.rxt.io/v1",
		Kind:       "CloudSpace",
		Metadata: client.ObjectMeta{
			Name:      cloudspaceName,
			Namespace: cloudspaceNamespace, // Use namespace flag
		},
		Spec: client.CloudSpaceSpec{
			Region:            cloudspaceRegion,
			KubernetesVersion: cloudspaceKubernetesVersion,
			Webhook:           cloudspaceWebhook,
			HAControlPlane:    cloudspaceHAControlPlane,
			Cloud:             "default", // API requires this to be set to "default"
			CNI:               cloudspaceCNI,
		},
	}

	// Validate required fields
	if cloudspaceName == "" {
		return fmt.Errorf("cloudspace name is required (use positional argument)")
	}
	if cloudspaceNamespace == "" {
		return fmt.Errorf("namespace is required (use --namespace flag)")
	}
	if cloudspaceRegion == "" {
		return fmt.Errorf("region is required (use --region flag)")
	}
	if cloudspaceKubernetesVersion == "" {
		return fmt.Errorf("kubernetes version is required (use --kubernetes-version flag)")
	}

	ctx := context.Background()
	createdCloudSpace, err := apiClient.CreateCloudSpace(ctx, cloudspaceNamespace, cloudSpace)
	if err != nil {
		return fmt.Errorf("failed to create cloudspace: %w", err)
	}

	// Output the created cloudspace
	return outputCreatedCloudSpace(createdCloudSpace, cloudspacesOutputFormat)
}

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

		// For single cloudspace output, we'll use a simplified table config
		tableConfig := getCreatedCloudSpaceTableConfig()
		return formatter.Output(cloudSpace, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// For single cloudspace output, we'll use a simplified table config
	tableConfig := getCreatedCloudSpaceTableConfig()
	return formatter.Output(cloudSpace, tableConfig)
}

func getCreatedCloudSpaceTableConfig() *output.TableConfig {
	return &output.TableConfig{
		Columns: []output.TableColumn{
			{Header: "NAME", Field: "metadata.name"},
			{Header: "NAMESPACE", Field: "metadata.namespace"},
			{Header: "REGION", Field: "spec.region"},
			{Header: "K8S VERSION", Field: "spec.kubernetesVersion", Default: "<none>"},
			{Header: "PHASE", Field: "status.phase", Default: "<none>"},
			{Header: "HEALTH", Field: "status.health", Default: "<none>"},
		},
	}
}

func init() {
	rootCmd.AddCommand(cloudspacesCmd)
	cloudspacesCmd.AddCommand(cloudspacesListCmd)
	cloudspacesCmd.AddCommand(cloudspacesCreateCmd)

	// Add flags for cloudspaces list command
	cloudspacesListCmd.Flags().StringVarP(&cloudspacesNamespace, "namespace", "n", "", "Namespace to list cloudspaces from (required)")
	cloudspacesListCmd.Flags().StringVarP(&cloudspacesOutputFormat, "output", "o", "table", "Output format (table, json, yaml)")
	cloudspacesListCmd.Flags().BoolVarP(&cloudspacesShowDetails, "details", "d", false, "Show detailed output")
	cloudspacesListCmd.Flags().BoolVarP(&cloudspacesWideOutput, "wide", "w", false, "Show additional columns")

	// Add flags for cloudspaces create command
	cloudspacesCreateCmd.Flags().StringVar(&cloudspaceNamespace, "namespace", "", "Namespace for the cloudspace (required)")
	cloudspacesCreateCmd.Flags().StringVarP(&cloudspaceRegion, "region", "r", "", "Region for the cloudspace (required)")
	cloudspacesCreateCmd.Flags().StringVarP(&cloudspaceKubernetesVersion, "kubernetes-version", "k", "1.31.1", "Kubernetes version (1.31.1, 1.30.10, 1.29.6)")
	cloudspacesCreateCmd.Flags().StringVarP(&cloudspaceWebhook, "webhook", "w", "", "Webhook URL for notifications")
	cloudspacesCreateCmd.Flags().BoolVar(&cloudspaceHAControlPlane, "ha-control-plane", false, "Enable HA control plane")
	cloudspacesCreateCmd.Flags().StringVar(&cloudspaceCNI, "cni", "", "CNI plugin (leave as default unless custom values needed)")
	cloudspacesCreateCmd.Flags().StringVarP(&cloudspacesOutputFormat, "output", "o", "table", "Output format (table, json, yaml)")

	// Mark required flags
	cloudspacesListCmd.MarkFlagRequired("namespace")
	cloudspacesCreateCmd.MarkFlagRequired("namespace")
	cloudspacesCreateCmd.MarkFlagRequired("region")
}
