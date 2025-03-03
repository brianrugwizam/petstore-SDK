# DefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**DefaultResponseMeta**](DefaultResponseMeta.md) |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** | Errors specifies a list of errors that occurred, can be filled using error handlers. | [optional] 

## Methods

### NewDefaultResponse

`func NewDefaultResponse() *DefaultResponse`

NewDefaultResponse instantiates a new DefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultResponseWithDefaults

`func NewDefaultResponseWithDefaults() *DefaultResponse`

NewDefaultResponseWithDefaults instantiates a new DefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *DefaultResponse) GetMeta() DefaultResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DefaultResponse) GetMetaOk() (*DefaultResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DefaultResponse) SetMeta(v DefaultResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DefaultResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *DefaultResponse) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DefaultResponse) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DefaultResponse) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DefaultResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


