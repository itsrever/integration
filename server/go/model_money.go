/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// Money - Model for an amount plus the currency in which it is expressed. 
type Money struct {

	// amount (optionally with decimals), without currency symbol and thousands separator
	Amount float64 `json:"amount"`

	// three-letter code as ISO 4217 currency code.  Examples: EUR, USD, JPY, GBP... The currency must be supported by REVER. 
	Currency string `json:"currency"`
}

// AssertMoneyRequired checks if the required fields are not zero-ed
func AssertMoneyRequired(obj Money) error {
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

// AssertRecurseMoneyRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Money (e.g. [][]Money), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseMoneyRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aMoney, ok := obj.(Money)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertMoneyRequired(aMoney)
	})
}