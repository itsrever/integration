/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// MultiMoneyAmountCustomer - Amount in the customer currency.  Used as the amount + currency that the customer is really going to use for paying, the currency that the customer selected in the website. Usually, this comes from applying a conversion rate to the shop currency.  
type MultiMoneyAmountCustomer struct {

	// amount (optionally with decimals), without currency symbol and thousands separator
	Amount float64 `json:"amount"`

	// three-letter code as ISO 4217 currency code.  Examples: EUR, USD, JPY, GBP... The currency must be supported by REVER. 
	Currency string `json:"currency"`
}

// AssertMultiMoneyAmountCustomerRequired checks if the required fields are not zero-ed
func AssertMultiMoneyAmountCustomerRequired(obj MultiMoneyAmountCustomer) error {
	elements := map[string]interface{}{
		"amount": obj.Amount,
		"currency": obj.Currency,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseMultiMoneyAmountCustomerRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of MultiMoneyAmountCustomer (e.g. [][]MultiMoneyAmountCustomer), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseMultiMoneyAmountCustomerRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aMultiMoneyAmountCustomer, ok := obj.(MultiMoneyAmountCustomer)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertMultiMoneyAmountCustomerRequired(aMultiMoneyAmountCustomer)
	})
}
