/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorNamespaceStatus struct for ConnectorNamespaceStatus
type ConnectorNamespaceStatus struct {
	State ConnectorNamespaceState `json:"state"`
	Version *string `json:"version,omitempty"`
	ConnectorsDeployed int32 `json:"connectors_deployed"`
	Error *string `json:"error,omitempty"`
}

// NewConnectorNamespaceStatus instantiates a new ConnectorNamespaceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorNamespaceStatus(state ConnectorNamespaceState, connectorsDeployed int32) *ConnectorNamespaceStatus {
	this := ConnectorNamespaceStatus{}
	this.State = state
	this.ConnectorsDeployed = connectorsDeployed
	return &this
}

// NewConnectorNamespaceStatusWithDefaults instantiates a new ConnectorNamespaceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorNamespaceStatusWithDefaults() *ConnectorNamespaceStatus {
	this := ConnectorNamespaceStatus{}
	return &this
}

// GetState returns the State field value
func (o *ConnectorNamespaceStatus) GetState() ConnectorNamespaceState {
	if o == nil {
		var ret ConnectorNamespaceState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceStatus) GetStateOk() (*ConnectorNamespaceState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ConnectorNamespaceStatus) SetState(v ConnectorNamespaceState) {
	o.State = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ConnectorNamespaceStatus) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceStatus) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ConnectorNamespaceStatus) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ConnectorNamespaceStatus) SetVersion(v string) {
	o.Version = &v
}

// GetConnectorsDeployed returns the ConnectorsDeployed field value
func (o *ConnectorNamespaceStatus) GetConnectorsDeployed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConnectorsDeployed
}

// GetConnectorsDeployedOk returns a tuple with the ConnectorsDeployed field value
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceStatus) GetConnectorsDeployedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectorsDeployed, true
}

// SetConnectorsDeployed sets field value
func (o *ConnectorNamespaceStatus) SetConnectorsDeployed(v int32) {
	o.ConnectorsDeployed = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ConnectorNamespaceStatus) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceStatus) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ConnectorNamespaceStatus) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ConnectorNamespaceStatus) SetError(v string) {
	o.Error = &v
}

func (o ConnectorNamespaceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["connectors_deployed"] = o.ConnectorsDeployed
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorNamespaceStatus struct {
	value *ConnectorNamespaceStatus
	isSet bool
}

func (v NullableConnectorNamespaceStatus) Get() *ConnectorNamespaceStatus {
	return v.value
}

func (v *NullableConnectorNamespaceStatus) Set(val *ConnectorNamespaceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorNamespaceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorNamespaceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorNamespaceStatus(val *ConnectorNamespaceStatus) *NullableConnectorNamespaceStatus {
	return &NullableConnectorNamespaceStatus{value: val, isSet: true}
}

func (v NullableConnectorNamespaceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorNamespaceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


