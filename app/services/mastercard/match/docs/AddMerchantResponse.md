# AddMerchantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantReferenceNumber** | Pointer to **string** | A number assigned by the MATCH system to uniquely identify the newly added merchant. Return of this number indicates a successful add. | [optional] 
**Name** | Pointer to **string** | Name of the merchant that was added to the MATCH system as verification of what was added. | [optional] 

## Methods

### NewAddMerchantResponse

`func NewAddMerchantResponse() *AddMerchantResponse`

NewAddMerchantResponse instantiates a new AddMerchantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMerchantResponseWithDefaults

`func NewAddMerchantResponseWithDefaults() *AddMerchantResponse`

NewAddMerchantResponseWithDefaults instantiates a new AddMerchantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantReferenceNumber

`func (o *AddMerchantResponse) GetMerchantReferenceNumber() string`

GetMerchantReferenceNumber returns the MerchantReferenceNumber field if non-nil, zero value otherwise.

### GetMerchantReferenceNumberOk

`func (o *AddMerchantResponse) GetMerchantReferenceNumberOk() (*string, bool)`

GetMerchantReferenceNumberOk returns a tuple with the MerchantReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReferenceNumber

`func (o *AddMerchantResponse) SetMerchantReferenceNumber(v string)`

SetMerchantReferenceNumber sets MerchantReferenceNumber field to given value.

### HasMerchantReferenceNumber

`func (o *AddMerchantResponse) HasMerchantReferenceNumber() bool`

HasMerchantReferenceNumber returns a boolean if a field has been set.

### GetName

`func (o *AddMerchantResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddMerchantResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddMerchantResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddMerchantResponse) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


