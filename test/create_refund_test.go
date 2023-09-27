package test

import (
	"testing"

	server "github.com/itsrever/integration/server/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Create_Refund(t *testing.T) {
	cfg, err := configFromEnv()
	require.NoError(t, err)
	c := NewClient(cfg.BaseURL).WithAuth(cfg.Auth)
	if cfg.Debug {
		c = c.Debug()
	}
	test := cfg.Test("CreateRefund")
	testFindOrder := cfg.Test("FindOrderByCustomerOrderPrintedId")
	val, err := NewJsonValidator(schemaLocation)
	require.NoError(t, err)

	refundBody := refundBody()

	t.Run("CREATEREFUND00", func(t *testing.T) {
		resp, err := c.WithNoAuth().Do("POST", test.UrlPattern, nonExistingOrderVars(), refundBody)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})

	t.Run("CREATEREFUND01", func(t *testing.T) {
		resp, err := c.WithAuth(&AuthenticationInfo{
			HeaderName: cfg.Auth.HeaderName,
			ApiKey:     "invalid-api-key",
		}).Do("POST", test.UrlPattern, nonExistingOrderVars(), refundBody)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})
	t.Run("CREATEREFUND02", func(t *testing.T) {
		resp, err := c.Do("POST", test.UrlPattern, nonExistingOrderVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("CREATEREFUND03", func(t *testing.T) {
		resp, err := c.Do("POST", test.UrlPattern, nonExistingOrderVars(), refundBody)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("CREATEREFUND04", func(t *testing.T) {
		test.SkipTestIfScenarioNotPresent(t, testName(t))
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("POST", test.UrlPattern, scenario.Vars(), refundBody)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)

		// recover the order and see that the note has been added
		resp, err = c.Do("GET", testFindOrder.UrlPattern, scenario.Vars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
		body := requireBodyFromResponse(t, resp)
		val.RequireModel(t, "order", body)
		order, err := orderFromBody(body)
		require.NoError(t, err)
		server.AssertSanity(t, order)
		assertRefundInsideOrder(t, order, refundBody)
	})
}

func refundBody() server.RefundRequest {
	return server.RefundRequest{
		Items: refundItems(),
	}
}

func refundItems() []server.RefundRequestItem {
	return []server.RefundRequestItem{
		{
			LineItemId: "1",
			Quantity:   1,
			Amount: server.RefundRequestItemAmount{
				Amount:   10,
				Currency: "EUR",
			},
		},
	}
}

func assertRefundInsideOrder(t *testing.T, order *server.Order, refundBody server.RefundRequest) {
	assert.NotEmpty(t, order.Refunds, "Refunds should not be empty")
	assert.NotEmpty(t, order.Refunds, "Refunds should not be empty")
	assert.Len(t, order.Refunds, 1, "Refunds should have one item")
	assert.Equal(t, len(refundBody.Items), len(order.Refunds[0].Refunds), "Refund items should be equal")
	assertRefundOrder(t, order, order.Refunds)
	assertItemsFromRefund(t, order, order.Refunds, refundBody)
}

//nolint:unparam
func assertRefundOrder(t *testing.T, order *server.Order, refunds []server.RefundOrder) {
	for _, refund := range refunds {
		assert.NotEmpty(t, refund.RefundId, "Refund ID should not be empty")
		assert.NotEmpty(t, refund.TransactionId, "Transaction ID should not be empty")
		assert.NotEmpty(t, refund.Date, "Date should not be empty")
		assert.NotEmpty(t, refund.Amount, "Amount should not be empty")
		assert.Positive(t, refund.Amount.Amount, "Amount should be positive")
		assert.NotEmpty(t, refund.Amount.Currency, "Currency should not be empty")
	}
}

//nolint:unparam
func assertItemsFromRefund(t *testing.T, order *server.Order, items []server.RefundOrder, itemsPayload server.RefundRequest) {
	for i, refund := range items {
		assert.NotEmpty(t, refund.Refunds, "Refunds should not be empty")
		assert.Len(t, refund.Refunds, 1, "Refunds should have one item")
		assert.Equal(t, itemsPayload.Items[i].LineItemId, refund.Refunds[0].LineItemId, "Line item ID should be equal")
		assert.Equal(t, itemsPayload.Items[i].Quantity, refund.Refunds[0].Quantity, "Quantity should be equal")
	}
}
