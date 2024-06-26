// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *CapacityProvider) ResolveReferences( // ResolveReferences of this CapacityProvider.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.AutoScalingGroupProvider != nil {
		{
			m, l, err = apisresolver.GetManagedResource("autoscaling.aws.upbound.io", "v1beta3", "AutoscalingGroup", "AutoscalingGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoScalingGroupProvider.AutoScalingGroupArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.AutoScalingGroupProvider.AutoScalingGroupArnRef,
				Selector:     mg.Spec.ForProvider.AutoScalingGroupProvider.AutoScalingGroupArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AutoScalingGroupProvider.AutoScalingGroupArn")
		}
		mg.Spec.ForProvider.AutoScalingGroupProvider.AutoScalingGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AutoScalingGroupProvider.AutoScalingGroupArnRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.AutoScalingGroupProvider != nil {
		{
			m, l, err = apisresolver.GetManagedResource("autoscaling.aws.upbound.io", "v1beta3", "AutoscalingGroup", "AutoscalingGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AutoScalingGroupProvider.AutoScalingGroupArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.InitProvider.AutoScalingGroupProvider.AutoScalingGroupArnRef,
				Selector:     mg.Spec.InitProvider.AutoScalingGroupProvider.AutoScalingGroupArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.AutoScalingGroupProvider.AutoScalingGroupArn")
		}
		mg.Spec.InitProvider.AutoScalingGroupProvider.AutoScalingGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.AutoScalingGroupProvider.AutoScalingGroupArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ecs.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Cluster),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClusterRef,
			Selector:     mg.Spec.ForProvider.ClusterSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Cluster")
	}
	mg.Spec.ForProvider.Cluster = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IAMRole),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.IAMRoleRef,
			Selector:     mg.Spec.ForProvider.IAMRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRole")
	}
	mg.Spec.ForProvider.IAMRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IAMRoleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LoadBalancer); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBTargetGroup", "LBTargetGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArn),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArnRef,
				Selector:     mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArn")
		}
		mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArnRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.NetworkConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.NetworkConfiguration.SecurityGroups),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.NetworkConfiguration.SecurityGroupRefs,
				Selector:      mg.Spec.ForProvider.NetworkConfiguration.SecurityGroupSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkConfiguration.SecurityGroups")
		}
		mg.Spec.ForProvider.NetworkConfiguration.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.NetworkConfiguration.SecurityGroupRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.ForProvider.NetworkConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.NetworkConfiguration.Subnets),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.NetworkConfiguration.SubnetRefs,
				Selector:      mg.Spec.ForProvider.NetworkConfiguration.SubnetSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkConfiguration.Subnets")
		}
		mg.Spec.ForProvider.NetworkConfiguration.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.NetworkConfiguration.SubnetRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("ecs.aws.upbound.io", "v1beta2", "TaskDefinition", "TaskDefinitionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TaskDefinition),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.TaskDefinitionRef,
			Selector:     mg.Spec.ForProvider.TaskDefinitionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TaskDefinition")
	}
	mg.Spec.ForProvider.TaskDefinition = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TaskDefinitionRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ecs.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Cluster),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClusterRef,
			Selector:     mg.Spec.InitProvider.ClusterSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Cluster")
	}
	mg.Spec.InitProvider.Cluster = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IAMRole),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.IAMRoleRef,
			Selector:     mg.Spec.InitProvider.IAMRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IAMRole")
	}
	mg.Spec.InitProvider.IAMRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IAMRoleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.LoadBalancer); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBTargetGroup", "LBTargetGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancer[i3].TargetGroupArn),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.LoadBalancer[i3].TargetGroupArnRef,
				Selector:     mg.Spec.InitProvider.LoadBalancer[i3].TargetGroupArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancer[i3].TargetGroupArn")
		}
		mg.Spec.InitProvider.LoadBalancer[i3].TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.LoadBalancer[i3].TargetGroupArnRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.NetworkConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.NetworkConfiguration.SecurityGroups),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.NetworkConfiguration.SecurityGroupRefs,
				Selector:      mg.Spec.InitProvider.NetworkConfiguration.SecurityGroupSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.NetworkConfiguration.SecurityGroups")
		}
		mg.Spec.InitProvider.NetworkConfiguration.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.NetworkConfiguration.SecurityGroupRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.InitProvider.NetworkConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.NetworkConfiguration.Subnets),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.NetworkConfiguration.SubnetRefs,
				Selector:      mg.Spec.InitProvider.NetworkConfiguration.SubnetSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.NetworkConfiguration.Subnets")
		}
		mg.Spec.InitProvider.NetworkConfiguration.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.NetworkConfiguration.SubnetRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("ecs.aws.upbound.io", "v1beta2", "TaskDefinition", "TaskDefinitionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TaskDefinition),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.TaskDefinitionRef,
			Selector:     mg.Spec.InitProvider.TaskDefinitionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TaskDefinition")
	}
	mg.Spec.InitProvider.TaskDefinition = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TaskDefinitionRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TaskDefinition.
func (mg *TaskDefinition) ResolveReferences(ctx context.Context, c client.Reader) error {
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
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExecutionRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ExecutionRoleArnRef,
			Selector:     mg.Spec.ForProvider.ExecutionRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ExecutionRoleArn")
	}
	mg.Spec.ForProvider.ExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ExecutionRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ExecutionRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.ExecutionRoleArnRef,
			Selector:     mg.Spec.InitProvider.ExecutionRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ExecutionRoleArn")
	}
	mg.Spec.InitProvider.ExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ExecutionRoleArnRef = rsp.ResolvedReference

	return nil
}
