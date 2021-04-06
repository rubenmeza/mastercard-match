# \AddTerminatedMerchantApi

All URIs are relative to *https://api.mastercard.com/fraud/merchant/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMerchantPost**](AddTerminatedMerchantApi.md#AddMerchantPost) | **Post** /add-merchant | ##### Adds a new terminated merchant to MATCH system.



## AddMerchantPost

> AddTerminatedMerchantResponseSchema AddMerchantPost(ctx).Format(format).AddTerminatedMerchantRequestSchema(addTerminatedMerchantRequestSchema).Execute()

##### Adds a new terminated merchant to MATCH system.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    format := "JSON" // string | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON.
    addTerminatedMerchantRequestSchema := *openapiclient.NewAddTerminatedMerchantRequestSchema() // AddTerminatedMerchantRequestSchema | Body of the Add Terminated Merchant Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddTerminatedMerchantApi.AddMerchantPost(context.Background()).Format(format).AddTerminatedMerchantRequestSchema(addTerminatedMerchantRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddTerminatedMerchantApi.AddMerchantPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMerchantPost`: AddTerminatedMerchantResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `AddTerminatedMerchantApi.AddMerchantPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMerchantPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON. | 
 **addTerminatedMerchantRequestSchema** | [**AddTerminatedMerchantRequestSchema**](AddTerminatedMerchantRequestSchema.md) | Body of the Add Terminated Merchant Request | 

### Return type

[**AddTerminatedMerchantResponseSchema**](AddTerminatedMerchantResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

