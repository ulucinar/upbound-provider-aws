// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Server.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta2", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Certificate),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.CertificateRef,
			Selector:     mg.Spec.ForProvider.CertificateSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Certificate")
	}
	mg.Spec.ForProvider.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ds.aws.upbound.io", "v1beta2", "Directory", "DirectoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DirectoryID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DirectoryIDRef,
			Selector:     mg.Spec.ForProvider.DirectoryIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DirectoryID")
	}
	mg.Spec.ForProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DirectoryIDRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.EndpointDetails != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EndpointDetails.VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.EndpointDetails.VPCIDRef,
				Selector:     mg.Spec.ForProvider.EndpointDetails.VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EndpointDetails.VPCID")
		}
		mg.Spec.ForProvider.EndpointDetails.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EndpointDetails.VPCIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoggingRole),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.LoggingRoleRef,
			Selector:     mg.Spec.ForProvider.LoggingRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoggingRole")
	}
	mg.Spec.ForProvider.LoggingRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoggingRoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta2", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Certificate),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.CertificateRef,
			Selector:     mg.Spec.InitProvider.CertificateSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Certificate")
	}
	mg.Spec.InitProvider.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ds.aws.upbound.io", "v1beta2", "Directory", "DirectoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DirectoryID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DirectoryIDRef,
			Selector:     mg.Spec.InitProvider.DirectoryIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DirectoryID")
	}
	mg.Spec.InitProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DirectoryIDRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.EndpointDetails != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EndpointDetails.VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.EndpointDetails.VPCIDRef,
				Selector:     mg.Spec.InitProvider.EndpointDetails.VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EndpointDetails.VPCID")
		}
		mg.Spec.InitProvider.EndpointDetails.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EndpointDetails.VPCIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoggingRole),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.LoggingRoleRef,
			Selector:     mg.Spec.InitProvider.LoggingRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoggingRole")
	}
	mg.Spec.InitProvider.LoggingRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoggingRoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleRef,
			Selector:     mg.Spec.ForProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("transfer.aws.upbound.io", "v1beta2", "Server", "ServerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ServerIDRef,
			Selector:     mg.Spec.ForProvider.ServerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleRef,
			Selector:     mg.Spec.InitProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("transfer.aws.upbound.io", "v1beta2", "Server", "ServerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServerID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ServerIDRef,
			Selector:     mg.Spec.InitProvider.ServerIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServerID")
	}
	mg.Spec.InitProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workflow.
func (mg *Workflow) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Steps); i3++ {
		if mg.Spec.ForProvider.Steps[i3].CustomStepDetails != nil {
			{
				m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Steps[i3].CustomStepDetails.Target),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.Steps[i3].CustomStepDetails.TargetRef,
					Selector:     mg.Spec.ForProvider.Steps[i3].CustomStepDetails.TargetSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Steps[i3].CustomStepDetails.Target")
			}
			mg.Spec.ForProvider.Steps[i3].CustomStepDetails.Target = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Steps[i3].CustomStepDetails.TargetRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Steps); i3++ {
		if mg.Spec.InitProvider.Steps[i3].CustomStepDetails != nil {
			{
				m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Steps[i3].CustomStepDetails.Target),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.Steps[i3].CustomStepDetails.TargetRef,
					Selector:     mg.Spec.InitProvider.Steps[i3].CustomStepDetails.TargetSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Steps[i3].CustomStepDetails.Target")
			}
			mg.Spec.InitProvider.Steps[i3].CustomStepDetails.Target = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Steps[i3].CustomStepDetails.TargetRef = rsp.ResolvedReference

		}
	}

	return nil
}