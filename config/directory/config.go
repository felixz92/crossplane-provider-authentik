package directory

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/vhdirk/provider-authentik/config/base"
)

const shortGroup = "directory"

// Configure configures the directory provider.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "User"
	})
	p.AddResourceConfigurator("authentik_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		r.References["users"] = config.Reference{
			Type:      "User",
			Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
	p.AddResourceConfigurator("authentik_source_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceLDAP"
	})
	p.AddResourceConfigurator("authentik_source_oauth", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceOAuth"

		r.References["authentication_flow"] = base.FlowUUIDRef
		r.References["enrollment_flow"] = base.FlowUUIDRef
	})
	p.AddResourceConfigurator("authentik_source_plex", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourcePlex"

		r.References["authentication_flow"] = base.FlowUUIDRef
		r.References["enrollment_flow"] = base.FlowUUIDRef
	})
	p.AddResourceConfigurator("authentik_source_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceSAML"

		r.References["authentication_flow"] = base.FlowUUIDRef
		r.References["enrollment_flow"] = base.FlowUUIDRef
		r.References["pre_authentication_flow"] = base.FlowUUIDRef
	})
}
