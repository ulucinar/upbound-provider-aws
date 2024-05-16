// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package mq

import (
	"encoding/json"
	"fmt"
	"strings"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/apis/mq/v1beta2"
)

const (
	annotBrokerBootstrapUsers = "conversion.upjet.crossplane.io/spec.forProvider.bootstrapUsers"

	argBrokerBootstrapUsers = "bootstrap_users"
	argBrokerUser           = "user"
)

// Configure adds configurations for the mq group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_mq_broker", func(r *config.Resource) {
		r.References["security_groups"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.UseAsync = true
		// TODO(aru): looks like currently angryjet cannot handle references
		//  for non-string struct fields. `configuration.revision` is a
		//  float64 field. Thus here we remove the automatically injected
		//  cross-resource reference from example manifests.
		delete(r.References, "configuration.revision")

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if instances, ok := attr["instances"].([]any); ok {
				for i, inst := range instances {
					if instance, ok := inst.(map[string]any); ok {
						if cu, ok := instance["console_url"].(string); ok {
							key := fmt.Sprintf("instance_%d_console_url", i)
							conn[key] = []byte(cu)
						}
						if ip, ok := instance["ip_address"].(string); ok {
							key := fmt.Sprintf("instance_%d_ip_address", i)
							conn[key] = []byte(ip)
						}
						if endpoints, ok := instance["endpoints"].([]any); ok && len(endpoints) > 0 {
							for j, endpoint := range endpoints {
								key := fmt.Sprintf("instance_%d_endpoint_%d", i, j)
								conn[key] = []byte(endpoint.(string))
							}
						}
					}
				}
			}
			return conn, nil
		}

		r.TerraformResource.Schema[argBrokerBootstrapUsers] = r.TerraformResource.Schema[argBrokerUser]
		r.TerraformResource.Schema[argBrokerUser].Required = false
		r.TerraformResource.Schema[argBrokerUser].Optional = true
		r.TerraformResource.Schema[argBrokerBootstrapUsers].Required = false
		r.TerraformResource.Schema[argBrokerBootstrapUsers].Optional = true
		r.Version = "v1beta2"
		r.PreviousVersions = []string{"v1beta1"}
		r.SetCRDStorageVersion("v1beta1")
		r.Conversions = append(r.Conversions,
			conversion.NewCustomConverter("v1beta2", "v1beta1", func(src, target xpresource.Managed) error {
				sBroker := src.(*v1beta2.Broker)
				an := target.GetAnnotations()
				if an == nil {
					an = make(map[string]string)
				}

				if sBroker.Spec.ForProvider.BootstrapUsers == nil {
					if an[annotBrokerBootstrapUsers] == "" {
						return nil
					}
					delete(an, annotBrokerBootstrapUsers)
				} else {
					buff, err := json.Marshal(sBroker.Spec.ForProvider.BootstrapUsers)
					if err != nil {
						return errors.Wrap(err, "failed to marshal spec.forProvider.boostrapUsers into JSON")
					}
					an[annotBrokerBootstrapUsers] = string(buff)
				}
				target.SetAnnotations(an)
				return nil
			}),
			conversion.NewCustomConverter("v1beta1", "v1beta2", func(src, target xpresource.Managed) error {
				an := src.GetAnnotations()
				if an == nil {
					an = make(map[string]string)
				}

				if an[annotBrokerBootstrapUsers] == "" {
					an[annotBrokerBootstrapUsers] = "[]"
				}
				return errors.Wrapf(json.Unmarshal([]byte(an[annotBrokerBootstrapUsers]), &target.(*v1beta2.Broker).Spec.ForProvider.BootstrapUsers), "failed to unmarshal the %s annotations value as JSON: %s", annotBrokerBootstrapUsers, an[annotBrokerBootstrapUsers])
			}))
		r.TerraformConversions = []config.TerraformConversion{&brokerBootstrapUserConversion{}}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil {
				return diff, nil
			}
			for k := range diff.Attributes {
				if strings.HasPrefix(k, fmt.Sprintf("%s.", argBrokerBootstrapUsers)) {
					delete(diff.Attributes, k)
				}
			}

			if state == nil || state.ID == "" || config == nil || config.Raw == nil || config.Raw[argBrokerBootstrapUsers] == nil || config.Raw[argBrokerUser] == nil {
				return diff, nil
			}

			userNames := make(map[string]struct{})
			for _, v := range config.Raw[argBrokerUser].([]any) {
				u := v.(map[string]any)
				userNames[u["username"].(string)] = struct{}{}
			}

			var bootstrapUserNames []string
			for _, v := range config.Raw[argBrokerBootstrapUsers].([]any) {
				u := v.(map[string]any)
				bootstrapUserNames = append(bootstrapUserNames, u["username"].(string))
			}

			for _, u := range bootstrapUserNames {
				delete(userNames, u)
			}
			if len(userNames) == 0 {
				// then all users are bootstrap users
				for k := range diff.Attributes {
					if strings.HasPrefix(k, "user.") {
						delete(diff.Attributes, k)
					}
				}
			}
			return diff, nil
		}

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{argBrokerUser},
		}
	})

	p.AddResourceConfigurator("aws_mq_user", func(r *config.Resource) {
		r.References["broker_id"] = config.Reference{
			TerraformName: "aws_mq_broker",
		}
		r.Version = "v1alpha1"
	})
}

type brokerBootstrapUserConversion struct{}

func (b *brokerBootstrapUserConversion) Convert(params map[string]any, _ *config.Resource, mode config.Mode) (map[string]any, error) {
	if mode != config.ToTerraform {
		return params, nil
	}

	bUsers, ok := params[argBrokerBootstrapUsers]
	if !ok {
		return params, nil
	}
	params[argBrokerUser] = bUsers
	return params, nil
}
