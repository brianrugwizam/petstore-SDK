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

// checks if the CreatePet201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePet201Response{}

// CreatePet201Response struct for CreatePet201Response
type CreatePet201Response struct {
	Meta *DefaultResponseMeta `json:"_meta,omitempty"`
	// Errors specifies a list of errors that occurred, can be filled using error handlers.
	Errors map[string]interface{} `json:"errors,omitempty"`
	Data *Pet `json:"data,omitempty"`
}

// NewCreatePet201Response instantiates a new CreatePet201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePet201Response() *CreatePet201Response {
	this := CreatePet201Response{}
	return &this
}

// NewCreatePet201ResponseWithDefaults instantiates a new CreatePet201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePet201ResponseWithDefaults() *CreatePet201Response {
	this := CreatePet201Response{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CreatePet201Response) GetMeta() DefaultResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret DefaultResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePet201Response) GetMetaOk() (*DefaultResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CreatePet201Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given DefaultResponseMeta and assigns it to the Meta field.
func (o *CreatePet201Response) SetMeta(v DefaultResponseMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreatePet201Response) GetErrors() map[string]interface{} {
	if o == nil || IsNil(o.Errors) {
		var ret map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePet201Response) GetErrorsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Errors) {
		return map[string]interface{}{}, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreatePet201Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]interface{} and assigns it to the Errors field.
func (o *CreatePet201Response) SetErrors(v map[string]interface{}) {
	o.Errors = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreatePet201Response) GetData() Pet {
	if o == nil || IsNil(o.Data) {
		var ret Pet
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePet201Response) GetDataOk() (*Pet, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreatePet201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Pet and assigns it to the Data field.
func (o *CreatePet201Response) SetData(v Pet) {
	o.Data = &v
}

func (o CreatePet201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePet201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["_meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreatePet201Response struct {
	value *CreatePet201Response
	isSet bool
}

func (v NullableCreatePet201Response) Get() *CreatePet201Response {
	return v.value
}

func (v *NullableCreatePet201Response) Set(val *CreatePet201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePet201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePet201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePet201Response(val *CreatePet201Response) *NullableCreatePet201Response {
	return &NullableCreatePet201Response{value: val, isSet: true}
}

func (v NullableCreatePet201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePet201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


