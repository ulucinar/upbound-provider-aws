// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type APIStagesInitParameters struct {

	// API Id of the associated API stage in a usage plan.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta2.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// API stage name of the associated API stage in a usage plan.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta2.Stage
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("stage_name",false)
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// Reference to a Stage in apigateway to populate stage.
	// +kubebuilder:validation:Optional
	StageRef *v1.Reference `json:"stageRef,omitempty" tf:"-"`

	// Selector for a Stage in apigateway to populate stage.
	// +kubebuilder:validation:Optional
	StageSelector *v1.Selector `json:"stageSelector,omitempty" tf:"-"`

	// The throttling limits of the usage plan.
	Throttle []ThrottleInitParameters `json:"throttle,omitempty" tf:"throttle,omitempty"`
}

type APIStagesObservation struct {

	// API Id of the associated API stage in a usage plan.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// API stage name of the associated API stage in a usage plan.
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// The throttling limits of the usage plan.
	Throttle []ThrottleObservation `json:"throttle,omitempty" tf:"throttle,omitempty"`
}

type APIStagesParameters struct {

	// API Id of the associated API stage in a usage plan.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta2.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// API stage name of the associated API stage in a usage plan.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta2.Stage
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("stage_name",false)
	// +kubebuilder:validation:Optional
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// Reference to a Stage in apigateway to populate stage.
	// +kubebuilder:validation:Optional
	StageRef *v1.Reference `json:"stageRef,omitempty" tf:"-"`

	// Selector for a Stage in apigateway to populate stage.
	// +kubebuilder:validation:Optional
	StageSelector *v1.Selector `json:"stageSelector,omitempty" tf:"-"`

	// The throttling limits of the usage plan.
	// +kubebuilder:validation:Optional
	Throttle []ThrottleParameters `json:"throttle,omitempty" tf:"throttle,omitempty"`
}

type QuotaSettingsInitParameters struct {

	// Maximum number of requests that can be made in a given time period.
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	// Number of requests subtracted from the given limit in the initial time period.
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// Time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
	Period *string `json:"period,omitempty" tf:"period,omitempty"`
}

type QuotaSettingsObservation struct {

	// Maximum number of requests that can be made in a given time period.
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	// Number of requests subtracted from the given limit in the initial time period.
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// Time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
	Period *string `json:"period,omitempty" tf:"period,omitempty"`
}

type QuotaSettingsParameters struct {

	// Maximum number of requests that can be made in a given time period.
	// +kubebuilder:validation:Optional
	Limit *float64 `json:"limit" tf:"limit,omitempty"`

	// Number of requests subtracted from the given limit in the initial time period.
	// +kubebuilder:validation:Optional
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// Time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
	// +kubebuilder:validation:Optional
	Period *string `json:"period" tf:"period,omitempty"`
}

type ThrottleInitParameters struct {

	// The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
	BurstLimit *float64 `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`

	// Method to apply the throttle settings for. Specfiy the path and method, for example /test/GET.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The API request steady-state rate limit.
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type ThrottleObservation struct {

	// The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
	BurstLimit *float64 `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`

	// Method to apply the throttle settings for. Specfiy the path and method, for example /test/GET.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The API request steady-state rate limit.
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type ThrottleParameters struct {

	// The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
	// +kubebuilder:validation:Optional
	BurstLimit *float64 `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`

	// Method to apply the throttle settings for. Specfiy the path and method, for example /test/GET.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`

	// The API request steady-state rate limit.
	// +kubebuilder:validation:Optional
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type ThrottleSettingsInitParameters struct {

	// The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
	BurstLimit *float64 `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`

	// The API request steady-state rate limit.
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type ThrottleSettingsObservation struct {

	// The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
	BurstLimit *float64 `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`

	// The API request steady-state rate limit.
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type ThrottleSettingsParameters struct {

	// The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
	// +kubebuilder:validation:Optional
	BurstLimit *float64 `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`

	// The API request steady-state rate limit.
	// +kubebuilder:validation:Optional
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type UsagePlanInitParameters struct {

	// Associated API stages of the usage plan.
	APIStages []APIStagesInitParameters `json:"apiStages,omitempty" tf:"api_stages,omitempty"`

	// Description of a usage plan.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the usage plan.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode *string `json:"productCode,omitempty" tf:"product_code,omitempty"`

	// The quota settings of the usage plan.
	QuotaSettings *QuotaSettingsInitParameters `json:"quotaSettings,omitempty" tf:"quota_settings,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The throttling limits of the usage plan.
	ThrottleSettings *ThrottleSettingsInitParameters `json:"throttleSettings,omitempty" tf:"throttle_settings,omitempty"`
}

type UsagePlanObservation struct {

	// Associated API stages of the usage plan.
	APIStages []APIStagesObservation `json:"apiStages,omitempty" tf:"api_stages,omitempty"`

	// ARN
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of a usage plan.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the API resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the usage plan.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	ProductCode *string `json:"productCode,omitempty" tf:"product_code,omitempty"`

	// The quota settings of the usage plan.
	QuotaSettings *QuotaSettingsObservation `json:"quotaSettings,omitempty" tf:"quota_settings,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The throttling limits of the usage plan.
	ThrottleSettings *ThrottleSettingsObservation `json:"throttleSettings,omitempty" tf:"throttle_settings,omitempty"`
}

type UsagePlanParameters struct {

	// Associated API stages of the usage plan.
	// +kubebuilder:validation:Optional
	APIStages []APIStagesParameters `json:"apiStages,omitempty" tf:"api_stages,omitempty"`

	// Description of a usage plan.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the usage plan.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// AWS Marketplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
	// +kubebuilder:validation:Optional
	ProductCode *string `json:"productCode,omitempty" tf:"product_code,omitempty"`

	// The quota settings of the usage plan.
	// +kubebuilder:validation:Optional
	QuotaSettings *QuotaSettingsParameters `json:"quotaSettings,omitempty" tf:"quota_settings,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The throttling limits of the usage plan.
	// +kubebuilder:validation:Optional
	ThrottleSettings *ThrottleSettingsParameters `json:"throttleSettings,omitempty" tf:"throttle_settings,omitempty"`
}

// UsagePlanSpec defines the desired state of UsagePlan
type UsagePlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UsagePlanParameters `json:"forProvider"`
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
	InitProvider UsagePlanInitParameters `json:"initProvider,omitempty"`
}

// UsagePlanStatus defines the observed state of UsagePlan.
type UsagePlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UsagePlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// UsagePlan is the Schema for the UsagePlans API. Provides an API Gateway Usage Plan.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UsagePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   UsagePlanSpec   `json:"spec"`
	Status UsagePlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsagePlanList contains a list of UsagePlans
type UsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UsagePlan `json:"items"`
}

// Repository type metadata.
var (
	UsagePlan_Kind             = "UsagePlan"
	UsagePlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UsagePlan_Kind}.String()
	UsagePlan_KindAPIVersion   = UsagePlan_Kind + "." + CRDGroupVersion.String()
	UsagePlan_GroupVersionKind = CRDGroupVersion.WithKind(UsagePlan_Kind)
)

func init() {
	SchemeBuilder.Register(&UsagePlan{}, &UsagePlanList{})
}