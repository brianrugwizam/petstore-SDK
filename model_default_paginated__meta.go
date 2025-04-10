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

// checks if the DefaultPaginatedMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultPaginatedMeta{}

// DefaultPaginatedMeta MetaData contains metadata of the response, such as record count, pagination and other additional information.
type DefaultPaginatedMeta struct {
	TraceId *string `json:"traceId,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Total *int64 `json:"total,omitempty"`
}

// NewDefaultPaginatedMeta instantiates a new DefaultPaginatedMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultPaginatedMeta() *DefaultPaginatedMeta {
	this := DefaultPaginatedMeta{}
	return &this
}

// NewDefaultPaginatedMetaWithDefaults instantiates a new DefaultPaginatedMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultPaginatedMetaWithDefaults() *DefaultPaginatedMeta {
	this := DefaultPaginatedMeta{}
	return &this
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *DefaultPaginatedMeta) GetTraceId() string {
	if o == nil || IsNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultPaginatedMeta) GetTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TraceId) {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *DefaultPaginatedMeta) HasTraceId() bool {
	if o != nil && !IsNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *DefaultPaginatedMeta) SetTraceId(v string) {
	o.TraceId = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *DefaultPaginatedMeta) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultPaginatedMeta) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *DefaultPaginatedMeta) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *DefaultPaginatedMeta) SetLimit(v int32) {
	o.Limit = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DefaultPaginatedMeta) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultPaginatedMeta) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DefaultPaginatedMeta) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *DefaultPaginatedMeta) SetTotal(v int64) {
	o.Total = &v
}

func (o DefaultPaginatedMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultPaginatedMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TraceId) {
		toSerialize["traceId"] = o.TraceId
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableDefaultPaginatedMeta struct {
	value *DefaultPaginatedMeta
	isSet bool
}

func (v NullableDefaultPaginatedMeta) Get() *DefaultPaginatedMeta {
	return v.value
}

func (v *NullableDefaultPaginatedMeta) Set(val *DefaultPaginatedMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultPaginatedMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultPaginatedMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultPaginatedMeta(val *DefaultPaginatedMeta) *NullableDefaultPaginatedMeta {
	return &NullableDefaultPaginatedMeta{value: val, isSet: true}
}

func (v NullableDefaultPaginatedMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultPaginatedMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


