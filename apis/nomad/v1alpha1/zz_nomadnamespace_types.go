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

type CapabilitiesObservation struct {

	// Task drivers disabled for the namespace.
	// Disabled task drivers for the namespace.
	DisabledTaskDrivers []*string `json:"disabledTaskDrivers,omitempty" tf:"disabled_task_drivers,omitempty"`

	// Task drivers enabled for the namespace.
	// Enabled task drivers for the namespace.
	EnabledTaskDrivers []*string `json:"enabledTaskDrivers,omitempty" tf:"enabled_task_drivers,omitempty"`
}

type CapabilitiesParameters struct {

	// Task drivers disabled for the namespace.
	// Disabled task drivers for the namespace.
	// +kubebuilder:validation:Optional
	DisabledTaskDrivers []*string `json:"disabledTaskDrivers,omitempty" tf:"disabled_task_drivers,omitempty"`

	// Task drivers enabled for the namespace.
	// Enabled task drivers for the namespace.
	// +kubebuilder:validation:Optional
	EnabledTaskDrivers []*string `json:"enabledTaskDrivers,omitempty" tf:"enabled_task_drivers,omitempty"`
}

type NomadNamespaceObservation struct {

	// A block of capabilities for the namespace. Can't
	// be repeated. See below for the structure of this block.
	// Capabilities of the namespace.
	Capabilities []CapabilitiesObservation `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// A description of the namespace.
	// Description for this namespace.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies arbitrary KV metadata to associate with the namespace.
	// Metadata associated with the namespace.
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// A resource quota to attach to the namespace.
	// Quota to set for this namespace.
	Quota *string `json:"quota,omitempty" tf:"quota,omitempty"`
}

type NomadNamespaceParameters struct {

	// A block of capabilities for the namespace. Can't
	// be repeated. See below for the structure of this block.
	// Capabilities of the namespace.
	// +kubebuilder:validation:Optional
	Capabilities []CapabilitiesParameters `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// A description of the namespace.
	// Description for this namespace.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies arbitrary KV metadata to associate with the namespace.
	// Metadata associated with the namespace.
	// +kubebuilder:validation:Optional
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// A resource quota to attach to the namespace.
	// Quota to set for this namespace.
	// +kubebuilder:validation:Optional
	Quota *string `json:"quota,omitempty" tf:"quota,omitempty"`
}

// NomadNamespaceSpec defines the desired state of NomadNamespace
type NomadNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NomadNamespaceParameters `json:"forProvider"`
}

// NomadNamespaceStatus defines the observed state of NomadNamespace.
type NomadNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NomadNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NomadNamespace is the Schema for the NomadNamespaces API. Provisions a namespace within a Nomad cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nomad}
type NomadNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NomadNamespaceSpec   `json:"spec"`
	Status            NomadNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NomadNamespaceList contains a list of NomadNamespaces
type NomadNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NomadNamespace `json:"items"`
}

// Repository type metadata.
var (
	NomadNamespace_Kind             = "NomadNamespace"
	NomadNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NomadNamespace_Kind}.String()
	NomadNamespace_KindAPIVersion   = NomadNamespace_Kind + "." + CRDGroupVersion.String()
	NomadNamespace_GroupVersionKind = CRDGroupVersion.WithKind(NomadNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&NomadNamespace{}, &NomadNamespaceList{})
}
