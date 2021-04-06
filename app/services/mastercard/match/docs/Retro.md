# Retro

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefNum** | Pointer to **string** | The reference number of retroactive inquiry. Using this reference number, the details can be obtained. | [optional] 
**Date** | Pointer to **string** | Date on which the retro inquiry was made. This will be in MM/DD/YYYY format. | [optional] 
**BusinessName** | Pointer to **string** | Name of Merchant or Business corresponds to retro inquiry. | [optional] 
**City** | Pointer to **string** | The name of the city for the merchant location corresponds to retro inquiry. | [optional] 
**State** | Pointer to **string** | The abbreviated state for a merchant location corresponds to retro inquiry. | [optional] 
**Country** | Pointer to **string** | The three digit country code corresponds to retro inquiry. Valid values are Three digit alpha country codes as defined in ISO 3166-1. | [optional] 

## Methods

### NewRetro

`func NewRetro() *Retro`

NewRetro instantiates a new Retro object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetroWithDefaults

`func NewRetroWithDefaults() *Retro`

NewRetroWithDefaults instantiates a new Retro object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefNum

`func (o *Retro) GetRefNum() string`

GetRefNum returns the RefNum field if non-nil, zero value otherwise.

### GetRefNumOk

`func (o *Retro) GetRefNumOk() (*string, bool)`

GetRefNumOk returns a tuple with the RefNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNum

`func (o *Retro) SetRefNum(v string)`

SetRefNum sets RefNum field to given value.

### HasRefNum

`func (o *Retro) HasRefNum() bool`

HasRefNum returns a boolean if a field has been set.

### GetDate

`func (o *Retro) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Retro) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Retro) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Retro) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetBusinessName

`func (o *Retro) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *Retro) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *Retro) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *Retro) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetCity

`func (o *Retro) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Retro) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Retro) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Retro) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *Retro) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Retro) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Retro) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Retro) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *Retro) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Retro) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Retro) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Retro) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


