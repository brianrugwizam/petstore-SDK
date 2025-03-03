# CreatePet201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**DefaultResponseMeta**](DefaultResponseMeta.md) |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** | Errors specifies a list of errors that occurred, can be filled using error handlers. | [optional] 
**Data** | Pointer to [**Pet**](Pet.md) |  | [optional] 

## Methods

### NewCreatePet201Response

`func NewCreatePet201Response() *CreatePet201Response`

NewCreatePet201Response instantiates a new CreatePet201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePet201ResponseWithDefaults

`func NewCreatePet201ResponseWithDefaults() *CreatePet201Response`

NewCreatePet201ResponseWithDefaults instantiates a new CreatePet201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *CreatePet201Response) GetMeta() DefaultResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreatePet201Response) GetMetaOk() (*DefaultResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreatePet201Response) SetMeta(v DefaultResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreatePet201Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *CreatePet201Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreatePet201Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreatePet201Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreatePet201Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *CreatePet201Response) GetData() Pet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreatePet201Response) GetDataOk() (*Pet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreatePet201Response) SetData(v Pet)`

SetData sets Data field to given value.

### HasData

`func (o *CreatePet201Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


