# PossibleInquiryMatches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalLength** | Pointer to **int32** | The total length of the result set from possible merchant matches of inquiry. | [optional] 
**InquiredMerchant** | Pointer to [**[]InquiredMerchant**](InquiredMerchant.md) |  | [optional] 

## Methods

### NewPossibleInquiryMatches

`func NewPossibleInquiryMatches() *PossibleInquiryMatches`

NewPossibleInquiryMatches instantiates a new PossibleInquiryMatches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPossibleInquiryMatchesWithDefaults

`func NewPossibleInquiryMatchesWithDefaults() *PossibleInquiryMatches`

NewPossibleInquiryMatchesWithDefaults instantiates a new PossibleInquiryMatches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalLength

`func (o *PossibleInquiryMatches) GetTotalLength() int32`

GetTotalLength returns the TotalLength field if non-nil, zero value otherwise.

### GetTotalLengthOk

`func (o *PossibleInquiryMatches) GetTotalLengthOk() (*int32, bool)`

GetTotalLengthOk returns a tuple with the TotalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLength

`func (o *PossibleInquiryMatches) SetTotalLength(v int32)`

SetTotalLength sets TotalLength field to given value.

### HasTotalLength

`func (o *PossibleInquiryMatches) HasTotalLength() bool`

HasTotalLength returns a boolean if a field has been set.

### GetInquiredMerchant

`func (o *PossibleInquiryMatches) GetInquiredMerchant() []InquiredMerchant`

GetInquiredMerchant returns the InquiredMerchant field if non-nil, zero value otherwise.

### GetInquiredMerchantOk

`func (o *PossibleInquiryMatches) GetInquiredMerchantOk() (*[]InquiredMerchant, bool)`

GetInquiredMerchantOk returns a tuple with the InquiredMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquiredMerchant

`func (o *PossibleInquiryMatches) SetInquiredMerchant(v []InquiredMerchant)`

SetInquiredMerchant sets InquiredMerchant field to given value.

### HasInquiredMerchant

`func (o *PossibleInquiryMatches) HasInquiredMerchant() bool`

HasInquiredMerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


