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
