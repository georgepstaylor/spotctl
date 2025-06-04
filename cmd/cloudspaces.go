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

// cloudspacesDeleteCmd represents the cloudspaces delete command
var cloudspacesDeleteCmd = &cobra.Command{
	Use:   "delete [NAME]",
	Short: "Delete a cloudspace",
	Long: `Delete a cloudspace by name in the specified namespace.

Examples:
  # Delete a cloudspace
  spotctl cloudspaces delete my-cloudspace --namespace org-abc123

  # Delete with confirmation
  spotctl cloudspaces delete my-cloudspace --namespace org-abc123 --confirm`,
	Args: cobra.ExactArgs(1),
	RunE: runCloudspacesDelete,
}

// cloudspacesGetCmd represents the cloudspaces get command
var cloudspacesGetCmd = &cobra.Command{
	Use:   "get [NAME]",
	Short: "Get a specific cloudspace",
	Long: `Get detailed information about a specific cloudspace by name in the specified namespace.

Examples:
  # Get a specific cloudspace
  spotctl cloudspaces get my-cloudspace --namespace org-abc123

  # Get cloudspace with JSON output
  spotctl cloudspaces get my-cloudspace --namespace org-abc123 --output json

  # Get cloudspace with YAML output
  spotctl cloudspaces get my-cloudspace --namespace org-abc123 --output yaml`,
	Args: cobra.ExactArgs(1),
	RunE: runCloudspacesGet,
}


func runCloudspacesList(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	if namespace == "" {
		return fmt.Errorf("namespace is required")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	cloudSpaceList, err := client.ListCloudSpaces(ctx, namespace)
	if err != nil {
		return fmt.Errorf("failed to list cloudspaces: %w", err)
	}

	outputFormat, _ := cmd.Flags().GetString("output")
	showDetails, _ := cmd.Flags().GetBool("details")
	wideOutput, _ := cmd.Flags().GetBool("wide")

	return outputCloudSpaces(cloudSpaceList, outputFormat, showDetails, wideOutput, namespace)
}

func outputCloudSpaces(cloudSpaceList *client.CloudSpaceList, format string, showDetails bool, wideOutput bool, namespace string) error {
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

	// Get flag values
	namespace, _ := cmd.Flags().GetString("namespace")
	region, _ := cmd.Flags().GetString("region")
	kubernetesVersion, _ := cmd.Flags().GetString("kubernetes-version")
	webhook, _ := cmd.Flags().GetString("webhook")
	haControlPlane, _ := cmd.Flags().GetBool("ha-control-plane")
	cni, _ := cmd.Flags().GetString("cni")
	outputFormat, _ := cmd.Flags().GetString("output")

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
			Namespace: namespace, // Use namespace flag
		},
		Spec: client.CloudSpaceSpec{
			Region:            region,
			KubernetesVersion: kubernetesVersion,
			Webhook:           webhook,
			HAControlPlane:    haControlPlane,
			Cloud:             "default", // API requires this to be set to "default"
			CNI:               cni,
		},
	}

	// Validate required fields
	if cloudspaceName == "" {
		return fmt.Errorf("cloudspace name is required (use positional argument)")
	}
	if namespace == "" {
		return fmt.Errorf("namespace is required (use --namespace flag)")
	}
	if region == "" {
		return fmt.Errorf("region is required (use --region flag)")
	}
	if kubernetesVersion == "" {
		return fmt.Errorf("kubernetes version is required (use --kubernetes-version flag)")
	}

	ctx := context.Background()
	createdCloudSpace, err := apiClient.CreateCloudSpace(ctx, namespace, cloudSpace)
	if err != nil {
		return fmt.Errorf("failed to create cloudspace: %w", err)
	}

	// Output the created cloudspace
	return outputCreatedCloudSpace(createdCloudSpace, outputFormat)
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

func runCloudspacesDelete(cmd *cobra.Command, args []string) error {
	cloudspaceName := args[0] // Get name from positional argument

	// Get flag values
	namespace, _ := cmd.Flags().GetString("namespace")
	confirm, _ := cmd.Flags().GetBool("confirm")

	if namespace == "" {
		return fmt.Errorf("namespace is required")
	}

	// Ask for confirmation unless --confirm flag is used
	if !confirm {
		fmt.Printf("Are you sure you want to delete cloudspace '%s' in namespace '%s'? (y/N): ", cloudspaceName, namespace)
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" && response != "yes" && response != "Yes" {
			fmt.Println("Delete cancelled")
			return nil
		}
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	apiClient := client.NewClient(cfg)

	ctx := context.Background()
	deleteResponse, err := apiClient.DeleteCloudSpace(ctx, namespace, cloudspaceName)
	if err != nil {
		return fmt.Errorf("failed to delete cloudspace: %w", err)
	}

	// Check if the deletion was successful
	if deleteResponse.Status == "Success" || deleteResponse.Status == "" {
		fmt.Printf("Cloudspace '%s' deleted successfully from namespace '%s'\n", cloudspaceName, namespace)
	} else {
		fmt.Printf("Delete operation completed with status: %s\n", deleteResponse.Status)
		if deleteResponse.Message != "" {
			fmt.Printf("Message: %s\n", deleteResponse.Message)
		}
	}
	return nil
}

func runCloudspacesGet(cmd *cobra.Command, args []string) error {
	cloudspaceName := args[0] // Get name from positional argument

	// Get flag values
	namespace, _ := cmd.Flags().GetString("namespace")
	outputFormat, _ := cmd.Flags().GetString("output")

	if namespace == "" {
		return fmt.Errorf("namespace is required")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := client.NewClient(cfg)

	ctx := context.Background()
	cloudSpace, err := client.GetCloudSpace(ctx, namespace, cloudspaceName)
	if err != nil {
		return fmt.Errorf("failed to get cloudspace: %w", err)
	}

	return outputCloudSpace(cloudSpace, outputFormat)
}

func outputCloudSpace(cloudSpace *client.CloudSpace, format string) error {
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
		tableConfig := getCloudSpaceTableConfig()
		return formatter.Output(cloudSpace, tableConfig)
	}

	formatter := output.NewFormatter(options)

	// For single cloudspace output, we'll use a simplified table config
	tableConfig := getCloudSpaceTableConfig()
	return formatter.Output(cloudSpace, tableConfig)
}

func getCloudSpaceTableConfig() *output.TableConfig {
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
	cloudspacesCmd.AddCommand(cloudspacesDeleteCmd)
	cloudspacesCmd.AddCommand(cloudspacesGetCmd)

	// Add flags for cloudspaces list command
	cloudspacesListCmd.Flags().StringP("namespace", "n", "", "Namespace to list cloudspaces from (required)")
	cloudspacesListCmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	cloudspacesListCmd.Flags().BoolP("details", "d", false, "Show detailed output")
	cloudspacesListCmd.Flags().BoolP("wide", "w", false, "Show additional columns")

	// Add flags for cloudspaces create command
	cloudspacesCreateCmd.Flags().String("namespace", "", "Namespace for the cloudspace (required)")
	cloudspacesCreateCmd.Flags().StringP("region", "r", "", "Region for the cloudspace (required)")
	cloudspacesCreateCmd.Flags().StringP("kubernetes-version", "k", "1.31.1", "Kubernetes version (1.31.1, 1.30.10, 1.29.6)")
	cloudspacesCreateCmd.Flags().StringP("webhook", "w", "", "Webhook URL for notifications")
	cloudspacesCreateCmd.Flags().Bool("ha-control-plane", false, "Enable HA control plane")
	cloudspacesCreateCmd.Flags().String("cni", "", "CNI plugin (leave as default unless custom values needed)")
	cloudspacesCreateCmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")

	// Add flags for cloudspaces delete command
	cloudspacesDeleteCmd.Flags().String("namespace", "", "Namespace for the cloudspace (required)")
	cloudspacesDeleteCmd.Flags().Bool("confirm", false, "Skip confirmation prompt")

	// Add flags for cloudspaces get command
	cloudspacesGetCmd.Flags().String("namespace", "", "Namespace for the cloudspace (required)")
	cloudspacesGetCmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")

	// Mark required flags
	cloudspacesListCmd.MarkFlagRequired("namespace")
	cloudspacesCreateCmd.MarkFlagRequired("namespace")
	cloudspacesCreateCmd.MarkFlagRequired("region")
	cloudspacesDeleteCmd.MarkFlagRequired("namespace")
	cloudspacesGetCmd.MarkFlagRequired("namespace")
}
