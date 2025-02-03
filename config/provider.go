/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/felixz92/crossplane-provider-authentik/config/base"
	"github.com/felixz92/crossplane-provider-authentik/config/directory"
	"github.com/felixz92/crossplane-provider-authentik/config/policy"
	"github.com/felixz92/crossplane-provider-authentik/config/propertymapping"
	"github.com/felixz92/crossplane-provider-authentik/config/provider"
	"github.com/felixz92/crossplane-provider-authentik/config/rbac"
	"github.com/felixz92/crossplane-provider-authentik/config/stage"
)

const (
	resourcePrefix = "authentik"
	modulePath     = "github.com/felixz92/crossplane-provider-authentik"
	rootGroup      = "authentik.crossplane.io"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider(generationProvider bool) (*ujconfig.Provider, error) {

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		base.Configure,
		directory.Configure,
		policy.Configure,
		propertymapping.Configure,
		provider.Configure,
		stage.Configure,
		rbac.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}
