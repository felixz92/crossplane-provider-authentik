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

type ExpiryInitParameters struct {

	// (Number)
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	DenyOnly *bool `json:"denyOnly,omitempty" tf:"deny_only,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	ExecutionLogging *bool `json:"executionLogging,omitempty" tf:"execution_logging,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ExpiryObservation struct {

	// (Number)
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	DenyOnly *bool `json:"denyOnly,omitempty" tf:"deny_only,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	ExecutionLogging *bool `json:"executionLogging,omitempty" tf:"execution_logging,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ExpiryParameters struct {

	// (Number)
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	DenyOnly *bool `json:"denyOnly,omitempty" tf:"deny_only,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	ExecutionLogging *bool `json:"executionLogging,omitempty" tf:"execution_logging,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// ExpirySpec defines the desired state of Expiry
type ExpirySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpiryParameters `json:"forProvider"`
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
	InitProvider ExpiryInitParameters `json:"initProvider,omitempty"`
}

// ExpiryStatus defines the observed state of Expiry.
type ExpiryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpiryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Expiry is the Schema for the Expirys API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type Expiry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.days) || (has(self.initProvider) && has(self.initProvider.days))",message="spec.forProvider.days is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ExpirySpec   `json:"spec"`
	Status ExpiryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpiryList contains a list of Expirys
type ExpiryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Expiry `json:"items"`
}

// Repository type metadata.
var (
	Expiry_Kind             = "Expiry"
	Expiry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Expiry_Kind}.String()
	Expiry_KindAPIVersion   = Expiry_Kind + "." + CRDGroupVersion.String()
	Expiry_GroupVersionKind = CRDGroupVersion.WithKind(Expiry_Kind)
)

func init() {
	SchemeBuilder.Register(&Expiry{}, &ExpiryList{})
}
