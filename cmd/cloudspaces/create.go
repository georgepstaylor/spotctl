package cloudspaces

import (
	"context"
	"fmt"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewCreateCommand returns the cloudspaces create command
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
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
		RunE: runCreate,
	}

	// Add flags for cloudspaces create command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	cmd.Flags().StringP("namespace", "n", "", "Namespace to create the cloudspace in (required)")
	cmd.Flags().StringP("region", "r", "", "Region to deploy the cloudspace in (required)")
	cmd.Flags().String("kubernetes-version", "", "Kubernetes version for the cloudspace (required)")
	cmd.Flags().String("webhook", "", "Webhook URL for notifications")
	cmd.Flags().Bool("ha-control-plane", false, "Enable high availability control plane")
	cmd.Flags().String("cni", "cilium", "Container Network Interface (CNI) to use")

	return cmd
}

func runCreate(cmd *cobra.Command, args []string) error {
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
