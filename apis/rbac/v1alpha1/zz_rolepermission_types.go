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

type RolePermissionInitParameters struct {

	// (String) Allowed values:
	// Allowed values:
	// - `authentik_tenants.domain`
	// - `authentik_crypto.certificatekeypair`
	// - `authentik_flows.flow`
	// - `authentik_flows.flowstagebinding`
	// - `authentik_outposts.dockerserviceconnection`
	// - `authentik_outposts.kubernetesserviceconnection`
	// - `authentik_outposts.outpost`
	// - `authentik_policies_dummy.dummypolicy`
	// - `authentik_policies_event_matcher.eventmatcherpolicy`
	// - `authentik_policies_expiry.passwordexpirypolicy`
	// - `authentik_policies_expression.expressionpolicy`
	// - `authentik_policies_geoip.geoippolicy`
	// - `authentik_policies_password.passwordpolicy`
	// - `authentik_policies_reputation.reputationpolicy`
	// - `authentik_policies.policybinding`
	// - `authentik_providers_ldap.ldapprovider`
	// - `authentik_providers_oauth2.scopemapping`
	// - `authentik_providers_oauth2.oauth2provider`
	// - `authentik_providers_proxy.proxyprovider`
	// - `authentik_providers_radius.radiusprovider`
	// - `authentik_providers_radius.radiusproviderpropertymapping`
	// - `authentik_providers_saml.samlprovider`
	// - `authentik_providers_saml.samlpropertymapping`
	// - `authentik_providers_scim.scimprovider`
	// - `authentik_providers_scim.scimmapping`
	// - `authentik_rbac.role`
	// - `authentik_sources_kerberos.kerberossource`
	// - `authentik_sources_kerberos.kerberossourcepropertymapping`
	// - `authentik_sources_kerberos.userkerberossourceconnection`
	// - `authentik_sources_kerberos.groupkerberossourceconnection`
	// - `authentik_sources_ldap.ldapsource`
	// - `authentik_sources_ldap.ldapsourcepropertymapping`
	// - `authentik_sources_oauth.oauthsource`
	// - `authentik_sources_oauth.oauthsourcepropertymapping`
	// - `authentik_sources_oauth.useroauthsourceconnection`
	// - `authentik_sources_oauth.groupoauthsourceconnection`
	// - `authentik_sources_plex.plexsource`
	// - `authentik_sources_plex.plexsourcepropertymapping`
	// - `authentik_sources_plex.userplexsourceconnection`
	// - `authentik_sources_plex.groupplexsourceconnection`
	// - `authentik_sources_saml.samlsource`
	// - `authentik_sources_saml.samlsourcepropertymapping`
	// - `authentik_sources_saml.usersamlsourceconnection`
	// - `authentik_sources_saml.groupsamlsourceconnection`
	// - `authentik_sources_scim.scimsource`
	// - `authentik_sources_scim.scimsourcepropertymapping`
	// - `authentik_stages_authenticator_duo.authenticatorduostage`
	// - `authentik_stages_authenticator_duo.duodevice`
	// - `authentik_stages_authenticator_sms.authenticatorsmsstage`
	// - `authentik_stages_authenticator_sms.smsdevice`
	// - `authentik_stages_authenticator_static.authenticatorstaticstage`
	// - `authentik_stages_authenticator_static.staticdevice`
	// - `authentik_stages_authenticator_totp.authenticatortotpstage`
	// - `authentik_stages_authenticator_totp.totpdevice`
	// - `authentik_stages_authenticator_validate.authenticatorvalidatestage`
	// - `authentik_stages_authenticator_webauthn.authenticatorwebauthnstage`
	// - `authentik_stages_authenticator_webauthn.webauthndevice`
	// - `authentik_stages_captcha.captchastage`
	// - `authentik_stages_consent.consentstage`
	// - `authentik_stages_consent.userconsent`
	// - `authentik_stages_deny.denystage`
	// - `authentik_stages_dummy.dummystage`
	// - `authentik_stages_email.emailstage`
	// - `authentik_stages_identification.identificationstage`
	// - `authentik_stages_invitation.invitationstage`
	// - `authentik_stages_invitation.invitation`
	// - `authentik_stages_password.passwordstage`
	// - `authentik_stages_prompt.prompt`
	// - `authentik_stages_prompt.promptstage`
	// - `authentik_stages_redirect.redirectstage`
	// - `authentik_stages_user_delete.userdeletestage`
	// - `authentik_stages_user_login.userloginstage`
	// - `authentik_stages_user_logout.userlogoutstage`
	// - `authentik_stages_user_write.userwritestage`
	// - `authentik_brands.brand`
	// - `authentik_blueprints.blueprintinstance`
	// - `authentik_core.group`
	// - `authentik_core.user`
	// - `authentik_core.application`
	// - `authentik_core.applicationentitlement`
	// - `authentik_core.token`
	// - `authentik_enterprise.license`
	// - `authentik_providers_google_workspace.googleworkspaceprovider`
	// - `authentik_providers_google_workspace.googleworkspaceprovidermapping`
	// - `authentik_providers_microsoft_entra.microsoftentraprovider`
	// - `authentik_providers_microsoft_entra.microsoftentraprovidermapping`
	// - `authentik_providers_rac.racprovider`
	// - `authentik_providers_rac.endpoint`
	// - `authentik_providers_rac.racpropertymapping`
	// - `authentik_stages_authenticator_endpoint_gdtc.authenticatorendpointgdtcstage`
	// - `authentik_stages_source.sourcestage`
	// - `authentik_events.event`
	// - `authentik_events.notificationtransport`
	// - `authentik_events.notification`
	// - `authentik_events.notificationrule`
	// - `authentik_events.notificationwebhookmapping`
	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	// (String)
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// (String)
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// (String)
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type RolePermissionObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `authentik_tenants.domain`
	// - `authentik_crypto.certificatekeypair`
	// - `authentik_flows.flow`
	// - `authentik_flows.flowstagebinding`
	// - `authentik_outposts.dockerserviceconnection`
	// - `authentik_outposts.kubernetesserviceconnection`
	// - `authentik_outposts.outpost`
	// - `authentik_policies_dummy.dummypolicy`
	// - `authentik_policies_event_matcher.eventmatcherpolicy`
	// - `authentik_policies_expiry.passwordexpirypolicy`
	// - `authentik_policies_expression.expressionpolicy`
	// - `authentik_policies_geoip.geoippolicy`
	// - `authentik_policies_password.passwordpolicy`
	// - `authentik_policies_reputation.reputationpolicy`
	// - `authentik_policies.policybinding`
	// - `authentik_providers_ldap.ldapprovider`
	// - `authentik_providers_oauth2.scopemapping`
	// - `authentik_providers_oauth2.oauth2provider`
	// - `authentik_providers_proxy.proxyprovider`
	// - `authentik_providers_radius.radiusprovider`
	// - `authentik_providers_radius.radiusproviderpropertymapping`
	// - `authentik_providers_saml.samlprovider`
	// - `authentik_providers_saml.samlpropertymapping`
	// - `authentik_providers_scim.scimprovider`
	// - `authentik_providers_scim.scimmapping`
	// - `authentik_rbac.role`
	// - `authentik_sources_kerberos.kerberossource`
	// - `authentik_sources_kerberos.kerberossourcepropertymapping`
	// - `authentik_sources_kerberos.userkerberossourceconnection`
	// - `authentik_sources_kerberos.groupkerberossourceconnection`
	// - `authentik_sources_ldap.ldapsource`
	// - `authentik_sources_ldap.ldapsourcepropertymapping`
	// - `authentik_sources_oauth.oauthsource`
	// - `authentik_sources_oauth.oauthsourcepropertymapping`
	// - `authentik_sources_oauth.useroauthsourceconnection`
	// - `authentik_sources_oauth.groupoauthsourceconnection`
	// - `authentik_sources_plex.plexsource`
	// - `authentik_sources_plex.plexsourcepropertymapping`
	// - `authentik_sources_plex.userplexsourceconnection`
	// - `authentik_sources_plex.groupplexsourceconnection`
	// - `authentik_sources_saml.samlsource`
	// - `authentik_sources_saml.samlsourcepropertymapping`
	// - `authentik_sources_saml.usersamlsourceconnection`
	// - `authentik_sources_saml.groupsamlsourceconnection`
	// - `authentik_sources_scim.scimsource`
	// - `authentik_sources_scim.scimsourcepropertymapping`
	// - `authentik_stages_authenticator_duo.authenticatorduostage`
	// - `authentik_stages_authenticator_duo.duodevice`
	// - `authentik_stages_authenticator_sms.authenticatorsmsstage`
	// - `authentik_stages_authenticator_sms.smsdevice`
	// - `authentik_stages_authenticator_static.authenticatorstaticstage`
	// - `authentik_stages_authenticator_static.staticdevice`
	// - `authentik_stages_authenticator_totp.authenticatortotpstage`
	// - `authentik_stages_authenticator_totp.totpdevice`
	// - `authentik_stages_authenticator_validate.authenticatorvalidatestage`
	// - `authentik_stages_authenticator_webauthn.authenticatorwebauthnstage`
	// - `authentik_stages_authenticator_webauthn.webauthndevice`
	// - `authentik_stages_captcha.captchastage`
	// - `authentik_stages_consent.consentstage`
	// - `authentik_stages_consent.userconsent`
	// - `authentik_stages_deny.denystage`
	// - `authentik_stages_dummy.dummystage`
	// - `authentik_stages_email.emailstage`
	// - `authentik_stages_identification.identificationstage`
	// - `authentik_stages_invitation.invitationstage`
	// - `authentik_stages_invitation.invitation`
	// - `authentik_stages_password.passwordstage`
	// - `authentik_stages_prompt.prompt`
	// - `authentik_stages_prompt.promptstage`
	// - `authentik_stages_redirect.redirectstage`
	// - `authentik_stages_user_delete.userdeletestage`
	// - `authentik_stages_user_login.userloginstage`
	// - `authentik_stages_user_logout.userlogoutstage`
	// - `authentik_stages_user_write.userwritestage`
	// - `authentik_brands.brand`
	// - `authentik_blueprints.blueprintinstance`
	// - `authentik_core.group`
	// - `authentik_core.user`
	// - `authentik_core.application`
	// - `authentik_core.applicationentitlement`
	// - `authentik_core.token`
	// - `authentik_enterprise.license`
	// - `authentik_providers_google_workspace.googleworkspaceprovider`
	// - `authentik_providers_google_workspace.googleworkspaceprovidermapping`
	// - `authentik_providers_microsoft_entra.microsoftentraprovider`
	// - `authentik_providers_microsoft_entra.microsoftentraprovidermapping`
	// - `authentik_providers_rac.racprovider`
	// - `authentik_providers_rac.endpoint`
	// - `authentik_providers_rac.racpropertymapping`
	// - `authentik_stages_authenticator_endpoint_gdtc.authenticatorendpointgdtcstage`
	// - `authentik_stages_source.sourcestage`
	// - `authentik_events.event`
	// - `authentik_events.notificationtransport`
	// - `authentik_events.notification`
	// - `authentik_events.notificationrule`
	// - `authentik_events.notificationwebhookmapping`
	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	// (String)
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// (String)
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// (String)
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type RolePermissionParameters struct {

	// (String) Allowed values:
	// Allowed values:
	// - `authentik_tenants.domain`
	// - `authentik_crypto.certificatekeypair`
	// - `authentik_flows.flow`
	// - `authentik_flows.flowstagebinding`
	// - `authentik_outposts.dockerserviceconnection`
	// - `authentik_outposts.kubernetesserviceconnection`
	// - `authentik_outposts.outpost`
	// - `authentik_policies_dummy.dummypolicy`
	// - `authentik_policies_event_matcher.eventmatcherpolicy`
	// - `authentik_policies_expiry.passwordexpirypolicy`
	// - `authentik_policies_expression.expressionpolicy`
	// - `authentik_policies_geoip.geoippolicy`
	// - `authentik_policies_password.passwordpolicy`
	// - `authentik_policies_reputation.reputationpolicy`
	// - `authentik_policies.policybinding`
	// - `authentik_providers_ldap.ldapprovider`
	// - `authentik_providers_oauth2.scopemapping`
	// - `authentik_providers_oauth2.oauth2provider`
	// - `authentik_providers_proxy.proxyprovider`
	// - `authentik_providers_radius.radiusprovider`
	// - `authentik_providers_radius.radiusproviderpropertymapping`
	// - `authentik_providers_saml.samlprovider`
	// - `authentik_providers_saml.samlpropertymapping`
	// - `authentik_providers_scim.scimprovider`
	// - `authentik_providers_scim.scimmapping`
	// - `authentik_rbac.role`
	// - `authentik_sources_kerberos.kerberossource`
	// - `authentik_sources_kerberos.kerberossourcepropertymapping`
	// - `authentik_sources_kerberos.userkerberossourceconnection`
	// - `authentik_sources_kerberos.groupkerberossourceconnection`
	// - `authentik_sources_ldap.ldapsource`
	// - `authentik_sources_ldap.ldapsourcepropertymapping`
	// - `authentik_sources_oauth.oauthsource`
	// - `authentik_sources_oauth.oauthsourcepropertymapping`
	// - `authentik_sources_oauth.useroauthsourceconnection`
	// - `authentik_sources_oauth.groupoauthsourceconnection`
	// - `authentik_sources_plex.plexsource`
	// - `authentik_sources_plex.plexsourcepropertymapping`
	// - `authentik_sources_plex.userplexsourceconnection`
	// - `authentik_sources_plex.groupplexsourceconnection`
	// - `authentik_sources_saml.samlsource`
	// - `authentik_sources_saml.samlsourcepropertymapping`
	// - `authentik_sources_saml.usersamlsourceconnection`
	// - `authentik_sources_saml.groupsamlsourceconnection`
	// - `authentik_sources_scim.scimsource`
	// - `authentik_sources_scim.scimsourcepropertymapping`
	// - `authentik_stages_authenticator_duo.authenticatorduostage`
	// - `authentik_stages_authenticator_duo.duodevice`
	// - `authentik_stages_authenticator_sms.authenticatorsmsstage`
	// - `authentik_stages_authenticator_sms.smsdevice`
	// - `authentik_stages_authenticator_static.authenticatorstaticstage`
	// - `authentik_stages_authenticator_static.staticdevice`
	// - `authentik_stages_authenticator_totp.authenticatortotpstage`
	// - `authentik_stages_authenticator_totp.totpdevice`
	// - `authentik_stages_authenticator_validate.authenticatorvalidatestage`
	// - `authentik_stages_authenticator_webauthn.authenticatorwebauthnstage`
	// - `authentik_stages_authenticator_webauthn.webauthndevice`
	// - `authentik_stages_captcha.captchastage`
	// - `authentik_stages_consent.consentstage`
	// - `authentik_stages_consent.userconsent`
	// - `authentik_stages_deny.denystage`
	// - `authentik_stages_dummy.dummystage`
	// - `authentik_stages_email.emailstage`
	// - `authentik_stages_identification.identificationstage`
	// - `authentik_stages_invitation.invitationstage`
	// - `authentik_stages_invitation.invitation`
	// - `authentik_stages_password.passwordstage`
	// - `authentik_stages_prompt.prompt`
	// - `authentik_stages_prompt.promptstage`
	// - `authentik_stages_redirect.redirectstage`
	// - `authentik_stages_user_delete.userdeletestage`
	// - `authentik_stages_user_login.userloginstage`
	// - `authentik_stages_user_logout.userlogoutstage`
	// - `authentik_stages_user_write.userwritestage`
	// - `authentik_brands.brand`
	// - `authentik_blueprints.blueprintinstance`
	// - `authentik_core.group`
	// - `authentik_core.user`
	// - `authentik_core.application`
	// - `authentik_core.applicationentitlement`
	// - `authentik_core.token`
	// - `authentik_enterprise.license`
	// - `authentik_providers_google_workspace.googleworkspaceprovider`
	// - `authentik_providers_google_workspace.googleworkspaceprovidermapping`
	// - `authentik_providers_microsoft_entra.microsoftentraprovider`
	// - `authentik_providers_microsoft_entra.microsoftentraprovidermapping`
	// - `authentik_providers_rac.racprovider`
	// - `authentik_providers_rac.endpoint`
	// - `authentik_providers_rac.racpropertymapping`
	// - `authentik_stages_authenticator_endpoint_gdtc.authenticatorendpointgdtcstage`
	// - `authentik_stages_source.sourcestage`
	// - `authentik_events.event`
	// - `authentik_events.notificationtransport`
	// - `authentik_events.notification`
	// - `authentik_events.notificationrule`
	// - `authentik_events.notificationwebhookmapping`
	// +kubebuilder:validation:Optional
	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// RolePermissionSpec defines the desired state of RolePermission
type RolePermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RolePermissionParameters `json:"forProvider"`
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
	InitProvider RolePermissionInitParameters `json:"initProvider,omitempty"`
}

// RolePermissionStatus defines the observed state of RolePermission.
type RolePermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RolePermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RolePermission is the Schema for the RolePermissions API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type RolePermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permission) || (has(self.initProvider) && has(self.initProvider.permission))",message="spec.forProvider.permission is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   RolePermissionSpec   `json:"spec"`
	Status RolePermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RolePermissionList contains a list of RolePermissions
type RolePermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RolePermission `json:"items"`
}

// Repository type metadata.
var (
	RolePermission_Kind             = "RolePermission"
	RolePermission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RolePermission_Kind}.String()
	RolePermission_KindAPIVersion   = RolePermission_Kind + "." + CRDGroupVersion.String()
	RolePermission_GroupVersionKind = CRDGroupVersion.WithKind(RolePermission_Kind)
)

func init() {
	SchemeBuilder.Register(&RolePermission{}, &RolePermissionList{})
}
