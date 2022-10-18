/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RestAPIEndpointConfigurationObservation struct {
}

type RestAPIEndpointConfigurationParameters struct {

	// A list of endpoint types. This resource currently only supports managing a single value. Valid values: EDGE, REGIONAL or PRIVATE. If unspecified, defaults to EDGE. Must be declared as REGIONAL in non-Commercial partitions. Refer to the documentation for more information on the difference between edge-optimized and regional APIs.
	// +kubebuilder:validation:Required
	Types []*string `json:"types" tf:"types,omitempty"`

	// Set of VPC Endpoint identifiers. It is only supported for PRIVATE endpoint type. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-endpoint-configuration extension vpcEndpointIds property. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	VPCEndpointIds []*string `json:"vpcEndpointIds,omitempty" tf:"vpc_endpoint_ids,omitempty"`
}

type RestAPIObservation struct {

	// Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The creation date of the REST API
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// The execution ARN part to be used in lambda_permission's source_arn
	// when allowing API Gateway to invoke a Lambda function,
	// e.g., arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j, which can be concatenated with allowed stage, method and resource path.
	ExecutionArn *string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`

	// The ID of the REST API
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// JSON formatted policy document that controls access to the API Gateway. It is recommended to use the aws_api_gateway_rest_api_policy resource instead. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-policy extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The resource ID of the REST API's root
	RootResourceID *string `json:"rootResourceId,omitempty" tf:"root_resource_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RestAPIParameters struct {

	// Source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-api-key-source extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	APIKeySource *string `json:"apiKeySource,omitempty" tf:"api_key_source,omitempty"`

	// List of binary media types supported by the REST API. By default, the REST API supports only UTF-8-encoded text payloads. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-binary-media-types extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	BinaryMediaTypes []*string `json:"binaryMediaTypes,omitempty" tf:"binary_media_types,omitempty"`

	// OpenAPI specification that defines the set of routes and integrations to create as part of the REST API. This configuration, and any updates to it, will replace all REST API configuration except values overridden in this resource configuration and other resource updates applied after this resource but before any aws_api_gateway_deployment creation. More information about REST API OpenAPI support can be found in the API Gateway Developer Guide.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Description of the REST API. If importing an OpenAPI specification via the body argument, this corresponds to the info.description field. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint. Defaults to false. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-endpoint-configuration extension disableExecuteApiEndpoint property. If the argument value is true and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	DisableExecuteAPIEndpoint *bool `json:"disableExecuteApiEndpoint,omitempty" tf:"disable_execute_api_endpoint,omitempty"`

	// Configuration block defining API endpoint configuration including endpoint type. Defined below.
	// +kubebuilder:validation:Optional
	EndpointConfiguration []RestAPIEndpointConfigurationParameters `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`

	// Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default). If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-minimum-compression-size extension. If the argument value (except -1) is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	MinimumCompressionSize *float64 `json:"minimumCompressionSize,omitempty" tf:"minimum_compression_size,omitempty"`

	// Name of the REST API. If importing an OpenAPI specification via the body argument, this corresponds to the info.title field. If the argument value is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Map of customizations for importing the specification in the body argument. For example, to exclude DocumentationParts from an imported API, set ignore equal to documentation. Additional documentation, including other parameters such as basepath, can be found in the API Gateway Developer Guide.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RestAPISpec defines the desired state of RestAPI
type RestAPISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RestAPIParameters `json:"forProvider"`
}

// RestAPIStatus defines the observed state of RestAPI.
type RestAPIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RestAPIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RestAPI is the Schema for the RestAPIs API. Manages an API Gateway REST API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RestAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RestAPISpec   `json:"spec"`
	Status            RestAPIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RestAPIList contains a list of RestAPIs
type RestAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestAPI `json:"items"`
}

// Repository type metadata.
var (
	RestAPI_Kind             = "RestAPI"
	RestAPI_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RestAPI_Kind}.String()
	RestAPI_KindAPIVersion   = RestAPI_Kind + "." + CRDGroupVersion.String()
	RestAPI_GroupVersionKind = CRDGroupVersion.WithKind(RestAPI_Kind)
)

func init() {
	SchemeBuilder.Register(&RestAPI{}, &RestAPIList{})
}