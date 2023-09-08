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

// Note - Note (text) associated to an order
type Note struct {

	// Date when the note was added to the order 
	Date time.Time `json:"date"`

	// User that added the note. Can be the customer or a third-party, like REVER. This is optional, but it's recommended. Integrators can guess the name of the app through the usage of the API key. 
	User string `json:"user,omitempty"`

	// Arbitrary text associated to the order 
	Text string `json:"text"`
}

// AssertNoteRequired checks if the required fields are not zero-ed
func AssertNoteRequired(obj Note) error {
	elements := map[string]interface{}{
		"date": obj.Date,
		"text": obj.Text,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseNoteRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Note (e.g. [][]Note), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNoteRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNote, ok := obj.(Note)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNoteRequired(aNote)
	})
}