package client

import "time"

// Standard Kubernetes-style metadata
type ObjectMeta struct {
	Annotations                map[string]string    `json:"annotations,omitempty"`
	CreationTimestamp          *time.Time           `json:"creationTimestamp,omitempty"`
	DeletionGracePeriodSeconds *int64               `json:"deletionGracePeriodSeconds,omitempty"`
	DeletionTimestamp          *time.Time           `json:"deletionTimestamp,omitempty"`
	Finalizers                 []string             `json:"finalizers,omitempty"`
	GenerateName               string               `json:"generateName,omitempty"`
	Generation                 *int64               `json:"generation,omitempty"`
	Labels                     map[string]string    `json:"labels,omitempty"`
	ManagedFields              []ManagedFieldsEntry `json:"managedFields,omitempty"`
	Name                       string               `json:"name,omitempty"`
	Namespace                  string               `json:"namespace,omitempty"`
	OwnerReferences            []OwnerReference     `json:"ownerReferences,omitempty"`
	ResourceVersion            string               `json:"resourceVersion,omitempty"`
	SelfLink                   string               `json:"selfLink,omitempty"`
	UID                        string               `json:"uid,omitempty"`
}

type ManagedFieldsEntry struct {
	APIVersion  string                 `json:"apiVersion,omitempty"`
	FieldsType  string                 `json:"fieldsType,omitempty"`
	FieldsV1    map[string]interface{} `json:"fieldsV1,omitempty"`
	Manager     string                 `json:"manager,omitempty"`
	Operation   string                 `json:"operation,omitempty"`
	Subresource string                 `json:"subresource,omitempty"`
	Time        *time.Time             `json:"time,omitempty"`
}

type OwnerReference struct {
	APIVersion         string `json:"apiVersion"`
	BlockOwnerDeletion *bool  `json:"blockOwnerDeletion,omitempty"`
	Controller         *bool  `json:"controller,omitempty"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	UID                string `json:"uid"`
}

// List metadata for Kubernetes-style lists
type ListMeta struct {
	Continue           string `json:"continue,omitempty"`
	RemainingItemCount *int64 `json:"remainingItemCount,omitempty"`
	ResourceVersion    string `json:"resourceVersion,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
}

// Region represents a Rackspace Spot region
type Region struct {
	APIVersion string     `json:"apiVersion,omitempty"`
	Kind       string     `json:"kind,omitempty"`
	Metadata   ObjectMeta `json:"metadata,omitempty"`
	Spec       RegionSpec `json:"spec,omitempty"`
}

type RegionSpec struct {
	Country     string         `json:"country,omitempty"`
	Description string         `json:"description,omitempty"`
	Provider    RegionProvider `json:"provider,omitempty"`
}

type RegionProvider struct {
	ProviderRegionName string `json:"providerRegionName,omitempty"`
	ProviderType       string `json:"providerType,omitempty"`
}

// RegionList represents a list of regions
type RegionList struct {
	APIVersion string   `json:"apiVersion,omitempty"`
	Items      []Region `json:"items"`
	Kind       string   `json:"kind,omitempty"`
	Metadata   ListMeta `json:"metadata,omitempty"`
}

// ServerClass represents a Rackspace Spot server class
type ServerClass struct {
	APIVersion string            `json:"apiVersion,omitempty"`
	Kind       string            `json:"kind,omitempty"`
	Metadata   ObjectMeta        `json:"metadata,omitempty"`
	Spec       ServerClassSpec   `json:"spec,omitempty"`
	Status     ServerClassStatus `json:"status,omitempty"`
}

type ServerClassSpec struct {
	Availability    string               `json:"availability,omitempty"`
	Category        string               `json:"category,omitempty"`
	DisplayName     string               `json:"displayName,omitempty"`
	FlavorType      string               `json:"flavorType,omitempty"`
	OnDemandPricing ServerClassPricing   `json:"onDemandPricing,omitempty"`
	Provider        ServerClassProvider  `json:"provider,omitempty"`
	Region          string               `json:"region,omitempty"`
	Resources       ServerClassResources `json:"resources,omitempty"`
}

type ServerClassPricing struct {
	Cost     string `json:"cost,omitempty"`
	Interval string `json:"interval,omitempty"`
}

type ServerClassProvider struct {
	ProviderFlavorID string `json:"providerFlavorID,omitempty"`
	ProviderType     string `json:"providerType,omitempty"`
}

type ServerClassResources struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type ServerClassStatus struct {
	Available   *int                   `json:"available,omitempty"`
	Capacity    *int                   `json:"capacity,omitempty"`
	LastAuction *int                   `json:"lastAuction,omitempty"`
	Reserved    *int                   `json:"reserved,omitempty"`
	SpotPricing ServerClassSpotPricing `json:"spotPricing,omitempty"`
}

type ServerClassSpotPricing struct {
	HammerPricePerHour string `json:"hammerPricePerHour,omitempty"`
	MarketPricePerHour string `json:"marketPricePerHour,omitempty"`
}

// ServerClassList represents a list of server classes
type ServerClassList struct {
	APIVersion string        `json:"apiVersion,omitempty"`
	Items      []ServerClass `json:"items"`
	Kind       string        `json:"kind,omitempty"`
	Metadata   ListMeta      `json:"metadata,omitempty"`
}

// Organization types - follows a different API pattern than Kubernetes-style
type Organization struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	DisplayName string               `json:"display_name"`
	Metadata    OrganizationMetadata `json:"metadata"`
}

type OrganizationMetadata struct {
	Namespace string `json:"namespace"`
}

// OrganizationList represents the paginated response for organizations
type OrganizationList struct {
	Start         int            `json:"start"`
	Limit         int            `json:"limit"`
	Length        int            `json:"length"`
	Total         int            `json:"total"`
	Next          string         `json:"next,omitempty"`
	Organizations []Organization `json:"organizations"`
}

// CloudSpace represents a Rackspace Spot cloudspace
type CloudSpace struct {
	APIVersion string           `json:"apiVersion,omitempty"`
	Kind       string           `json:"kind,omitempty"`
	Metadata   ObjectMeta       `json:"metadata,omitempty"`
	Spec       CloudSpaceSpec   `json:"spec,omitempty"`
	Status     CloudSpaceStatus `json:"status,omitempty"`
}

type CloudSpaceSpec struct {
	HAControlPlane    bool                `json:"HAControlPlane,omitempty"`
	BidRequests       []string            `json:"bidRequests,omitempty"`
	Cloud             string              `json:"cloud,omitempty"`
	ClusterRef        *ObjectReference    `json:"clusterRef,omitempty"`
	CNI               string              `json:"cni,omitempty"`
	DeploymentType    string              `json:"deploymentType,omitempty"`
	KubernetesVersion string              `json:"kubernetesVersion,omitempty"`
	Networks          []CloudSpaceNetwork `json:"networks,omitempty"`
	OnDemandRequests  []string            `json:"onDemandRequests,omitempty"`
	Region            string              `json:"region,omitempty"`
	Servers           []CloudSpaceServer  `json:"servers,omitempty"`
	Type              string              `json:"type,omitempty"`
	Webhook           string              `json:"webhook,omitempty"`
}

type CloudSpaceNetwork struct {
	Name   string `json:"name,omitempty"`
	Subnet string `json:"subnet,omitempty"`
}

type CloudSpaceServer struct {
	Class string `json:"class,omitempty"`
	Count int    `json:"count,omitempty"`
}

type ObjectReference struct {
	APIVersion      string `json:"apiVersion,omitempty"`
	FieldPath       string `json:"fieldPath,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	UID             string `json:"uid,omitempty"`
}

type CloudSpaceStatus struct {
	APIServerEndpoint        string                 `json:"APIServerEndpoint,omitempty"`
	AssignedServers          map[string]interface{} `json:"assignedServers,omitempty"`
	Bids                     map[string]interface{} `json:"bids,omitempty"`
	CloudspaceClassName      string                 `json:"cloudspaceClassName,omitempty"`
	ClusterRef               *CloudSpaceClusterRef  `json:"clusterRef,omitempty"`
	Conditions               []CloudSpaceCondition  `json:"conditions,omitempty"`
	CurrentKubernetesVersion string                 `json:"currentKubernetesVersion,omitempty"`
	FirstReadyTimestamp      *time.Time             `json:"firstReadyTimestamp,omitempty"`
	Health                   string                 `json:"health,omitempty"`
	PendingAllocations       map[string]interface{} `json:"pendingAllocations,omitempty"`
	Phase                    string                 `json:"phase,omitempty"`
	Reason                   string                 `json:"reason,omitempty"`
	SSHSecretName            string                 `json:"sshSecretName,omitempty"`
	UpgradePhase             string                 `json:"upgradePhase,omitempty"`
}

type CloudSpaceClusterRef struct {
	Cluster *ObjectReference `json:"cluster,omitempty"`
	Reason  string           `json:"reason,omitempty"`
	Status  string           `json:"status,omitempty"`
}

type CloudSpaceCondition struct {
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	Message            string     `json:"message,omitempty"`
	Reason             string     `json:"reason,omitempty"`
	Severity           string     `json:"severity,omitempty"`
	Status             string     `json:"status,omitempty"`
	Type               string     `json:"type,omitempty"`
}

// CloudSpaceList represents a list of cloudspaces
type CloudSpaceList struct {
	APIVersion string       `json:"apiVersion,omitempty"`
	Items      []CloudSpace `json:"items"`
	Kind       string       `json:"kind,omitempty"`
	Metadata   ListMeta     `json:"metadata,omitempty"`
}
