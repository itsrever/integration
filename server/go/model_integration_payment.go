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

// IntegrationPayment - Payment information for the order
type IntegrationPayment struct {

	// Date when the payment for the whole order was made. This field should be present with a valid value if the order has been fully paid (not just partially). 
	Date time.Time `json:"date"`

	// List of transactions executed as payment for the order.  If the order is `fully_paid`, then this list should have at least one element. 
	Transactions []IntegrationTransaction `json:"transactions"`
}

// AssertIntegrationPaymentRequired checks if the required fields are not zero-ed
func AssertIntegrationPaymentRequired(obj IntegrationPayment) error {
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
		if err := AssertIntegrationTransactionRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseIntegrationPaymentRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IntegrationPayment (e.g. [][]IntegrationPayment), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIntegrationPaymentRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIntegrationPayment, ok := obj.(IntegrationPayment)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIntegrationPaymentRequired(aIntegrationPayment)
	})
}