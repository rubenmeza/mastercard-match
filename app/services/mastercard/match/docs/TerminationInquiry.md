# TerminationInquiry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageOffset** | Pointer to **int32** | PageOffset for the inquiry done | [optional] 
**Ref** | Pointer to **string** | rReference URL to get inquiry | [optional] 
**TransactionReferenceNumber** | Pointer to **string** | User-defined identifier for the inquiry submitted. | [optional] 
**PossibleMerchantMatches** | Pointer to [**[]PossibleMerchantMatches**](PossibleMerchantMatches.md) |  | [optional] 
**PossibleInquiryMatches** | Pointer to [**[]PossibleInquiryMatches**](PossibleInquiryMatches.md) |  | [optional] 

## Methods

### NewTerminationInquiry

`func NewTerminationInquiry() *TerminationInquiry`

NewTerminationInquiry instantiates a new TerminationInquiry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminationInquiryWithDefaults

`func NewTerminationInquiryWithDefaults() *TerminationInquiry`

NewTerminationInquiryWithDefaults instantiates a new TerminationInquiry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageOffset

`func (o *TerminationInquiry) GetPageOffset() int32`

GetPageOffset returns the PageOffset field if non-nil, zero value otherwise.

### GetPageOffsetOk

`func (o *TerminationInquiry) GetPageOffsetOk() (*int32, bool)`

GetPageOffsetOk returns a tuple with the PageOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageOffset

`func (o *TerminationInquiry) SetPageOffset(v int32)`

SetPageOffset sets PageOffset field to given value.

### HasPageOffset

`func (o *TerminationInquiry) HasPageOffset() bool`

HasPageOffset returns a boolean if a field has been set.

### GetRef

`func (o *TerminationInquiry) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *TerminationInquiry) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *TerminationInquiry) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *TerminationInquiry) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetTransactionReferenceNumber

`func (o *TerminationInquiry) GetTransactionReferenceNumber() string`

GetTransactionReferenceNumber returns the TransactionReferenceNumber field if non-nil, zero value otherwise.

### GetTransactionReferenceNumberOk

`func (o *TerminationInquiry) GetTransactionReferenceNumberOk() (*string, bool)`

GetTransactionReferenceNumberOk returns a tuple with the TransactionReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReferenceNumber

`func (o *TerminationInquiry) SetTransactionReferenceNumber(v string)`

SetTransactionReferenceNumber sets TransactionReferenceNumber field to given value.

### HasTransactionReferenceNumber

`func (o *TerminationInquiry) HasTransactionReferenceNumber() bool`

HasTransactionReferenceNumber returns a boolean if a field has been set.

### GetPossibleMerchantMatches

`func (o *TerminationInquiry) GetPossibleMerchantMatches() []PossibleMerchantMatches`

GetPossibleMerchantMatches returns the PossibleMerchantMatches field if non-nil, zero value otherwise.

### GetPossibleMerchantMatchesOk

`func (o *TerminationInquiry) GetPossibleMerchantMatchesOk() (*[]PossibleMerchantMatches, bool)`

GetPossibleMerchantMatchesOk returns a tuple with the PossibleMerchantMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleMerchantMatches

`func (o *TerminationInquiry) SetPossibleMerchantMatches(v []PossibleMerchantMatches)`

SetPossibleMerchantMatches sets PossibleMerchantMatches field to given value.

### HasPossibleMerchantMatches

`func (o *TerminationInquiry) HasPossibleMerchantMatches() bool`

HasPossibleMerchantMatches returns a boolean if a field has been set.

### GetPossibleInquiryMatches

`func (o *TerminationInquiry) GetPossibleInquiryMatches() []PossibleInquiryMatches`

GetPossibleInquiryMatches returns the PossibleInquiryMatches field if non-nil, zero value otherwise.

### GetPossibleInquiryMatchesOk

`func (o *TerminationInquiry) GetPossibleInquiryMatchesOk() (*[]PossibleInquiryMatches, bool)`

GetPossibleInquiryMatchesOk returns a tuple with the PossibleInquiryMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleInquiryMatches

`func (o *TerminationInquiry) SetPossibleInquiryMatches(v []PossibleInquiryMatches)`

SetPossibleInquiryMatches sets PossibleInquiryMatches field to given value.

### HasPossibleInquiryMatches

`func (o *TerminationInquiry) HasPossibleInquiryMatches() bool`

HasPossibleInquiryMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


