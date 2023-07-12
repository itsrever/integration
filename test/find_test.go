package test

import (
	"encoding/json"
	"fmt"
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
		require.Equal(t, 401, resp.StatusCode)
	})

	t.Run("FIND01", func(t *testing.T) {
		resp, err := c.WithAuth(&AuthenticationInfo{
			HeaderName: cfg.Auth.HeaderName,
			ApiKey:     "invalid-api-key",
		}).Do("GET", test.UrlPattern, emptyVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})

	t.Run("FIND02", func(t *testing.T) {
		resp, err := c.Do("GET", test.UrlPattern, emptyVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("FIND03", func(t *testing.T) {
		resp, err := c.Do("GET", test.UrlPattern, map[string]string{
			"customer_printed_order_id": "non-existing-order",
		}, nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("FIND04", func(t *testing.T) {
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("GET", test.UrlPattern, scenario.Vars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
		order, err := orderFromResponse(resp)
		require.NoError(t, err)
		assertSanity(t, order)
		assertOrderWithoutVariants(t, order)
		assert.Equal(t, scenario.Vars()["customer_printed_order_id"], order.Identification.CustomerPrintedOrderId)
	})
	t.Run("FIND05", func(t *testing.T) {
		fmt.Print(testName(t))
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("GET", test.UrlPattern, scenario.Vars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
		order, err := orderFromResponse(resp)
		require.NoError(t, err)
		assertSanity(t, order)
		assertOrderWitVariants(t, order)
		assert.Equal(t, scenario.Vars()["customer_printed_order_id"], order.Identification.CustomerPrintedOrderId)
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

func hasDiscountOrder(order *server.IntegrationOrder) bool {
	for _, lineItem := range order.LineItems {
		if hasDiscountLineItem(lineItem) {
			return true
		}
	}
	return false
}

func hasDiscountLineItem(order server.IntegrationLineItem) bool {
	return order.TotalDiscounts.AmountShop.Amount > 0
}
