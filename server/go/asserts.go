package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertSanity asserts that the order has data that complies with basic requirements
func AssertSanity(t *testing.T, order *IntegrationOrder) {
	assertPositiveAmount(t, order)
	assertTaxes(t, order)
	assertShippingCosts(t, order)
	assertIsFulfilled(t, order)
	assertIsPaid(t, order)
	assertCustomer(t, order)
	assertAmountsDoMatch(t, order)
	assertSameCurrencies(t, order)
}

// AssertOrderWithVariants validates order with multiple `line_items`, referring products/services **with variants**.
// Product variants are a requirement for supporting exchange orders as compensation method.
// The order must have a positive amount in EUR, with taxes and shipping costs.
// Regarding the payment method, must be paid with a non-cash, non-cash on delivery, non-BNPL payment method.
// It should have a discount applied. It must be associated with a valid customer.
// It must be fulfilled and paid
func AssertOrderWithVariants(t *testing.T, order *IntegrationOrder) {
	assertRefundablePaymentMethod(t, order)
	assertDiscountApplied(t, order)

	assert.GreaterOrEqual(t, len(order.LineItems), 1)
	for _, lineItem := range order.LineItems {
		assertHasProduct(t, &lineItem)
		assertVariants(t, &lineItem)
	}
}

// AssertOrderWithoutVariants checks a valid order with multiple `line_items`,
// referring products/services **without variants**.
// Implement this case if your e-commerce supports products but has no support for Variants.
// Product variants are a requirement for supporting exchange orders as compensation method.
// The order must have a positive amount in EUR, with taxes and shipping costs.
// Regarding the payment method, must be paid with a non-cash, non-cash on delivery, non-BNPL payment method.
// It should have a discount applied. It must be associated with a valid customer.
// It must be fulfilled and paid
func AssertOrderWithoutVariants(t *testing.T, order *IntegrationOrder) {
	assertRefundablePaymentMethod(t, order)
	assertDiscountApplied(t, order)

	assert.GreaterOrEqual(t, len(order.LineItems), 1)
	for _, lineItem := range order.LineItems {
		assertHasProduct(t, &lineItem)
		assertNoVariants(t, &lineItem)
	}
}

// ***************
// INTERNAL METHODS
// ***************

func assertPositiveAmount(t *testing.T, order *IntegrationOrder) {
	assert.Greater(t, order.TotalAmount.AmountCustomer.Amount, float64(0),
		"order customer total amount is not greater than 0")
	assert.Greater(t, order.TotalAmount.AmountShop.Amount, float64(0),
		"order shop total amount is not greater than 0")
	isValidCurrency(t, order.TotalAmount.AmountCustomer.Currency)
	isValidCurrency(t, order.TotalAmount.AmountShop.Currency)
}

func assertTaxes(t *testing.T, order *IntegrationOrder) {
	assert.GreaterOrEqual(t, order.TotalTaxes.AmountCustomer.Amount, float64(0),
		"order customer amount total taxes are not greater or equal to 0")
	assert.GreaterOrEqual(t, order.TotalTaxes.AmountShop.Amount, float64(0),
		"order shop amount total taxes are not greater or equal to 0")
	assert.GreaterOrEqual(t, order.Shipping.Taxes.AmountShop.Amount, float64(0),
		"order shipping taxes are not greater or equal to 0")

	isValidCurrency(t, order.TotalTaxes.AmountCustomer.Currency)
	isValidCurrency(t, order.TotalTaxes.AmountShop.Currency)
	isValidCurrency(t, order.Shipping.Taxes.AmountCustomer.Currency)
	isValidCurrency(t, order.Shipping.Taxes.AmountShop.Currency)
}

func assertShippingCosts(t *testing.T, order *IntegrationOrder) {
	assert.GreaterOrEqual(t, order.Shipping.Amount.AmountCustomer.Amount, float64(0),
		"order customer shipping amount is not greater or equal to 0")
	assert.GreaterOrEqual(t, order.Shipping.Amount.AmountShop.Amount, float64(0),
		"order shop shipping amount is not greater or equal to 0")
	isValidCurrency(t, order.Shipping.Amount.AmountCustomer.Currency)
	isValidCurrency(t, order.Shipping.Amount.AmountShop.Currency)
}

func assertRefundablePaymentMethod(t *testing.T, order *IntegrationOrder) {
	assert.NotNil(t, order.Payment)
	for _, transaction := range order.Payment.Transactions {
		assert.True(t, isRefundablePaymentMethod(transaction.PaymentMethodType),
			fmt.Sprintf("the payment method type %v is not refundable", transaction.PaymentMethodType))
	}
}

func isRefundablePaymentMethod(paymentMethodType string) bool {
	// TODO: decide enums for payment method types
	return paymentMethodType != "cash" && paymentMethodType != "CoD"
}

func assertDiscountApplied(t *testing.T, order *IntegrationOrder) {
	for _, lineItem := range order.LineItems {
		assert.Greater(t, lineItem.TotalDiscounts.AmountCustomer.Amount, float64(0),
			"a discount was not applied to the line item but it was expected")
	}
}

func assertIsFulfilled(t *testing.T, order *IntegrationOrder) {
	assert.Greater(t, len(order.FulfillmentOrders), 0,
		"the order has not at least one fulfillment order")
}

func assertIsPaid(t *testing.T, order *IntegrationOrder) {
	assert.Greater(t, len(order.Payment.Transactions), 0,
		"the order has not at least one transaction")
}

func assertCustomer(t *testing.T, order *IntegrationOrder) {
	assert.NotNil(t, order.Customer, "missing customer data")
	assert.NotEmpty(t, order.Customer.Email, "missing customer email")
	assert.NotEmpty(t, order.Customer.FirstName, "the customer first name is empty")
}

func assertHasProduct(t *testing.T, lineItem *IntegrationLineItem) {
	assert.NotEmpty(t, lineItem.Product, "missing product data")
	assert.NotEmpty(t, lineItem.Product.Id, "missing product id")
	assert.NotEmpty(t, lineItem.Product.Name, "missing product name")
	assert.NotEmpty(t, lineItem.Product.Description, "missing product description")
	assert.Greater(t, len(lineItem.Product.Images), 0,
		"missing product image")
}

func assertNoVariants(t *testing.T, lineItem *IntegrationLineItem) {
	assert.Empty(t, lineItem.Product.Variants, "product variants are not empty")
}

func assertAmountsDoMatch(t *testing.T, order *IntegrationOrder) {
	var totalAmountShop, totalAmountCustomer float64
	var totalTaxShop, totalTaxCustomer float64

	for _, lineItem := range order.LineItems {
		totalCustomerLine := lineItem.Subtotal.AmountCustomer.Amount +
			lineItem.TotalTaxes.AmountCustomer.Amount - lineItem.TotalDiscounts.AmountCustomer.Amount
		assert.Equal(t, lineItem.Total.AmountCustomer.Amount, totalCustomerLine,
			"line item customer total amount does not match ")
		totalAmountCustomer += totalCustomerLine
		totalTaxCustomer += lineItem.TotalTaxes.AmountCustomer.Amount
	}

	for _, lineItem := range order.LineItems {
		totalShopLine := lineItem.Subtotal.AmountShop.Amount +
			lineItem.TotalTaxes.AmountShop.Amount - lineItem.TotalDiscounts.AmountShop.Amount
		assert.Equal(t, lineItem.Total.AmountShop.Amount, totalShopLine,
			"line item shop total amount does not match ")
		totalAmountShop += totalShopLine
		totalTaxShop += lineItem.TotalTaxes.AmountShop.Amount
	}

	totalAmountShop += order.Shipping.Amount.AmountShop.Amount
	totalAmountCustomer += order.Shipping.Amount.AmountCustomer.Amount

	assert.Equal(t, totalAmountCustomer, order.TotalAmount.AmountCustomer.Amount,
		"order total customer amount does not match")
	assert.Equal(t, totalAmountShop, order.TotalAmount.AmountShop.Amount,
		"order total shop amount does not match")
}

func assertVariants(t *testing.T, lineItem *IntegrationLineItem) {
	assert.NotEmpty(t, lineItem.Product.Variants)
	for _, variant := range lineItem.Product.Variants {
		assert.NotEmpty(t, variant.Id, "variant id is empty")
		assert.NotEmpty(t, variant.Name, "variant name is empty")
		assert.NotEmpty(t, variant.Description, "variant description is empty")
		assert.Greater(t, len(variant.Images), 0, "variant images is empty")
		assert.NotEmpty(t, variant.Sku, "variant sku is empty")
		assert.Greater(t, variant.UnitPrice.Amount, float64(0), "variant unit price is empty")
		assert.NotEmpty(t, variant.UnitPrice.Currency, "variant unit price currency is empty")
		isValidCurrency(t, variant.UnitPrice.Currency)
	}
}

// assertSameCurrencies asserts that all currencies in the order are the same (in terms of shop and customer)
func assertSameCurrencies(t *testing.T, order *IntegrationOrder) {
	shopCurrency := order.TotalAmount.AmountShop.Currency
	custCurrency := order.TotalAmount.AmountCustomer.Currency

	for _, lineItem := range order.LineItems {
		if lineItem.Subtotal.AmountShop.Currency != shopCurrency {
			assert.Fail(t, "line item subtotal currency does not match shop currency")
		}
		if lineItem.Subtotal.AmountCustomer.Currency != custCurrency {
			assert.Fail(t, "line item subtotal currency does not match customer currency")
		}
		if lineItem.Total.AmountShop.Currency != shopCurrency {
			assert.Fail(t, "line item total currency does not match shop currency")
		}
		if lineItem.Total.AmountCustomer.Currency != custCurrency {
			assert.Fail(t, "line item total currency does not match customer currency")
		}
		if lineItem.TotalTaxes.AmountShop.Currency != shopCurrency {
			assert.Fail(t, "line item taxes currency does not match shop currency")
		}
		if lineItem.TotalTaxes.AmountCustomer.Currency != custCurrency {
			assert.Fail(t, "line item total taxes currency does not match customer currency")
		}
		if hasDiscountLineItem(lineItem) {
			if lineItem.TotalDiscounts.AmountShop.Currency != shopCurrency {
				assert.Fail(t, "line item discount currency does not match shop currency")
			}
			if lineItem.TotalDiscounts.AmountCustomer.Currency != custCurrency {
				assert.Fail(t, "line item discount taxes currency does not match customer currency")
			}
		}
	}
}

func hasDiscountLineItem(order IntegrationLineItem) bool {
	return order.TotalDiscounts.AmountShop.Amount > 0
}

func isValidCurrency(t *testing.T, currency string) {
	assert.Len(t, currency, 3, fmt.Sprintf("currency '%v' is empty or not valid", currency))
}
