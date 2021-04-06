# InquiredMerchantMerchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Business which has been terminated. | [optional] 
**DoingBusinessAsName** | Pointer to **string** | The name used by a merchant that could be different from the legal name of the business. Such as Bait R Us instead of the legal name, The Bait Shop | [optional] 
**AddedByAcquirerID** | Pointer to **string** | The Member ICA that has added the merchant to the MATCH system | [optional] 
**AddedOnDate** | Pointer to **string** | The date on which the merchant was added to the MATCH system. Format MM/DD/YYYY | [optional] 
**PhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s phone number. | [optional] 
**AltPhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s alternate phone number. | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**CountrySubdivisionTaxId** | Pointer to **string** | The Merchant&#39;s state tax ID; for the U.S region only. Return value will be hidden. | [optional] 
**NationalTaxId** | Pointer to **string** | The National tax ID or business registration number. Return value will be hidden. | [optional] 
**ServiceProvLegal** | Pointer to **string** | The name of the service provider associated with the merchant listed in the MATCH. | [optional] 
**ServiceProvDBA** | Pointer to **string** | The name of the service provider associated with the merchant listed in the MATCH. | [optional] 
**TerminationReasonCode** | Pointer to **string** | A two digit numeric code indication why a particular merchant was terminated.  01   Account Data Compromise, 02   Common Points of Purchase, 03   Laundering, 04   Excessive Chargebacks, 05   Excessive Fraud, 06   Reserved for Future Use, 07   Fraud Conviction, 08   MasterCard Questionable Merchant Audit Program, 09   Bankruptcy/Liquidation/Insolvency, 10   Violation of MasterCard Standards, 11   Merchant collusion, 12   PCI Data Security Standard, Noncompliance, 13   Illegal Transactions, 14   Identity Theft | [optional] 
**UrlGroup** | Pointer to [**[]UrlGroup**](UrlGroup.md) |  | [optional] 
**Url** | Pointer to [**[]Url**](Url.md) |  | [optional] 
**Principal** | Pointer to [**[]Principal**](Principal.md) |  | [optional] 
**SearchCriteria** | Pointer to [**SearchCriteria**](SearchCriteria.md) |  | [optional] 
**MerchantMatch** | Pointer to [**MerchantMatch**](MerchantMatch.md) |  | [optional] 

## Methods

### NewInquiredMerchantMerchant

`func NewInquiredMerchantMerchant() *InquiredMerchantMerchant`

NewInquiredMerchantMerchant instantiates a new InquiredMerchantMerchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInquiredMerchantMerchantWithDefaults

`func NewInquiredMerchantMerchantWithDefaults() *InquiredMerchantMerchant`

NewInquiredMerchantMerchantWithDefaults instantiates a new InquiredMerchantMerchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InquiredMerchantMerchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InquiredMerchantMerchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InquiredMerchantMerchant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InquiredMerchantMerchant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDoingBusinessAsName

`func (o *InquiredMerchantMerchant) GetDoingBusinessAsName() string`

GetDoingBusinessAsName returns the DoingBusinessAsName field if non-nil, zero value otherwise.

### GetDoingBusinessAsNameOk

`func (o *InquiredMerchantMerchant) GetDoingBusinessAsNameOk() (*string, bool)`

GetDoingBusinessAsNameOk returns a tuple with the DoingBusinessAsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoingBusinessAsName

`func (o *InquiredMerchantMerchant) SetDoingBusinessAsName(v string)`

SetDoingBusinessAsName sets DoingBusinessAsName field to given value.

### HasDoingBusinessAsName

`func (o *InquiredMerchantMerchant) HasDoingBusinessAsName() bool`

HasDoingBusinessAsName returns a boolean if a field has been set.

### GetAddedByAcquirerID

`func (o *InquiredMerchantMerchant) GetAddedByAcquirerID() string`

GetAddedByAcquirerID returns the AddedByAcquirerID field if non-nil, zero value otherwise.

### GetAddedByAcquirerIDOk

`func (o *InquiredMerchantMerchant) GetAddedByAcquirerIDOk() (*string, bool)`

GetAddedByAcquirerIDOk returns a tuple with the AddedByAcquirerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedByAcquirerID

`func (o *InquiredMerchantMerchant) SetAddedByAcquirerID(v string)`

SetAddedByAcquirerID sets AddedByAcquirerID field to given value.

### HasAddedByAcquirerID

`func (o *InquiredMerchantMerchant) HasAddedByAcquirerID() bool`

HasAddedByAcquirerID returns a boolean if a field has been set.

### GetAddedOnDate

`func (o *InquiredMerchantMerchant) GetAddedOnDate() string`

GetAddedOnDate returns the AddedOnDate field if non-nil, zero value otherwise.

### GetAddedOnDateOk

`func (o *InquiredMerchantMerchant) GetAddedOnDateOk() (*string, bool)`

GetAddedOnDateOk returns a tuple with the AddedOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOnDate

`func (o *InquiredMerchantMerchant) SetAddedOnDate(v string)`

SetAddedOnDate sets AddedOnDate field to given value.

### HasAddedOnDate

`func (o *InquiredMerchantMerchant) HasAddedOnDate() bool`

HasAddedOnDate returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *InquiredMerchantMerchant) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InquiredMerchantMerchant) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InquiredMerchantMerchant) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *InquiredMerchantMerchant) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAltPhoneNumber

`func (o *InquiredMerchantMerchant) GetAltPhoneNumber() string`

GetAltPhoneNumber returns the AltPhoneNumber field if non-nil, zero value otherwise.

### GetAltPhoneNumberOk

`func (o *InquiredMerchantMerchant) GetAltPhoneNumberOk() (*string, bool)`

GetAltPhoneNumberOk returns a tuple with the AltPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltPhoneNumber

`func (o *InquiredMerchantMerchant) SetAltPhoneNumber(v string)`

SetAltPhoneNumber sets AltPhoneNumber field to given value.

### HasAltPhoneNumber

`func (o *InquiredMerchantMerchant) HasAltPhoneNumber() bool`

HasAltPhoneNumber returns a boolean if a field has been set.

### GetAddress

`func (o *InquiredMerchantMerchant) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InquiredMerchantMerchant) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InquiredMerchantMerchant) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InquiredMerchantMerchant) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCountrySubdivisionTaxId

`func (o *InquiredMerchantMerchant) GetCountrySubdivisionTaxId() string`

GetCountrySubdivisionTaxId returns the CountrySubdivisionTaxId field if non-nil, zero value otherwise.

### GetCountrySubdivisionTaxIdOk

`func (o *InquiredMerchantMerchant) GetCountrySubdivisionTaxIdOk() (*string, bool)`

GetCountrySubdivisionTaxIdOk returns a tuple with the CountrySubdivisionTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivisionTaxId

`func (o *InquiredMerchantMerchant) SetCountrySubdivisionTaxId(v string)`

SetCountrySubdivisionTaxId sets CountrySubdivisionTaxId field to given value.

### HasCountrySubdivisionTaxId

`func (o *InquiredMerchantMerchant) HasCountrySubdivisionTaxId() bool`

HasCountrySubdivisionTaxId returns a boolean if a field has been set.

### GetNationalTaxId

`func (o *InquiredMerchantMerchant) GetNationalTaxId() string`

GetNationalTaxId returns the NationalTaxId field if non-nil, zero value otherwise.

### GetNationalTaxIdOk

`func (o *InquiredMerchantMerchant) GetNationalTaxIdOk() (*string, bool)`

GetNationalTaxIdOk returns a tuple with the NationalTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalTaxId

`func (o *InquiredMerchantMerchant) SetNationalTaxId(v string)`

SetNationalTaxId sets NationalTaxId field to given value.

### HasNationalTaxId

`func (o *InquiredMerchantMerchant) HasNationalTaxId() bool`

HasNationalTaxId returns a boolean if a field has been set.

### GetServiceProvLegal

`func (o *InquiredMerchantMerchant) GetServiceProvLegal() string`

GetServiceProvLegal returns the ServiceProvLegal field if non-nil, zero value otherwise.

### GetServiceProvLegalOk

`func (o *InquiredMerchantMerchant) GetServiceProvLegalOk() (*string, bool)`

GetServiceProvLegalOk returns a tuple with the ServiceProvLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvLegal

`func (o *InquiredMerchantMerchant) SetServiceProvLegal(v string)`

SetServiceProvLegal sets ServiceProvLegal field to given value.

### HasServiceProvLegal

`func (o *InquiredMerchantMerchant) HasServiceProvLegal() bool`

HasServiceProvLegal returns a boolean if a field has been set.

### GetServiceProvDBA

`func (o *InquiredMerchantMerchant) GetServiceProvDBA() string`

GetServiceProvDBA returns the ServiceProvDBA field if non-nil, zero value otherwise.

### GetServiceProvDBAOk

`func (o *InquiredMerchantMerchant) GetServiceProvDBAOk() (*string, bool)`

GetServiceProvDBAOk returns a tuple with the ServiceProvDBA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvDBA

`func (o *InquiredMerchantMerchant) SetServiceProvDBA(v string)`

SetServiceProvDBA sets ServiceProvDBA field to given value.

### HasServiceProvDBA

`func (o *InquiredMerchantMerchant) HasServiceProvDBA() bool`

HasServiceProvDBA returns a boolean if a field has been set.

### GetTerminationReasonCode

`func (o *InquiredMerchantMerchant) GetTerminationReasonCode() string`

GetTerminationReasonCode returns the TerminationReasonCode field if non-nil, zero value otherwise.

### GetTerminationReasonCodeOk

`func (o *InquiredMerchantMerchant) GetTerminationReasonCodeOk() (*string, bool)`

GetTerminationReasonCodeOk returns a tuple with the TerminationReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationReasonCode

`func (o *InquiredMerchantMerchant) SetTerminationReasonCode(v string)`

SetTerminationReasonCode sets TerminationReasonCode field to given value.

### HasTerminationReasonCode

`func (o *InquiredMerchantMerchant) HasTerminationReasonCode() bool`

HasTerminationReasonCode returns a boolean if a field has been set.

### GetUrlGroup

`func (o *InquiredMerchantMerchant) GetUrlGroup() []UrlGroup`

GetUrlGroup returns the UrlGroup field if non-nil, zero value otherwise.

### GetUrlGroupOk

`func (o *InquiredMerchantMerchant) GetUrlGroupOk() (*[]UrlGroup, bool)`

GetUrlGroupOk returns a tuple with the UrlGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlGroup

`func (o *InquiredMerchantMerchant) SetUrlGroup(v []UrlGroup)`

SetUrlGroup sets UrlGroup field to given value.

### HasUrlGroup

`func (o *InquiredMerchantMerchant) HasUrlGroup() bool`

HasUrlGroup returns a boolean if a field has been set.

### GetUrl

`func (o *InquiredMerchantMerchant) GetUrl() []Url`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InquiredMerchantMerchant) GetUrlOk() (*[]Url, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InquiredMerchantMerchant) SetUrl(v []Url)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InquiredMerchantMerchant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPrincipal

`func (o *InquiredMerchantMerchant) GetPrincipal() []Principal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *InquiredMerchantMerchant) GetPrincipalOk() (*[]Principal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *InquiredMerchantMerchant) SetPrincipal(v []Principal)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *InquiredMerchantMerchant) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetSearchCriteria

`func (o *InquiredMerchantMerchant) GetSearchCriteria() SearchCriteria`

GetSearchCriteria returns the SearchCriteria field if non-nil, zero value otherwise.

### GetSearchCriteriaOk

`func (o *InquiredMerchantMerchant) GetSearchCriteriaOk() (*SearchCriteria, bool)`

GetSearchCriteriaOk returns a tuple with the SearchCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCriteria

`func (o *InquiredMerchantMerchant) SetSearchCriteria(v SearchCriteria)`

SetSearchCriteria sets SearchCriteria field to given value.

### HasSearchCriteria

`func (o *InquiredMerchantMerchant) HasSearchCriteria() bool`

HasSearchCriteria returns a boolean if a field has been set.

### GetMerchantMatch

`func (o *InquiredMerchantMerchant) GetMerchantMatch() MerchantMatch`

GetMerchantMatch returns the MerchantMatch field if non-nil, zero value otherwise.

### GetMerchantMatchOk

`func (o *InquiredMerchantMerchant) GetMerchantMatchOk() (*MerchantMatch, bool)`

GetMerchantMatchOk returns a tuple with the MerchantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMatch

`func (o *InquiredMerchantMerchant) SetMerchantMatch(v MerchantMatch)`

SetMerchantMatch sets MerchantMatch field to given value.

### HasMerchantMatch

`func (o *InquiredMerchantMerchant) HasMerchantMatch() bool`

HasMerchantMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


