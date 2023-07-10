/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// IntegrationProductUnitPrice - Price per unit of the product in the shop currency. Must be present if there are no variants. Does not include taxes nor discounts. Those will be applied later on.
type IntegrationProductUnitPrice struct {

	// amount (optionally with decimals), without currency symbol and thousands separator
	Amount float32 `json:"amount"`

	// three-letter code as ISO 4217 currency code.  Examples: EUR, USD, JPY, GBP... The currency must be supported by REVER. 
	Currency string `json:"currency"`
}

// AssertIntegrationProductUnitPriceRequired checks if the required fields are not zero-ed
func AssertIntegrationProductUnitPriceRequired(obj IntegrationProductUnitPrice) error {
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

// AssertRecurseIntegrationProductUnitPriceRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IntegrationProductUnitPrice (e.g. [][]IntegrationProductUnitPrice), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIntegrationProductUnitPriceRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIntegrationProductUnitPrice, ok := obj.(IntegrationProductUnitPrice)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIntegrationProductUnitPriceRequired(aIntegrationProductUnitPrice)
	})
}