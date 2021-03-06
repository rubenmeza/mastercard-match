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

// AddMerchantRequest struct for AddMerchantRequest
type AddMerchantRequest struct {
	// The ID of the acquirer.
	AcquirerId string `json:"AcquirerId"`
	Merchant AddMerchantMerchant `json:"Merchant"`
}

// NewAddMerchantRequest instantiates a new AddMerchantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMerchantRequest(acquirerId string, merchant AddMerchantMerchant) *AddMerchantRequest {
	this := AddMerchantRequest{}
	this.AcquirerId = acquirerId
	this.Merchant = merchant
	return &this
}

// NewAddMerchantRequestWithDefaults instantiates a new AddMerchantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMerchantRequestWithDefaults() *AddMerchantRequest {
	this := AddMerchantRequest{}
	return &this
}

// GetAcquirerId returns the AcquirerId field value
func (o *AddMerchantRequest) GetAcquirerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcquirerId
}

// GetAcquirerIdOk returns a tuple with the AcquirerId field value
// and a boolean to check if the value has been set.
func (o *AddMerchantRequest) GetAcquirerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AcquirerId, true
}

// SetAcquirerId sets field value
func (o *AddMerchantRequest) SetAcquirerId(v string) {
	o.AcquirerId = v
}

// GetMerchant returns the Merchant field value
func (o *AddMerchantRequest) GetMerchant() AddMerchantMerchant {
	if o == nil {
		var ret AddMerchantMerchant
		return ret
	}

	return o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value
// and a boolean to check if the value has been set.
func (o *AddMerchantRequest) GetMerchantOk() (*AddMerchantMerchant, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Merchant, true
}

// SetMerchant sets field value
func (o *AddMerchantRequest) SetMerchant(v AddMerchantMerchant) {
	o.Merchant = v
}

func (o AddMerchantRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AcquirerId"] = o.AcquirerId
	}
	if true {
		toSerialize["Merchant"] = o.Merchant
	}
	return json.Marshal(toSerialize)
}

type NullableAddMerchantRequest struct {
	value *AddMerchantRequest
	isSet bool
}

func (v NullableAddMerchantRequest) Get() *AddMerchantRequest {
	return v.value
}

func (v *NullableAddMerchantRequest) Set(val *AddMerchantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMerchantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMerchantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMerchantRequest(val *AddMerchantRequest) *NullableAddMerchantRequest {
	return &NullableAddMerchantRequest{value: val, isSet: true}
}

func (v NullableAddMerchantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMerchantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


