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

type TokenInitParameters struct {

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	Expiring *bool `json:"expiring,omitempty" tf:"expiring,omitempty"`

	// (String)
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `verification`
	// - `api`
	// - `recovery`
	// - `app_password`
	// Defaults to `api`.
	Intent *string `json:"intent,omitempty" tf:"intent,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	RetrieveKey *bool `json:"retrieveKey,omitempty" tf:"retrieve_key,omitempty"`

	// (Number)
	// +crossplane:generate:reference:type=github.com/felixz92/crossplane-provider-authentik/apis/directory/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)
	User *float64 `json:"user,omitempty" tf:"user,omitempty"`

	// Reference to a User in directory to populate user.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User in directory to populate user.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

type TokenObservation struct {

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// (Number) Generated.
	// Generated.
	ExpiresIn *float64 `json:"expiresIn,omitempty" tf:"expires_in,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	Expiring *bool `json:"expiring,omitempty" tf:"expiring,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `verification`
	// - `api`
	// - `recovery`
	// - `app_password`
	// Defaults to `api`.
	Intent *string `json:"intent,omitempty" tf:"intent,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	RetrieveKey *bool `json:"retrieveKey,omitempty" tf:"retrieve_key,omitempty"`

	// (Number)
	User *float64 `json:"user,omitempty" tf:"user,omitempty"`
}

type TokenParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	Expiring *bool `json:"expiring,omitempty" tf:"expiring,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `verification`
	// - `api`
	// - `recovery`
	// - `app_password`
	// Defaults to `api`.
	// +kubebuilder:validation:Optional
	Intent *string `json:"intent,omitempty" tf:"intent,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	RetrieveKey *bool `json:"retrieveKey,omitempty" tf:"retrieve_key,omitempty"`

	// (Number)
	// +crossplane:generate:reference:type=github.com/felixz92/crossplane-provider-authentik/apis/directory/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)
	// +kubebuilder:validation:Optional
	User *float64 `json:"user,omitempty" tf:"user,omitempty"`

	// Reference to a User in directory to populate user.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User in directory to populate user.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// TokenSpec defines the desired state of Token
type TokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TokenParameters `json:"forProvider"`
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
	InitProvider TokenInitParameters `json:"initProvider,omitempty"`
}

// TokenStatus defines the observed state of Token.
type TokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Token is the Schema for the Tokens API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identifier) || (has(self.initProvider) && has(self.initProvider.identifier))",message="spec.forProvider.identifier is a required parameter"
	Spec   TokenSpec   `json:"spec"`
	Status TokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TokenList contains a list of Tokens
type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Token `json:"items"`
}

// Repository type metadata.
var (
	Token_Kind             = "Token"
	Token_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Token_Kind}.String()
	Token_KindAPIVersion   = Token_Kind + "." + CRDGroupVersion.String()
	Token_GroupVersionKind = CRDGroupVersion.WithKind(Token_Kind)
)

func init() {
	SchemeBuilder.Register(&Token{}, &TokenList{})
}
