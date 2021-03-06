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

// PossibleInquiryMatches struct for PossibleInquiryMatches
type PossibleInquiryMatches struct {
	// The total length of the result set from possible merchant matches of inquiry.
	TotalLength *int32 `json:"TotalLength,omitempty"`
	InquiredMerchant *[]InquiredMerchant `json:"InquiredMerchant,omitempty"`
}

// NewPossibleInquiryMatches instantiates a new PossibleInquiryMatches object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPossibleInquiryMatches() *PossibleInquiryMatches {
	this := PossibleInquiryMatches{}
	return &this
}

// NewPossibleInquiryMatchesWithDefaults instantiates a new PossibleInquiryMatches object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPossibleInquiryMatchesWithDefaults() *PossibleInquiryMatches {
	this := PossibleInquiryMatches{}
	return &this
}

// GetTotalLength returns the TotalLength field value if set, zero value otherwise.
func (o *PossibleInquiryMatches) GetTotalLength() int32 {
	if o == nil || o.TotalLength == nil {
		var ret int32
		return ret
	}
	return *o.TotalLength
}

// GetTotalLengthOk returns a tuple with the TotalLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossibleInquiryMatches) GetTotalLengthOk() (*int32, bool) {
	if o == nil || o.TotalLength == nil {
		return nil, false
	}
	return o.TotalLength, true
}

// HasTotalLength returns a boolean if a field has been set.
func (o *PossibleInquiryMatches) HasTotalLength() bool {
	if o != nil && o.TotalLength != nil {
		return true
	}

	return false
}

// SetTotalLength gets a reference to the given int32 and assigns it to the TotalLength field.
func (o *PossibleInquiryMatches) SetTotalLength(v int32) {
	o.TotalLength = &v
}

// GetInquiredMerchant returns the InquiredMerchant field value if set, zero value otherwise.
func (o *PossibleInquiryMatches) GetInquiredMerchant() []InquiredMerchant {
	if o == nil || o.InquiredMerchant == nil {
		var ret []InquiredMerchant
		return ret
	}
	return *o.InquiredMerchant
}

// GetInquiredMerchantOk returns a tuple with the InquiredMerchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossibleInquiryMatches) GetInquiredMerchantOk() (*[]InquiredMerchant, bool) {
	if o == nil || o.InquiredMerchant == nil {
		return nil, false
	}
	return o.InquiredMerchant, true
}

// HasInquiredMerchant returns a boolean if a field has been set.
func (o *PossibleInquiryMatches) HasInquiredMerchant() bool {
	if o != nil && o.InquiredMerchant != nil {
		return true
	}

	return false
}

// SetInquiredMerchant gets a reference to the given []InquiredMerchant and assigns it to the InquiredMerchant field.
func (o *PossibleInquiryMatches) SetInquiredMerchant(v []InquiredMerchant) {
	o.InquiredMerchant = &v
}

func (o PossibleInquiryMatches) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalLength != nil {
		toSerialize["TotalLength"] = o.TotalLength
	}
	if o.InquiredMerchant != nil {
		toSerialize["InquiredMerchant"] = o.InquiredMerchant
	}
	return json.Marshal(toSerialize)
}

type NullablePossibleInquiryMatches struct {
	value *PossibleInquiryMatches
	isSet bool
}

func (v NullablePossibleInquiryMatches) Get() *PossibleInquiryMatches {
	return v.value
}

func (v *NullablePossibleInquiryMatches) Set(val *PossibleInquiryMatches) {
	v.value = val
	v.isSet = true
}

func (v NullablePossibleInquiryMatches) IsSet() bool {
	return v.isSet
}

func (v *NullablePossibleInquiryMatches) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePossibleInquiryMatches(val *PossibleInquiryMatches) *NullablePossibleInquiryMatches {
	return &NullablePossibleInquiryMatches{value: val, isSet: true}
}

func (v NullablePossibleInquiryMatches) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePossibleInquiryMatches) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


