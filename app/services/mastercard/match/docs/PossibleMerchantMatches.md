# PossibleMerchantMatches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalLength** | Pointer to **int32** | The total length of the result set from possible merchant matches of inquiry. | [optional] 
**TerminatedMerchant** | Pointer to [**[]TerminatedMerchant**](TerminatedMerchant.md) |  | [optional] 

## Methods

### NewPossibleMerchantMatches

`func NewPossibleMerchantMatches() *PossibleMerchantMatches`

NewPossibleMerchantMatches instantiates a new PossibleMerchantMatches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPossibleMerchantMatchesWithDefaults

`func NewPossibleMerchantMatchesWithDefaults() *PossibleMerchantMatches`

NewPossibleMerchantMatchesWithDefaults instantiates a new PossibleMerchantMatches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalLength

`func (o *PossibleMerchantMatches) GetTotalLength() int32`

GetTotalLength returns the TotalLength field if non-nil, zero value otherwise.

### GetTotalLengthOk

`func (o *PossibleMerchantMatches) GetTotalLengthOk() (*int32, bool)`

GetTotalLengthOk returns a tuple with the TotalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLength

`func (o *PossibleMerchantMatches) SetTotalLength(v int32)`

SetTotalLength sets TotalLength field to given value.

### HasTotalLength

`func (o *PossibleMerchantMatches) HasTotalLength() bool`

HasTotalLength returns a boolean if a field has been set.

### GetTerminatedMerchant

`func (o *PossibleMerchantMatches) GetTerminatedMerchant() []TerminatedMerchant`

GetTerminatedMerchant returns the TerminatedMerchant field if non-nil, zero value otherwise.

### GetTerminatedMerchantOk

`func (o *PossibleMerchantMatches) GetTerminatedMerchantOk() (*[]TerminatedMerchant, bool)`

GetTerminatedMerchantOk returns a tuple with the TerminatedMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedMerchant

`func (o *PossibleMerchantMatches) SetTerminatedMerchant(v []TerminatedMerchant)`

SetTerminatedMerchant sets TerminatedMerchant field to given value.

### HasTerminatedMerchant

`func (o *PossibleMerchantMatches) HasTerminatedMerchant() bool`

HasTerminatedMerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


