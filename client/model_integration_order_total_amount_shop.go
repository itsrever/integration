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

// IntegrationOrderTotalAmountShop Total amount of the order, including taxes, discounts and shipping costs in the shop currency. This amount is not displayed to the customer but used for accounting purposes. Must match the sum of the `line_items`  using the shop currency, plus the shipping costs. 
type IntegrationOrderTotalAmountShop struct {
	// amount (optionally with decimals), without currency symbol and thousands separator
	Amount float32 `json:"amount"`
	// three-letter code as ISO 4217 currency code.  Examples: EUR, USD, JPY, GBP... The currency must be supported by REVER. 
	Currency string `json:"currency"`
}

// NewIntegrationOrderTotalAmountShop instantiates a new IntegrationOrderTotalAmountShop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationOrderTotalAmountShop(amount float32, currency string) *IntegrationOrderTotalAmountShop {
	this := IntegrationOrderTotalAmountShop{}
	this.Amount = amount
	this.Currency = currency
	return &this
}

// NewIntegrationOrderTotalAmountShopWithDefaults instantiates a new IntegrationOrderTotalAmountShop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationOrderTotalAmountShopWithDefaults() *IntegrationOrderTotalAmountShop {
	this := IntegrationOrderTotalAmountShop{}
	return &this
}

// GetAmount returns the Amount field value
func (o *IntegrationOrderTotalAmountShop) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *IntegrationOrderTotalAmountShop) GetAmountOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *IntegrationOrderTotalAmountShop) SetAmount(v float32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *IntegrationOrderTotalAmountShop) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *IntegrationOrderTotalAmountShop) GetCurrencyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *IntegrationOrderTotalAmountShop) SetCurrency(v string) {
	o.Currency = v
}

func (o IntegrationOrderTotalAmountShop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationOrderTotalAmountShop struct {
	value *IntegrationOrderTotalAmountShop
	isSet bool
}

func (v NullableIntegrationOrderTotalAmountShop) Get() *IntegrationOrderTotalAmountShop {
	return v.value
}

func (v *NullableIntegrationOrderTotalAmountShop) Set(val *IntegrationOrderTotalAmountShop) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationOrderTotalAmountShop) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationOrderTotalAmountShop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationOrderTotalAmountShop(val *IntegrationOrderTotalAmountShop) *NullableIntegrationOrderTotalAmountShop {
	return &NullableIntegrationOrderTotalAmountShop{value: val, isSet: true}
}

func (v NullableIntegrationOrderTotalAmountShop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationOrderTotalAmountShop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

