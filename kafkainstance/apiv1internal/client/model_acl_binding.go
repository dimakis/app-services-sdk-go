/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.11.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// AclBinding struct for AclBinding
type AclBinding struct {
	// Unique identifier for the object. Not supported for all object kinds.
	Id *string `json:"id,omitempty"`
	Kind string `json:"kind"`
	// Link path to request the object. Not supported for all object kinds.
	Href *string `json:"href,omitempty"`
	ResourceType AclResourceType `json:"resourceType"`
	ResourceName string `json:"resourceName"`
	PatternType AclPatternType `json:"patternType"`
	// Identifies the user or service account to which an ACL entry is bound. The literal prefix value of `User:` is required. May be used to specify all users with value `User:*`.
	Principal string `json:"principal"`
	Operation AclOperation `json:"operation"`
	Permission AclPermissionType `json:"permission"`
}

// NewAclBinding instantiates a new AclBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAclBinding(kind string, resourceType AclResourceType, resourceName string, patternType AclPatternType, principal string, operation AclOperation, permission AclPermissionType) *AclBinding {
	this := AclBinding{}
	this.Kind = kind
	this.ResourceType = resourceType
	this.ResourceName = resourceName
	this.PatternType = patternType
	this.Principal = principal
	this.Operation = operation
	this.Permission = permission
	return &this
}

// NewAclBindingWithDefaults instantiates a new AclBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAclBindingWithDefaults() *AclBinding {
	this := AclBinding{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AclBinding) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclBinding) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AclBinding) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AclBinding) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value
func (o *AclBinding) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AclBinding) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AclBinding) SetKind(v string) {
	o.Kind = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *AclBinding) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclBinding) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *AclBinding) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *AclBinding) SetHref(v string) {
	o.Href = &v
}

// GetResourceType returns the ResourceType field value
func (o *AclBinding) GetResourceType() AclResourceType {
	if o == nil {
		var ret AclResourceType
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *AclBinding) GetResourceTypeOk() (*AclResourceType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *AclBinding) SetResourceType(v AclResourceType) {
	o.ResourceType = v
}

// GetResourceName returns the ResourceName field value
func (o *AclBinding) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *AclBinding) GetResourceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *AclBinding) SetResourceName(v string) {
	o.ResourceName = v
}

// GetPatternType returns the PatternType field value
func (o *AclBinding) GetPatternType() AclPatternType {
	if o == nil {
		var ret AclPatternType
		return ret
	}

	return o.PatternType
}

// GetPatternTypeOk returns a tuple with the PatternType field value
// and a boolean to check if the value has been set.
func (o *AclBinding) GetPatternTypeOk() (*AclPatternType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PatternType, true
}

// SetPatternType sets field value
func (o *AclBinding) SetPatternType(v AclPatternType) {
	o.PatternType = v
}

// GetPrincipal returns the Principal field value
func (o *AclBinding) GetPrincipal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *AclBinding) GetPrincipalOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *AclBinding) SetPrincipal(v string) {
	o.Principal = v
}

// GetOperation returns the Operation field value
func (o *AclBinding) GetOperation() AclOperation {
	if o == nil {
		var ret AclOperation
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *AclBinding) GetOperationOk() (*AclOperation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *AclBinding) SetOperation(v AclOperation) {
	o.Operation = v
}

// GetPermission returns the Permission field value
func (o *AclBinding) GetPermission() AclPermissionType {
	if o == nil {
		var ret AclPermissionType
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *AclBinding) GetPermissionOk() (*AclPermissionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *AclBinding) SetPermission(v AclPermissionType) {
	o.Permission = v
}

func (o AclBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["resourceType"] = o.ResourceType
	}
	if true {
		toSerialize["resourceName"] = o.ResourceName
	}
	if true {
		toSerialize["patternType"] = o.PatternType
	}
	if true {
		toSerialize["principal"] = o.Principal
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	if true {
		toSerialize["permission"] = o.Permission
	}
	return json.Marshal(toSerialize)
}

type NullableAclBinding struct {
	value *AclBinding
	isSet bool
}

func (v NullableAclBinding) Get() *AclBinding {
	return v.value
}

func (v *NullableAclBinding) Set(val *AclBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableAclBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableAclBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclBinding(val *AclBinding) *NullableAclBinding {
	return &NullableAclBinding{value: val, isSet: true}
}

func (v NullableAclBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


