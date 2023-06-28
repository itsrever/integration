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

// IntegrationOrderTotalAmountCustomer Total amount of the order, including taxes, discounts and shipping costs in the customer currency. This is the amount displayed to the customer as total price of the order. Must match the sum of the `line_items` using the  customer currency, plus the shipping costs. 
type IntegrationOrderTotalAmountCustomer struct {
	// amount (optionally with decimals), without currency symbol and thousands separator
	Amount float32 `json:"amount"`
	// three-letter code as ISO 4217 currency code.  Examples: EUR, USD, JPY, GBP... The currency must be supported by REVER. 
	Currency string `json:"currency"`
}

// NewIntegrationOrderTotalAmountCustomer instantiates a new IntegrationOrderTotalAmountCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationOrderTotalAmountCustomer(amount float32, currency string) *IntegrationOrderTotalAmountCustomer {
	this := IntegrationOrderTotalAmountCustomer{}
	this.Amount = amount
	this.Currency = currency
	return &this
}

// NewIntegrationOrderTotalAmountCustomerWithDefaults instantiates a new IntegrationOrderTotalAmountCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationOrderTotalAmountCustomerWithDefaults() *IntegrationOrderTotalAmountCustomer {
	this := IntegrationOrderTotalAmountCustomer{}
	return &this
}

// GetAmount returns the Amount field value
func (o *IntegrationOrderTotalAmountCustomer) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *IntegrationOrderTotalAmountCustomer) GetAmountOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *IntegrationOrderTotalAmountCustomer) SetAmount(v float32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *IntegrationOrderTotalAmountCustomer) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *IntegrationOrderTotalAmountCustomer) GetCurrencyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *IntegrationOrderTotalAmountCustomer) SetCurrency(v string) {
	o.Currency = v
}

func (o IntegrationOrderTotalAmountCustomer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationOrderTotalAmountCustomer struct {
	value *IntegrationOrderTotalAmountCustomer
	isSet bool
}

func (v NullableIntegrationOrderTotalAmountCustomer) Get() *IntegrationOrderTotalAmountCustomer {
	return v.value
}

func (v *NullableIntegrationOrderTotalAmountCustomer) Set(val *IntegrationOrderTotalAmountCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationOrderTotalAmountCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationOrderTotalAmountCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationOrderTotalAmountCustomer(val *IntegrationOrderTotalAmountCustomer) *NullableIntegrationOrderTotalAmountCustomer {
	return &NullableIntegrationOrderTotalAmountCustomer{value: val, isSet: true}
}

func (v NullableIntegrationOrderTotalAmountCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationOrderTotalAmountCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

