# Principal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | The first name of the principal owner of the business. | 
**MiddleInitial** | Pointer to **string** | THe middle initial of the name of the principal owner of the business. | [optional] 
**LastName** | **string** | The last name of the principal owner of the business. | 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s phone number, including the area code. Within the USA, phone numbers have a length of 10, and for outside the US, it can be any length with a maximum of 12 digits. Within the U.S. phone numbers can not start with 0 or 1. If the number is outside the U.S. region; do not include the country code. The phone number must be digits only, with no format characters such as parenthesis or dashes. | [optional] 
**AltPhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s alternate phone number, including the area code. Within the USA, phone numbers have a length of 10, and for out outside the US, a max length of 25. Within the U.S. phone numbers can not start with 0 or 1. If the number is outside the U.S. region; do not include the country code. The number must be digits only, with no format characters such as parenthesis or dashes. | [optional] 
**NationalId** | Pointer to **string** | The Social Security number of a principal owner. If the principal owner is not from the U.S. Region, then use their national ID card number. | [optional] 
**DriversLicense** | Pointer to [**DriversLicense**](DriversLicense.md) |  | [optional] 

## Methods

### NewPrincipal

`func NewPrincipal(firstName string, lastName string, ) *Principal`

NewPrincipal instantiates a new Principal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalWithDefaults

`func NewPrincipalWithDefaults() *Principal`

NewPrincipalWithDefaults instantiates a new Principal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *Principal) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Principal) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Principal) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleInitial

`func (o *Principal) GetMiddleInitial() string`

GetMiddleInitial returns the MiddleInitial field if non-nil, zero value otherwise.

### GetMiddleInitialOk

`func (o *Principal) GetMiddleInitialOk() (*string, bool)`

GetMiddleInitialOk returns a tuple with the MiddleInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleInitial

`func (o *Principal) SetMiddleInitial(v string)`

SetMiddleInitial sets MiddleInitial field to given value.

### HasMiddleInitial

`func (o *Principal) HasMiddleInitial() bool`

HasMiddleInitial returns a boolean if a field has been set.

### GetLastName

`func (o *Principal) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Principal) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Principal) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetAddress

`func (o *Principal) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Principal) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Principal) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Principal) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Principal) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Principal) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Principal) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Principal) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAltPhoneNumber

`func (o *Principal) GetAltPhoneNumber() string`

GetAltPhoneNumber returns the AltPhoneNumber field if non-nil, zero value otherwise.

### GetAltPhoneNumberOk

`func (o *Principal) GetAltPhoneNumberOk() (*string, bool)`

GetAltPhoneNumberOk returns a tuple with the AltPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltPhoneNumber

`func (o *Principal) SetAltPhoneNumber(v string)`

SetAltPhoneNumber sets AltPhoneNumber field to given value.

### HasAltPhoneNumber

`func (o *Principal) HasAltPhoneNumber() bool`

HasAltPhoneNumber returns a boolean if a field has been set.

### GetNationalId

`func (o *Principal) GetNationalId() string`

GetNationalId returns the NationalId field if non-nil, zero value otherwise.

### GetNationalIdOk

`func (o *Principal) GetNationalIdOk() (*string, bool)`

GetNationalIdOk returns a tuple with the NationalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalId

`func (o *Principal) SetNationalId(v string)`

SetNationalId sets NationalId field to given value.

### HasNationalId

`func (o *Principal) HasNationalId() bool`

HasNationalId returns a boolean if a field has been set.

### GetDriversLicense

`func (o *Principal) GetDriversLicense() DriversLicense`

GetDriversLicense returns the DriversLicense field if non-nil, zero value otherwise.

### GetDriversLicenseOk

`func (o *Principal) GetDriversLicenseOk() (*DriversLicense, bool)`

GetDriversLicenseOk returns a tuple with the DriversLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriversLicense

`func (o *Principal) SetDriversLicense(v DriversLicense)`

SetDriversLicense sets DriversLicense field to given value.

### HasDriversLicense

`func (o *Principal) HasDriversLicense() bool`

HasDriversLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


