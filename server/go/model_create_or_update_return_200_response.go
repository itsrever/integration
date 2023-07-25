/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

type CreateOrUpdateReturn200Response struct {

	// the id of the return generated in the platform
	ReturnId string `json:"return_id"`
}

// AssertCreateOrUpdateReturn200ResponseRequired checks if the required fields are not zero-ed
func AssertCreateOrUpdateReturn200ResponseRequired(obj CreateOrUpdateReturn200Response) error {
	elements := map[string]interface{}{
		"return_id": obj.ReturnId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCreateOrUpdateReturn200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CreateOrUpdateReturn200Response (e.g. [][]CreateOrUpdateReturn200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCreateOrUpdateReturn200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCreateOrUpdateReturn200Response, ok := obj.(CreateOrUpdateReturn200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCreateOrUpdateReturn200ResponseRequired(aCreateOrUpdateReturn200Response)
	})
}
