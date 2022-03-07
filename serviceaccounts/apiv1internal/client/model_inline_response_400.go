/*
 * Service Accounts API Documentation
 *
 * This is the API documentation for Service Accounts
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccountsclient

import (
	"encoding/json"
)

// InlineResponse400 struct for InlineResponse400
type InlineResponse400 struct {

	Cause *InlineResponse400Cause `json:"cause,omitempty"`

	StackTrace *[]InlineResponse400CauseStackTrace `json:"stackTrace,omitempty"`

	Response *InlineResponse400Response `json:"response,omitempty"`

	Message *string `json:"message,omitempty"`

	LocalizedMessage *string `json:"localizedMessage,omitempty"`

	Suppressed *[]InlineResponse400CauseSuppressed `json:"suppressed,omitempty"`

}

// NewInlineResponse400 instantiates a new InlineResponse400 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse400() *InlineResponse400 {
	this := InlineResponse400{}
	return &this
}

// NewInlineResponse400WithDefaults instantiates a new InlineResponse400 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse400WithDefaults() *InlineResponse400 {
	this := InlineResponse400{}







	return &this
}


// GetCause returns the Cause field value if set, zero value otherwise.
func (o *InlineResponse400) GetCause() InlineResponse400Cause {
	if o == nil || o.Cause == nil {
		var ret InlineResponse400Cause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetCauseOk() (*InlineResponse400Cause, bool) {
	if o == nil || o.Cause == nil {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *InlineResponse400) HasCause() bool {
	if o != nil && o.Cause != nil {
		return true
	}

	return false
}

// SetCause gets a reference to the given InlineResponse400Cause and assigns it to the Cause field.
func (o *InlineResponse400) SetCause(v InlineResponse400Cause) {
	o.Cause = &v
}


// GetStackTrace returns the StackTrace field value if set, zero value otherwise.
func (o *InlineResponse400) GetStackTrace() []InlineResponse400CauseStackTrace {
	if o == nil || o.StackTrace == nil {
		var ret []InlineResponse400CauseStackTrace
		return ret
	}
	return *o.StackTrace
}

// GetStackTraceOk returns a tuple with the StackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetStackTraceOk() (*[]InlineResponse400CauseStackTrace, bool) {
	if o == nil || o.StackTrace == nil {
		return nil, false
	}
	return o.StackTrace, true
}

// HasStackTrace returns a boolean if a field has been set.
func (o *InlineResponse400) HasStackTrace() bool {
	if o != nil && o.StackTrace != nil {
		return true
	}

	return false
}

// SetStackTrace gets a reference to the given []InlineResponse400CauseStackTrace and assigns it to the StackTrace field.
func (o *InlineResponse400) SetStackTrace(v []InlineResponse400CauseStackTrace) {
	o.StackTrace = &v
}


// GetResponse returns the Response field value if set, zero value otherwise.
func (o *InlineResponse400) GetResponse() InlineResponse400Response {
	if o == nil || o.Response == nil {
		var ret InlineResponse400Response
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetResponseOk() (*InlineResponse400Response, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *InlineResponse400) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given InlineResponse400Response and assigns it to the Response field.
func (o *InlineResponse400) SetResponse(v InlineResponse400Response) {
	o.Response = &v
}


// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse400) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse400) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse400) SetMessage(v string) {
	o.Message = &v
}


// GetLocalizedMessage returns the LocalizedMessage field value if set, zero value otherwise.
func (o *InlineResponse400) GetLocalizedMessage() string {
	if o == nil || o.LocalizedMessage == nil {
		var ret string
		return ret
	}
	return *o.LocalizedMessage
}

// GetLocalizedMessageOk returns a tuple with the LocalizedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetLocalizedMessageOk() (*string, bool) {
	if o == nil || o.LocalizedMessage == nil {
		return nil, false
	}
	return o.LocalizedMessage, true
}

// HasLocalizedMessage returns a boolean if a field has been set.
func (o *InlineResponse400) HasLocalizedMessage() bool {
	if o != nil && o.LocalizedMessage != nil {
		return true
	}

	return false
}

// SetLocalizedMessage gets a reference to the given string and assigns it to the LocalizedMessage field.
func (o *InlineResponse400) SetLocalizedMessage(v string) {
	o.LocalizedMessage = &v
}


// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *InlineResponse400) GetSuppressed() []InlineResponse400CauseSuppressed {
	if o == nil || o.Suppressed == nil {
		var ret []InlineResponse400CauseSuppressed
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetSuppressedOk() (*[]InlineResponse400CauseSuppressed, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *InlineResponse400) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given []InlineResponse400CauseSuppressed and assigns it to the Suppressed field.
func (o *InlineResponse400) SetSuppressed(v []InlineResponse400CauseSuppressed) {
	o.Suppressed = &v
}


func (o InlineResponse400) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Cause != nil {
		toSerialize["cause"] = o.Cause
	}
    
	if o.StackTrace != nil {
		toSerialize["stackTrace"] = o.StackTrace
	}
    
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
    
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
    
	if o.LocalizedMessage != nil {
		toSerialize["localizedMessage"] = o.LocalizedMessage
	}
    
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
    
	return json.Marshal(toSerialize)
}

type NullableInlineResponse400 struct {
	value *InlineResponse400
	isSet bool
}

func (v NullableInlineResponse400) Get() *InlineResponse400 {
	return v.value
}

func (v *NullableInlineResponse400) Set(val *InlineResponse400) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse400) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse400(val *InlineResponse400) *NullableInlineResponse400 {
	return &NullableInlineResponse400{value: val, isSet: true}
}

func (v NullableInlineResponse400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

