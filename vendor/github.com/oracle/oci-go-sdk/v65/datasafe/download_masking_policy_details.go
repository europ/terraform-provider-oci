// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DownloadMaskingPolicyDetails Details to download a masking policy.
type DownloadMaskingPolicyDetails struct {

	// The format of the masking policy file.
	PolicyFormat PolicyFormatEnum `mandatory:"false" json:"policyFormat,omitempty"`
}

func (m DownloadMaskingPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DownloadMaskingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPolicyFormatEnum(string(m.PolicyFormat)); !ok && m.PolicyFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PolicyFormat: %s. Supported values are: %s.", m.PolicyFormat, strings.Join(GetPolicyFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
