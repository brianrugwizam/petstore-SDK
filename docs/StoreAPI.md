# \StoreAPI

All URIs are relative to *http://127.0.0.1:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrderById**](StoreAPI.md#GetOrderById) | **Get** /store/order/{orderId} | Find purchase order by ID
[**PlaceOrder**](StoreAPI.md#PlaceOrder) | **Post** /store/order | Place an order for a pet



## GetOrderById

> DefaultResponse GetOrderById(ctx, orderId).DefaultResponse(defaultResponse).Execute()

Find purchase order by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brianrugwizam/petstore2"
)

func main() {
	orderId := int64(789) // int64 | ID of order that needs to be fetched
	defaultResponse := map[string][]openapiclient.DefaultResponse{ ... } // DefaultResponse |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.GetOrderById(context.Background(), orderId).DefaultResponse(defaultResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetOrderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderById`: DefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.GetOrderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int64** | ID of order that needs to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultResponse** | [**DefaultResponse**](DefaultResponse.md) |  | 

### Return type

[**DefaultResponse**](DefaultResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceOrder

> PlaceOrder(ctx).DefaultResponse(defaultResponse).Execute()

Place an order for a pet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brianrugwizam/petstore2"
)

func main() {
	defaultResponse := map[string][]openapiclient.DefaultResponse{ ... } // DefaultResponse |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PlaceOrder(context.Background()).DefaultResponse(defaultResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PlaceOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaceOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **defaultResponse** | [**DefaultResponse**](DefaultResponse.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

