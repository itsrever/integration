/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// RefundResponse - Items from an order being to be refunded.
type RefundResponse struct {

	// The unique identifier for the refund in the e-commerce.
	RefundId string `json:"refund_id"`

	// The unique identifier for the transaction in the e-commerce.
	TransactionId string `json:"transaction_id,omitempty"`

	// List of items to be refunded.
	Items []RefundRequestItem `json:"items"`
}

// AssertRefundResponseRequired checks if the required fields are not zero-ed
func AssertRefundResponseRequired(obj RefundResponse) error {
	elements := map[string]interface{}{
		"refund_id": obj.RefundId,
		"items": obj.Items,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Items {
		if err := AssertRefundRequestItemRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseRefundResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RefundResponse (e.g. [][]RefundResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRefundResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRefundResponse, ok := obj.(RefundResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRefundResponseRequired(aRefundResponse)
	})
}