# AddMerchantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquirerId** | **string** | The ID of the acquirer. | 
**Merchant** | [**AddMerchantMerchant**](AddMerchantMerchant.md) |  | 

## Methods

### NewAddMerchantRequest

`func NewAddMerchantRequest(acquirerId string, merchant AddMerchantMerchant, ) *AddMerchantRequest`

NewAddMerchantRequest instantiates a new AddMerchantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMerchantRequestWithDefaults

`func NewAddMerchantRequestWithDefaults() *AddMerchantRequest`

NewAddMerchantRequestWithDefaults instantiates a new AddMerchantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquirerId

`func (o *AddMerchantRequest) GetAcquirerId() string`

GetAcquirerId returns the AcquirerId field if non-nil, zero value otherwise.

### GetAcquirerIdOk

`func (o *AddMerchantRequest) GetAcquirerIdOk() (*string, bool)`

GetAcquirerIdOk returns a tuple with the AcquirerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerId

`func (o *AddMerchantRequest) SetAcquirerId(v string)`

SetAcquirerId sets AcquirerId field to given value.


### GetMerchant

`func (o *AddMerchantRequest) GetMerchant() AddMerchantMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *AddMerchantRequest) GetMerchantOk() (*AddMerchantMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *AddMerchantRequest) SetMerchant(v AddMerchantMerchant)`

SetMerchant sets Merchant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


