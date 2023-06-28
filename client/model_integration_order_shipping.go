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

// IntegrationOrderShipping Shipping information for the order 
type IntegrationOrderShipping struct {
	AmountShop *IntegrationShippingAmountShop `json:"amount_shop,omitempty"`
	AmountCustomer *IntegrationShippingAmountCustomer `json:"amount_customer,omitempty"`
}

// NewIntegrationOrderShipping instantiates a new IntegrationOrderShipping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationOrderShipping() *IntegrationOrderShipping {
	this := IntegrationOrderShipping{}
	return &this
}

// NewIntegrationOrderShippingWithDefaults instantiates a new IntegrationOrderShipping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationOrderShippingWithDefaults() *IntegrationOrderShipping {
	this := IntegrationOrderShipping{}
	return &this
}

// GetAmountShop returns the AmountShop field value if set, zero value otherwise.
func (o *IntegrationOrderShipping) GetAmountShop() IntegrationShippingAmountShop {
	if o == nil || isNil(o.AmountShop) {
		var ret IntegrationShippingAmountShop
		return ret
	}
	return *o.AmountShop
}

// GetAmountShopOk returns a tuple with the AmountShop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOrderShipping) GetAmountShopOk() (*IntegrationShippingAmountShop, bool) {
	if o == nil || isNil(o.AmountShop) {
    return nil, false
	}
	return o.AmountShop, true
}

// HasAmountShop returns a boolean if a field has been set.
func (o *IntegrationOrderShipping) HasAmountShop() bool {
	if o != nil && !isNil(o.AmountShop) {
		return true
	}

	return false
}

// SetAmountShop gets a reference to the given IntegrationShippingAmountShop and assigns it to the AmountShop field.
func (o *IntegrationOrderShipping) SetAmountShop(v IntegrationShippingAmountShop) {
	o.AmountShop = &v
}

// GetAmountCustomer returns the AmountCustomer field value if set, zero value otherwise.
func (o *IntegrationOrderShipping) GetAmountCustomer() IntegrationShippingAmountCustomer {
	if o == nil || isNil(o.AmountCustomer) {
		var ret IntegrationShippingAmountCustomer
		return ret
	}
	return *o.AmountCustomer
}

// GetAmountCustomerOk returns a tuple with the AmountCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOrderShipping) GetAmountCustomerOk() (*IntegrationShippingAmountCustomer, bool) {
	if o == nil || isNil(o.AmountCustomer) {
    return nil, false
	}
	return o.AmountCustomer, true
}

// HasAmountCustomer returns a boolean if a field has been set.
func (o *IntegrationOrderShipping) HasAmountCustomer() bool {
	if o != nil && !isNil(o.AmountCustomer) {
		return true
	}

	return false
}

// SetAmountCustomer gets a reference to the given IntegrationShippingAmountCustomer and assigns it to the AmountCustomer field.
func (o *IntegrationOrderShipping) SetAmountCustomer(v IntegrationShippingAmountCustomer) {
	o.AmountCustomer = &v
}

func (o IntegrationOrderShipping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AmountShop) {
		toSerialize["amount_shop"] = o.AmountShop
	}
	if !isNil(o.AmountCustomer) {
		toSerialize["amount_customer"] = o.AmountCustomer
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationOrderShipping struct {
	value *IntegrationOrderShipping
	isSet bool
}

func (v NullableIntegrationOrderShipping) Get() *IntegrationOrderShipping {
	return v.value
}

func (v *NullableIntegrationOrderShipping) Set(val *IntegrationOrderShipping) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationOrderShipping) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationOrderShipping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationOrderShipping(val *IntegrationOrderShipping) *NullableIntegrationOrderShipping {
	return &NullableIntegrationOrderShipping{value: val, isSet: true}
}

func (v NullableIntegrationOrderShipping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationOrderShipping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


