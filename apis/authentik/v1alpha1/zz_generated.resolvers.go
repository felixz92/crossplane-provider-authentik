/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Outpost.
func (mg *Outpost) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceConnection),
		Extract:      resource.ExtractParamPath("id", true),
		Reference:    mg.Spec.ForProvider.ServiceConnectionRef,
		Selector:     mg.Spec.ForProvider.ServiceConnectionSelector,
		To: reference.To{
			List:    &ServiceConnectionKubernetesList{},
			Managed: &ServiceConnectionKubernetes{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceConnection")
	}
	mg.Spec.ForProvider.ServiceConnection = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceConnectionRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceConnection),
		Extract:      resource.ExtractParamPath("id", true),
		Reference:    mg.Spec.InitProvider.ServiceConnectionRef,
		Selector:     mg.Spec.InitProvider.ServiceConnectionSelector,
		To: reference.To{
			List:    &ServiceConnectionKubernetesList{},
			Managed: &ServiceConnectionKubernetes{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceConnection")
	}
	mg.Spec.InitProvider.ServiceConnection = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceConnectionRef = rsp.ResolvedReference

	return nil
}
