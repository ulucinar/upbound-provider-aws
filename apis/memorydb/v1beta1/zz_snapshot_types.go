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

type ClusterConfigurationInitParameters struct {
}

type ClusterConfigurationObservation struct {

	// Description for the cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The engine that will run on cluster nodes.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Version number of the engine used by the cluster.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The weekly time range during which maintenance on the cluster is performed.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// Name of the snapshot. Conflicts with name_prefix.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Compute and memory capacity of the nodes in the cluster.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Number of shards in the cluster.
	NumShards *float64 `json:"numShards,omitempty" tf:"num_shards,omitempty"`

	// Name of the parameter group associated with the cluster.
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// Port number on which the cluster accepts connections.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Number of days for which MemoryDB retains automatic snapshots before deleting them.
	SnapshotRetentionLimit *float64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of the shard.
	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`

	// Name of the subnet group used by the cluster.
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// ARN of the SNS topic to which cluster notifications are sent.
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`

	// The VPC in which the cluster exists.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type ClusterConfigurationParameters struct {
}

type SnapshotInitParameters struct {

	// Name of the MemoryDB cluster to take a snapshot of.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/memorydb/v1beta1.Cluster
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in memorydb to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in memorydb to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// ARN of the KMS key used to encrypt the snapshot at rest.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SnapshotObservation struct {

	// The ARN of the snapshot.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The configuration of the cluster from which the snapshot was taken.
	ClusterConfiguration []ClusterConfigurationObservation `json:"clusterConfiguration,omitempty" tf:"cluster_configuration,omitempty"`

	// Name of the MemoryDB cluster to take a snapshot of.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// The name of the snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the KMS key used to encrypt the snapshot at rest.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Indicates whether the snapshot is from an automatic backup (automated) or was created manually (manual).
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SnapshotParameters struct {

	// Name of the MemoryDB cluster to take a snapshot of.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/memorydb/v1beta1.Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in memorydb to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in memorydb to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// ARN of the KMS key used to encrypt the snapshot at rest.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
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
	InitProvider SnapshotInitParameters `json:"initProvider,omitempty"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Snapshot is the Schema for the Snapshots API. Provides a MemoryDB Snapshot.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
