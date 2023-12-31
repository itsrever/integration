/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// Identification - Extended data for identifying an order in any e-commerce
type Identification struct {

	// This is the `order_id` as seen by the customer (for example, in the email confirmation) 
	CustomerPrintedOrderId string `json:"customer_printed_order_id"`

	// Unique identifier in the origin platform. Might not be customer-friendly. It can be the same as the `customer_printed_order_id``
	Id string `json:"id,omitempty"`
}

// AssertIdentificationRequired checks if the required fields are not zero-ed
func AssertIdentificationRequired(obj Identification) error {
	elements := map[string]interface{}{
		"customer_printed_order_id": obj.CustomerPrintedOrderId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseIdentificationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Identification (e.g. [][]Identification), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIdentificationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIdentification, ok := obj.(Identification)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIdentificationRequired(aIdentification)
	})
}
