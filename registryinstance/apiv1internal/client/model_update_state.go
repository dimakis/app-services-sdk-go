/*
 * Apicurio Registry API [v2]
 *
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * API version: 2.1.0-SNAPSHOT
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// UpdateState struct for UpdateState
type UpdateState struct {
	State ArtifactState `json:"state"`
}

// NewUpdateState instantiates a new UpdateState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateState(state ArtifactState) *UpdateState {
	this := UpdateState{}
	this.State = state
	return &this
}

// NewUpdateStateWithDefaults instantiates a new UpdateState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStateWithDefaults() *UpdateState {
	this := UpdateState{}
	return &this
}

// GetState returns the State field value
func (o *UpdateState) GetState() ArtifactState {
	if o == nil {
		var ret ArtifactState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *UpdateState) GetStateOk() (*ArtifactState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *UpdateState) SetState(v ArtifactState) {
	o.State = v
}

func (o UpdateState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateState struct {
	value *UpdateState
	isSet bool
}

func (v NullableUpdateState) Get() *UpdateState {
	return v.value
}

func (v *NullableUpdateState) Set(val *UpdateState) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateState) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateState(val *UpdateState) *NullableUpdateState {
	return &NullableUpdateState{value: val, isSet: true}
}

func (v NullableUpdateState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


