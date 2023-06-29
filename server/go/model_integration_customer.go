/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// IntegrationCustomer - Customer associated to the order
type IntegrationCustomer struct {

	// two-letter code as per the ISO 639-1 codes, preferred by the customer 
	PreferredLang string `json:"preferred_lang,omitempty"`

	// Email address of the customer. This field will be used to match the customer when starting a return process 
	Email string `json:"email"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`
}

// AssertIntegrationCustomerRequired checks if the required fields are not zero-ed
func AssertIntegrationCustomerRequired(obj IntegrationCustomer) error {
	elements := map[string]interface{}{
		"email": obj.Email,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseIntegrationCustomerRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IntegrationCustomer (e.g. [][]IntegrationCustomer), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIntegrationCustomerRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIntegrationCustomer, ok := obj.(IntegrationCustomer)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIntegrationCustomerRequired(aIntegrationCustomer)
	})
}