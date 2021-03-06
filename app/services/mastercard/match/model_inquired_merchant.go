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

// InquiredMerchant struct for InquiredMerchant
type InquiredMerchant struct {
	Merchant *InquiredMerchantMerchant `json:"Merchant,omitempty"`
}

// NewInquiredMerchant instantiates a new InquiredMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInquiredMerchant() *InquiredMerchant {
	this := InquiredMerchant{}
	return &this
}

// NewInquiredMerchantWithDefaults instantiates a new InquiredMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInquiredMerchantWithDefaults() *InquiredMerchant {
	this := InquiredMerchant{}
	return &this
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *InquiredMerchant) GetMerchant() InquiredMerchantMerchant {
	if o == nil || o.Merchant == nil {
		var ret InquiredMerchantMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InquiredMerchant) GetMerchantOk() (*InquiredMerchantMerchant, bool) {
	if o == nil || o.Merchant == nil {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *InquiredMerchant) HasMerchant() bool {
	if o != nil && o.Merchant != nil {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given InquiredMerchantMerchant and assigns it to the Merchant field.
func (o *InquiredMerchant) SetMerchant(v InquiredMerchantMerchant) {
	o.Merchant = &v
}

func (o InquiredMerchant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Merchant != nil {
		toSerialize["Merchant"] = o.Merchant
	}
	return json.Marshal(toSerialize)
}

type NullableInquiredMerchant struct {
	value *InquiredMerchant
	isSet bool
}

func (v NullableInquiredMerchant) Get() *InquiredMerchant {
	return v.value
}

func (v *NullableInquiredMerchant) Set(val *InquiredMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableInquiredMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableInquiredMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInquiredMerchant(val *InquiredMerchant) *NullableInquiredMerchant {
	return &NullableInquiredMerchant{value: val, isSet: true}
}

func (v NullableInquiredMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInquiredMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


