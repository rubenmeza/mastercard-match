/*
 * MATCH API
 *
 * Helps acquirers identify potentially high-risk merchants before entering to a merchant agreement.
 *
 * API version: 1.0.0
 * Contact: apisupport@mastercard.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package match

import (
	"encoding/json"
)

// Merchant struct for Merchant
type Merchant struct {
	// The name of the business assigned by the principal owner(s)
	Name string `json:"Name"`
	// The name used by a merchant that could be different from the legal name of the business. Such as Bait R Us instead of the legal name, The Bait Shop.
	DoingBusinessAsName *string `json:"DoingBusinessAsName,omitempty"`
	Address *Address `json:"Address,omitempty"`
	// The Business or Merchant's phone number, including the area code. Within the USA, phone numbers have a length of 10, and for outside the US, it can be any length with a maximum of 12 digits. Within the U.S. phone numbers can not start with 0 or 1. If the number is outside the U.S. region; do not include the country code. The phone number must be digits only, with no format characters such as parenthesis or dashes.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The Business or Merchant's alternate phone number, including the area code. Within the USA, phone numbers have a length of 10, and for out outside the US, a max length of 25. Within the U.S. phone numbers can not start with 0 or 1. If the number is outside the U.S. region; do not include the country code. The number must be digits only, with no format characters such as parenthesis or dashes.
	AltPhoneNumber *string `json:"AltPhoneNumber,omitempty"`
	// The Merchant national tax ID, leave blank if not in the U.S region.
	NationalTaxId *string `json:"NationalTaxId,omitempty"`
	// The Merchant Country Subdivision tax ID, leave blank if not in the U.S region.
	CountrySubdivisionTaxId *string `json:"CountrySubdivisionTaxId,omitempty"`
	// The name of the service provider associated with the merchant listed in the MATCH
	ServiceProvLegal *string `json:"ServiceProvLegal,omitempty"`
	// The name of the service provider associated with the merchant listed in the MATCH
	ServiceProvDBA *string `json:"ServiceProvDBA,omitempty"`
	// Website address of the merchant. A request may include multiple URLs . The total cumulative size of the URLs cannot exceed 20000 bytes.
	Url *[]string `json:"Url,omitempty"`
	// The details for the principal owner of the business.  A maximum of 5 principals may be submitted.
	Principal *[]Principal `json:"Principal,omitempty"`
	SearchCriteria *SearchCriteria `json:"SearchCriteria,omitempty"`
	// Date the merchant was added to the MATCH database.
	AddedOnDate *string `json:"AddedOnDate,omitempty"`
	// A two digit numeric code indication why a particular merchant was terminated.  01   Account Data Compromise, 02   Common Points of Purchase, 03   Laundering, 04   Excessive Chargebacks, 05   Excessive Fraud, 06   Reserved for Future Use, 07   Fraud Conviction, 08   MasterCard Questionable Merchant Audit Program, 09   Bankruptcy/Liquidation/Insolvency, 10   Violation of MasterCard Standards, 11   Merchant collusion, 12   PCI Data Security Standard, Noncompliance, 13   Illegal Transactions, 14   Identity Theft
	TerminationReasonCode *string `json:"TerminationReasonCode,omitempty"`
	// The Member ICA that has added the merchant to the MATCH system
	AddedByAcquirerID *string `json:"AddedByAcquirerID,omitempty"`
	UrlGroup *[]UrlGroup `json:"UrlGroup,omitempty"`
}

// NewMerchant instantiates a new Merchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchant(name string) *Merchant {
	this := Merchant{}
	this.Name = name
	return &this
}

// NewMerchantWithDefaults instantiates a new Merchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantWithDefaults() *Merchant {
	this := Merchant{}
	return &this
}

// GetName returns the Name field value
func (o *Merchant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Merchant) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Merchant) SetName(v string) {
	o.Name = v
}

// GetDoingBusinessAsName returns the DoingBusinessAsName field value if set, zero value otherwise.
func (o *Merchant) GetDoingBusinessAsName() string {
	if o == nil || o.DoingBusinessAsName == nil {
		var ret string
		return ret
	}
	return *o.DoingBusinessAsName
}

// GetDoingBusinessAsNameOk returns a tuple with the DoingBusinessAsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetDoingBusinessAsNameOk() (*string, bool) {
	if o == nil || o.DoingBusinessAsName == nil {
		return nil, false
	}
	return o.DoingBusinessAsName, true
}

// HasDoingBusinessAsName returns a boolean if a field has been set.
func (o *Merchant) HasDoingBusinessAsName() bool {
	if o != nil && o.DoingBusinessAsName != nil {
		return true
	}

	return false
}

// SetDoingBusinessAsName gets a reference to the given string and assigns it to the DoingBusinessAsName field.
func (o *Merchant) SetDoingBusinessAsName(v string) {
	o.DoingBusinessAsName = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Merchant) GetAddress() Address {
	if o == nil || o.Address == nil {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetAddressOk() (*Address, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Merchant) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Merchant) SetAddress(v Address) {
	o.Address = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Merchant) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Merchant) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Merchant) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetAltPhoneNumber returns the AltPhoneNumber field value if set, zero value otherwise.
func (o *Merchant) GetAltPhoneNumber() string {
	if o == nil || o.AltPhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.AltPhoneNumber
}

// GetAltPhoneNumberOk returns a tuple with the AltPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetAltPhoneNumberOk() (*string, bool) {
	if o == nil || o.AltPhoneNumber == nil {
		return nil, false
	}
	return o.AltPhoneNumber, true
}

// HasAltPhoneNumber returns a boolean if a field has been set.
func (o *Merchant) HasAltPhoneNumber() bool {
	if o != nil && o.AltPhoneNumber != nil {
		return true
	}

	return false
}

// SetAltPhoneNumber gets a reference to the given string and assigns it to the AltPhoneNumber field.
func (o *Merchant) SetAltPhoneNumber(v string) {
	o.AltPhoneNumber = &v
}

// GetNationalTaxId returns the NationalTaxId field value if set, zero value otherwise.
func (o *Merchant) GetNationalTaxId() string {
	if o == nil || o.NationalTaxId == nil {
		var ret string
		return ret
	}
	return *o.NationalTaxId
}

// GetNationalTaxIdOk returns a tuple with the NationalTaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetNationalTaxIdOk() (*string, bool) {
	if o == nil || o.NationalTaxId == nil {
		return nil, false
	}
	return o.NationalTaxId, true
}

// HasNationalTaxId returns a boolean if a field has been set.
func (o *Merchant) HasNationalTaxId() bool {
	if o != nil && o.NationalTaxId != nil {
		return true
	}

	return false
}

// SetNationalTaxId gets a reference to the given string and assigns it to the NationalTaxId field.
func (o *Merchant) SetNationalTaxId(v string) {
	o.NationalTaxId = &v
}

// GetCountrySubdivisionTaxId returns the CountrySubdivisionTaxId field value if set, zero value otherwise.
func (o *Merchant) GetCountrySubdivisionTaxId() string {
	if o == nil || o.CountrySubdivisionTaxId == nil {
		var ret string
		return ret
	}
	return *o.CountrySubdivisionTaxId
}

// GetCountrySubdivisionTaxIdOk returns a tuple with the CountrySubdivisionTaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetCountrySubdivisionTaxIdOk() (*string, bool) {
	if o == nil || o.CountrySubdivisionTaxId == nil {
		return nil, false
	}
	return o.CountrySubdivisionTaxId, true
}

// HasCountrySubdivisionTaxId returns a boolean if a field has been set.
func (o *Merchant) HasCountrySubdivisionTaxId() bool {
	if o != nil && o.CountrySubdivisionTaxId != nil {
		return true
	}

	return false
}

// SetCountrySubdivisionTaxId gets a reference to the given string and assigns it to the CountrySubdivisionTaxId field.
func (o *Merchant) SetCountrySubdivisionTaxId(v string) {
	o.CountrySubdivisionTaxId = &v
}

// GetServiceProvLegal returns the ServiceProvLegal field value if set, zero value otherwise.
func (o *Merchant) GetServiceProvLegal() string {
	if o == nil || o.ServiceProvLegal == nil {
		var ret string
		return ret
	}
	return *o.ServiceProvLegal
}

// GetServiceProvLegalOk returns a tuple with the ServiceProvLegal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetServiceProvLegalOk() (*string, bool) {
	if o == nil || o.ServiceProvLegal == nil {
		return nil, false
	}
	return o.ServiceProvLegal, true
}

// HasServiceProvLegal returns a boolean if a field has been set.
func (o *Merchant) HasServiceProvLegal() bool {
	if o != nil && o.ServiceProvLegal != nil {
		return true
	}

	return false
}

// SetServiceProvLegal gets a reference to the given string and assigns it to the ServiceProvLegal field.
func (o *Merchant) SetServiceProvLegal(v string) {
	o.ServiceProvLegal = &v
}

// GetServiceProvDBA returns the ServiceProvDBA field value if set, zero value otherwise.
func (o *Merchant) GetServiceProvDBA() string {
	if o == nil || o.ServiceProvDBA == nil {
		var ret string
		return ret
	}
	return *o.ServiceProvDBA
}

// GetServiceProvDBAOk returns a tuple with the ServiceProvDBA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetServiceProvDBAOk() (*string, bool) {
	if o == nil || o.ServiceProvDBA == nil {
		return nil, false
	}
	return o.ServiceProvDBA, true
}

// HasServiceProvDBA returns a boolean if a field has been set.
func (o *Merchant) HasServiceProvDBA() bool {
	if o != nil && o.ServiceProvDBA != nil {
		return true
	}

	return false
}

// SetServiceProvDBA gets a reference to the given string and assigns it to the ServiceProvDBA field.
func (o *Merchant) SetServiceProvDBA(v string) {
	o.ServiceProvDBA = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Merchant) GetUrl() []string {
	if o == nil || o.Url == nil {
		var ret []string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetUrlOk() (*[]string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Merchant) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given []string and assigns it to the Url field.
func (o *Merchant) SetUrl(v []string) {
	o.Url = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *Merchant) GetPrincipal() []Principal {
	if o == nil || o.Principal == nil {
		var ret []Principal
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetPrincipalOk() (*[]Principal, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *Merchant) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given []Principal and assigns it to the Principal field.
func (o *Merchant) SetPrincipal(v []Principal) {
	o.Principal = &v
}

// GetSearchCriteria returns the SearchCriteria field value if set, zero value otherwise.
func (o *Merchant) GetSearchCriteria() SearchCriteria {
	if o == nil || o.SearchCriteria == nil {
		var ret SearchCriteria
		return ret
	}
	return *o.SearchCriteria
}

// GetSearchCriteriaOk returns a tuple with the SearchCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetSearchCriteriaOk() (*SearchCriteria, bool) {
	if o == nil || o.SearchCriteria == nil {
		return nil, false
	}
	return o.SearchCriteria, true
}

// HasSearchCriteria returns a boolean if a field has been set.
func (o *Merchant) HasSearchCriteria() bool {
	if o != nil && o.SearchCriteria != nil {
		return true
	}

	return false
}

// SetSearchCriteria gets a reference to the given SearchCriteria and assigns it to the SearchCriteria field.
func (o *Merchant) SetSearchCriteria(v SearchCriteria) {
	o.SearchCriteria = &v
}

// GetAddedOnDate returns the AddedOnDate field value if set, zero value otherwise.
func (o *Merchant) GetAddedOnDate() string {
	if o == nil || o.AddedOnDate == nil {
		var ret string
		return ret
	}
	return *o.AddedOnDate
}

// GetAddedOnDateOk returns a tuple with the AddedOnDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetAddedOnDateOk() (*string, bool) {
	if o == nil || o.AddedOnDate == nil {
		return nil, false
	}
	return o.AddedOnDate, true
}

// HasAddedOnDate returns a boolean if a field has been set.
func (o *Merchant) HasAddedOnDate() bool {
	if o != nil && o.AddedOnDate != nil {
		return true
	}

	return false
}

// SetAddedOnDate gets a reference to the given string and assigns it to the AddedOnDate field.
func (o *Merchant) SetAddedOnDate(v string) {
	o.AddedOnDate = &v
}

// GetTerminationReasonCode returns the TerminationReasonCode field value if set, zero value otherwise.
func (o *Merchant) GetTerminationReasonCode() string {
	if o == nil || o.TerminationReasonCode == nil {
		var ret string
		return ret
	}
	return *o.TerminationReasonCode
}

// GetTerminationReasonCodeOk returns a tuple with the TerminationReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetTerminationReasonCodeOk() (*string, bool) {
	if o == nil || o.TerminationReasonCode == nil {
		return nil, false
	}
	return o.TerminationReasonCode, true
}

// HasTerminationReasonCode returns a boolean if a field has been set.
func (o *Merchant) HasTerminationReasonCode() bool {
	if o != nil && o.TerminationReasonCode != nil {
		return true
	}

	return false
}

// SetTerminationReasonCode gets a reference to the given string and assigns it to the TerminationReasonCode field.
func (o *Merchant) SetTerminationReasonCode(v string) {
	o.TerminationReasonCode = &v
}

// GetAddedByAcquirerID returns the AddedByAcquirerID field value if set, zero value otherwise.
func (o *Merchant) GetAddedByAcquirerID() string {
	if o == nil || o.AddedByAcquirerID == nil {
		var ret string
		return ret
	}
	return *o.AddedByAcquirerID
}

// GetAddedByAcquirerIDOk returns a tuple with the AddedByAcquirerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetAddedByAcquirerIDOk() (*string, bool) {
	if o == nil || o.AddedByAcquirerID == nil {
		return nil, false
	}
	return o.AddedByAcquirerID, true
}

// HasAddedByAcquirerID returns a boolean if a field has been set.
func (o *Merchant) HasAddedByAcquirerID() bool {
	if o != nil && o.AddedByAcquirerID != nil {
		return true
	}

	return false
}

// SetAddedByAcquirerID gets a reference to the given string and assigns it to the AddedByAcquirerID field.
func (o *Merchant) SetAddedByAcquirerID(v string) {
	o.AddedByAcquirerID = &v
}

// GetUrlGroup returns the UrlGroup field value if set, zero value otherwise.
func (o *Merchant) GetUrlGroup() []UrlGroup {
	if o == nil || o.UrlGroup == nil {
		var ret []UrlGroup
		return ret
	}
	return *o.UrlGroup
}

// GetUrlGroupOk returns a tuple with the UrlGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetUrlGroupOk() (*[]UrlGroup, bool) {
	if o == nil || o.UrlGroup == nil {
		return nil, false
	}
	return o.UrlGroup, true
}

// HasUrlGroup returns a boolean if a field has been set.
func (o *Merchant) HasUrlGroup() bool {
	if o != nil && o.UrlGroup != nil {
		return true
	}

	return false
}

// SetUrlGroup gets a reference to the given []UrlGroup and assigns it to the UrlGroup field.
func (o *Merchant) SetUrlGroup(v []UrlGroup) {
	o.UrlGroup = &v
}

func (o Merchant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Name"] = o.Name
	}
	if o.DoingBusinessAsName != nil {
		toSerialize["DoingBusinessAsName"] = o.DoingBusinessAsName
	}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.PhoneNumber != nil {
		toSerialize["PhoneNumber"] = o.PhoneNumber
	}
	if o.AltPhoneNumber != nil {
		toSerialize["AltPhoneNumber"] = o.AltPhoneNumber
	}
	if o.NationalTaxId != nil {
		toSerialize["NationalTaxId"] = o.NationalTaxId
	}
	if o.CountrySubdivisionTaxId != nil {
		toSerialize["CountrySubdivisionTaxId"] = o.CountrySubdivisionTaxId
	}
	if o.ServiceProvLegal != nil {
		toSerialize["ServiceProvLegal"] = o.ServiceProvLegal
	}
	if o.ServiceProvDBA != nil {
		toSerialize["ServiceProvDBA"] = o.ServiceProvDBA
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}
	if o.Principal != nil {
		toSerialize["Principal"] = o.Principal
	}
	if o.SearchCriteria != nil {
		toSerialize["SearchCriteria"] = o.SearchCriteria
	}
	if o.AddedOnDate != nil {
		toSerialize["AddedOnDate"] = o.AddedOnDate
	}
	if o.TerminationReasonCode != nil {
		toSerialize["TerminationReasonCode"] = o.TerminationReasonCode
	}
	if o.AddedByAcquirerID != nil {
		toSerialize["AddedByAcquirerID"] = o.AddedByAcquirerID
	}
	if o.UrlGroup != nil {
		toSerialize["UrlGroup"] = o.UrlGroup
	}
	return json.Marshal(toSerialize)
}

type NullableMerchant struct {
	value *Merchant
	isSet bool
}

func (v NullableMerchant) Get() *Merchant {
	return v.value
}

func (v *NullableMerchant) Set(val *Merchant) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchant(val *Merchant) *NullableMerchant {
	return &NullableMerchant{value: val, isSet: true}
}

func (v NullableMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


