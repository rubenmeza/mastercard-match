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

// UrlGroup struct for UrlGroup
type UrlGroup struct {
	ExactMatchUrls *Url `json:"ExactMatchUrls,omitempty"`
	CloseMatchUrls *Url `json:"CloseMatchUrls,omitempty"`
	NoMatchUrls *Url `json:"NoMatchUrls,omitempty"`
}

// NewUrlGroup instantiates a new UrlGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrlGroup() *UrlGroup {
	this := UrlGroup{}
	return &this
}

// NewUrlGroupWithDefaults instantiates a new UrlGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlGroupWithDefaults() *UrlGroup {
	this := UrlGroup{}
	return &this
}

// GetExactMatchUrls returns the ExactMatchUrls field value if set, zero value otherwise.
func (o *UrlGroup) GetExactMatchUrls() Url {
	if o == nil || o.ExactMatchUrls == nil {
		var ret Url
		return ret
	}
	return *o.ExactMatchUrls
}

// GetExactMatchUrlsOk returns a tuple with the ExactMatchUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlGroup) GetExactMatchUrlsOk() (*Url, bool) {
	if o == nil || o.ExactMatchUrls == nil {
		return nil, false
	}
	return o.ExactMatchUrls, true
}

// HasExactMatchUrls returns a boolean if a field has been set.
func (o *UrlGroup) HasExactMatchUrls() bool {
	if o != nil && o.ExactMatchUrls != nil {
		return true
	}

	return false
}

// SetExactMatchUrls gets a reference to the given Url and assigns it to the ExactMatchUrls field.
func (o *UrlGroup) SetExactMatchUrls(v Url) {
	o.ExactMatchUrls = &v
}

// GetCloseMatchUrls returns the CloseMatchUrls field value if set, zero value otherwise.
func (o *UrlGroup) GetCloseMatchUrls() Url {
	if o == nil || o.CloseMatchUrls == nil {
		var ret Url
		return ret
	}
	return *o.CloseMatchUrls
}

// GetCloseMatchUrlsOk returns a tuple with the CloseMatchUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlGroup) GetCloseMatchUrlsOk() (*Url, bool) {
	if o == nil || o.CloseMatchUrls == nil {
		return nil, false
	}
	return o.CloseMatchUrls, true
}

// HasCloseMatchUrls returns a boolean if a field has been set.
func (o *UrlGroup) HasCloseMatchUrls() bool {
	if o != nil && o.CloseMatchUrls != nil {
		return true
	}

	return false
}

// SetCloseMatchUrls gets a reference to the given Url and assigns it to the CloseMatchUrls field.
func (o *UrlGroup) SetCloseMatchUrls(v Url) {
	o.CloseMatchUrls = &v
}

// GetNoMatchUrls returns the NoMatchUrls field value if set, zero value otherwise.
func (o *UrlGroup) GetNoMatchUrls() Url {
	if o == nil || o.NoMatchUrls == nil {
		var ret Url
		return ret
	}
	return *o.NoMatchUrls
}

// GetNoMatchUrlsOk returns a tuple with the NoMatchUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlGroup) GetNoMatchUrlsOk() (*Url, bool) {
	if o == nil || o.NoMatchUrls == nil {
		return nil, false
	}
	return o.NoMatchUrls, true
}

// HasNoMatchUrls returns a boolean if a field has been set.
func (o *UrlGroup) HasNoMatchUrls() bool {
	if o != nil && o.NoMatchUrls != nil {
		return true
	}

	return false
}

// SetNoMatchUrls gets a reference to the given Url and assigns it to the NoMatchUrls field.
func (o *UrlGroup) SetNoMatchUrls(v Url) {
	o.NoMatchUrls = &v
}

func (o UrlGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExactMatchUrls != nil {
		toSerialize["ExactMatchUrls"] = o.ExactMatchUrls
	}
	if o.CloseMatchUrls != nil {
		toSerialize["CloseMatchUrls"] = o.CloseMatchUrls
	}
	if o.NoMatchUrls != nil {
		toSerialize["NoMatchUrls"] = o.NoMatchUrls
	}
	return json.Marshal(toSerialize)
}

type NullableUrlGroup struct {
	value *UrlGroup
	isSet bool
}

func (v NullableUrlGroup) Get() *UrlGroup {
	return v.value
}

func (v *NullableUrlGroup) Set(val *UrlGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUrlGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlGroup(val *UrlGroup) *NullableUrlGroup {
	return &NullableUrlGroup{value: val, isSet: true}
}

func (v NullableUrlGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


