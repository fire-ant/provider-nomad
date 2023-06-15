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

type LimitsObservation struct {

	// The region these limits should apply to.
	// Region in which this limit has affect.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The limits to enforce. This block
	// may only be specified once in the limits block. Its structure is
	// documented below.
	// The limit applied to this region.
	RegionLimit []RegionLimitObservation `json:"regionLimit,omitempty" tf:"region_limit,omitempty"`
}

type LimitsParameters struct {

	// The region these limits should apply to.
	// Region in which this limit has affect.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The limits to enforce. This block
	// may only be specified once in the limits block. Its structure is
	// documented below.
	// The limit applied to this region.
	// +kubebuilder:validation:Required
	RegionLimit []RegionLimitParameters `json:"regionLimit" tf:"region_limit,omitempty"`
}

type RegionLimitObservation struct {

	// The amount of CPU to limit allocations to. A value of zero
	// is treated as unlimited, and a negative value is treated as fully disallowed.
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// The amount of memory (in megabytes) to limit
	// allocations to. A value of zero is treated as unlimited, and a negative value
	// is treated as fully disallowed.
	MemoryMb *float64 `json:"memoryMb,omitempty" tf:"memory_mb,omitempty"`
}

type RegionLimitParameters struct {

	// The amount of CPU to limit allocations to. A value of zero
	// is treated as unlimited, and a negative value is treated as fully disallowed.
	// +kubebuilder:validation:Optional
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// The amount of memory (in megabytes) to limit
	// allocations to. A value of zero is treated as unlimited, and a negative value
	// is treated as fully disallowed.
	// +kubebuilder:validation:Optional
	MemoryMb *float64 `json:"memoryMb,omitempty" tf:"memory_mb,omitempty"`
}

type SpecificationObservation struct {

	// A description of the quota specification.
	// Description for this quota specification.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A block of quota limits to enforce. Can
	// be repeated. See below for the structure of this block.
	// Limits encapsulated by this quota specification.
	Limits []LimitsObservation `json:"limits,omitempty" tf:"limits,omitempty"`
}

type SpecificationParameters struct {

	// A description of the quota specification.
	// Description for this quota specification.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A block of quota limits to enforce. Can
	// be repeated. See below for the structure of this block.
	// Limits encapsulated by this quota specification.
	// +kubebuilder:validation:Optional
	Limits []LimitsParameters `json:"limits,omitempty" tf:"limits,omitempty"`
}

// SpecificationSpec defines the desired state of Specification
type SpecificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpecificationParameters `json:"forProvider"`
}

// SpecificationStatus defines the observed state of Specification.
type SpecificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpecificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Specification is the Schema for the Specifications API. Manages a quota specification in a Nomad cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nomad}
type Specification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.limits)",message="limits is a required parameter"
	Spec   SpecificationSpec   `json:"spec"`
	Status SpecificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpecificationList contains a list of Specifications
type SpecificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Specification `json:"items"`
}

// Repository type metadata.
var (
	Specification_Kind             = "Specification"
	Specification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Specification_Kind}.String()
	Specification_KindAPIVersion   = Specification_Kind + "." + CRDGroupVersion.String()
	Specification_GroupVersionKind = CRDGroupVersion.WithKind(Specification_Kind)
)

func init() {
	SchemeBuilder.Register(&Specification{}, &SpecificationList{})
}