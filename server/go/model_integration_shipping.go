/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// IntegrationShipping - Shipping information
type IntegrationShipping struct {

	Amount IntegrationShippingAmount `json:"amount"`
}

// AssertIntegrationShippingRequired checks if the required fields are not zero-ed
func AssertIntegrationShippingRequired(obj IntegrationShipping) error {
	elements := map[string]interface{}{
		"amount": obj.Amount,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertIntegrationShippingAmountRequired(obj.Amount); err != nil {
		return err
	}
	return nil
}

// AssertRecurseIntegrationShippingRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IntegrationShipping (e.g. [][]IntegrationShipping), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIntegrationShippingRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIntegrationShipping, ok := obj.(IntegrationShipping)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIntegrationShippingRequired(aIntegrationShipping)
	})
}
