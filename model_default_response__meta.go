/*
PetStore API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testsdkgo

import (
	"encoding/json"
)

// checks if the DefaultResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultResponseMeta{}

// DefaultResponseMeta MetaData contains metadata of the response, such as record count, pagination and other additional information.
type DefaultResponseMeta struct {
	TraceId *string `json:"traceId,omitempty"`
}

// NewDefaultResponseMeta instantiates a new DefaultResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultResponseMeta() *DefaultResponseMeta {
	this := DefaultResponseMeta{}
	return &this
}

// NewDefaultResponseMetaWithDefaults instantiates a new DefaultResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultResponseMetaWithDefaults() *DefaultResponseMeta {
	this := DefaultResponseMeta{}
	return &this
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *DefaultResponseMeta) GetTraceId() string {
	if o == nil || IsNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultResponseMeta) GetTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TraceId) {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *DefaultResponseMeta) HasTraceId() bool {
	if o != nil && !IsNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *DefaultResponseMeta) SetTraceId(v string) {
	o.TraceId = &v
}

func (o DefaultResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TraceId) {
		toSerialize["traceId"] = o.TraceId
	}
	return toSerialize, nil
}

type NullableDefaultResponseMeta struct {
	value *DefaultResponseMeta
	isSet bool
}

func (v NullableDefaultResponseMeta) Get() *DefaultResponseMeta {
	return v.value
}

func (v *NullableDefaultResponseMeta) Set(val *DefaultResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultResponseMeta(val *DefaultResponseMeta) *NullableDefaultResponseMeta {
	return &NullableDefaultResponseMeta{value: val, isSet: true}
}

func (v NullableDefaultResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


