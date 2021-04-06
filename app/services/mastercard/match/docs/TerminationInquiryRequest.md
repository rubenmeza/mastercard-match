# TerminationInquiryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquirerId** | **string** | The Member ICA number which is used to validate that the application has permission to hit the MATCH database. This number must be obtained from a participating MATCH acquiring bank or entity before a developer can access the MATCH resource. | 
**Merchant** | Pointer to [**Merchant**](Merchant.md) |  | [optional] 

## Methods

### NewTerminationInquiryRequest

`func NewTerminationInquiryRequest(acquirerId string, ) *TerminationInquiryRequest`

NewTerminationInquiryRequest instantiates a new TerminationInquiryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminationInquiryRequestWithDefaults

`func NewTerminationInquiryRequestWithDefaults() *TerminationInquiryRequest`

NewTerminationInquiryRequestWithDefaults instantiates a new TerminationInquiryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquirerId

`func (o *TerminationInquiryRequest) GetAcquirerId() string`

GetAcquirerId returns the AcquirerId field if non-nil, zero value otherwise.

### GetAcquirerIdOk

`func (o *TerminationInquiryRequest) GetAcquirerIdOk() (*string, bool)`

GetAcquirerIdOk returns a tuple with the AcquirerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerId

`func (o *TerminationInquiryRequest) SetAcquirerId(v string)`

SetAcquirerId sets AcquirerId field to given value.


### GetMerchant

`func (o *TerminationInquiryRequest) GetMerchant() Merchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TerminationInquiryRequest) GetMerchantOk() (*Merchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TerminationInquiryRequest) SetMerchant(v Merchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *TerminationInquiryRequest) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


