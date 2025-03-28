// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PrivateDNSNameConfigurationInitParameters struct {
}

type PrivateDNSNameConfigurationObservation struct {

	// Name of the record subdomain the service provider needs to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The state of the VPC endpoint service.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Endpoint service verification type, for example TXT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Value the service provider adds to the private DNS name domain record before verification.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PrivateDNSNameConfigurationParameters struct {
}

type VPCEndpointServiceInitParameters struct {

	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - true or false.
	AcceptanceRequired *bool `json:"acceptanceRequired,omitempty" tf:"acceptance_required,omitempty"`

	// Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta2.LB
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +listType=set
	GatewayLoadBalancerArns []*string `json:"gatewayLoadBalancerArns,omitempty" tf:"gateway_load_balancer_arns,omitempty"`

	// References to LB in elbv2 to populate gatewayLoadBalancerArns.
	// +kubebuilder:validation:Optional
	GatewayLoadBalancerArnsRefs []v1.Reference `json:"gatewayLoadBalancerArnsRefs,omitempty" tf:"-"`

	// Selector for a list of LB in elbv2 to populate gatewayLoadBalancerArns.
	// +kubebuilder:validation:Optional
	GatewayLoadBalancerArnsSelector *v1.Selector `json:"gatewayLoadBalancerArnsSelector,omitempty" tf:"-"`

	// Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta2.LB
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +listType=set
	NetworkLoadBalancerArns []*string `json:"networkLoadBalancerArns,omitempty" tf:"network_load_balancer_arns,omitempty"`

	// References to LB in elbv2 to populate networkLoadBalancerArns.
	// +kubebuilder:validation:Optional
	NetworkLoadBalancerArnsRefs []v1.Reference `json:"networkLoadBalancerArnsRefs,omitempty" tf:"-"`

	// Selector for a list of LB in elbv2 to populate networkLoadBalancerArns.
	// +kubebuilder:validation:Optional
	NetworkLoadBalancerArnsSelector *v1.Selector `json:"networkLoadBalancerArnsSelector,omitempty" tf:"-"`

	// The private DNS name for the service.
	PrivateDNSName *string `json:"privateDnsName,omitempty" tf:"private_dns_name,omitempty"`

	// The supported IP address types. The possible values are ipv4 and ipv6.
	// +listType=set
	SupportedIPAddressTypes []*string `json:"supportedIpAddressTypes,omitempty" tf:"supported_ip_address_types,omitempty"`

	// The set of regions from which service consumers can access the service.
	// +listType=set
	SupportedRegions []*string `json:"supportedRegions,omitempty" tf:"supported_regions,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VPCEndpointServiceObservation struct {

	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - true or false.
	AcceptanceRequired *bool `json:"acceptanceRequired,omitempty" tf:"acceptance_required,omitempty"`

	// The ARNs of one or more principals allowed to discover the endpoint service.
	// +listType=set
	AllowedPrincipals []*string `json:"allowedPrincipals,omitempty" tf:"allowed_principals,omitempty"`

	// The Amazon Resource Name (ARN) of the VPC endpoint service.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A set of Availability Zones in which the service is available.
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// A set of DNS names for the service.
	// +listType=set
	BaseEndpointDNSNames []*string `json:"baseEndpointDnsNames,omitempty" tf:"base_endpoint_dns_names,omitempty"`

	// Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
	// +listType=set
	GatewayLoadBalancerArns []*string `json:"gatewayLoadBalancerArns,omitempty" tf:"gateway_load_balancer_arns,omitempty"`

	// The ID of the VPC endpoint service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether or not the service manages its VPC endpoints - true or false.
	ManagesVPCEndpoints *bool `json:"managesVpcEndpoints,omitempty" tf:"manages_vpc_endpoints,omitempty"`

	// Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
	// +listType=set
	NetworkLoadBalancerArns []*string `json:"networkLoadBalancerArns,omitempty" tf:"network_load_balancer_arns,omitempty"`

	// The private DNS name for the service.
	PrivateDNSName *string `json:"privateDnsName,omitempty" tf:"private_dns_name,omitempty"`

	// List of objects containing information about the endpoint service private DNS name configuration.
	PrivateDNSNameConfiguration []PrivateDNSNameConfigurationObservation `json:"privateDnsNameConfiguration,omitempty" tf:"private_dns_name_configuration,omitempty"`

	// The service name.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The service type, Gateway or Interface.
	ServiceType *string `json:"serviceType,omitempty" tf:"service_type,omitempty"`

	// The state of the VPC endpoint service.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The supported IP address types. The possible values are ipv4 and ipv6.
	// +listType=set
	SupportedIPAddressTypes []*string `json:"supportedIpAddressTypes,omitempty" tf:"supported_ip_address_types,omitempty"`

	// The set of regions from which service consumers can access the service.
	// +listType=set
	SupportedRegions []*string `json:"supportedRegions,omitempty" tf:"supported_regions,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VPCEndpointServiceParameters struct {

	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - true or false.
	// +kubebuilder:validation:Optional
	AcceptanceRequired *bool `json:"acceptanceRequired,omitempty" tf:"acceptance_required,omitempty"`

	// Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta2.LB
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	GatewayLoadBalancerArns []*string `json:"gatewayLoadBalancerArns,omitempty" tf:"gateway_load_balancer_arns,omitempty"`

	// References to LB in elbv2 to populate gatewayLoadBalancerArns.
	// +kubebuilder:validation:Optional
	GatewayLoadBalancerArnsRefs []v1.Reference `json:"gatewayLoadBalancerArnsRefs,omitempty" tf:"-"`

	// Selector for a list of LB in elbv2 to populate gatewayLoadBalancerArns.
	// +kubebuilder:validation:Optional
	GatewayLoadBalancerArnsSelector *v1.Selector `json:"gatewayLoadBalancerArnsSelector,omitempty" tf:"-"`

	// Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta2.LB
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	NetworkLoadBalancerArns []*string `json:"networkLoadBalancerArns,omitempty" tf:"network_load_balancer_arns,omitempty"`

	// References to LB in elbv2 to populate networkLoadBalancerArns.
	// +kubebuilder:validation:Optional
	NetworkLoadBalancerArnsRefs []v1.Reference `json:"networkLoadBalancerArnsRefs,omitempty" tf:"-"`

	// Selector for a list of LB in elbv2 to populate networkLoadBalancerArns.
	// +kubebuilder:validation:Optional
	NetworkLoadBalancerArnsSelector *v1.Selector `json:"networkLoadBalancerArnsSelector,omitempty" tf:"-"`

	// The private DNS name for the service.
	// +kubebuilder:validation:Optional
	PrivateDNSName *string `json:"privateDnsName,omitempty" tf:"private_dns_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The supported IP address types. The possible values are ipv4 and ipv6.
	// +kubebuilder:validation:Optional
	// +listType=set
	SupportedIPAddressTypes []*string `json:"supportedIpAddressTypes,omitempty" tf:"supported_ip_address_types,omitempty"`

	// The set of regions from which service consumers can access the service.
	// +kubebuilder:validation:Optional
	// +listType=set
	SupportedRegions []*string `json:"supportedRegions,omitempty" tf:"supported_regions,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCEndpointServiceSpec defines the desired state of VPCEndpointService
type VPCEndpointServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointServiceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VPCEndpointServiceInitParameters `json:"initProvider,omitempty"`
}

// VPCEndpointServiceStatus defines the observed state of VPCEndpointService.
type VPCEndpointServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPCEndpointService is the Schema for the VPCEndpointServices API. Provides a VPC Endpoint Service resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpointService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.acceptanceRequired) || (has(self.initProvider) && has(self.initProvider.acceptanceRequired))",message="spec.forProvider.acceptanceRequired is a required parameter"
	Spec   VPCEndpointServiceSpec   `json:"spec"`
	Status VPCEndpointServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointServiceList contains a list of VPCEndpointServices
type VPCEndpointServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointService `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpointService_Kind             = "VPCEndpointService"
	VPCEndpointService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpointService_Kind}.String()
	VPCEndpointService_KindAPIVersion   = VPCEndpointService_Kind + "." + CRDGroupVersion.String()
	VPCEndpointService_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpointService_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpointService{}, &VPCEndpointServiceList{})
}
