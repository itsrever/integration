/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

type Status struct {

	// Status of the return which could be set to accepted, rejected or missing by the ecommerce after the review.
	Status string `json:"status,omitempty"`
}

// AssertStatusRequired checks if the required fields are not zero-ed
func AssertStatusRequired(obj Status) error {
	return nil
}

// AssertRecurseStatusRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Status (e.g. [][]Status), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseStatusRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aStatus, ok := obj.(Status)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertStatusRequired(aStatus)
	})
}
