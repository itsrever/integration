/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// Category - A category from the e-commerce catalog. It can be used to filter products or decide if a product is returnable or not.
type Category struct {

	Id string `json:"id"`
}

// AssertCategoryRequired checks if the required fields are not zero-ed
func AssertCategoryRequired(obj Category) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCategoryRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Category (e.g. [][]Category), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCategoryRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCategory, ok := obj.(Category)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCategoryRequired(aCategory)
	})
}