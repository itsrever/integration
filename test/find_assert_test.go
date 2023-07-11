package test

import (
	"testing"

	server "github.com/itsrever/integration/server/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// assertSanity asserts that the order has data that complies with basic requirements
func assertSanity(t *testing.T, order *server.IntegrationOrder) {
	require.NoError(t, server.AssertIntegrationOrderRequired(*order))
	assertPositiveAmount(t, order)
	assertTaxes(t, order)
	assertShippingCosts(t, order)
	assertIsFulfilled(t, order)
	assertIsPaid(t, order)
	assertCustomer(t, order)
	assertAmountsDoMatch(t, order)
	assertSameCurrencies(t, order) // TODO: maybe this is a list
}

// Valid order with multiple `line_items`, referring products/services **without variants**.
// Implement this case if your e-commerce supports products but has no support for Variants.
// Product variants are a requirement for supporting exchange orders as compensation method.
// The order must have a positive amount in EUR, with taxes and shipping costs.
// Regarding the payment method, must be paid with a non-cash, non-cash on delivery, non-BNPL payment method.
// It should have a discount applied. It must be associated with a valid customer.
// It must be fulfilled and paid
func assertOrderWithoutVariants(t *testing.T, order *server.IntegrationOrder) {
	assertRefundablePaymentMethod(t, order)
	assertDiscountApplied(t, order)

	assert.GreaterOrEqual(t, len(order.LineItems), 1)
	for _, lineItem := range order.LineItems {
		assertHasProduct(t, &lineItem)
		assertNoVariants(t, &lineItem)
	}
}

func assertPositiveAmount(t *testing.T, order *server.IntegrationOrder) {
	assert.Greater(t, order.TotalAmount.AmountCustomer.Amount, float32(0))
	assert.Greater(t, order.TotalAmount.AmountShop.Amount, float32(0))
	isValidCurrency(t, order.TotalAmount.AmountCustomer.Currency)
	isValidCurrency(t, order.TotalAmount.AmountShop.Currency)
}

func assertTaxes(t *testing.T, order *server.IntegrationOrder) {
	assert.Greater(t, order.TotalTaxes.AmountCustomer.Amount, float32(0))
	assert.Greater(t, order.TotalTaxes.AmountShop.Amount, float32(0))
	isValidCurrency(t, order.TotalTaxes.AmountCustomer.Currency)
	isValidCurrency(t, order.TotalTaxes.AmountShop.Currency)
}

func assertShippingCosts(t *testing.T, order *server.IntegrationOrder) {
	assert.Greater(t, order.Shipping.Amount.AmountCustomer.Amount, float32(0))
	assert.Greater(t, order.Shipping.Amount.AmountShop.Amount, float32(0))
	isValidCurrency(t, order.Shipping.Amount.AmountCustomer.Currency)
	isValidCurrency(t, order.Shipping.Amount.AmountShop.Currency)
}

func assertRefundablePaymentMethod(t *testing.T, order *server.IntegrationOrder) {
	assert.NotNil(t, order.Payment)
	for _, transaction := range order.Payment.Transactions {
		assert.True(t, isRefundablePaymentMethod(transaction.PaymentMethodType))
	}
}

func isRefundablePaymentMethod(paymentMethodType string) bool {
	return paymentMethodType == "non-cash" || paymentMethodType == "non-cash-on-delivery" || paymentMethodType == "non-bnpl"
}

func assertDiscountApplied(t *testing.T, order *server.IntegrationOrder) {
	for _, lineItem := range order.LineItems {
		assert.Greater(t, lineItem.TotalDiscounts.AmountCustomer.Amount, float32(0))
	}
}

func assertIsFulfilled(t *testing.T, order *server.IntegrationOrder) {
	assert.Greater(t, len(order.FulfillmentOrders), 0)
}

func assertIsPaid(t *testing.T, order *server.IntegrationOrder) {
	assert.Greater(t, len(order.Payment.Transactions), 0)
}

func assertCustomer(t *testing.T, order *server.IntegrationOrder) {
	assert.NotNil(t, order.Customer)
	assert.NotEmpty(t, order.Customer.Email)
	assert.NotEmpty(t, order.Customer.FirstName)
	assert.NotEmpty(t, order.Customer.LastName)
}

func assertHasProduct(t *testing.T, lineItem *server.IntegrationLineItem) {
	assert.NotEmpty(t, lineItem.Product)
	assert.NotEmpty(t, lineItem.Product.Id)
	assert.NotEmpty(t, lineItem.Product.Name)
	assert.NotEmpty(t, lineItem.Product.Description)
	assert.Greater(t, len(lineItem.Product.Images), 0)
}

func assertNoVariants(t *testing.T, lineItem *server.IntegrationLineItem) {
	assert.Empty(t, lineItem.Product.Variants)
}

func assertAmountsDoMatch(t *testing.T, order *server.IntegrationOrder) {
}

// assertSameCurrencies asserts that all currencies in the order are the same (in terms of shop and customer)
func assertSameCurrencies(t *testing.T, order *server.IntegrationOrder) {
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

func isValidCurrency(t *testing.T, currency string) {
	assert.Len(t, currency, 3)
}
