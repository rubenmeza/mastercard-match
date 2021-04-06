# AddMerchantMerchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Business which has been terminated. | 
**DoingBusinessAsName** | Pointer to **string** | The name used by a merchant that could be different from the legal name of the business. Such as Bait R Us instead of the legal name, The Bait Shop | [optional] 
**MerchantId** | **string** | The identifier assigned to a merchant by an Acquirer. An Acquirer Id and Merchant Id combination must be unique and Merchant Id should be less than 15 characters. | 
**MerchantCategory** | **string** | A classification code used in authorization, clearing, and other transactions or reports to identify the type of merchant. | 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**PhoneNumber** | **string** | The Business or Merchant&#39;s phone number. | 
**AltPhoneNumber** | Pointer to **string** | The Business or Merchant&#39;s alternate phone number. | [optional] 
**NationalTaxId** | Pointer to **string** | The National tax ID or business registration number. Return value will be hidden. | [optional] 
**CountrySubdivisionTaxId** | Pointer to **string** | The Merchant&#39;s state tax ID; for the U.S region only. Return value will be hidden. | [optional] 
**CATFlag** | **string** | Cardholder-activated terminal indicator. | 
**DateOpened** | **string** | Date the merchant entered into agreement with the acquirer | 
**DateClosed** | **string** | Date the agreement was terminated with the merchant | 
**ServiceProvLegal** | Pointer to **string** | The name of the service provider associated with the merchant listed in the MATCH. | [optional] 
**ServiceProvDBA** | Pointer to **string** | The name of the service provider associated with the merchant listed in the MATCH. | [optional] 
**Url** | Pointer to **[]string** | Website address of the merchant. A request may include multiple URLs . The total cumulative size of the URLs cannot exceed 20000 bytes. | [optional] 
**Principal** | Pointer to [**[]Principal**](Principal.md) |  | [optional] 
**ReasonCode** | **string** | A two digit numeric code indication why a particular merchant was terminated.  01   Account Data Compromise, 02   Common Points of Purchase, 03   Laundering, 04   Excessive Chargebacks, 05   Excessive Fraud, 06   Reserved for Future Use, 07   Fraud Conviction, 08   MasterCard Questionable Merchant Audit Program, 09   Bankruptcy/Liquidation/Insolvency, 10   Violation of MasterCard Standards, 11   Merchant collusion, 12   PCI Data Security Standard, Noncompliance, 13   Illegal Transactions, 14   Identity Theft | 
**Comments** | Pointer to **string** | Brief comments on why Merchant is added | [optional] 

## Methods

### NewAddMerchantMerchant

`func NewAddMerchantMerchant(name string, merchantId string, merchantCategory string, phoneNumber string, cATFlag string, dateOpened string, dateClosed string, reasonCode string, ) *AddMerchantMerchant`

NewAddMerchantMerchant instantiates a new AddMerchantMerchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMerchantMerchantWithDefaults

`func NewAddMerchantMerchantWithDefaults() *AddMerchantMerchant`

NewAddMerchantMerchantWithDefaults instantiates a new AddMerchantMerchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddMerchantMerchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddMerchantMerchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddMerchantMerchant) SetName(v string)`

SetName sets Name field to given value.


### GetDoingBusinessAsName

`func (o *AddMerchantMerchant) GetDoingBusinessAsName() string`

GetDoingBusinessAsName returns the DoingBusinessAsName field if non-nil, zero value otherwise.

### GetDoingBusinessAsNameOk

`func (o *AddMerchantMerchant) GetDoingBusinessAsNameOk() (*string, bool)`

GetDoingBusinessAsNameOk returns a tuple with the DoingBusinessAsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoingBusinessAsName

`func (o *AddMerchantMerchant) SetDoingBusinessAsName(v string)`

SetDoingBusinessAsName sets DoingBusinessAsName field to given value.

### HasDoingBusinessAsName

`func (o *AddMerchantMerchant) HasDoingBusinessAsName() bool`

HasDoingBusinessAsName returns a boolean if a field has been set.

### GetMerchantId

`func (o *AddMerchantMerchant) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *AddMerchantMerchant) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *AddMerchantMerchant) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantCategory

`func (o *AddMerchantMerchant) GetMerchantCategory() string`

GetMerchantCategory returns the MerchantCategory field if non-nil, zero value otherwise.

### GetMerchantCategoryOk

`func (o *AddMerchantMerchant) GetMerchantCategoryOk() (*string, bool)`

GetMerchantCategoryOk returns a tuple with the MerchantCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategory

`func (o *AddMerchantMerchant) SetMerchantCategory(v string)`

SetMerchantCategory sets MerchantCategory field to given value.


### GetAddress

`func (o *AddMerchantMerchant) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddMerchantMerchant) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddMerchantMerchant) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddMerchantMerchant) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *AddMerchantMerchant) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AddMerchantMerchant) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AddMerchantMerchant) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetAltPhoneNumber

`func (o *AddMerchantMerchant) GetAltPhoneNumber() string`

GetAltPhoneNumber returns the AltPhoneNumber field if non-nil, zero value otherwise.

### GetAltPhoneNumberOk

`func (o *AddMerchantMerchant) GetAltPhoneNumberOk() (*string, bool)`

GetAltPhoneNumberOk returns a tuple with the AltPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltPhoneNumber

`func (o *AddMerchantMerchant) SetAltPhoneNumber(v string)`

SetAltPhoneNumber sets AltPhoneNumber field to given value.

### HasAltPhoneNumber

`func (o *AddMerchantMerchant) HasAltPhoneNumber() bool`

HasAltPhoneNumber returns a boolean if a field has been set.

### GetNationalTaxId

`func (o *AddMerchantMerchant) GetNationalTaxId() string`

GetNationalTaxId returns the NationalTaxId field if non-nil, zero value otherwise.

### GetNationalTaxIdOk

`func (o *AddMerchantMerchant) GetNationalTaxIdOk() (*string, bool)`

GetNationalTaxIdOk returns a tuple with the NationalTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalTaxId

`func (o *AddMerchantMerchant) SetNationalTaxId(v string)`

SetNationalTaxId sets NationalTaxId field to given value.

### HasNationalTaxId

`func (o *AddMerchantMerchant) HasNationalTaxId() bool`

HasNationalTaxId returns a boolean if a field has been set.

### GetCountrySubdivisionTaxId

`func (o *AddMerchantMerchant) GetCountrySubdivisionTaxId() string`

GetCountrySubdivisionTaxId returns the CountrySubdivisionTaxId field if non-nil, zero value otherwise.

### GetCountrySubdivisionTaxIdOk

`func (o *AddMerchantMerchant) GetCountrySubdivisionTaxIdOk() (*string, bool)`

GetCountrySubdivisionTaxIdOk returns a tuple with the CountrySubdivisionTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivisionTaxId

`func (o *AddMerchantMerchant) SetCountrySubdivisionTaxId(v string)`

SetCountrySubdivisionTaxId sets CountrySubdivisionTaxId field to given value.

### HasCountrySubdivisionTaxId

`func (o *AddMerchantMerchant) HasCountrySubdivisionTaxId() bool`

HasCountrySubdivisionTaxId returns a boolean if a field has been set.

### GetCATFlag

`func (o *AddMerchantMerchant) GetCATFlag() string`

GetCATFlag returns the CATFlag field if non-nil, zero value otherwise.

### GetCATFlagOk

`func (o *AddMerchantMerchant) GetCATFlagOk() (*string, bool)`

GetCATFlagOk returns a tuple with the CATFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCATFlag

`func (o *AddMerchantMerchant) SetCATFlag(v string)`

SetCATFlag sets CATFlag field to given value.


### GetDateOpened

`func (o *AddMerchantMerchant) GetDateOpened() string`

GetDateOpened returns the DateOpened field if non-nil, zero value otherwise.

### GetDateOpenedOk

`func (o *AddMerchantMerchant) GetDateOpenedOk() (*string, bool)`

GetDateOpenedOk returns a tuple with the DateOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOpened

`func (o *AddMerchantMerchant) SetDateOpened(v string)`

SetDateOpened sets DateOpened field to given value.


### GetDateClosed

`func (o *AddMerchantMerchant) GetDateClosed() string`

GetDateClosed returns the DateClosed field if non-nil, zero value otherwise.

### GetDateClosedOk

`func (o *AddMerchantMerchant) GetDateClosedOk() (*string, bool)`

GetDateClosedOk returns a tuple with the DateClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClosed

`func (o *AddMerchantMerchant) SetDateClosed(v string)`

SetDateClosed sets DateClosed field to given value.


### GetServiceProvLegal

`func (o *AddMerchantMerchant) GetServiceProvLegal() string`

GetServiceProvLegal returns the ServiceProvLegal field if non-nil, zero value otherwise.

### GetServiceProvLegalOk

`func (o *AddMerchantMerchant) GetServiceProvLegalOk() (*string, bool)`

GetServiceProvLegalOk returns a tuple with the ServiceProvLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvLegal

`func (o *AddMerchantMerchant) SetServiceProvLegal(v string)`

SetServiceProvLegal sets ServiceProvLegal field to given value.

### HasServiceProvLegal

`func (o *AddMerchantMerchant) HasServiceProvLegal() bool`

HasServiceProvLegal returns a boolean if a field has been set.

### GetServiceProvDBA

`func (o *AddMerchantMerchant) GetServiceProvDBA() string`

GetServiceProvDBA returns the ServiceProvDBA field if non-nil, zero value otherwise.

### GetServiceProvDBAOk

`func (o *AddMerchantMerchant) GetServiceProvDBAOk() (*string, bool)`

GetServiceProvDBAOk returns a tuple with the ServiceProvDBA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvDBA

`func (o *AddMerchantMerchant) SetServiceProvDBA(v string)`

SetServiceProvDBA sets ServiceProvDBA field to given value.

### HasServiceProvDBA

`func (o *AddMerchantMerchant) HasServiceProvDBA() bool`

HasServiceProvDBA returns a boolean if a field has been set.

### GetUrl

`func (o *AddMerchantMerchant) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddMerchantMerchant) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddMerchantMerchant) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AddMerchantMerchant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPrincipal

`func (o *AddMerchantMerchant) GetPrincipal() []Principal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *AddMerchantMerchant) GetPrincipalOk() (*[]Principal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *AddMerchantMerchant) SetPrincipal(v []Principal)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *AddMerchantMerchant) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetReasonCode

`func (o *AddMerchantMerchant) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *AddMerchantMerchant) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *AddMerchantMerchant) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetComments

`func (o *AddMerchantMerchant) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AddMerchantMerchant) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AddMerchantMerchant) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AddMerchantMerchant) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


