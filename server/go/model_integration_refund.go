/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// IntegrationRefund - A refund operation executed over one or more items of an order 
type IntegrationRefund struct {

	// ID of the line item returned. Must exist in the `line_items` array of the order.
	LineItemId string `json:"line_item_id"`

	// Number of products returned. The sum of quantities per `line_item_id` must match the total quantity of the line item.
	Quantity float32 `json:"quantity"`
}

// AssertIntegrationRefundRequired checks if the required fields are not zero-ed
func AssertIntegrationRefundRequired(obj IntegrationRefund) error {
	elements := map[string]interface{}{
		"line_item_id": obj.LineItemId,
		"quantity": obj.Quantity,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseIntegrationRefundRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IntegrationRefund (e.g. [][]IntegrationRefund), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIntegrationRefundRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIntegrationRefund, ok := obj.(IntegrationRefund)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIntegrationRefundRequired(aIntegrationRefund)
	})
}