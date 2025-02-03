package rbac

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/felixz92/crossplane-provider-authentik/config/base"
)

const shortGroup = "rbac"

// Configure configures the policy provider.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_rbac_permission_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UserPermission"

		r.References["user"] = base.UserRef
	})
	p.AddResourceConfigurator("authentik_rbac_permission_role", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RolePermission"
	})
	p.AddResourceConfigurator("authentik_rbac_role", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Role"
	})
}
