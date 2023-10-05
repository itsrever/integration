package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	server "github.com/itsrever/integration/server/go"
)

func Test_Update_Return(t *testing.T) {
	cfg, err := configFromEnv()
	require.NoError(t, err)
	c := NewClient(cfg.BaseURL).WithAuth(cfg.Auth)
	if cfg.Debug {
		c = c.Debug()
	}
	test := cfg.Test("UpdateReturn")
	returnRequest := server.ReturnRequest{
		Returns: []server.ReturnRequestItem{
			{
				LineItemId: "testing1",
				Quantity:   1,
				Status:     server.Status{Status: "APPROVED"},
			},
		},
	}

	t.Run("UPDATERETURN001", func(t *testing.T) {
		scenario := test.Scenario(testName(t))
		resp,err := c.Do("PUT", test.UrlPattern, scenario.Vars(), returnRequest)
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("UPDATERETURN002", func(t *testing.T) {
		resp, err := c.WithNoAuth().Do("PUT", test.UrlPattern, nonExistingOrder(), returnRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})

	t.Run("UPDATERETURN003", func(t *testing.T) {
		resp, err := c.WithAuth(&AuthenticationInfo{
			HeaderName: cfg.Auth.HeaderName,
			ApiKey:     "invalid-api-key",
		}).Do("PUT", test.UrlPattern, nonExistingOrderVars(), returnRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})

	t.Run("UPDATERETURN004", func(t *testing.T) {
		resp, err := c.Do("PUT", test.UrlPattern, nonExistingOrderVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("UPDATERETURN005", func(t *testing.T) {
		resp, err := c.Do("PUT", test.UrlPattern, nonExistingOrderVars(), returnRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("UPDATERETURN006", func(t *testing.T) {
		resp, err := c.WithNoAuth().Do("PUT", test.UrlPattern, nonExistingReturnId(), returnRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})
}

func nonExistingReturnId() map[string]string {
	return map[string]string{
		"customer_printed_order_id": "simple_order_1",
		"return_id": "non-existing-return",
	}
}