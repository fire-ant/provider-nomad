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

type RoleObservation struct {

	// A description of the ACL Role.
	// Description for this ACL role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of policy names to associate with this
	// ACL Role. It may be used multiple times.
	// The policies that should be applied to the role. It may be used multiple times.
	Policy []RolePolicyObservation `json:"policy,omitempty" tf:"policy,omitempty"`
}

type RoleParameters struct {

	// A description of the ACL Role.
	// Description for this ACL role.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of policy names to associate with this
	// ACL Role. It may be used multiple times.
	// The policies that should be applied to the role. It may be used multiple times.
	// +kubebuilder:validation:Optional
	Policy []RolePolicyParameters `json:"policy,omitempty" tf:"policy,omitempty"`
}

type RolePolicyObservation struct {

	// A human-friendly name for this ACL Role.
	// The name of the ACL policy to link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RolePolicyParameters struct {

	// A human-friendly name for this ACL Role.
	// The name of the ACL policy to link.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Role is the Schema for the Roles API. Manages an ACL Role in Nomad.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nomad}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.policy)",message="policy is a required parameter"
	Spec   RoleSpec   `json:"spec"`
	Status RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}
