// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MigrationAsset Description of the migration asset.
type MigrationAsset struct {

	// Asset ID generated by mirgration service. It is used in the mirgration service pipeline.
	Id *string `mandatory:"true" json:"id"`

	// The type of asset referenced for inventory.
	Type *string `mandatory:"true" json:"type"`

	// The current state of the migration asset.
	LifecycleState MigrationAssetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the migration asset was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// OCID of the associated migration.
	MigrationId *string `mandatory:"true" json:"migrationId"`

	// Availability domain
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// Replication compartment identifier
	ReplicationCompartmentId *string `mandatory:"true" json:"replicationCompartmentId"`

	// Name of snapshot bucket
	SnapShotBucketName *string `mandatory:"true" json:"snapShotBucketName"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time when the migration asset was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Key-value pair representing disks ID mapped to the OCIDs of replicated or hydration server volume snapshots.
	// Example: `{"bar-key": "value"}`
	Snapshots map[string]HydratedVolume `mandatory:"false" json:"snapshots"`

	// The parent snapshot of the migration asset to be used by the replication task.
	ParentSnapshot *string `mandatory:"false" json:"parentSnapshot"`

	// Key-value pair representing asset metadata keys and values scoped to a namespace.
	// Example: `{"bar-key": "value"}`
	SourceAssetData map[string]interface{} `mandatory:"false" json:"sourceAssetData"`

	// List of notifications
	Notifications []MigrationAssetNotificationsEnum `mandatory:"false" json:"notifications,omitempty"`

	// OCID that is referenced to an asset for an inventory.
	SourceAssetId *string `mandatory:"false" json:"sourceAssetId"`

	// Replication schedule identifier
	ReplicationScheduleId *string `mandatory:"false" json:"replicationScheduleId"`

	// Tenancy identifier
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// List of migration assets that depend on the asset.
	DependedOnBy []string `mandatory:"false" json:"dependedOnBy"`

	// List of migration assets that depends on the asset.
	DependsOn []string `mandatory:"false" json:"dependsOn"`
}

func (m MigrationAsset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MigrationAsset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMigrationAssetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMigrationAssetLifecycleStateEnumStringValues(), ",")))
	}

	for _, val := range m.Notifications {
		if _, ok := GetMappingMigrationAssetNotificationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Notifications: %s. Supported values are: %s.", val, strings.Join(GetMigrationAssetNotificationsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MigrationAssetLifecycleStateEnum Enum with underlying type: string
type MigrationAssetLifecycleStateEnum string

// Set of constants representing the allowable values for MigrationAssetLifecycleStateEnum
const (
	MigrationAssetLifecycleStateCreating       MigrationAssetLifecycleStateEnum = "CREATING"
	MigrationAssetLifecycleStateUpdating       MigrationAssetLifecycleStateEnum = "UPDATING"
	MigrationAssetLifecycleStateNeedsAttention MigrationAssetLifecycleStateEnum = "NEEDS_ATTENTION"
	MigrationAssetLifecycleStateActive         MigrationAssetLifecycleStateEnum = "ACTIVE"
	MigrationAssetLifecycleStateDeleting       MigrationAssetLifecycleStateEnum = "DELETING"
	MigrationAssetLifecycleStateDeleted        MigrationAssetLifecycleStateEnum = "DELETED"
	MigrationAssetLifecycleStateFailed         MigrationAssetLifecycleStateEnum = "FAILED"
)

var mappingMigrationAssetLifecycleStateEnum = map[string]MigrationAssetLifecycleStateEnum{
	"CREATING":        MigrationAssetLifecycleStateCreating,
	"UPDATING":        MigrationAssetLifecycleStateUpdating,
	"NEEDS_ATTENTION": MigrationAssetLifecycleStateNeedsAttention,
	"ACTIVE":          MigrationAssetLifecycleStateActive,
	"DELETING":        MigrationAssetLifecycleStateDeleting,
	"DELETED":         MigrationAssetLifecycleStateDeleted,
	"FAILED":          MigrationAssetLifecycleStateFailed,
}

var mappingMigrationAssetLifecycleStateEnumLowerCase = map[string]MigrationAssetLifecycleStateEnum{
	"creating":        MigrationAssetLifecycleStateCreating,
	"updating":        MigrationAssetLifecycleStateUpdating,
	"needs_attention": MigrationAssetLifecycleStateNeedsAttention,
	"active":          MigrationAssetLifecycleStateActive,
	"deleting":        MigrationAssetLifecycleStateDeleting,
	"deleted":         MigrationAssetLifecycleStateDeleted,
	"failed":          MigrationAssetLifecycleStateFailed,
}

// GetMigrationAssetLifecycleStateEnumValues Enumerates the set of values for MigrationAssetLifecycleStateEnum
func GetMigrationAssetLifecycleStateEnumValues() []MigrationAssetLifecycleStateEnum {
	values := make([]MigrationAssetLifecycleStateEnum, 0)
	for _, v := range mappingMigrationAssetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationAssetLifecycleStateEnumStringValues Enumerates the set of values in String for MigrationAssetLifecycleStateEnum
func GetMigrationAssetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"NEEDS_ATTENTION",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMigrationAssetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrationAssetLifecycleStateEnum(val string) (MigrationAssetLifecycleStateEnum, bool) {
	enum, ok := mappingMigrationAssetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MigrationAssetNotificationsEnum Enum with underlying type: string
type MigrationAssetNotificationsEnum string

// Set of constants representing the allowable values for MigrationAssetNotificationsEnum
const (
	MigrationAssetNotificationsOutOfDate     MigrationAssetNotificationsEnum = "OUT_OF_DATE"
	MigrationAssetNotificationsSourceRemoved MigrationAssetNotificationsEnum = "SOURCE_REMOVED"
)

var mappingMigrationAssetNotificationsEnum = map[string]MigrationAssetNotificationsEnum{
	"OUT_OF_DATE":    MigrationAssetNotificationsOutOfDate,
	"SOURCE_REMOVED": MigrationAssetNotificationsSourceRemoved,
}

var mappingMigrationAssetNotificationsEnumLowerCase = map[string]MigrationAssetNotificationsEnum{
	"out_of_date":    MigrationAssetNotificationsOutOfDate,
	"source_removed": MigrationAssetNotificationsSourceRemoved,
}

// GetMigrationAssetNotificationsEnumValues Enumerates the set of values for MigrationAssetNotificationsEnum
func GetMigrationAssetNotificationsEnumValues() []MigrationAssetNotificationsEnum {
	values := make([]MigrationAssetNotificationsEnum, 0)
	for _, v := range mappingMigrationAssetNotificationsEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationAssetNotificationsEnumStringValues Enumerates the set of values in String for MigrationAssetNotificationsEnum
func GetMigrationAssetNotificationsEnumStringValues() []string {
	return []string{
		"OUT_OF_DATE",
		"SOURCE_REMOVED",
	}
}

// GetMappingMigrationAssetNotificationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrationAssetNotificationsEnum(val string) (MigrationAssetNotificationsEnum, bool) {
	enum, ok := mappingMigrationAssetNotificationsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
