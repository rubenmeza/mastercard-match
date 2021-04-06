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

// TerminationInquiryRequest struct for TerminationInquiryRequest
type TerminationInquiryRequest struct {
	// The Member ICA number which is used to validate that the application has permission to hit the MATCH database. This number must be obtained from a participating MATCH acquiring bank or entity before a developer can access the MATCH resource.
	AcquirerId string `json:"AcquirerId"`
	Merchant *Merchant `json:"Merchant,omitempty"`
}

// NewTerminationInquiryRequest instantiates a new TerminationInquiryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminationInquiryRequest(acquirerId string) *TerminationInquiryRequest {
	this := TerminationInquiryRequest{}
	this.AcquirerId = acquirerId
	return &this
}

// NewTerminationInquiryRequestWithDefaults instantiates a new TerminationInquiryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminationInquiryRequestWithDefaults() *TerminationInquiryRequest {
	this := TerminationInquiryRequest{}
	return &this
}

// GetAcquirerId returns the AcquirerId field value
func (o *TerminationInquiryRequest) GetAcquirerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcquirerId
}

// GetAcquirerIdOk returns a tuple with the AcquirerId field value
// and a boolean to check if the value has been set.
func (o *TerminationInquiryRequest) GetAcquirerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AcquirerId, true
}

// SetAcquirerId sets field value
func (o *TerminationInquiryRequest) SetAcquirerId(v string) {
	o.AcquirerId = v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *TerminationInquiryRequest) GetMerchant() Merchant {
	if o == nil || o.Merchant == nil {
		var ret Merchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminationInquiryRequest) GetMerchantOk() (*Merchant, bool) {
	if o == nil || o.Merchant == nil {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *TerminationInquiryRequest) HasMerchant() bool {
	if o != nil && o.Merchant != nil {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given Merchant and assigns it to the Merchant field.
func (o *TerminationInquiryRequest) SetMerchant(v Merchant) {
	o.Merchant = &v
}

func (o TerminationInquiryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AcquirerId"] = o.AcquirerId
	}
	if o.Merchant != nil {
		toSerialize["Merchant"] = o.Merchant
	}
	return json.Marshal(toSerialize)
}

type NullableTerminationInquiryRequest struct {
	value *TerminationInquiryRequest
	isSet bool
}

func (v NullableTerminationInquiryRequest) Get() *TerminationInquiryRequest {
	return v.value
}

func (v *NullableTerminationInquiryRequest) Set(val *TerminationInquiryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminationInquiryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminationInquiryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminationInquiryRequest(val *TerminationInquiryRequest) *NullableTerminationInquiryRequest {
	return &NullableTerminationInquiryRequest{value: val, isSet: true}
}

func (v NullableTerminationInquiryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminationInquiryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


