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

type ConnectionInitParameters struct {
	ConnectionProperties map[string]*string `json:"connectionPropertiesSecretRef,omitempty" tf:"-"`

	// –  Type of the connection. Valid values: AZURECOSMOS, AZURESQL, BIGQUERY, CUSTOM, JDBC, KAFKA, MARKETPLACE, MONGODB, NETWORK, OPENSEARCH, SNOWFLAKE. Defaults to JDBC.
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// –  Description of the connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  List of criteria that can be used in selecting this connection.
	MatchCriteria []*string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	// Map of physical connection requirements, such as VPC and SecurityGroup. See physical_connection_requirements Block for details.
	PhysicalConnectionRequirements *PhysicalConnectionRequirementsInitParameters `json:"physicalConnectionRequirements,omitempty" tf:"physical_connection_requirements,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConnectionObservation struct {

	// ARN of the Glue Connection.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// –  ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// –  Type of the connection. Valid values: AZURECOSMOS, AZURESQL, BIGQUERY, CUSTOM, JDBC, KAFKA, MARKETPLACE, MONGODB, NETWORK, OPENSEARCH, SNOWFLAKE. Defaults to JDBC.
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// –  Description of the connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Catalog ID and name of the connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// –  List of criteria that can be used in selecting this connection.
	MatchCriteria []*string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	// Map of physical connection requirements, such as VPC and SecurityGroup. See physical_connection_requirements Block for details.
	PhysicalConnectionRequirements *PhysicalConnectionRequirementsObservation `json:"physicalConnectionRequirements,omitempty" tf:"physical_connection_requirements,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ConnectionParameters struct {

	// –  ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	// +kubebuilder:validation:Required
	CatalogID *string `json:"catalogId" tf:"catalog_id,omitempty"`

	// value pairs used as parameters for this connection. For more information, see the AWS Documentation.
	// +kubebuilder:validation:Optional
	ConnectionPropertiesSecretRef *v1.SecretReference `json:"connectionPropertiesSecretRef,omitempty" tf:"-"`

	// –  Type of the connection. Valid values: AZURECOSMOS, AZURESQL, BIGQUERY, CUSTOM, JDBC, KAFKA, MARKETPLACE, MONGODB, NETWORK, OPENSEARCH, SNOWFLAKE. Defaults to JDBC.
	// +kubebuilder:validation:Optional
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// –  Description of the connection.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  List of criteria that can be used in selecting this connection.
	// +kubebuilder:validation:Optional
	MatchCriteria []*string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	// Map of physical connection requirements, such as VPC and SecurityGroup. See physical_connection_requirements Block for details.
	// +kubebuilder:validation:Optional
	PhysicalConnectionRequirements *PhysicalConnectionRequirementsParameters `json:"physicalConnectionRequirements,omitempty" tf:"physical_connection_requirements,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PhysicalConnectionRequirementsInitParameters struct {

	// The availability zone of the connection. This field is redundant and implied by subnet_id, but is currently an api requirement.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("availability_zone",false)
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Reference to a Subnet in ec2 to populate availabilityZone.
	// +kubebuilder:validation:Optional
	AvailabilityZoneRef *v1.Reference `json:"availabilityZoneRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate availabilityZone.
	// +kubebuilder:validation:Optional
	AvailabilityZoneSelector *v1.Selector `json:"availabilityZoneSelector,omitempty" tf:"-"`

	// The security group ID list used by the connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +listType=set
	SecurityGroupIDList []*string `json:"securityGroupIdList,omitempty" tf:"security_group_id_list,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIdList.
	// +kubebuilder:validation:Optional
	SecurityGroupIDListRefs []v1.Reference `json:"securityGroupIdListRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIdList.
	// +kubebuilder:validation:Optional
	SecurityGroupIDListSelector *v1.Selector `json:"securityGroupIdListSelector,omitempty" tf:"-"`

	// The subnet ID used by the connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type PhysicalConnectionRequirementsObservation struct {

	// The availability zone of the connection. This field is redundant and implied by subnet_id, but is currently an api requirement.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The security group ID list used by the connection.
	// +listType=set
	SecurityGroupIDList []*string `json:"securityGroupIdList,omitempty" tf:"security_group_id_list,omitempty"`

	// The subnet ID used by the connection.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type PhysicalConnectionRequirementsParameters struct {

	// The availability zone of the connection. This field is redundant and implied by subnet_id, but is currently an api requirement.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("availability_zone",false)
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Reference to a Subnet in ec2 to populate availabilityZone.
	// +kubebuilder:validation:Optional
	AvailabilityZoneRef *v1.Reference `json:"availabilityZoneRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate availabilityZone.
	// +kubebuilder:validation:Optional
	AvailabilityZoneSelector *v1.Selector `json:"availabilityZoneSelector,omitempty" tf:"-"`

	// The security group ID list used by the connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIDList []*string `json:"securityGroupIdList,omitempty" tf:"security_group_id_list,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIdList.
	// +kubebuilder:validation:Optional
	SecurityGroupIDListRefs []v1.Reference `json:"securityGroupIdListRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIdList.
	// +kubebuilder:validation:Optional
	SecurityGroupIDListSelector *v1.Selector `json:"securityGroupIdListSelector,omitempty" tf:"-"`

	// The subnet ID used by the connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
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
	InitProvider ConnectionInitParameters `json:"initProvider,omitempty"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Connection is the Schema for the Connections API. Provides an Glue Connection resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
