/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// IntegrationMoney - Model for an amount plus the currency in which it is expressed. 
type IntegrationMoney struct {

	// amount (optionally with decimals), without currency symbol and thousands separator
	Amount float64 `json:"amount"`

	// three-letter code as ISO 4217 currency code.  Examples: EUR, USD, JPY, GBP... The currency must be supported by REVER. 
	Currency string `json:"currency"`
}

// AssertIntegrationMoneyRequired checks if the required fields are not zero-ed
func AssertIntegrationMoneyRequired(obj IntegrationMoney) error {
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

// AssertRecurseIntegrationMoneyRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IntegrationMoney (e.g. [][]IntegrationMoney), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIntegrationMoneyRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIntegrationMoney, ok := obj.(IntegrationMoney)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIntegrationMoneyRequired(aIntegrationMoney)
	})
}
