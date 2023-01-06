// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// CreateEntityShapeRequest wrapper for the CreateEntityShape operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataconnectivity/CreateEntityShape.go.html to see an example of how to use CreateEntityShapeRequest.
type CreateEntityShapeRequest struct {

	// The registry OCID.
	RegistryId *string `mandatory:"true" contributesTo:"path" name:"registryId"`

	// The connection key.
	ConnectionKey *string `mandatory:"true" contributesTo:"path" name:"connectionKey"`

	// The schema resource name used for retrieving schemas.
	SchemaResourceName *string `mandatory:"true" contributesTo:"path" name:"schemaResourceName"`

	// The details required to create the data entity shape.
	CreateEntityShapeDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or server error without the risk of executing that same action again.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match` parameter to the value of the `etag` from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the `etag` you provide matches the resource's current `etag` value.
	// When 'if-match' is provided and its value does not exactly match the 'etag' of the resource on the server, the request fails with the 412 response code.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Authorization mode for communicating with another OCI service relevant for the API.
	AuthorizationMode CreateEntityShapeAuthorizationModeEnum `mandatory:"false" contributesTo:"query" name:"authorizationMode" omitEmpty:"true"`

	// Endpoint ID used for getDataAssetFullDetails.
	EndpointId *string `mandatory:"false" contributesTo:"query" name:"endpointId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateEntityShapeRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateEntityShapeRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateEntityShapeRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateEntityShapeRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateEntityShapeRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateEntityShapeAuthorizationModeEnum(string(request.AuthorizationMode)); !ok && request.AuthorizationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthorizationMode: %s. Supported values are: %s.", request.AuthorizationMode, strings.Join(GetCreateEntityShapeAuthorizationModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateEntityShapeResponse wrapper for the CreateEntityShape operation
type CreateEntityShapeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The EntityShape instance
	EntityShape `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateEntityShapeResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateEntityShapeResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// CreateEntityShapeAuthorizationModeEnum Enum with underlying type: string
type CreateEntityShapeAuthorizationModeEnum string

// Set of constants representing the allowable values for CreateEntityShapeAuthorizationModeEnum
const (
	CreateEntityShapeAuthorizationModeObo               CreateEntityShapeAuthorizationModeEnum = "OBO"
	CreateEntityShapeAuthorizationModeUserPrincipal     CreateEntityShapeAuthorizationModeEnum = "USER_PRINCIPAL"
	CreateEntityShapeAuthorizationModeResourcePrincipal CreateEntityShapeAuthorizationModeEnum = "RESOURCE_PRINCIPAL"
	CreateEntityShapeAuthorizationModeInstancePrincipal CreateEntityShapeAuthorizationModeEnum = "INSTANCE_PRINCIPAL"
	CreateEntityShapeAuthorizationModeUndefined         CreateEntityShapeAuthorizationModeEnum = "UNDEFINED"
)

var mappingCreateEntityShapeAuthorizationModeEnum = map[string]CreateEntityShapeAuthorizationModeEnum{
	"OBO":                CreateEntityShapeAuthorizationModeObo,
	"USER_PRINCIPAL":     CreateEntityShapeAuthorizationModeUserPrincipal,
	"RESOURCE_PRINCIPAL": CreateEntityShapeAuthorizationModeResourcePrincipal,
	"INSTANCE_PRINCIPAL": CreateEntityShapeAuthorizationModeInstancePrincipal,
	"UNDEFINED":          CreateEntityShapeAuthorizationModeUndefined,
}

var mappingCreateEntityShapeAuthorizationModeEnumLowerCase = map[string]CreateEntityShapeAuthorizationModeEnum{
	"obo":                CreateEntityShapeAuthorizationModeObo,
	"user_principal":     CreateEntityShapeAuthorizationModeUserPrincipal,
	"resource_principal": CreateEntityShapeAuthorizationModeResourcePrincipal,
	"instance_principal": CreateEntityShapeAuthorizationModeInstancePrincipal,
	"undefined":          CreateEntityShapeAuthorizationModeUndefined,
}

// GetCreateEntityShapeAuthorizationModeEnumValues Enumerates the set of values for CreateEntityShapeAuthorizationModeEnum
func GetCreateEntityShapeAuthorizationModeEnumValues() []CreateEntityShapeAuthorizationModeEnum {
	values := make([]CreateEntityShapeAuthorizationModeEnum, 0)
	for _, v := range mappingCreateEntityShapeAuthorizationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateEntityShapeAuthorizationModeEnumStringValues Enumerates the set of values in String for CreateEntityShapeAuthorizationModeEnum
func GetCreateEntityShapeAuthorizationModeEnumStringValues() []string {
	return []string{
		"OBO",
		"USER_PRINCIPAL",
		"RESOURCE_PRINCIPAL",
		"INSTANCE_PRINCIPAL",
		"UNDEFINED",
	}
}

// GetMappingCreateEntityShapeAuthorizationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateEntityShapeAuthorizationModeEnum(val string) (CreateEntityShapeAuthorizationModeEnum, bool) {
	enum, ok := mappingCreateEntityShapeAuthorizationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
