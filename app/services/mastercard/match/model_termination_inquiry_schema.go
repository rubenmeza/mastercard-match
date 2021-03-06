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

// TerminationInquirySchema struct for TerminationInquirySchema
type TerminationInquirySchema struct {
	TerminationInquiry *TerminationInquiry `json:"TerminationInquiry,omitempty"`
}

// NewTerminationInquirySchema instantiates a new TerminationInquirySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminationInquirySchema() *TerminationInquirySchema {
	this := TerminationInquirySchema{}
	return &this
}

// NewTerminationInquirySchemaWithDefaults instantiates a new TerminationInquirySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminationInquirySchemaWithDefaults() *TerminationInquirySchema {
	this := TerminationInquirySchema{}
	return &this
}

// GetTerminationInquiry returns the TerminationInquiry field value if set, zero value otherwise.
func (o *TerminationInquirySchema) GetTerminationInquiry() TerminationInquiry {
	if o == nil || o.TerminationInquiry == nil {
		var ret TerminationInquiry
		return ret
	}
	return *o.TerminationInquiry
}

// GetTerminationInquiryOk returns a tuple with the TerminationInquiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminationInquirySchema) GetTerminationInquiryOk() (*TerminationInquiry, bool) {
	if o == nil || o.TerminationInquiry == nil {
		return nil, false
	}
	return o.TerminationInquiry, true
}

// HasTerminationInquiry returns a boolean if a field has been set.
func (o *TerminationInquirySchema) HasTerminationInquiry() bool {
	if o != nil && o.TerminationInquiry != nil {
		return true
	}

	return false
}

// SetTerminationInquiry gets a reference to the given TerminationInquiry and assigns it to the TerminationInquiry field.
func (o *TerminationInquirySchema) SetTerminationInquiry(v TerminationInquiry) {
	o.TerminationInquiry = &v
}

func (o TerminationInquirySchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TerminationInquiry != nil {
		toSerialize["TerminationInquiry"] = o.TerminationInquiry
	}
	return json.Marshal(toSerialize)
}

type NullableTerminationInquirySchema struct {
	value *TerminationInquirySchema
	isSet bool
}

func (v NullableTerminationInquirySchema) Get() *TerminationInquirySchema {
	return v.value
}

func (v *NullableTerminationInquirySchema) Set(val *TerminationInquirySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminationInquirySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminationInquirySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminationInquirySchema(val *TerminationInquirySchema) *NullableTerminationInquirySchema {
	return &NullableTerminationInquirySchema{value: val, isSet: true}
}

func (v NullableTerminationInquirySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminationInquirySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


