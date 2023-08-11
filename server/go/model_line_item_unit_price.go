/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// LineItemUnitPrice - unit price of a single item (product as listed, no discounts, no taxes, quantity = 1)
type LineItemUnitPrice struct {

	AmountShop MultiMoneyAmountShop `json:"amount_shop"`

	AmountCustomer MultiMoneyAmountCustomer `json:"amount_customer"`
}

// AssertLineItemUnitPriceRequired checks if the required fields are not zero-ed
func AssertLineItemUnitPriceRequired(obj LineItemUnitPrice) error {
	elements := map[string]interface{}{
		"amount_shop": obj.AmountShop,
		"amount_customer": obj.AmountCustomer,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertMultiMoneyAmountShopRequired(obj.AmountShop); err != nil {
		return err
	}
	if err := AssertMultiMoneyAmountCustomerRequired(obj.AmountCustomer); err != nil {
		return err
	}
	return nil
}

// AssertRecurseLineItemUnitPriceRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LineItemUnitPrice (e.g. [][]LineItemUnitPrice), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLineItemUnitPriceRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLineItemUnitPrice, ok := obj.(LineItemUnitPrice)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLineItemUnitPriceRequired(aLineItemUnitPrice)
	})
}
