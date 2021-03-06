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

// ContactRequestSchema struct for ContactRequestSchema
type ContactRequestSchema struct {
	ContactRequest *ContactRequest `json:"ContactRequest,omitempty"`
}

// NewContactRequestSchema instantiates a new ContactRequestSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactRequestSchema() *ContactRequestSchema {
	this := ContactRequestSchema{}
	return &this
}

// NewContactRequestSchemaWithDefaults instantiates a new ContactRequestSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactRequestSchemaWithDefaults() *ContactRequestSchema {
	this := ContactRequestSchema{}
	return &this
}

// GetContactRequest returns the ContactRequest field value if set, zero value otherwise.
func (o *ContactRequestSchema) GetContactRequest() ContactRequest {
	if o == nil || o.ContactRequest == nil {
		var ret ContactRequest
		return ret
	}
	return *o.ContactRequest
}

// GetContactRequestOk returns a tuple with the ContactRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequestSchema) GetContactRequestOk() (*ContactRequest, bool) {
	if o == nil || o.ContactRequest == nil {
		return nil, false
	}
	return o.ContactRequest, true
}

// HasContactRequest returns a boolean if a field has been set.
func (o *ContactRequestSchema) HasContactRequest() bool {
	if o != nil && o.ContactRequest != nil {
		return true
	}

	return false
}

// SetContactRequest gets a reference to the given ContactRequest and assigns it to the ContactRequest field.
func (o *ContactRequestSchema) SetContactRequest(v ContactRequest) {
	o.ContactRequest = &v
}

func (o ContactRequestSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactRequest != nil {
		toSerialize["ContactRequest"] = o.ContactRequest
	}
	return json.Marshal(toSerialize)
}

type NullableContactRequestSchema struct {
	value *ContactRequestSchema
	isSet bool
}

func (v NullableContactRequestSchema) Get() *ContactRequestSchema {
	return v.value
}

func (v *NullableContactRequestSchema) Set(val *ContactRequestSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableContactRequestSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableContactRequestSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactRequestSchema(val *ContactRequestSchema) *NullableContactRequestSchema {
	return &NullableContactRequestSchema{value: val, isSet: true}
}

func (v NullableContactRequestSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactRequestSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


