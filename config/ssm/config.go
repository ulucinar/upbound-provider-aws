package ssm

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for ssm group.
func Configure(p *config.Provider) {
	p.Resources["aws_ssm_parameter"] = &config.Resource{
		Kind:         "Parameter",
		Name:         "aws_ssm_parameter",
		ShortGroup:   "ssm",
		Version:      "v1beta1",
		UseAsync:     true,
		ExternalName: config.NameAsIdentifier,
		Sensitive: config.Sensitive{
			AdditionalConnectionDetailsFn: config.NopAdditionalConnectionDetails,
		},
	}
}
