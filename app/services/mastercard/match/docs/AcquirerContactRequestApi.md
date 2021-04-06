# \AcquirerContactRequestApi

All URIs are relative to *https://api.mastercard.com/fraud/merchant/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommonContactDetailsPost**](AcquirerContactRequestApi.md#CommonContactDetailsPost) | **Post** /common/contact-details | ##### Retrieve contact information of another acquirer.



## CommonContactDetailsPost

> ContactResponseSchema CommonContactDetailsPost(ctx).Format(format).ContactRequestSchema(contactRequestSchema).Execute()

##### Retrieve contact information of another acquirer.



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
    contactRequestSchema := *openapiclient.NewContactRequestSchema() // ContactRequestSchema | The contact request object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AcquirerContactRequestApi.CommonContactDetailsPost(context.Background()).Format(format).ContactRequestSchema(contactRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcquirerContactRequestApi.CommonContactDetailsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommonContactDetailsPost`: ContactResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `AcquirerContactRequestApi.CommonContactDetailsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommonContactDetailsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Describes format of the response you wants to serverd, response can be delevired either as XML or JSON. | 
 **contactRequestSchema** | [**ContactRequestSchema**](ContactRequestSchema.md) | The contact request object | 

### Return type

[**ContactResponseSchema**](ContactResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

