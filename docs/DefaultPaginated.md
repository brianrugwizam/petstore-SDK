# DefaultPaginated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**DefaultPaginatedMeta**](DefaultPaginatedMeta.md) |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** | Errors specifies a list of errors that occurred, can be filled using error handlers. | [optional] 

## Methods

### NewDefaultPaginated

`func NewDefaultPaginated() *DefaultPaginated`

NewDefaultPaginated instantiates a new DefaultPaginated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultPaginatedWithDefaults

`func NewDefaultPaginatedWithDefaults() *DefaultPaginated`

NewDefaultPaginatedWithDefaults instantiates a new DefaultPaginated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *DefaultPaginated) GetMeta() DefaultPaginatedMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DefaultPaginated) GetMetaOk() (*DefaultPaginatedMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DefaultPaginated) SetMeta(v DefaultPaginatedMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DefaultPaginated) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *DefaultPaginated) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DefaultPaginated) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DefaultPaginated) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DefaultPaginated) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


