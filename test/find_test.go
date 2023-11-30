package test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	server "github.com/itsrever/integration/server/go"
)

func Test_FindOrderByCustomerOrderPrintedId(t *testing.T) {
	cfg, err := configFromEnv()
	require.NoError(t, err)
	c := NewClient(cfg.BaseURL).WithAuth(cfg.Auth)
	if cfg.Debug {
		c = c.Debug()
	}
	val, err := NewJsonValidator(schemaLocation)
	require.NoError(t, err)
	test := cfg.Test("FindOrderByCustomerOrderPrintedId")

	t.Run("FIND00", func(t *testing.T) {
		resp, err := c.WithNoAuth().Do("GET", test.UrlPattern, emptyVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})

	t.Run("FIND01", func(t *testing.T) {
		resp, err := c.WithAuth(&ApiKeyAuthInfo{
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
		resp, err := c.Do("GET", test.UrlPattern, nonExistingOrderVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("FIND04", func(t *testing.T) {
		test.SkipTestIfScenarioNotPresent(t, testName(t))
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("GET", test.UrlPattern, scenario.Vars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
		body := requireBodyFromResponse(t, resp)
		val.RequireModel(t, "order", body)
		order, err := orderFromBody(body)
		require.NoError(t, err)
		server.AssertSanity(t, order)
		server.AssertOrderWithoutVariants(t, order)
		assert.Equal(t, scenario.Vars()["customer_printed_order_id"], order.Identification.CustomerPrintedOrderId,
			"the customer printed order id does not match")
	})
	t.Run("FIND05", func(t *testing.T) {
		test.SkipTestIfScenarioNotPresent(t, testName(t))
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("GET", test.UrlPattern, scenario.Vars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
		body := requireBodyFromResponse(t, resp)
		val.RequireModel(t, "order", body)
		order, err := orderFromBody(body)
		require.NoError(t, err)
		server.AssertSanity(t, order)
		server.AssertOrderWithVariants(t, order)
		assert.Equal(t, scenario.Vars()["customer_printed_order_id"], order.Identification.CustomerPrintedOrderId,
			"the customer printed order id does not match")
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

func requireBodyFromResponse(t *testing.T, resp *http.Response) []byte {
	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return data
}

// orderFromBody decodes the response body into an order
func orderFromBody(body []byte) (*server.Order, error) {
	result := &server.Order{}
	err := json.Unmarshal(body, result)
	return result, err
}

func nonExistingOrderVars() map[string]string {
	return map[string]string{
		"customer_printed_order_id": "non-existing-order",
	}
}
