package spotnodepool

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/georgetaylor/spotctl/pkg/client"
	"github.com/georgetaylor/spotctl/pkg/config"
	"github.com/spf13/cobra"
)

// NewCreateCommand returns the spotnodepool create command
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [NAME]",
		Short: "Create a new spot node pool",
		Long: `Create a new spot node pool in the specified namespace.

Examples:
  # Create a spot node pool with required parameters
  spotctl spotnodepool create my-nodepool --namespace org-abc123 --server-class gp.vs1.large-lon --cloudspace my-cloudspace --desired 3

  # Create a spot node pool with autoscaling
  spotctl spotnodepool create my-nodepool --namespace org-abc123 --server-class gp.vs1.large-lon --cloudspace my-cloudspace --autoscaling --autoscaling-min-nodes 1 --autoscaling-max-nodes 10

  # Create a spot node pool with bid price
  spotctl spotnodepool create my-nodepool --namespace org-abc123 --server-class gp.vs1.large-lon --cloudspace my-cloudspace --desired 3 --bid-price 0.50

  # Create a spot node pool from a spec file
  spotctl spotnodepool create my-nodepool --namespace org-abc123 --file spec.json`,
		Args: cobra.ExactArgs(1),
		RunE: runCreate,
	}

	// Add flags for spotnodepool create command
	cmd.Flags().StringP("output", "o", "table", "Output format (table, json, yaml)")
	cmd.Flags().StringP("namespace", "n", "", "Namespace to create the spot node pool in (required)")
	cmd.Flags().StringP("file", "f", "", "Path to JSON file containing spot node pool spec")
	cmd.Flags().String("server-class", "", "Server class for the spot node pool (required unless using --file)")
	cmd.Flags().String("cloudspace", "", "Cloud space for the spot node pool (required unless using --file)")
	cmd.Flags().Int("desired", 0, "Desired number of nodes (required unless using --file)")
	cmd.Flags().Bool("autoscaling", false, "Enable autoscaling")
	cmd.Flags().Int("autoscaling-min-nodes", 0, "Minimum number of nodes for autoscaling")
	cmd.Flags().Int("autoscaling-max-nodes", 0, "Maximum number of nodes for autoscaling")
	cmd.Flags().String("bid-price", "", "Bid price for spot instances")

	return cmd
}

func runCreate(cmd *cobra.Command, args []string) error {
	spotNodePoolName := args[0] // Get name from positional argument

	// Get flag values
	namespace, _ := cmd.Flags().GetString("namespace")
	file, _ := cmd.Flags().GetString("file")
	serverClass, _ := cmd.Flags().GetString("server-class")
	cloudSpace, _ := cmd.Flags().GetString("cloudspace")
	desired, _ := cmd.Flags().GetInt("desired")
	autoscaling, _ := cmd.Flags().GetBool("autoscaling")
	autoscalingMinNodes, _ := cmd.Flags().GetInt("autoscaling-min-nodes")
	autoscalingMaxNodes, _ := cmd.Flags().GetInt("autoscaling-max-nodes")
	bidPrice, _ := cmd.Flags().GetString("bid-price")
	outputFormat, _ := cmd.Flags().GetString("output")

	// Validate required fields
	if spotNodePoolName == "" {
		return fmt.Errorf("spot node pool name is required (use positional argument)")
	}
	if namespace == "" {
		return fmt.Errorf("namespace is required (use --namespace flag)")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	apiClient := client.NewClient(cfg)

	var spotNodePool *client.SpotNodePool

	if file != "" {
		// Load spec from file
		spec, err := loadSpecFromFile(file)
		if err != nil {
			return fmt.Errorf("failed to load spec from file: %w", err)
		}

		spotNodePool = &client.SpotNodePool{
			APIVersion: "ngpc.rxt.io/v1",
			Kind:       "SpotNodePool",
			Metadata: client.ObjectMeta{
				Name:      spotNodePoolName,
				Namespace: namespace,
			},
			Spec: *spec,
		}
	} else {
		// Build from flags
		if serverClass == "" {
			return fmt.Errorf("server-class is required (use --server-class flag or --file)")
		}
		if cloudSpace == "" {
			return fmt.Errorf("cloudspace is required (use --cloudspace flag or --file)")
		}
		if desired == 0 {
			return fmt.Errorf("desired is required and must be greater than 0 (use --desired flag or --file)")
		}

		spec := client.SpotNodePoolSpec{
			ServerClass: serverClass,
			CloudSpace:  cloudSpace,
			Desired:     &desired,
		}

		if bidPrice != "" {
			spec.BidPrice = bidPrice
		}

		if autoscaling {
			spec.Autoscaling = &client.SpotNodePoolAutoscaling{
				Enabled: true,
			}
			if autoscalingMinNodes > 0 {
				spec.Autoscaling.MinNodes = &autoscalingMinNodes
			}
			if autoscalingMaxNodes > 0 {
				spec.Autoscaling.MaxNodes = &autoscalingMaxNodes
			}
		}

		spotNodePool = &client.SpotNodePool{
			APIVersion: "ngpc.rxt.io/v1",
			Kind:       "SpotNodePool",
			Metadata: client.ObjectMeta{
				Name:      spotNodePoolName,
				Namespace: namespace,
			},
			Spec: spec,
		}
	}

	ctx := context.Background()
	createdSpotNodePool, err := apiClient.CreateSpotNodePool(ctx, namespace, spotNodePool)
	if err != nil {
		return fmt.Errorf("failed to create spot node pool: %w", err)
	}

	// Output the created spot node pool
	return outputCreatedSpotNodePool(createdSpotNodePool, outputFormat)
}

// loadSpecFromFile loads a SpotNodePoolSpec from a JSON file
func loadSpecFromFile(filename string) (*client.SpotNodePoolSpec, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var spec client.SpotNodePoolSpec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &spec, nil
}
