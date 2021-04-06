# DriversLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | The drivers license number of a principal owner. Required when Drivers License Country is provided. Should not be provided if Drivers License Country is not provided. | [optional] 
**CountrySubdivision** | Pointer to **string** | The abbreviated state or province code for a merchant location (only supported for US and Canada merchants).  Required when Drivers License Country is &#39;USA&#39; . Should not be provided for non-US Drivers License Country. | [optional] 
**Country** | Pointer to **string** | The three digit country code of the principal owner. Valid values are Three digit alpha country codes as defined in ISO 3166-1. Required when Drivers License Number is provided. Should not be provided if Drivers License Number is not provided. | [optional] 

## Methods

### NewDriversLicense

`func NewDriversLicense() *DriversLicense`

NewDriversLicense instantiates a new DriversLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriversLicenseWithDefaults

`func NewDriversLicenseWithDefaults() *DriversLicense`

NewDriversLicenseWithDefaults instantiates a new DriversLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *DriversLicense) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *DriversLicense) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *DriversLicense) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *DriversLicense) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetCountrySubdivision

`func (o *DriversLicense) GetCountrySubdivision() string`

GetCountrySubdivision returns the CountrySubdivision field if non-nil, zero value otherwise.

### GetCountrySubdivisionOk

`func (o *DriversLicense) GetCountrySubdivisionOk() (*string, bool)`

GetCountrySubdivisionOk returns a tuple with the CountrySubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivision

`func (o *DriversLicense) SetCountrySubdivision(v string)`

SetCountrySubdivision sets CountrySubdivision field to given value.

### HasCountrySubdivision

`func (o *DriversLicense) HasCountrySubdivision() bool`

HasCountrySubdivision returns a boolean if a field has been set.

### GetCountry

`func (o *DriversLicense) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DriversLicense) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DriversLicense) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DriversLicense) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


