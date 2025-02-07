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

type SourcePlexInitParameters struct {

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	AllowFriends *bool `json:"allowFriends,omitempty" tf:"allow_friends,omitempty"`

	// (List of String)
	AllowedServers []*string `json:"allowedServers,omitempty" tf:"allowed_servers,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/felixz92/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	AuthenticationFlow *string `json:"authenticationFlow,omitempty" tf:"authentication_flow,omitempty"`

	// Reference to a Flow in authentik to populate authenticationFlow.
	// +kubebuilder:validation:Optional
	AuthenticationFlowRef *v1.Reference `json:"authenticationFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate authenticationFlow.
	// +kubebuilder:validation:Optional
	AuthenticationFlowSelector *v1.Selector `json:"authenticationFlowSelector,omitempty" tf:"-"`

	// (String)
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/felixz92/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	EnrollmentFlow *string `json:"enrollmentFlow,omitempty" tf:"enrollment_flow,omitempty"`

	// Reference to a Flow in authentik to populate enrollmentFlow.
	// +kubebuilder:validation:Optional
	EnrollmentFlowRef *v1.Reference `json:"enrollmentFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate enrollmentFlow.
	// +kubebuilder:validation:Optional
	EnrollmentFlowSelector *v1.Selector `json:"enrollmentFlowSelector,omitempty" tf:"-"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `name_link`
	// - `name_deny`
	// Defaults to `identifier`.
	GroupMatchingMode *string `json:"groupMatchingMode,omitempty" tf:"group_matching_mode,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive)
	PlexTokenSecretRef v1.SecretKeySelector `json:"plexTokenSecretRef" tf:"-"`

	// (String) Allowed values:
	// Allowed values:
	// - `all`
	// - `any`
	// Defaults to `any`.
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (String) Generated.
	// Generated.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `email_link`
	// - `email_deny`
	// - `username_link`
	// - `username_deny`
	// Defaults to `identifier`.
	UserMatchingMode *string `json:"userMatchingMode,omitempty" tf:"user_matching_mode,omitempty"`

	// (String) Defaults to goauthentik.io/sources/%(slug)s.
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate *string `json:"userPathTemplate,omitempty" tf:"user_path_template,omitempty"`
}

type SourcePlexObservation struct {

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	AllowFriends *bool `json:"allowFriends,omitempty" tf:"allow_friends,omitempty"`

	// (List of String)
	AllowedServers []*string `json:"allowedServers,omitempty" tf:"allowed_servers,omitempty"`

	// (String)
	AuthenticationFlow *string `json:"authenticationFlow,omitempty" tf:"authentication_flow,omitempty"`

	// (String)
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String)
	EnrollmentFlow *string `json:"enrollmentFlow,omitempty" tf:"enrollment_flow,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `name_link`
	// - `name_deny`
	// Defaults to `identifier`.
	GroupMatchingMode *string `json:"groupMatchingMode,omitempty" tf:"group_matching_mode,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `all`
	// - `any`
	// Defaults to `any`.
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (String) Generated.
	// Generated.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `email_link`
	// - `email_deny`
	// - `username_link`
	// - `username_deny`
	// Defaults to `identifier`.
	UserMatchingMode *string `json:"userMatchingMode,omitempty" tf:"user_matching_mode,omitempty"`

	// (String) Defaults to goauthentik.io/sources/%(slug)s.
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate *string `json:"userPathTemplate,omitempty" tf:"user_path_template,omitempty"`
}

type SourcePlexParameters struct {

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	AllowFriends *bool `json:"allowFriends,omitempty" tf:"allow_friends,omitempty"`

	// (List of String)
	// +kubebuilder:validation:Optional
	AllowedServers []*string `json:"allowedServers,omitempty" tf:"allowed_servers,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/felixz92/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	// +kubebuilder:validation:Optional
	AuthenticationFlow *string `json:"authenticationFlow,omitempty" tf:"authentication_flow,omitempty"`

	// Reference to a Flow in authentik to populate authenticationFlow.
	// +kubebuilder:validation:Optional
	AuthenticationFlowRef *v1.Reference `json:"authenticationFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate authenticationFlow.
	// +kubebuilder:validation:Optional
	AuthenticationFlowSelector *v1.Selector `json:"authenticationFlowSelector,omitempty" tf:"-"`

	// (String)
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/felixz92/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	// +kubebuilder:validation:Optional
	EnrollmentFlow *string `json:"enrollmentFlow,omitempty" tf:"enrollment_flow,omitempty"`

	// Reference to a Flow in authentik to populate enrollmentFlow.
	// +kubebuilder:validation:Optional
	EnrollmentFlowRef *v1.Reference `json:"enrollmentFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate enrollmentFlow.
	// +kubebuilder:validation:Optional
	EnrollmentFlowSelector *v1.Selector `json:"enrollmentFlowSelector,omitempty" tf:"-"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `name_link`
	// - `name_deny`
	// Defaults to `identifier`.
	// +kubebuilder:validation:Optional
	GroupMatchingMode *string `json:"groupMatchingMode,omitempty" tf:"group_matching_mode,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive)
	// +kubebuilder:validation:Optional
	PlexTokenSecretRef v1.SecretKeySelector `json:"plexTokenSecretRef" tf:"-"`

	// (String) Allowed values:
	// Allowed values:
	// - `all`
	// - `any`
	// Defaults to `any`.
	// +kubebuilder:validation:Optional
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (String) Generated.
	// Generated.
	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `email_link`
	// - `email_deny`
	// - `username_link`
	// - `username_deny`
	// Defaults to `identifier`.
	// +kubebuilder:validation:Optional
	UserMatchingMode *string `json:"userMatchingMode,omitempty" tf:"user_matching_mode,omitempty"`

	// (String) Defaults to goauthentik.io/sources/%(slug)s.
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	// +kubebuilder:validation:Optional
	UserPathTemplate *string `json:"userPathTemplate,omitempty" tf:"user_path_template,omitempty"`
}

// SourcePlexSpec defines the desired state of SourcePlex
type SourcePlexSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SourcePlexParameters `json:"forProvider"`
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
	InitProvider SourcePlexInitParameters `json:"initProvider,omitempty"`
}

// SourcePlexStatus defines the observed state of SourcePlex.
type SourcePlexStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SourcePlexObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SourcePlex is the Schema for the SourcePlexs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type SourcePlex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || (has(self.initProvider) && has(self.initProvider.clientId))",message="spec.forProvider.clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.plexTokenSecretRef)",message="spec.forProvider.plexTokenSecretRef is a required parameter"
	Spec   SourcePlexSpec   `json:"spec"`
	Status SourcePlexStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SourcePlexList contains a list of SourcePlexs
type SourcePlexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SourcePlex `json:"items"`
}

// Repository type metadata.
var (
	SourcePlex_Kind             = "SourcePlex"
	SourcePlex_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SourcePlex_Kind}.String()
	SourcePlex_KindAPIVersion   = SourcePlex_Kind + "." + CRDGroupVersion.String()
	SourcePlex_GroupVersionKind = CRDGroupVersion.WithKind(SourcePlex_Kind)
)

func init() {
	SchemeBuilder.Register(&SourcePlex{}, &SourcePlexList{})
}
