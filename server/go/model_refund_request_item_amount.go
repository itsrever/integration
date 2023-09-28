/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// RefundRequestItemAmount - The total amount to be refunded for this line item.
type RefundRequestItemAmount struct {

	// amount (optionally with decimals), without currency symbol and thousands separator
	Amount float64 `json:"amount"`

	// three-letter code as ISO 4217 currency code.  Examples: EUR, USD, JPY, GBP... The currency must be supported by REVER. 
	Currency string `json:"currency"`
}

// AssertRefundRequestItemAmountRequired checks if the required fields are not zero-ed
func AssertRefundRequestItemAmountRequired(obj RefundRequestItemAmount) error {
	elements := map[string]interface{}{
		"amount": obj.Amount,
		"currency": obj.Currency,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseRefundRequestItemAmountRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RefundRequestItemAmount (e.g. [][]RefundRequestItemAmount), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRefundRequestItemAmountRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRefundRequestItemAmount, ok := obj.(RefundRequestItemAmount)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRefundRequestItemAmountRequired(aRefundRequestItemAmount)
	})
}