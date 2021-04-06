# \TerminationInquiryRequestApi

All URIs are relative to *https://api.mastercard.com/fraud/merchant/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TerminationInquiryPost**](TerminationInquiryRequestApi.md#TerminationInquiryPost) | **Post** /termination-inquiry | ##### Retrieves terminated merchant information based on the criteria defined in the API request.



## TerminationInquiryPost

> TerminationInquirySchema TerminationInquiryPost(ctx).PageOffset(pageOffset).PageLength(pageLength).Format(format).TerminationInquiryRequestSchema(terminationInquiryRequestSchema).Execute()

##### Retrieves terminated merchant information based on the criteria defined in the API request.



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
    pageOffset := int32(0) // int32 | The zero-based offset to start at. The actual start position is this value +1. An offset of 10 starts at item 11. Combined with the PageLength option this allows pagination to be supported through the service requests.
    pageLength := int32(10) // int32 | The maximum number of items to retrieve within the current \"page\" of results.
    format := "JSON" // string | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON.
    terminationInquiryRequestSchema := *openapiclient.NewTerminationInquiryRequestSchema() // TerminationInquiryRequestSchema | Body of the Termination Inquiry Request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TerminationInquiryRequestApi.TerminationInquiryPost(context.Background()).PageOffset(pageOffset).PageLength(pageLength).Format(format).TerminationInquiryRequestSchema(terminationInquiryRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminationInquiryRequestApi.TerminationInquiryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminationInquiryPost`: TerminationInquirySchema
    fmt.Fprintf(os.Stdout, "Response from `TerminationInquiryRequestApi.TerminationInquiryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminationInquiryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageOffset** | **int32** | The zero-based offset to start at. The actual start position is this value +1. An offset of 10 starts at item 11. Combined with the PageLength option this allows pagination to be supported through the service requests. | 
 **pageLength** | **int32** | The maximum number of items to retrieve within the current \&quot;page\&quot; of results. | 
 **format** | **string** | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON. | 
 **terminationInquiryRequestSchema** | [**TerminationInquiryRequestSchema**](TerminationInquiryRequestSchema.md) | Body of the Termination Inquiry Request | 

### Return type

[**TerminationInquirySchema**](TerminationInquirySchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

