/*
Integration stubs

Stubs for implementing a REVER integration

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IntegrationProductOption struct for IntegrationProductOption
type IntegrationProductOption struct {
	// Unique identifier for the resource in the source platform
	Id *string `json:"id,omitempty"`
	// Name of the option
	Name *string `json:"name,omitempty"`
	// Value of the option
	Value *string `json:"value,omitempty"`
}

// NewIntegrationProductOption instantiates a new IntegrationProductOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationProductOption() *IntegrationProductOption {
	this := IntegrationProductOption{}
	return &this
}

// NewIntegrationProductOptionWithDefaults instantiates a new IntegrationProductOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationProductOptionWithDefaults() *IntegrationProductOption {
	this := IntegrationProductOption{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationProductOption) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProductOption) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationProductOption) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegrationProductOption) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationProductOption) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProductOption) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationProductOption) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationProductOption) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IntegrationProductOption) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationProductOption) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IntegrationProductOption) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IntegrationProductOption) SetValue(v string) {
	o.Value = &v
}

func (o IntegrationProductOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationProductOption struct {
	value *IntegrationProductOption
	isSet bool
}

func (v NullableIntegrationProductOption) Get() *IntegrationProductOption {
	return v.value
}

func (v *NullableIntegrationProductOption) Set(val *IntegrationProductOption) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationProductOption) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationProductOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationProductOption(val *IntegrationProductOption) *NullableIntegrationProductOption {
	return &NullableIntegrationProductOption{value: val, isSet: true}
}

func (v NullableIntegrationProductOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationProductOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

