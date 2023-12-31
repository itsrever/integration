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

// Payment - Payment information for the order
type Payment struct {

	// Date when the payment for the whole order was made. This field should be present with a valid value if the order has been fully paid (not just partially). 
	Date time.Time `json:"date"`

	// List of transactions executed as payment for the order.  If the order is `fully_paid`, then this list should have at least one element. 
	Transactions []Transaction `json:"transactions"`
}

// AssertPaymentRequired checks if the required fields are not zero-ed
func AssertPaymentRequired(obj Payment) error {
	elements := map[string]interface{}{
		"date": obj.Date,
		"transactions": obj.Transactions,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Transactions {
		if err := AssertTransactionRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecursePaymentRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Payment (e.g. [][]Payment), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePaymentRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPayment, ok := obj.(Payment)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPaymentRequired(aPayment)
	})
}
