# \RetroactiveInquiryRequestApi

All URIs are relative to *https://api.mastercard.com/fraud/merchant/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetroRetroListPost**](RetroactiveInquiryRequestApi.md#RetroRetroListPost) | **Post** /retro/retro-list | ##### The retroactive inquiry helps acquirer to retrieve list of termination inquiry matches made within 360 days of inquiry initiation.



## RetroRetroListPost

> RetroResponseSchema RetroRetroListPost(ctx).Format(format).RetroRequestSchema(retroRequestSchema).Execute()

##### The retroactive inquiry helps acquirer to retrieve list of termination inquiry matches made within 360 days of inquiry initiation.



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
    retroRequestSchema := *openapiclient.NewRetroRequestSchema() // RetroRequestSchema | The retro request object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RetroactiveInquiryRequestApi.RetroRetroListPost(context.Background()).Format(format).RetroRequestSchema(retroRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetroactiveInquiryRequestApi.RetroRetroListPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetroRetroListPost`: RetroResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `RetroactiveInquiryRequestApi.RetroRetroListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetroRetroListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON. | 
 **retroRequestSchema** | [**RetroRequestSchema**](RetroRequestSchema.md) | The retro request object | 

### Return type

[**RetroResponseSchema**](RetroResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

