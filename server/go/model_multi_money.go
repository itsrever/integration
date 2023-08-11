/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// MultiMoney - Total amount charged as shipping costs. Taxes are included. Can be zero. 
type MultiMoney struct {

	AmountShop MultiMoneyAmountShop `json:"amount_shop"`

	AmountCustomer MultiMoneyAmountCustomer `json:"amount_customer"`
}

// AssertMultiMoneyRequired checks if the required fields are not zero-ed
func AssertMultiMoneyRequired(obj MultiMoney) error {
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

// AssertRecurseMultiMoneyRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of MultiMoney (e.g. [][]MultiMoney), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseMultiMoneyRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aMultiMoney, ok := obj.(MultiMoney)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertMultiMoneyRequired(aMultiMoney)
	})
}
