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

// RetroInquiryResponse struct for RetroInquiryResponse
type RetroInquiryResponse struct {
	PossibleMerchantMatches *[]PossibleMerchantMatches `json:"PossibleMerchantMatches,omitempty"`
}

// NewRetroInquiryResponse instantiates a new RetroInquiryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetroInquiryResponse() *RetroInquiryResponse {
	this := RetroInquiryResponse{}
	return &this
}

// NewRetroInquiryResponseWithDefaults instantiates a new RetroInquiryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetroInquiryResponseWithDefaults() *RetroInquiryResponse {
	this := RetroInquiryResponse{}
	return &this
}

// GetPossibleMerchantMatches returns the PossibleMerchantMatches field value if set, zero value otherwise.
func (o *RetroInquiryResponse) GetPossibleMerchantMatches() []PossibleMerchantMatches {
	if o == nil || o.PossibleMerchantMatches == nil {
		var ret []PossibleMerchantMatches
		return ret
	}
	return *o.PossibleMerchantMatches
}

// GetPossibleMerchantMatchesOk returns a tuple with the PossibleMerchantMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetroInquiryResponse) GetPossibleMerchantMatchesOk() (*[]PossibleMerchantMatches, bool) {
	if o == nil || o.PossibleMerchantMatches == nil {
		return nil, false
	}
	return o.PossibleMerchantMatches, true
}

// HasPossibleMerchantMatches returns a boolean if a field has been set.
func (o *RetroInquiryResponse) HasPossibleMerchantMatches() bool {
	if o != nil && o.PossibleMerchantMatches != nil {
		return true
	}

	return false
}

// SetPossibleMerchantMatches gets a reference to the given []PossibleMerchantMatches and assigns it to the PossibleMerchantMatches field.
func (o *RetroInquiryResponse) SetPossibleMerchantMatches(v []PossibleMerchantMatches) {
	o.PossibleMerchantMatches = &v
}

func (o RetroInquiryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PossibleMerchantMatches != nil {
		toSerialize["PossibleMerchantMatches"] = o.PossibleMerchantMatches
	}
	return json.Marshal(toSerialize)
}

type NullableRetroInquiryResponse struct {
	value *RetroInquiryResponse
	isSet bool
}

func (v NullableRetroInquiryResponse) Get() *RetroInquiryResponse {
	return v.value
}

func (v *NullableRetroInquiryResponse) Set(val *RetroInquiryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRetroInquiryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRetroInquiryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetroInquiryResponse(val *RetroInquiryResponse) *NullableRetroInquiryResponse {
	return &NullableRetroInquiryResponse{value: val, isSet: true}
}

func (v NullableRetroInquiryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetroInquiryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

