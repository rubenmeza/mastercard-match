# Merchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the business assigned by the principal owner(s) | 
**DoingBusinessAsName** | Pointer to **string** | The name used by a merchant that could be different from the legal name of the business. Such as Bait R Us instead of the legal name, The Bait Shop. | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s phone number, including the area code. Within the USA, phone numbers have a length of 10, and for outside the US, it can be any length with a maximum of 12 digits. Within the U.S. phone numbers can not start with 0 or 1. If the number is outside the U.S. region; do not include the country code. The phone number must be digits only, with no format characters such as parenthesis or dashes. | [optional] 
**AltPhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s alternate phone number, including the area code. Within the USA, phone numbers have a length of 10, and for out outside the US, a max length of 25. Within the U.S. phone numbers can not start with 0 or 1. If the number is outside the U.S. region; do not include the country code. The number must be digits only, with no format characters such as parenthesis or dashes. | [optional] 
**NationalTaxId** | Pointer to **string** | The Merchant national tax ID, leave blank if not in the U.S region. | [optional] 
**CountrySubdivisionTaxId** | Pointer to **string** | The Merchant Country Subdivision tax ID, leave blank if not in the U.S region. | [optional] 
**ServiceProvLegal** | Pointer to **string** | The name of the service provider associated with the merchant listed in the MATCH | [optional] 
**ServiceProvDBA** | Pointer to **string** | The name of the service provider associated with the merchant listed in the MATCH | [optional] 
**Url** | Pointer to **[]string** | Website address of the merchant. A request may include multiple URLs . The total cumulative size of the URLs cannot exceed 20000 bytes. | [optional] 
**Principal** | Pointer to [**[]Principal**](Principal.md) | The details for the principal owner of the business.  A maximum of 5 principals may be submitted. | [optional] 
**SearchCriteria** | Pointer to [**SearchCriteria**](SearchCriteria.md) |  | [optional] 
**AddedOnDate** | Pointer to **string** | Date the merchant was added to the MATCH database. | [optional] 
**TerminationReasonCode** | Pointer to **string** | A two digit numeric code indication why a particular merchant was terminated.  01   Account Data Compromise, 02   Common Points of Purchase, 03   Laundering, 04   Excessive Chargebacks, 05   Excessive Fraud, 06   Reserved for Future Use, 07   Fraud Conviction, 08   MasterCard Questionable Merchant Audit Program, 09   Bankruptcy/Liquidation/Insolvency, 10   Violation of MasterCard Standards, 11   Merchant collusion, 12   PCI Data Security Standard, Noncompliance, 13   Illegal Transactions, 14   Identity Theft | [optional] 
**AddedByAcquirerID** | Pointer to **string** | The Member ICA that has added the merchant to the MATCH system | [optional] 
**UrlGroup** | Pointer to [**[]UrlGroup**](UrlGroup.md) |  | [optional] 

## Methods

### NewMerchant

`func NewMerchant(name string, ) *Merchant`

NewMerchant instantiates a new Merchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantWithDefaults

`func NewMerchantWithDefaults() *Merchant`

NewMerchantWithDefaults instantiates a new Merchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Merchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Merchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Merchant) SetName(v string)`

SetName sets Name field to given value.


### GetDoingBusinessAsName

`func (o *Merchant) GetDoingBusinessAsName() string`

GetDoingBusinessAsName returns the DoingBusinessAsName field if non-nil, zero value otherwise.

### GetDoingBusinessAsNameOk

`func (o *Merchant) GetDoingBusinessAsNameOk() (*string, bool)`

GetDoingBusinessAsNameOk returns a tuple with the DoingBusinessAsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoingBusinessAsName

`func (o *Merchant) SetDoingBusinessAsName(v string)`

SetDoingBusinessAsName sets DoingBusinessAsName field to given value.

### HasDoingBusinessAsName

`func (o *Merchant) HasDoingBusinessAsName() bool`

HasDoingBusinessAsName returns a boolean if a field has been set.

### GetAddress

`func (o *Merchant) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Merchant) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Merchant) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Merchant) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Merchant) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Merchant) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Merchant) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Merchant) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAltPhoneNumber

`func (o *Merchant) GetAltPhoneNumber() string`

GetAltPhoneNumber returns the AltPhoneNumber field if non-nil, zero value otherwise.

### GetAltPhoneNumberOk

`func (o *Merchant) GetAltPhoneNumberOk() (*string, bool)`

GetAltPhoneNumberOk returns a tuple with the AltPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltPhoneNumber

`func (o *Merchant) SetAltPhoneNumber(v string)`

SetAltPhoneNumber sets AltPhoneNumber field to given value.

### HasAltPhoneNumber

`func (o *Merchant) HasAltPhoneNumber() bool`

HasAltPhoneNumber returns a boolean if a field has been set.

### GetNationalTaxId

`func (o *Merchant) GetNationalTaxId() string`

GetNationalTaxId returns the NationalTaxId field if non-nil, zero value otherwise.

### GetNationalTaxIdOk

`func (o *Merchant) GetNationalTaxIdOk() (*string, bool)`

GetNationalTaxIdOk returns a tuple with the NationalTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalTaxId

`func (o *Merchant) SetNationalTaxId(v string)`

SetNationalTaxId sets NationalTaxId field to given value.

### HasNationalTaxId

`func (o *Merchant) HasNationalTaxId() bool`

HasNationalTaxId returns a boolean if a field has been set.

### GetCountrySubdivisionTaxId

`func (o *Merchant) GetCountrySubdivisionTaxId() string`

GetCountrySubdivisionTaxId returns the CountrySubdivisionTaxId field if non-nil, zero value otherwise.

### GetCountrySubdivisionTaxIdOk

`func (o *Merchant) GetCountrySubdivisionTaxIdOk() (*string, bool)`

GetCountrySubdivisionTaxIdOk returns a tuple with the CountrySubdivisionTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivisionTaxId

`func (o *Merchant) SetCountrySubdivisionTaxId(v string)`

SetCountrySubdivisionTaxId sets CountrySubdivisionTaxId field to given value.

### HasCountrySubdivisionTaxId

`func (o *Merchant) HasCountrySubdivisionTaxId() bool`

HasCountrySubdivisionTaxId returns a boolean if a field has been set.

### GetServiceProvLegal

`func (o *Merchant) GetServiceProvLegal() string`

GetServiceProvLegal returns the ServiceProvLegal field if non-nil, zero value otherwise.

### GetServiceProvLegalOk

`func (o *Merchant) GetServiceProvLegalOk() (*string, bool)`

GetServiceProvLegalOk returns a tuple with the ServiceProvLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvLegal

`func (o *Merchant) SetServiceProvLegal(v string)`

SetServiceProvLegal sets ServiceProvLegal field to given value.

### HasServiceProvLegal

`func (o *Merchant) HasServiceProvLegal() bool`

HasServiceProvLegal returns a boolean if a field has been set.

### GetServiceProvDBA

`func (o *Merchant) GetServiceProvDBA() string`

GetServiceProvDBA returns the ServiceProvDBA field if non-nil, zero value otherwise.

### GetServiceProvDBAOk

`func (o *Merchant) GetServiceProvDBAOk() (*string, bool)`

GetServiceProvDBAOk returns a tuple with the ServiceProvDBA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvDBA

`func (o *Merchant) SetServiceProvDBA(v string)`

SetServiceProvDBA sets ServiceProvDBA field to given value.

### HasServiceProvDBA

`func (o *Merchant) HasServiceProvDBA() bool`

HasServiceProvDBA returns a boolean if a field has been set.

### GetUrl

`func (o *Merchant) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Merchant) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Merchant) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Merchant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPrincipal

`func (o *Merchant) GetPrincipal() []Principal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *Merchant) GetPrincipalOk() (*[]Principal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *Merchant) SetPrincipal(v []Principal)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *Merchant) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetSearchCriteria

`func (o *Merchant) GetSearchCriteria() SearchCriteria`

GetSearchCriteria returns the SearchCriteria field if non-nil, zero value otherwise.

### GetSearchCriteriaOk

`func (o *Merchant) GetSearchCriteriaOk() (*SearchCriteria, bool)`

GetSearchCriteriaOk returns a tuple with the SearchCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCriteria

`func (o *Merchant) SetSearchCriteria(v SearchCriteria)`

SetSearchCriteria sets SearchCriteria field to given value.

### HasSearchCriteria

`func (o *Merchant) HasSearchCriteria() bool`

HasSearchCriteria returns a boolean if a field has been set.

### GetAddedOnDate

`func (o *Merchant) GetAddedOnDate() string`

GetAddedOnDate returns the AddedOnDate field if non-nil, zero value otherwise.

### GetAddedOnDateOk

`func (o *Merchant) GetAddedOnDateOk() (*string, bool)`

GetAddedOnDateOk returns a tuple with the AddedOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOnDate

`func (o *Merchant) SetAddedOnDate(v string)`

SetAddedOnDate sets AddedOnDate field to given value.

### HasAddedOnDate

`func (o *Merchant) HasAddedOnDate() bool`

HasAddedOnDate returns a boolean if a field has been set.

### GetTerminationReasonCode

`func (o *Merchant) GetTerminationReasonCode() string`

GetTerminationReasonCode returns the TerminationReasonCode field if non-nil, zero value otherwise.

### GetTerminationReasonCodeOk

`func (o *Merchant) GetTerminationReasonCodeOk() (*string, bool)`

GetTerminationReasonCodeOk returns a tuple with the TerminationReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationReasonCode

`func (o *Merchant) SetTerminationReasonCode(v string)`

SetTerminationReasonCode sets TerminationReasonCode field to given value.

### HasTerminationReasonCode

`func (o *Merchant) HasTerminationReasonCode() bool`

HasTerminationReasonCode returns a boolean if a field has been set.

### GetAddedByAcquirerID

`func (o *Merchant) GetAddedByAcquirerID() string`

GetAddedByAcquirerID returns the AddedByAcquirerID field if non-nil, zero value otherwise.

### GetAddedByAcquirerIDOk

`func (o *Merchant) GetAddedByAcquirerIDOk() (*string, bool)`

GetAddedByAcquirerIDOk returns a tuple with the AddedByAcquirerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedByAcquirerID

`func (o *Merchant) SetAddedByAcquirerID(v string)`

SetAddedByAcquirerID sets AddedByAcquirerID field to given value.

### HasAddedByAcquirerID

`func (o *Merchant) HasAddedByAcquirerID() bool`

HasAddedByAcquirerID returns a boolean if a field has been set.

### GetUrlGroup

`func (o *Merchant) GetUrlGroup() []UrlGroup`

GetUrlGroup returns the UrlGroup field if non-nil, zero value otherwise.

### GetUrlGroupOk

`func (o *Merchant) GetUrlGroupOk() (*[]UrlGroup, bool)`

GetUrlGroupOk returns a tuple with the UrlGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlGroup

`func (o *Merchant) SetUrlGroup(v []UrlGroup)`

SetUrlGroup sets UrlGroup field to given value.

### HasUrlGroup

`func (o *Merchant) HasUrlGroup() bool`

HasUrlGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


