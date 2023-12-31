/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"time"
)

// Transaction - Transaction executed as payment for an order
type Transaction struct {

	// Payment method used for this transaction. Any string can be returned in here, but some do have a special meaning and should have be preferred if possible:   - `manual`: the payment was made manually, outside of the e-commerce   - `cash-on-delivery`: the payment was made in cash when the order was delivered   - `bnpl`: the payment was made using a Buy Now Pay Later method   - `credit-card`: the payment was made using a credit card   - `debit-card`: the payment was made using a debit card   - `paypal`: the payment was made using PayPal   - `gift`: the payment was made using a gift card 
	PaymentMethodType string `json:"payment_method_type"`

	// Identifier of the transaction in the payment gateway 
	TransactionId string `json:"transaction_id"`

	Amount TransactionAmount `json:"amount"`

	// Date when the transaction was executed 
	Date time.Time `json:"date"`
}

// AssertTransactionRequired checks if the required fields are not zero-ed
func AssertTransactionRequired(obj Transaction) error {
	elements := map[string]interface{}{
		"payment_method_type": obj.PaymentMethodType,
		"transaction_id": obj.TransactionId,
		"amount": obj.Amount,
		"date": obj.Date,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertTransactionAmountRequired(obj.Amount); err != nil {
		return err
	}
	return nil
}

// AssertRecurseTransactionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Transaction (e.g. [][]Transaction), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTransactionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTransaction, ok := obj.(Transaction)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTransactionRequired(aTransaction)
	})
}
