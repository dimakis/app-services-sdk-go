/*
 * Red Hat Openshift SmartEvents Fleet Manager
 *
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * API version: 0.0.1
 * Contact: openbridge-dev@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsclient

import (
	"encoding/json"
)

// CloudRegionResponse struct for CloudRegionResponse
type CloudRegionResponse struct {
	Kind *string `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCloudRegionResponse instantiates a new CloudRegionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRegionResponse() *CloudRegionResponse {
	this := CloudRegionResponse{}
	return &this
}

// NewCloudRegionResponseWithDefaults instantiates a new CloudRegionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRegionResponseWithDefaults() *CloudRegionResponse {
	this := CloudRegionResponse{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CloudRegionResponse) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionResponse) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CloudRegionResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CloudRegionResponse) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudRegionResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudRegionResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudRegionResponse) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CloudRegionResponse) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CloudRegionResponse) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CloudRegionResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudRegionResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudRegionResponse) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudRegionResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CloudRegionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableCloudRegionResponse struct {
	value *CloudRegionResponse
	isSet bool
}

func (v NullableCloudRegionResponse) Get() *CloudRegionResponse {
	return v.value
}

func (v *NullableCloudRegionResponse) Set(val *CloudRegionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRegionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRegionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRegionResponse(val *CloudRegionResponse) *NullableCloudRegionResponse {
	return &NullableCloudRegionResponse{value: val, isSet: true}
}

func (v NullableCloudRegionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRegionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


