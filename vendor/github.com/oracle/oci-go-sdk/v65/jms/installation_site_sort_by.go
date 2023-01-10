// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// InstallationSiteSortByEnum Enum with underlying type: string
type InstallationSiteSortByEnum string

// Set of constants representing the allowable values for InstallationSiteSortByEnum
const (
	InstallationSiteSortByManagedInstanceId           InstallationSiteSortByEnum = "managedInstanceId"
	InstallationSiteSortByJreDistribution             InstallationSiteSortByEnum = "jreDistribution"
	InstallationSiteSortByJreVendor                   InstallationSiteSortByEnum = "jreVendor"
	InstallationSiteSortByJreVersion                  InstallationSiteSortByEnum = "jreVersion"
	InstallationSiteSortByPath                        InstallationSiteSortByEnum = "path"
	InstallationSiteSortByApproximateApplicationCount InstallationSiteSortByEnum = "approximateApplicationCount"
	InstallationSiteSortByOsName                      InstallationSiteSortByEnum = "osName"
	InstallationSiteSortBySecurityStatus              InstallationSiteSortByEnum = "securityStatus"
)

var mappingInstallationSiteSortByEnum = map[string]InstallationSiteSortByEnum{
	"managedInstanceId":           InstallationSiteSortByManagedInstanceId,
	"jreDistribution":             InstallationSiteSortByJreDistribution,
	"jreVendor":                   InstallationSiteSortByJreVendor,
	"jreVersion":                  InstallationSiteSortByJreVersion,
	"path":                        InstallationSiteSortByPath,
	"approximateApplicationCount": InstallationSiteSortByApproximateApplicationCount,
	"osName":                      InstallationSiteSortByOsName,
	"securityStatus":              InstallationSiteSortBySecurityStatus,
}

var mappingInstallationSiteSortByEnumLowerCase = map[string]InstallationSiteSortByEnum{
	"managedinstanceid":           InstallationSiteSortByManagedInstanceId,
	"jredistribution":             InstallationSiteSortByJreDistribution,
	"jrevendor":                   InstallationSiteSortByJreVendor,
	"jreversion":                  InstallationSiteSortByJreVersion,
	"path":                        InstallationSiteSortByPath,
	"approximateapplicationcount": InstallationSiteSortByApproximateApplicationCount,
	"osname":                      InstallationSiteSortByOsName,
	"securitystatus":              InstallationSiteSortBySecurityStatus,
}

// GetInstallationSiteSortByEnumValues Enumerates the set of values for InstallationSiteSortByEnum
func GetInstallationSiteSortByEnumValues() []InstallationSiteSortByEnum {
	values := make([]InstallationSiteSortByEnum, 0)
	for _, v := range mappingInstallationSiteSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetInstallationSiteSortByEnumStringValues Enumerates the set of values in String for InstallationSiteSortByEnum
func GetInstallationSiteSortByEnumStringValues() []string {
	return []string{
		"managedInstanceId",
		"jreDistribution",
		"jreVendor",
		"jreVersion",
		"path",
		"approximateApplicationCount",
		"osName",
		"securityStatus",
	}
}

// GetMappingInstallationSiteSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstallationSiteSortByEnum(val string) (InstallationSiteSortByEnum, bool) {
	enum, ok := mappingInstallationSiteSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
