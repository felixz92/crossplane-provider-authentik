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

type DenyInitParameters struct {

	// (String)
	DenyMessage *string `json:"denyMessage,omitempty" tf:"deny_message,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DenyObservation struct {

	// (String)
	DenyMessage *string `json:"denyMessage,omitempty" tf:"deny_message,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DenyParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	DenyMessage *string `json:"denyMessage,omitempty" tf:"deny_message,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// DenySpec defines the desired state of Deny
type DenySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DenyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DenyInitParameters `json:"initProvider,omitempty"`
}

// DenyStatus defines the observed state of Deny.
type DenyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DenyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Deny is the Schema for the Denys API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type Deny struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   DenySpec   `json:"spec"`
	Status DenyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DenyList contains a list of Denys
type DenyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deny `json:"items"`
}

// Repository type metadata.
var (
	Deny_Kind             = "Deny"
	Deny_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Deny_Kind}.String()
	Deny_KindAPIVersion   = Deny_Kind + "." + CRDGroupVersion.String()
	Deny_GroupVersionKind = CRDGroupVersion.WithKind(Deny_Kind)
)

func init() {
	SchemeBuilder.Register(&Deny{}, &DenyList{})
}
