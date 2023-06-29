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

// IntegrationFulfillmentOrder - Represents either an item or a group of items in an order that are to be fulfilled from the same location.  There can be more than one fulfillment order for an order at a given location 
type IntegrationFulfillmentOrder struct {

	// An identifier of the e-commerce location that shipped the items in this fulfillment order. 
	LocationId string `json:"location_id"`

	// Date when the fulfillment was executed 
	Date time.Time `json:"date"`

	// Information of the `line_items` fulfilled in this order  
	Fulfillments []IntegrationFulfillment `json:"fulfillments"`
}

// AssertIntegrationFulfillmentOrderRequired checks if the required fields are not zero-ed
func AssertIntegrationFulfillmentOrderRequired(obj IntegrationFulfillmentOrder) error {
	elements := map[string]interface{}{
		"location_id": obj.LocationId,
		"date": obj.Date,
		"fulfillments": obj.Fulfillments,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Fulfillments {
		if err := AssertIntegrationFulfillmentRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseIntegrationFulfillmentOrderRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IntegrationFulfillmentOrder (e.g. [][]IntegrationFulfillmentOrder), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIntegrationFulfillmentOrderRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIntegrationFulfillmentOrder, ok := obj.(IntegrationFulfillmentOrder)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIntegrationFulfillmentOrderRequired(aIntegrationFulfillmentOrder)
	})
}
