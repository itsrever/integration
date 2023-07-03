package test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	server "github.com/itsrever/integration/server/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_FindOrderByCustomerOrderPrintedId(t *testing.T) {
	cfg, err := configFromEnv()
	require.NoError(t, err)
	c := NewClient(cfg.BaseURL).WithAuth(cfg.Auth)
	test := cfg.Test("FindOrderByCustomerOrderPrintedId")

	t.Run("FIND00", func(t *testing.T) {
		resp, err := c.WithNoAuth().Do("GET", test.UrlPattern, emptyVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.StatusCode, 401)
	})

	t.Run("FIND01", func(t *testing.T) {
		resp, err := c.WithAuth(&AuthenticationInfo{
			HeaderName: cfg.Auth.HeaderName,
			ApiKey:     "invalid-api-key",
		}).Do("GET", test.UrlPattern, emptyVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.StatusCode, 401)
	})

	t.Run("FIND02", func(t *testing.T) {
		resp, err := c.Do("GET", test.UrlPattern, emptyVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.StatusCode, 400)
	})

	t.Run("FIND03", func(t *testing.T) {
		resp, err := c.Do("GET", test.UrlPattern, map[string]string{
			"customer_printed_order_id": "non-existing-order",
		}, nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.StatusCode, 404)
	})

	t.Run("FIND04", func(t *testing.T) {
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("GET", test.UrlPattern, scenario.Vars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.StatusCode, 200)
		order, err := orderFromResponse(resp)
		require.NoError(t, err)
		require.NoError(t, server.AssertIntegrationOrderRequired(*order))
		assertOrderWithoutVariants(t, order)
		assert.Equal(t, order.Identification.CustomerPrintedOrderId, scenario.Vars()["customer_printed_order_id"])
	})
}

func emptyVars() map[string]string {
	return map[string]string{
		"customer_printed_order_id": "",
	}
}

func testName(t *testing.T) string {
	return strings.Split(t.Name(), "/")[1]
}

// orderFromResponse decodes the response body into an order
func orderFromResponse(resp *http.Response) (*server.IntegrationOrder, error) {
	result := &server.IntegrationOrder{}
	err := json.NewDecoder(resp.Body).Decode(result)
	return result, err
}

// Valid order with multiple `line_items`, referring products/services **without variants**.
// Implement this case if your e-commerce supports products but has no support for Variants.
// Product variants are a requirement for supporting exchange orders as compensation method.
// The order must have a positive amount in EUR, with taxes and shipping costs.
// Regarding the payment method, must be paid with a non-cash, non-cash on delivery, non-BNPL payment method.
// It should have a discount applied. It must be associated with a valid customer.
// It must be fulfilled and paid
func assertOrderWithoutVariants(t *testing.T, order *server.IntegrationOrder) {
	hasPositiveAmount(t, order)
	hasTaxes(t, order)
	hasShippingCosts(t, order)
	hasRefundablePaymentMethod(t, order)
	hasDiscountApplied(t, order)
	isFulfilled(t, order)
	isPaid(t, order)
	hasCustomer(t, order)
	amountsDoMatch(t, order)
	// TODO: maybe this is a list
	sameCurrencies(t, order)

	assert.GreaterOrEqual(t, len(order.LineItems), 1)
	for _, lineItem := range order.LineItems {
		isProduct(t, &lineItem)
		hasNoVariants(t, &lineItem)
	}
}

func hasPositiveAmount(t *testing.T, order *server.IntegrationOrder) {
	assert.GreaterOrEqual(t, order.TotalAmount.AmountCustomer.Amount, 0)
	assert.GreaterOrEqual(t, order.TotalAmount.AmountShop.Amount, 0)
	isValidCurrency(t, order.TotalAmount.AmountCustomer.Currency)
	isValidCurrency(t, order.TotalAmount.AmountShop.Currency)
}

func hasTaxes(t *testing.T, order *server.IntegrationOrder) {

}

func hasShippingCosts(t *testing.T, order *server.IntegrationOrder) {

}

func hasRefundablePaymentMethod(t *testing.T, order *server.IntegrationOrder) {

}

func hasDiscountApplied(t *testing.T, order *server.IntegrationOrder) {

}

func isFulfilled(t *testing.T, order *server.IntegrationOrder) {

}

func isPaid(t *testing.T, order *server.IntegrationOrder) {

}

func hasCustomer(t *testing.T, order *server.IntegrationOrder) {

}

func isProduct(t *testing.T, lineItem *server.IntegrationLineItem) {

}

func hasNoVariants(t *testing.T, lineItem *server.IntegrationLineItem) {

}

func amountsDoMatch(t *testing.T, order *server.IntegrationOrder) {

}

func sameCurrencies(t *testing.T, order *server.IntegrationOrder) {

}

func isValidCurrency(t *testing.T, currency string) {
	assert.Len(t, currency, 3)
}
