/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/application"
	blueprint "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/blueprint"
	certificatekeypair "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/certificatekeypair"
	eventrule "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/eventrule"
	eventtransport "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/eventtransport"
	flow "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/flow"
	flowstagebinding "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/flowstagebinding"
	outpost "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/outpost"
	scopemapping "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/scopemapping"
	serviceconnectionkubernetes "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/serviceconnectionkubernetes"
	token "github.com/felixz92/crossplane-provider-authentik/internal/controller/authentik/token"
	group "github.com/felixz92/crossplane-provider-authentik/internal/controller/directory/group"
	sourcekerberos "github.com/felixz92/crossplane-provider-authentik/internal/controller/directory/sourcekerberos"
	sourceldap "github.com/felixz92/crossplane-provider-authentik/internal/controller/directory/sourceldap"
	sourceoauth "github.com/felixz92/crossplane-provider-authentik/internal/controller/directory/sourceoauth"
	sourceplex "github.com/felixz92/crossplane-provider-authentik/internal/controller/directory/sourceplex"
	sourcesaml "github.com/felixz92/crossplane-provider-authentik/internal/controller/directory/sourcesaml"
	user "github.com/felixz92/crossplane-provider-authentik/internal/controller/directory/user"
	binding "github.com/felixz92/crossplane-provider-authentik/internal/controller/policy/binding"
	dummy "github.com/felixz92/crossplane-provider-authentik/internal/controller/policy/dummy"
	eventmatcher "github.com/felixz92/crossplane-provider-authentik/internal/controller/policy/eventmatcher"
	expiry "github.com/felixz92/crossplane-provider-authentik/internal/controller/policy/expiry"
	expression "github.com/felixz92/crossplane-provider-authentik/internal/controller/policy/expression"
	password "github.com/felixz92/crossplane-provider-authentik/internal/controller/policy/password"
	reputation "github.com/felixz92/crossplane-provider-authentik/internal/controller/policy/reputation"
	googleworkspace "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/googleworkspace"
	ldap "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/ldap"
	microsoftentra "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/microsoftentra"
	notification "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/notification"
	providergoogleworkspace "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/providergoogleworkspace"
	providermicrosoftentra "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/providermicrosoftentra"
	providerrac "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/providerrac"
	providerradius "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/providerradius"
	providersaml "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/providersaml"
	providerscim "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/providerscim"
	providerscope "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/providerscope"
	saml "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/saml"
	scim "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/scim"
	sourcekerberospropertymapping "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/sourcekerberos"
	sourceldappropertymapping "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/sourceldap"
	sourceoauthpropertymapping "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/sourceoauth"
	sourceplexpropertymapping "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/sourceplex"
	sourcesamlpropertymapping "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/sourcesaml"
	sourcescim "github.com/felixz92/crossplane-provider-authentik/internal/controller/propertymapping/sourcescim"
	googleworkspaceprovider "github.com/felixz92/crossplane-provider-authentik/internal/controller/provider/googleworkspace"
	ldapprovider "github.com/felixz92/crossplane-provider-authentik/internal/controller/provider/ldap"
	microsoftentraprovider "github.com/felixz92/crossplane-provider-authentik/internal/controller/provider/microsoftentra"
	oauth2 "github.com/felixz92/crossplane-provider-authentik/internal/controller/provider/oauth2"
	proxy "github.com/felixz92/crossplane-provider-authentik/internal/controller/provider/proxy"
	rac "github.com/felixz92/crossplane-provider-authentik/internal/controller/provider/rac"
	radius "github.com/felixz92/crossplane-provider-authentik/internal/controller/provider/radius"
	samlprovider "github.com/felixz92/crossplane-provider-authentik/internal/controller/provider/saml"
	scimprovider "github.com/felixz92/crossplane-provider-authentik/internal/controller/provider/scim"
	providerconfig "github.com/felixz92/crossplane-provider-authentik/internal/controller/providerconfig"
	authenticatorduo "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/authenticatorduo"
	authenticatorendpointgdtc "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/authenticatorendpointgdtc"
	authenticatorsms "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/authenticatorsms"
	authenticatorstatic "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/authenticatorstatic"
	authenticatortotp "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/authenticatortotp"
	authenticatorvalidate "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/authenticatorvalidate"
	authenticatorwebauthn "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/authenticatorwebauthn"
	captcha "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/captcha"
	consent "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/consent"
	deny "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/deny"
	dummystage "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/dummy"
	email "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/email"
	identification "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/identification"
	invitation "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/invitation"
	passwordstage "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/password"
	prompt "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/prompt"
	promptfield "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/promptfield"
	userdelete "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/userdelete"
	userlogin "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/userlogin"
	userlogout "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/userlogout"
	userwrite "github.com/felixz92/crossplane-provider-authentik/internal/controller/stage/userwrite"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		blueprint.Setup,
		certificatekeypair.Setup,
		eventrule.Setup,
		eventtransport.Setup,
		flow.Setup,
		flowstagebinding.Setup,
		outpost.Setup,
		scopemapping.Setup,
		serviceconnectionkubernetes.Setup,
		token.Setup,
		group.Setup,
		sourcekerberos.Setup,
		sourceldap.Setup,
		sourceoauth.Setup,
		sourceplex.Setup,
		sourcesaml.Setup,
		user.Setup,
		binding.Setup,
		dummy.Setup,
		eventmatcher.Setup,
		expiry.Setup,
		expression.Setup,
		password.Setup,
		reputation.Setup,
		googleworkspace.Setup,
		ldap.Setup,
		microsoftentra.Setup,
		notification.Setup,
		providergoogleworkspace.Setup,
		providermicrosoftentra.Setup,
		providerrac.Setup,
		providerradius.Setup,
		providersaml.Setup,
		providerscim.Setup,
		providerscope.Setup,
		saml.Setup,
		scim.Setup,
		sourcekerberospropertymapping.Setup,
		sourceldappropertymapping.Setup,
		sourceoauthpropertymapping.Setup,
		sourceplexpropertymapping.Setup,
		sourcesamlpropertymapping.Setup,
		sourcescim.Setup,
		googleworkspaceprovider.Setup,
		ldapprovider.Setup,
		microsoftentraprovider.Setup,
		oauth2.Setup,
		proxy.Setup,
		rac.Setup,
		radius.Setup,
		samlprovider.Setup,
		scimprovider.Setup,
		providerconfig.Setup,
		authenticatorduo.Setup,
		authenticatorendpointgdtc.Setup,
		authenticatorsms.Setup,
		authenticatorstatic.Setup,
		authenticatortotp.Setup,
		authenticatorvalidate.Setup,
		authenticatorwebauthn.Setup,
		captcha.Setup,
		consent.Setup,
		deny.Setup,
		dummystage.Setup,
		email.Setup,
		identification.Setup,
		invitation.Setup,
		passwordstage.Setup,
		prompt.Setup,
		promptfield.Setup,
		userdelete.Setup,
		userlogin.Setup,
		userlogout.Setup,
		userwrite.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
