/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CapabilityObservation struct {

	// Deprecated. Use capability block instead. Defines whether a volume should be available concurrently. Possible values are:
	// Defines whether a volume should be available concurrently.
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// Deprecated. Use capability block instead. The storage API that will be used by the volume.
	// The storage API that will be used by the volume.
	AttachmentMode *string `json:"attachmentMode,omitempty" tf:"attachment_mode,omitempty"`
}

type CapabilityParameters struct {

	// Deprecated. Use capability block instead. Defines whether a volume should be available concurrently. Possible values are:
	// Defines whether a volume should be available concurrently.
	// +kubebuilder:validation:Required
	AccessMode *string `json:"accessMode" tf:"access_mode,omitempty"`

	// Deprecated. Use capability block instead. The storage API that will be used by the volume.
	// The storage API that will be used by the volume.
	// +kubebuilder:validation:Required
	AttachmentMode *string `json:"attachmentMode" tf:"attachment_mode,omitempty"`
}

type MountOptionsObservation struct {

	// The file system type.
	// The file system type.
	FsType *string `json:"fsType,omitempty" tf:"fs_type,omitempty"`

	// The flags passed to mount.
	// The flags passed to mount.
	MountFlags []*string `json:"mountFlags,omitempty" tf:"mount_flags,omitempty"`
}

type MountOptionsParameters struct {

	// The file system type.
	// The file system type.
	// +kubebuilder:validation:Optional
	FsType *string `json:"fsType,omitempty" tf:"fs_type,omitempty"`

	// The flags passed to mount.
	// The flags passed to mount.
	// +kubebuilder:validation:Optional
	MountFlags []*string `json:"mountFlags,omitempty" tf:"mount_flags,omitempty"`
}

type RequiredObservation struct {

	// Defines the location for the volume.
	// Defines the location for the volume.
	Topology []TopologyObservation `json:"topology,omitempty" tf:"topology,omitempty"`
}

type RequiredParameters struct {

	// Defines the location for the volume.
	// Defines the location for the volume.
	// +kubebuilder:validation:Required
	Topology []TopologyParameters `json:"topology" tf:"topology,omitempty"`
}

type TopologiesObservation struct {

	// Define the attributes for the topology request.
	Segments map[string]*string `json:"segments,omitempty" tf:"segments,omitempty"`
}

type TopologiesParameters struct {
}

type TopologyObservation struct {

	// Define the attributes for the topology request.
	// Define attributes for the topology request.
	Segments map[string]*string `json:"segments,omitempty" tf:"segments,omitempty"`
}

type TopologyParameters struct {

	// Define the attributes for the topology request.
	// Define attributes for the topology request.
	// +kubebuilder:validation:Required
	Segments map[string]*string `json:"segments" tf:"segments,omitempty"`
}

type TopologyRequestObservation struct {

	// Required topologies indicate that the volume must be created in a location accessible from all the listed topologies.
	// Required topologies indicate that the volume must be created in a location accessible from all the listed topologies.
	Required []RequiredObservation `json:"required,omitempty" tf:"required,omitempty"`
}

type TopologyRequestParameters struct {

	// Required topologies indicate that the volume must be created in a location accessible from all the listed topologies.
	// Required topologies indicate that the volume must be created in a location accessible from all the listed topologies.
	// +kubebuilder:validation:Optional
	Required []RequiredParameters `json:"required,omitempty" tf:"required,omitempty"`
}

type VolumeObservation struct {

	// Deprecated. Use capability block instead. Defines whether a volume should be available concurrently. Possible values are:
	// Defines whether a volume should be available concurrently.
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// Deprecated. Use capability block instead. The storage API that will be used by the volume.
	// The storage API that will be used by the volume.
	AttachmentMode *string `json:"attachmentMode,omitempty" tf:"attachment_mode,omitempty"`

	// Options for validating the capability of a volume.
	// Capabilities intended to be used in a job. At least one capability must be provided.
	Capability []CapabilityObservation `json:"capability,omitempty" tf:"capability,omitempty"`

	// An optional key-value map of strings passed directly to the CSI plugin to validate the volume.
	// An optional key-value map of strings passed directly to the CSI plugin to validate the volume.
	Context map[string]*string `json:"context,omitempty" tf:"context,omitempty"`

	// : (boolean)
	ControllerRequired *bool `json:"controllerRequired,omitempty" tf:"controller_required,omitempty"`

	// : (integer)
	ControllersExpected *float64 `json:"controllersExpected,omitempty" tf:"controllers_expected,omitempty"`

	// : (integer)
	ControllersHealthy *float64 `json:"controllersHealthy,omitempty" tf:"controllers_healthy,omitempty"`

	// If true, the volume will be deregistered on destroy.
	// If true, the volume will be deregistered on destroy.
	DeregisterOnDestroy *bool `json:"deregisterOnDestroy,omitempty" tf:"deregister_on_destroy,omitempty"`

	// The ID of the physical volume from the storage provider.
	// The ID of the physical volume from the storage provider.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// device volumes without a pre-formatted file system.
	// Options for mounting 'block-device' volumes without a pre-formatted file system.
	MountOptions []MountOptionsObservation `json:"mountOptions,omitempty" tf:"mount_options,omitempty"`

	// The namespace in which to register the volume.
	// The namespace in which to create the volume.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// : (integer)
	NodesExpected *float64 `json:"nodesExpected,omitempty" tf:"nodes_expected,omitempty"`

	// : (integer)
	NodesHealthy *float64 `json:"nodesHealthy,omitempty" tf:"nodes_healthy,omitempty"`

	// An optional key-value map of strings passed directly to the CSI plugin to configure the volume.
	// An optional key-value map of strings passed directly to the CSI plugin to configure the volume.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The ID of the Nomad plugin for registering this volume.
	// The ID of the CSI plugin that manages this volume.
	PluginID *string `json:"pluginId,omitempty" tf:"plugin_id,omitempty"`

	// : (string)
	PluginProvider *string `json:"pluginProvider,omitempty" tf:"plugin_provider,omitempty"`

	// : (string)
	PluginProviderVersion *string `json:"pluginProviderVersion,omitempty" tf:"plugin_provider_version,omitempty"`

	// : (boolean)
	Schedulable *bool `json:"schedulable,omitempty" tf:"schedulable,omitempty"`

	// : (List of topologies)
	Topologies []TopologiesObservation `json:"topologies,omitempty" tf:"topologies,omitempty"`

	// Specify locations (region, zone, rack, etc.) where the provisioned volume is accessible from.
	// Specify locations (region, zone, rack, etc.) where the provisioned volume is accessible from.
	TopologyRequest []TopologyRequestObservation `json:"topologyRequest,omitempty" tf:"topology_request,omitempty"`

	// The type of the volume. Currently, only csi is supported.
	// The type of the volume. Currently, only 'csi' is supported.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The unique ID of the volume.
	// The unique ID of the volume, how jobs will refer to the volume.
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type VolumeParameters struct {

	// Deprecated. Use capability block instead. Defines whether a volume should be available concurrently. Possible values are:
	// Defines whether a volume should be available concurrently.
	// +kubebuilder:validation:Optional
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// Deprecated. Use capability block instead. The storage API that will be used by the volume.
	// The storage API that will be used by the volume.
	// +kubebuilder:validation:Optional
	AttachmentMode *string `json:"attachmentMode,omitempty" tf:"attachment_mode,omitempty"`

	// Options for validating the capability of a volume.
	// Capabilities intended to be used in a job. At least one capability must be provided.
	// +kubebuilder:validation:Optional
	Capability []CapabilityParameters `json:"capability,omitempty" tf:"capability,omitempty"`

	// An optional key-value map of strings passed directly to the CSI plugin to validate the volume.
	// An optional key-value map of strings passed directly to the CSI plugin to validate the volume.
	// +kubebuilder:validation:Optional
	Context map[string]*string `json:"context,omitempty" tf:"context,omitempty"`

	// If true, the volume will be deregistered on destroy.
	// If true, the volume will be deregistered on destroy.
	// +kubebuilder:validation:Optional
	DeregisterOnDestroy *bool `json:"deregisterOnDestroy,omitempty" tf:"deregister_on_destroy,omitempty"`

	// The ID of the physical volume from the storage provider.
	// The ID of the physical volume from the storage provider.
	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// device volumes without a pre-formatted file system.
	// Options for mounting 'block-device' volumes without a pre-formatted file system.
	// +kubebuilder:validation:Optional
	MountOptions []MountOptionsParameters `json:"mountOptions,omitempty" tf:"mount_options,omitempty"`

	// The namespace in which to register the volume.
	// The namespace in which to create the volume.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// An optional key-value map of strings passed directly to the CSI plugin to configure the volume.
	// An optional key-value map of strings passed directly to the CSI plugin to configure the volume.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The ID of the Nomad plugin for registering this volume.
	// The ID of the CSI plugin that manages this volume.
	// +kubebuilder:validation:Optional
	PluginID *string `json:"pluginId,omitempty" tf:"plugin_id,omitempty"`

	// An optional key-value map of strings used as credentials for publishing and unpublishing volumes.
	// An optional key-value map of strings used as credentials for publishing and unpublishing volumes.
	// +kubebuilder:validation:Optional
	SecretsSecretRef *v1.SecretReference `json:"secretsSecretRef,omitempty" tf:"-"`

	// Specify locations (region, zone, rack, etc.) where the provisioned volume is accessible from.
	// Specify locations (region, zone, rack, etc.) where the provisioned volume is accessible from.
	// +kubebuilder:validation:Optional
	TopologyRequest []TopologyRequestParameters `json:"topologyRequest,omitempty" tf:"topology_request,omitempty"`

	// The type of the volume. Currently, only csi is supported.
	// The type of the volume. Currently, only 'csi' is supported.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The unique ID of the volume.
	// The unique ID of the volume, how jobs will refer to the volume.
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

// VolumeSpec defines the desired state of Volume
type VolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeParameters `json:"forProvider"`
}

// VolumeStatus defines the observed state of Volume.
type VolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Volume is the Schema for the Volumes API. Manages the lifecycle of registering and deregistering Nomad volumes.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nomad}
type Volume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.externalId)",message="externalId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.pluginId)",message="pluginId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.volumeId)",message="volumeId is a required parameter"
	Spec   VolumeSpec   `json:"spec"`
	Status VolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeList contains a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Volume `json:"items"`
}

// Repository type metadata.
var (
	Volume_Kind             = "Volume"
	Volume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Volume_Kind}.String()
	Volume_KindAPIVersion   = Volume_Kind + "." + CRDGroupVersion.String()
	Volume_GroupVersionKind = CRDGroupVersion.WithKind(Volume_Kind)
)

func init() {
	SchemeBuilder.Register(&Volume{}, &VolumeList{})
}