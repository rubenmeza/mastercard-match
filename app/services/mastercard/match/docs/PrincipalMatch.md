# PrincipalMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Business which has been terminated. | [optional] 
**Address** | Pointer to **string** | Address of the merchant location. | [optional] 
**PhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s phone number, including the area code. Within the USA, phone numbers have a length of 10, and for outside the US, it can be any length with a maximum of 12 digits. Within the U.S. phone numbers can not start with 0 or 1. If the number is outside the U.S. region; do not include the country code. The phone number must be digits only, with no format characters such as parenthesis or dashes. | [optional] 
**AltPhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s alternate phone number, including the area code. Within the USA, phone numbers have a length of 10, and for out outside the US, a max length of 25. Within the U.S. phone numbers can not start with 0 or 1. If the number is outside the U.S. region; do not include the country code. The number must be digits only, with no format characters such as parenthesis or dashes. | [optional] 
**NationalId** | Pointer to **string** | The Social Security number of a principal owner. If the principal owner is not from the U.S. Region, then use their national ID card number. | [optional] 
**DriversLicense** | Pointer to **string** | The drivers license number of a principal owner. owner is not from the U.S. Region, then use their national ID card number. | [optional] 

## Methods

### NewPrincipalMatch

`func NewPrincipalMatch() *PrincipalMatch`

NewPrincipalMatch instantiates a new PrincipalMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalMatchWithDefaults

`func NewPrincipalMatchWithDefaults() *PrincipalMatch`

NewPrincipalMatchWithDefaults instantiates a new PrincipalMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrincipalMatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalMatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalMatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrincipalMatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *PrincipalMatch) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PrincipalMatch) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PrincipalMatch) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PrincipalMatch) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PrincipalMatch) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PrincipalMatch) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PrincipalMatch) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PrincipalMatch) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAltPhoneNumber

`func (o *PrincipalMatch) GetAltPhoneNumber() string`

GetAltPhoneNumber returns the AltPhoneNumber field if non-nil, zero value otherwise.

### GetAltPhoneNumberOk

`func (o *PrincipalMatch) GetAltPhoneNumberOk() (*string, bool)`

GetAltPhoneNumberOk returns a tuple with the AltPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltPhoneNumber

`func (o *PrincipalMatch) SetAltPhoneNumber(v string)`

SetAltPhoneNumber sets AltPhoneNumber field to given value.

### HasAltPhoneNumber

`func (o *PrincipalMatch) HasAltPhoneNumber() bool`

HasAltPhoneNumber returns a boolean if a field has been set.

### GetNationalId

`func (o *PrincipalMatch) GetNationalId() string`

GetNationalId returns the NationalId field if non-nil, zero value otherwise.

### GetNationalIdOk

`func (o *PrincipalMatch) GetNationalIdOk() (*string, bool)`

GetNationalIdOk returns a tuple with the NationalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalId

`func (o *PrincipalMatch) SetNationalId(v string)`

SetNationalId sets NationalId field to given value.

### HasNationalId

`func (o *PrincipalMatch) HasNationalId() bool`

HasNationalId returns a boolean if a field has been set.

### GetDriversLicense

`func (o *PrincipalMatch) GetDriversLicense() string`

GetDriversLicense returns the DriversLicense field if non-nil, zero value otherwise.

### GetDriversLicenseOk

`func (o *PrincipalMatch) GetDriversLicenseOk() (*string, bool)`

GetDriversLicenseOk returns a tuple with the DriversLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriversLicense

`func (o *PrincipalMatch) SetDriversLicense(v string)`

SetDriversLicense sets DriversLicense field to given value.

### HasDriversLicense

`func (o *PrincipalMatch) HasDriversLicense() bool`

HasDriversLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


