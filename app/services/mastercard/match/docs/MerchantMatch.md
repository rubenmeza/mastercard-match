# MerchantMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Business which has been terminated. | [optional] 
**DoingBusinessAsName** | Pointer to **string** | The name used by a merchant that could be different from the legal name of the business. Such as Bait R Us instead of the legal name, The Bait Shop | [optional] 
**PhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s phone number. | [optional] 
**Address** | Pointer to **string** | Address of the merchant location | [optional] 
**AltPhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s alternate phone number. | [optional] 
**CountrySubdivisionTaxId** | Pointer to **string** | The Merchant&#39;s state tax ID; for the U.S region only. Return value will be hidden. | [optional] 
**NationalTaxId** | Pointer to **string** | The National tax ID or business registration number. Return value will be hidden. | [optional] 
**ServiceProvLegal** | Pointer to **string** | The name of the service provider associated with the merchant listed in the MATCH. | [optional] 
**ServiceProvDBA** | Pointer to **string** | The name of the service provider associated with the merchant listed in the MATCH. | [optional] 
**PrincipalMatch** | Pointer to [**[]PrincipalMatch**](PrincipalMatch.md) |  | [optional] 
**UrlMatch** | Pointer to [**[]UrlMatch**](UrlMatch.md) |  | [optional] 

## Methods

### NewMerchantMatch

`func NewMerchantMatch() *MerchantMatch`

NewMerchantMatch instantiates a new MerchantMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantMatchWithDefaults

`func NewMerchantMatchWithDefaults() *MerchantMatch`

NewMerchantMatchWithDefaults instantiates a new MerchantMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MerchantMatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantMatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantMatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantMatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDoingBusinessAsName

`func (o *MerchantMatch) GetDoingBusinessAsName() string`

GetDoingBusinessAsName returns the DoingBusinessAsName field if non-nil, zero value otherwise.

### GetDoingBusinessAsNameOk

`func (o *MerchantMatch) GetDoingBusinessAsNameOk() (*string, bool)`

GetDoingBusinessAsNameOk returns a tuple with the DoingBusinessAsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoingBusinessAsName

`func (o *MerchantMatch) SetDoingBusinessAsName(v string)`

SetDoingBusinessAsName sets DoingBusinessAsName field to given value.

### HasDoingBusinessAsName

`func (o *MerchantMatch) HasDoingBusinessAsName() bool`

HasDoingBusinessAsName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *MerchantMatch) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MerchantMatch) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MerchantMatch) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MerchantMatch) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAddress

`func (o *MerchantMatch) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MerchantMatch) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MerchantMatch) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MerchantMatch) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAltPhoneNumber

`func (o *MerchantMatch) GetAltPhoneNumber() string`

GetAltPhoneNumber returns the AltPhoneNumber field if non-nil, zero value otherwise.

### GetAltPhoneNumberOk

`func (o *MerchantMatch) GetAltPhoneNumberOk() (*string, bool)`

GetAltPhoneNumberOk returns a tuple with the AltPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltPhoneNumber

`func (o *MerchantMatch) SetAltPhoneNumber(v string)`

SetAltPhoneNumber sets AltPhoneNumber field to given value.

### HasAltPhoneNumber

`func (o *MerchantMatch) HasAltPhoneNumber() bool`

HasAltPhoneNumber returns a boolean if a field has been set.

### GetCountrySubdivisionTaxId

`func (o *MerchantMatch) GetCountrySubdivisionTaxId() string`

GetCountrySubdivisionTaxId returns the CountrySubdivisionTaxId field if non-nil, zero value otherwise.

### GetCountrySubdivisionTaxIdOk

`func (o *MerchantMatch) GetCountrySubdivisionTaxIdOk() (*string, bool)`

GetCountrySubdivisionTaxIdOk returns a tuple with the CountrySubdivisionTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivisionTaxId

`func (o *MerchantMatch) SetCountrySubdivisionTaxId(v string)`

SetCountrySubdivisionTaxId sets CountrySubdivisionTaxId field to given value.

### HasCountrySubdivisionTaxId

`func (o *MerchantMatch) HasCountrySubdivisionTaxId() bool`

HasCountrySubdivisionTaxId returns a boolean if a field has been set.

### GetNationalTaxId

`func (o *MerchantMatch) GetNationalTaxId() string`

GetNationalTaxId returns the NationalTaxId field if non-nil, zero value otherwise.

### GetNationalTaxIdOk

`func (o *MerchantMatch) GetNationalTaxIdOk() (*string, bool)`

GetNationalTaxIdOk returns a tuple with the NationalTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalTaxId

`func (o *MerchantMatch) SetNationalTaxId(v string)`

SetNationalTaxId sets NationalTaxId field to given value.

### HasNationalTaxId

`func (o *MerchantMatch) HasNationalTaxId() bool`

HasNationalTaxId returns a boolean if a field has been set.

### GetServiceProvLegal

`func (o *MerchantMatch) GetServiceProvLegal() string`

GetServiceProvLegal returns the ServiceProvLegal field if non-nil, zero value otherwise.

### GetServiceProvLegalOk

`func (o *MerchantMatch) GetServiceProvLegalOk() (*string, bool)`

GetServiceProvLegalOk returns a tuple with the ServiceProvLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvLegal

`func (o *MerchantMatch) SetServiceProvLegal(v string)`

SetServiceProvLegal sets ServiceProvLegal field to given value.

### HasServiceProvLegal

`func (o *MerchantMatch) HasServiceProvLegal() bool`

HasServiceProvLegal returns a boolean if a field has been set.

### GetServiceProvDBA

`func (o *MerchantMatch) GetServiceProvDBA() string`

GetServiceProvDBA returns the ServiceProvDBA field if non-nil, zero value otherwise.

### GetServiceProvDBAOk

`func (o *MerchantMatch) GetServiceProvDBAOk() (*string, bool)`

GetServiceProvDBAOk returns a tuple with the ServiceProvDBA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvDBA

`func (o *MerchantMatch) SetServiceProvDBA(v string)`

SetServiceProvDBA sets ServiceProvDBA field to given value.

### HasServiceProvDBA

`func (o *MerchantMatch) HasServiceProvDBA() bool`

HasServiceProvDBA returns a boolean if a field has been set.

### GetPrincipalMatch

`func (o *MerchantMatch) GetPrincipalMatch() []PrincipalMatch`

GetPrincipalMatch returns the PrincipalMatch field if non-nil, zero value otherwise.

### GetPrincipalMatchOk

`func (o *MerchantMatch) GetPrincipalMatchOk() (*[]PrincipalMatch, bool)`

GetPrincipalMatchOk returns a tuple with the PrincipalMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalMatch

`func (o *MerchantMatch) SetPrincipalMatch(v []PrincipalMatch)`

SetPrincipalMatch sets PrincipalMatch field to given value.

### HasPrincipalMatch

`func (o *MerchantMatch) HasPrincipalMatch() bool`

HasPrincipalMatch returns a boolean if a field has been set.

### GetUrlMatch

`func (o *MerchantMatch) GetUrlMatch() []UrlMatch`

GetUrlMatch returns the UrlMatch field if non-nil, zero value otherwise.

### GetUrlMatchOk

`func (o *MerchantMatch) GetUrlMatchOk() (*[]UrlMatch, bool)`

GetUrlMatchOk returns a tuple with the UrlMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlMatch

`func (o *MerchantMatch) SetUrlMatch(v []UrlMatch)`

SetUrlMatch sets UrlMatch field to given value.

### HasUrlMatch

`func (o *MerchantMatch) HasUrlMatch() bool`

HasUrlMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


