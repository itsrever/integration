/*
Integration stubs

Stubs for implementing a REVER integration

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// IntegrationPaymentTransaction Transaction executed as payment for the order
type IntegrationPaymentTransaction struct {
	// Payment method used for this transaction. Any string can be returned in here, but some do have a special meaning and should have be preferred if possible:   - `manual`: the payment was made manually, outside of the e-commerce   - `cash-on-delivery`: the payment was made in cash when the order was delivered   - `bnpl`: the payment was made using a Buy Now Pay Later method   - `credit-card`: the payment was made using a credit card   - `debit-card`: the payment was made using a debit card   - `paypal`: the payment was made using PayPal   - `gift`: the payment was made using a gift card 
	PaymentMethodType *string `json:"payment_method_type,omitempty"`
	// Identifier of the transaction in the payment gateway 
	TransactionId *string `json:"transaction_id,omitempty"`
	Amount *IntegrationPaymentTransactionAmount `json:"amount,omitempty"`
	// Date when the transaction was executed 
	Date *time.Time `json:"date,omitempty"`
}

// NewIntegrationPaymentTransaction instantiates a new IntegrationPaymentTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationPaymentTransaction() *IntegrationPaymentTransaction {
	this := IntegrationPaymentTransaction{}
	return &this
}

// NewIntegrationPaymentTransactionWithDefaults instantiates a new IntegrationPaymentTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationPaymentTransactionWithDefaults() *IntegrationPaymentTransaction {
	this := IntegrationPaymentTransaction{}
	return &this
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *IntegrationPaymentTransaction) GetPaymentMethodType() string {
	if o == nil || isNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationPaymentTransaction) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || isNil(o.PaymentMethodType) {
    return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *IntegrationPaymentTransaction) HasPaymentMethodType() bool {
	if o != nil && !isNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *IntegrationPaymentTransaction) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *IntegrationPaymentTransaction) GetTransactionId() string {
	if o == nil || isNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationPaymentTransaction) GetTransactionIdOk() (*string, bool) {
	if o == nil || isNil(o.TransactionId) {
    return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *IntegrationPaymentTransaction) HasTransactionId() bool {
	if o != nil && !isNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *IntegrationPaymentTransaction) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *IntegrationPaymentTransaction) GetAmount() IntegrationPaymentTransactionAmount {
	if o == nil || isNil(o.Amount) {
		var ret IntegrationPaymentTransactionAmount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationPaymentTransaction) GetAmountOk() (*IntegrationPaymentTransactionAmount, bool) {
	if o == nil || isNil(o.Amount) {
    return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *IntegrationPaymentTransaction) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given IntegrationPaymentTransactionAmount and assigns it to the Amount field.
func (o *IntegrationPaymentTransaction) SetAmount(v IntegrationPaymentTransactionAmount) {
	o.Amount = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *IntegrationPaymentTransaction) GetDate() time.Time {
	if o == nil || isNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationPaymentTransaction) GetDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.Date) {
    return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *IntegrationPaymentTransaction) HasDate() bool {
	if o != nil && !isNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *IntegrationPaymentTransaction) SetDate(v time.Time) {
	o.Date = &v
}

func (o IntegrationPaymentTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PaymentMethodType) {
		toSerialize["payment_method_type"] = o.PaymentMethodType
	}
	if !isNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationPaymentTransaction struct {
	value *IntegrationPaymentTransaction
	isSet bool
}

func (v NullableIntegrationPaymentTransaction) Get() *IntegrationPaymentTransaction {
	return v.value
}

func (v *NullableIntegrationPaymentTransaction) Set(val *IntegrationPaymentTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationPaymentTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationPaymentTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationPaymentTransaction(val *IntegrationPaymentTransaction) *NullableIntegrationPaymentTransaction {
	return &NullableIntegrationPaymentTransaction{value: val, isSet: true}
}

func (v NullableIntegrationPaymentTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationPaymentTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

