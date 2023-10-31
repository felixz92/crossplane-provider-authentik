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

type EmailInitParameters struct {

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	ActivateUserOnSuccess *bool `json:"activateUserOnSuccess,omitempty" tf:"activate_user_on_success,omitempty"`

	// (String) Defaults to system@authentik.local.
	// Defaults to `system@authentik.local`.
	FromAddress *string `json:"fromAddress,omitempty" tf:"from_address,omitempty"`

	// (String) Defaults to localhost.
	// Defaults to `localhost`.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Defaults to 25.
	// Defaults to `25`.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (String) Defaults to authentik.
	// Defaults to `authentik`.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// (String) Defaults to email/password_reset.html.
	// Defaults to `email/password_reset.html`.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Defaults to 30.
	// Defaults to `30`.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// (Number) Defaults to 30.
	// Defaults to `30`.
	TokenExpiry *float64 `json:"tokenExpiry,omitempty" tf:"token_expiry,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	UseGlobalSettings *bool `json:"useGlobalSettings,omitempty" tf:"use_global_settings,omitempty"`

	// (Boolean)
	UseSSL *bool `json:"useSsl,omitempty" tf:"use_ssl,omitempty"`

	// (Boolean)
	UseTLS *bool `json:"useTls,omitempty" tf:"use_tls,omitempty"`

	// (String)
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type EmailObservation struct {

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	ActivateUserOnSuccess *bool `json:"activateUserOnSuccess,omitempty" tf:"activate_user_on_success,omitempty"`

	// (String) Defaults to system@authentik.local.
	// Defaults to `system@authentik.local`.
	FromAddress *string `json:"fromAddress,omitempty" tf:"from_address,omitempty"`

	// (String) Defaults to localhost.
	// Defaults to `localhost`.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Defaults to 25.
	// Defaults to `25`.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (String) Defaults to authentik.
	// Defaults to `authentik`.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// (String) Defaults to email/password_reset.html.
	// Defaults to `email/password_reset.html`.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Defaults to 30.
	// Defaults to `30`.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// (Number) Defaults to 30.
	// Defaults to `30`.
	TokenExpiry *float64 `json:"tokenExpiry,omitempty" tf:"token_expiry,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	UseGlobalSettings *bool `json:"useGlobalSettings,omitempty" tf:"use_global_settings,omitempty"`

	// (Boolean)
	UseSSL *bool `json:"useSsl,omitempty" tf:"use_ssl,omitempty"`

	// (Boolean)
	UseTLS *bool `json:"useTls,omitempty" tf:"use_tls,omitempty"`

	// (String)
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type EmailParameters struct {

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	ActivateUserOnSuccess *bool `json:"activateUserOnSuccess,omitempty" tf:"activate_user_on_success,omitempty"`

	// (String) Defaults to system@authentik.local.
	// Defaults to `system@authentik.local`.
	// +kubebuilder:validation:Optional
	FromAddress *string `json:"fromAddress,omitempty" tf:"from_address,omitempty"`

	// (String) Defaults to localhost.
	// Defaults to `localhost`.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive)
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (Number) Defaults to 25.
	// Defaults to `25`.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (String) Defaults to authentik.
	// Defaults to `authentik`.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// (String) Defaults to email/password_reset.html.
	// Defaults to `email/password_reset.html`.
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Defaults to 30.
	// Defaults to `30`.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// (Number) Defaults to 30.
	// Defaults to `30`.
	// +kubebuilder:validation:Optional
	TokenExpiry *float64 `json:"tokenExpiry,omitempty" tf:"token_expiry,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	UseGlobalSettings *bool `json:"useGlobalSettings,omitempty" tf:"use_global_settings,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	UseSSL *bool `json:"useSsl,omitempty" tf:"use_ssl,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	UseTLS *bool `json:"useTls,omitempty" tf:"use_tls,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// EmailSpec defines the desired state of Email
type EmailSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmailParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider EmailInitParameters `json:"initProvider,omitempty"`
}

// EmailStatus defines the observed state of Email.
type EmailStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Email is the Schema for the Emails API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type Email struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   EmailSpec   `json:"spec"`
	Status EmailStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailList contains a list of Emails
type EmailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Email `json:"items"`
}

// Repository type metadata.
var (
	Email_Kind             = "Email"
	Email_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Email_Kind}.String()
	Email_KindAPIVersion   = Email_Kind + "." + CRDGroupVersion.String()
	Email_GroupVersionKind = CRDGroupVersion.WithKind(Email_Kind)
)

func init() {
	SchemeBuilder.Register(&Email{}, &EmailList{})
}
