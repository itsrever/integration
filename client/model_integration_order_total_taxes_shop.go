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

// IntegrationOrderTotalTaxesShop Total taxes of the order, in the shop currency 
type IntegrationOrderTotalTaxesShop struct {
	// amount (optionally with decimals), without currency symbol and thousands separator
	Amount float32 `json:"amount"`
	// three-letter code as ISO 4217 currency code.  Examples: EUR, USD, JPY, GBP... The currency must be supported by REVER. 
	Currency string `json:"currency"`
}

// NewIntegrationOrderTotalTaxesShop instantiates a new IntegrationOrderTotalTaxesShop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationOrderTotalTaxesShop(amount float32, currency string) *IntegrationOrderTotalTaxesShop {
	this := IntegrationOrderTotalTaxesShop{}
	this.Amount = amount
	this.Currency = currency
	return &this
}

// NewIntegrationOrderTotalTaxesShopWithDefaults instantiates a new IntegrationOrderTotalTaxesShop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationOrderTotalTaxesShopWithDefaults() *IntegrationOrderTotalTaxesShop {
	this := IntegrationOrderTotalTaxesShop{}
	return &this
}

// GetAmount returns the Amount field value
func (o *IntegrationOrderTotalTaxesShop) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *IntegrationOrderTotalTaxesShop) GetAmountOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *IntegrationOrderTotalTaxesShop) SetAmount(v float32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *IntegrationOrderTotalTaxesShop) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *IntegrationOrderTotalTaxesShop) GetCurrencyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *IntegrationOrderTotalTaxesShop) SetCurrency(v string) {
	o.Currency = v
}

func (o IntegrationOrderTotalTaxesShop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationOrderTotalTaxesShop struct {
	value *IntegrationOrderTotalTaxesShop
	isSet bool
}

func (v NullableIntegrationOrderTotalTaxesShop) Get() *IntegrationOrderTotalTaxesShop {
	return v.value
}

func (v *NullableIntegrationOrderTotalTaxesShop) Set(val *IntegrationOrderTotalTaxesShop) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationOrderTotalTaxesShop) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationOrderTotalTaxesShop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationOrderTotalTaxesShop(val *IntegrationOrderTotalTaxesShop) *NullableIntegrationOrderTotalTaxesShop {
	return &NullableIntegrationOrderTotalTaxesShop{value: val, isSet: true}
}

func (v NullableIntegrationOrderTotalTaxesShop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationOrderTotalTaxesShop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


