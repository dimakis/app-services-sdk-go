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

// BridgeListResponse struct for BridgeListResponse
type BridgeListResponse struct {
	Kind *string `json:"kind,omitempty"`
	Items *[]BridgeResponse `json:"items,omitempty"`
	Page *int64 `json:"page,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Total *int64 `json:"total,omitempty"`
}

// NewBridgeListResponse instantiates a new BridgeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBridgeListResponse() *BridgeListResponse {
	this := BridgeListResponse{}
	return &this
}

// NewBridgeListResponseWithDefaults instantiates a new BridgeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBridgeListResponseWithDefaults() *BridgeListResponse {
	this := BridgeListResponse{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *BridgeListResponse) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BridgeListResponse) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *BridgeListResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *BridgeListResponse) SetKind(v string) {
	o.Kind = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BridgeListResponse) GetItems() []BridgeResponse {
	if o == nil || o.Items == nil {
		var ret []BridgeResponse
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BridgeListResponse) GetItemsOk() (*[]BridgeResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BridgeListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BridgeResponse and assigns it to the Items field.
func (o *BridgeListResponse) SetItems(v []BridgeResponse) {
	o.Items = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *BridgeListResponse) GetPage() int64 {
	if o == nil || o.Page == nil {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BridgeListResponse) GetPageOk() (*int64, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *BridgeListResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *BridgeListResponse) SetPage(v int64) {
	o.Page = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BridgeListResponse) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BridgeListResponse) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BridgeListResponse) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *BridgeListResponse) SetSize(v int64) {
	o.Size = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BridgeListResponse) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BridgeListResponse) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BridgeListResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *BridgeListResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o BridgeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableBridgeListResponse struct {
	value *BridgeListResponse
	isSet bool
}

func (v NullableBridgeListResponse) Get() *BridgeListResponse {
	return v.value
}

func (v *NullableBridgeListResponse) Set(val *BridgeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBridgeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBridgeListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBridgeListResponse(val *BridgeListResponse) *NullableBridgeListResponse {
	return &NullableBridgeListResponse{value: val, isSet: true}
}

func (v NullableBridgeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBridgeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


