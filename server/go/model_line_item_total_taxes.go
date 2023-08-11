/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// LineItemTotalTaxes - total amount of taxes for this line. Each item of this line item has proportional taxes.
type LineItemTotalTaxes struct {

	AmountShop MultiMoneyAmountShop `json:"amount_shop"`

	AmountCustomer MultiMoneyAmountCustomer `json:"amount_customer"`
}

// AssertLineItemTotalTaxesRequired checks if the required fields are not zero-ed
func AssertLineItemTotalTaxesRequired(obj LineItemTotalTaxes) error {
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

// AssertRecurseLineItemTotalTaxesRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LineItemTotalTaxes (e.g. [][]LineItemTotalTaxes), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLineItemTotalTaxesRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLineItemTotalTaxes, ok := obj.(LineItemTotalTaxes)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLineItemTotalTaxesRequired(aLineItemTotalTaxes)
	})
}
