# \RetroactiveInquiryDetailsRequestApi

All URIs are relative to *https://api.mastercard.com/fraud/merchant/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetroRetroInquiryDetailsPost**](RetroactiveInquiryDetailsRequestApi.md#RetroRetroInquiryDetailsPost) | **Post** /retro/retro-inquiry-details | ##### Retrieve retro inquiry search results



## RetroRetroInquiryDetailsPost

> RetroInquiryResponseSchema RetroRetroInquiryDetailsPost(ctx).Format(format).AcquirerId(acquirerId).RetroInquiryRequestSchema(retroInquiryRequestSchema).Execute()

##### Retrieve retro inquiry search results



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
    acquirerId := "1996" // string | The Member ICA number which is used to validate that the application has permission to hit the MATCH database. This number must be obtained from a participating MATCH acquiring bank or entity before a developer can access the MATCH resource.
    retroInquiryRequestSchema := *openapiclient.NewRetroInquiryRequestSchema() // RetroInquiryRequestSchema | The retro inquiry request object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RetroactiveInquiryDetailsRequestApi.RetroRetroInquiryDetailsPost(context.Background()).Format(format).AcquirerId(acquirerId).RetroInquiryRequestSchema(retroInquiryRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetroactiveInquiryDetailsRequestApi.RetroRetroInquiryDetailsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetroRetroInquiryDetailsPost`: RetroInquiryResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `RetroactiveInquiryDetailsRequestApi.RetroRetroInquiryDetailsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetroRetroInquiryDetailsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON. | 
 **acquirerId** | **string** | The Member ICA number which is used to validate that the application has permission to hit the MATCH database. This number must be obtained from a participating MATCH acquiring bank or entity before a developer can access the MATCH resource. | 
 **retroInquiryRequestSchema** | [**RetroInquiryRequestSchema**](RetroInquiryRequestSchema.md) | The retro inquiry request object | 

### Return type

[**RetroInquiryResponseSchema**](RetroInquiryResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

