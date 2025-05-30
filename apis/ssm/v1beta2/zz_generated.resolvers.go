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
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Association) ResolveReferences( // ResolveReferences of this Association.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "Document", "DocumentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NameRef,
			Selector:     mg.Spec.ForProvider.NameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Name")
	}
	mg.Spec.ForProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Targets); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "Instance", "InstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Targets[i3].Values),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.ForProvider.Targets[i3].ValuesRefs,
				Selector:      mg.Spec.ForProvider.Targets[i3].ValuesSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Targets[i3].Values")
		}
		mg.Spec.ForProvider.Targets[i3].Values = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Targets[i3].ValuesRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "Document", "DocumentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NameRef,
			Selector:     mg.Spec.InitProvider.NameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Name")
	}
	mg.Spec.InitProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Targets); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "Instance", "InstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Targets[i3].Values),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.InitProvider.Targets[i3].ValuesRefs,
				Selector:      mg.Spec.InitProvider.Targets[i3].ValuesSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Targets[i3].Values")
		}
		mg.Spec.InitProvider.Targets[i3].Values = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Targets[i3].ValuesRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this MaintenanceWindowTask.
func (mg *MaintenanceWindowTask) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ServiceRoleArnRef,
			Selector:     mg.Spec.ForProvider.ServiceRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceRoleArn")
	}
	mg.Spec.ForProvider.ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Targets); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "Instance", "InstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Targets[i3].Values),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.ForProvider.Targets[i3].ValuesRefs,
				Selector:      mg.Spec.ForProvider.Targets[i3].ValuesSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Targets[i3].Values")
		}
		mg.Spec.ForProvider.Targets[i3].Values = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Targets[i3].ValuesRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TaskArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.TaskArnRef,
			Selector:     mg.Spec.ForProvider.TaskArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TaskArn")
	}
	mg.Spec.ForProvider.TaskArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TaskArnRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.TaskInvocationParameters != nil {
		if mg.Spec.ForProvider.TaskInvocationParameters.AutomationParameters != nil {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.TaskInvocationParameters.AutomationParameters.Parameter); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "Instance", "InstanceList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
						CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].Values),
						Extract:       resource.ExtractResourceID(),
						References:    mg.Spec.ForProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].ValuesRefs,
						Selector:      mg.Spec.ForProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].ValuesSelector,
						To:            reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].Values")
				}
				mg.Spec.ForProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].Values = reference.ToPtrValues(mrsp.ResolvedValues)
				mg.Spec.ForProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].ValuesRefs = mrsp.ResolvedReferences

			}
		}
	}
	if mg.Spec.ForProvider.TaskInvocationParameters != nil {
		if mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters != nil {
			if mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig != nil {
				{
					m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArnRef,
						Selector:     mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArn")
				}
				mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArnRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.ForProvider.TaskInvocationParameters != nil {
		if mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters != nil {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.OutputS3Bucket),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.OutputS3BucketRef,
					Selector:     mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.OutputS3BucketSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.OutputS3Bucket")
			}
			mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.OutputS3Bucket = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.OutputS3BucketRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.TaskInvocationParameters != nil {
		if mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters != nil {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArnRef,
					Selector:     mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArn")
			}
			mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArnRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "MaintenanceWindow", "MaintenanceWindowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WindowID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WindowIDRef,
			Selector:     mg.Spec.ForProvider.WindowIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WindowID")
	}
	mg.Spec.ForProvider.WindowID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WindowIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.ServiceRoleArnRef,
			Selector:     mg.Spec.InitProvider.ServiceRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceRoleArn")
	}
	mg.Spec.InitProvider.ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Targets); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "Instance", "InstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Targets[i3].Values),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.InitProvider.Targets[i3].ValuesRefs,
				Selector:      mg.Spec.InitProvider.Targets[i3].ValuesSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Targets[i3].Values")
		}
		mg.Spec.InitProvider.Targets[i3].Values = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Targets[i3].ValuesRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TaskArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.TaskArnRef,
			Selector:     mg.Spec.InitProvider.TaskArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TaskArn")
	}
	mg.Spec.InitProvider.TaskArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TaskArnRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.TaskInvocationParameters != nil {
		if mg.Spec.InitProvider.TaskInvocationParameters.AutomationParameters != nil {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.TaskInvocationParameters.AutomationParameters.Parameter); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "Instance", "InstanceList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
						CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].Values),
						Extract:       resource.ExtractResourceID(),
						References:    mg.Spec.InitProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].ValuesRefs,
						Selector:      mg.Spec.InitProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].ValuesSelector,
						To:            reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].Values")
				}
				mg.Spec.InitProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].Values = reference.ToPtrValues(mrsp.ResolvedValues)
				mg.Spec.InitProvider.TaskInvocationParameters.AutomationParameters.Parameter[i5].ValuesRefs = mrsp.ResolvedReferences

			}
		}
	}
	if mg.Spec.InitProvider.TaskInvocationParameters != nil {
		if mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters != nil {
			if mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig != nil {
				{
					m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArnRef,
						Selector:     mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArn")
				}
				mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.NotificationConfig.NotificationArnRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.InitProvider.TaskInvocationParameters != nil {
		if mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters != nil {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.OutputS3Bucket),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.OutputS3BucketRef,
					Selector:     mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.OutputS3BucketSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.OutputS3Bucket")
			}
			mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.OutputS3Bucket = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.OutputS3BucketRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.TaskInvocationParameters != nil {
		if mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters != nil {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArnRef,
					Selector:     mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArn")
			}
			mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.TaskInvocationParameters.RunCommandParameters.ServiceRoleArnRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("ssm.aws.upbound.io", "v1beta1", "MaintenanceWindow", "MaintenanceWindowList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WindowID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.WindowIDRef,
			Selector:     mg.Spec.InitProvider.WindowIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WindowID")
	}
	mg.Spec.InitProvider.WindowID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WindowIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourceDataSync.
func (mg *ResourceDataSync) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.S3Destination != nil {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3Destination.BucketName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.S3Destination.BucketNameRef,
				Selector:     mg.Spec.ForProvider.S3Destination.BucketNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.S3Destination.BucketName")
		}
		mg.Spec.ForProvider.S3Destination.BucketName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.S3Destination.BucketNameRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.S3Destination != nil {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3Destination.Region),
				Extract:      resource.ExtractParamPath("region", false),
				Reference:    mg.Spec.ForProvider.S3Destination.RegionRef,
				Selector:     mg.Spec.ForProvider.S3Destination.RegionSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.S3Destination.Region")
		}
		mg.Spec.ForProvider.S3Destination.Region = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.S3Destination.RegionRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.S3Destination != nil {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.S3Destination.BucketName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.S3Destination.BucketNameRef,
				Selector:     mg.Spec.InitProvider.S3Destination.BucketNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.S3Destination.BucketName")
		}
		mg.Spec.InitProvider.S3Destination.BucketName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.S3Destination.BucketNameRef = rsp.ResolvedReference

	}

	return nil
}
