// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Use the Bastion API to provide restricted and time-limited access to target resources that don't have public endpoints. Bastions let authorized users connect from specific IP addresses to target resources using Secure Shell (SSH) sessions. For more information, see the Bastion documentation (https://docs.cloud.oracle.com/iaas/Content/Bastion/home.htm).
//

package bastion

import (
	"strings"
)

// ActionTypeEnum Enum with underlying type: string
type ActionTypeEnum string

// Set of constants representing the allowable values for ActionTypeEnum
const (
	ActionTypeCreated    ActionTypeEnum = "CREATED"
	ActionTypeUpdated    ActionTypeEnum = "UPDATED"
	ActionTypeDeleted    ActionTypeEnum = "DELETED"
	ActionTypeInProgress ActionTypeEnum = "IN_PROGRESS"
	ActionTypeRelated    ActionTypeEnum = "RELATED"
)

var mappingActionTypeEnum = map[string]ActionTypeEnum{
	"CREATED":     ActionTypeCreated,
	"UPDATED":     ActionTypeUpdated,
	"DELETED":     ActionTypeDeleted,
	"IN_PROGRESS": ActionTypeInProgress,
	"RELATED":     ActionTypeRelated,
}

var mappingActionTypeEnumLowerCase = map[string]ActionTypeEnum{
	"created":     ActionTypeCreated,
	"updated":     ActionTypeUpdated,
	"deleted":     ActionTypeDeleted,
	"in_progress": ActionTypeInProgress,
	"related":     ActionTypeRelated,
}

// GetActionTypeEnumValues Enumerates the set of values for ActionTypeEnum
func GetActionTypeEnumValues() []ActionTypeEnum {
	values := make([]ActionTypeEnum, 0)
	for _, v := range mappingActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetActionTypeEnumStringValues Enumerates the set of values in String for ActionTypeEnum
func GetActionTypeEnumStringValues() []string {
	return []string{
		"CREATED",
		"UPDATED",
		"DELETED",
		"IN_PROGRESS",
		"RELATED",
	}
}

// GetMappingActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionTypeEnum(val string) (ActionTypeEnum, bool) {
	enum, ok := mappingActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
