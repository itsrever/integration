package test

import (
	"strings"
	"testing"

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

		// TODO: assert order
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
