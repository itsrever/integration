package test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FindOrderByCustomerOrderPrintedId(t *testing.T) {
	cfg, err := configFromEnv()
	require.NoError(t, err)
	c := NewClient(cfg.BaseURL).WithAuth(cfg.Auth).Debug()
	test := cfg.Test("FindOrderByCustomerOrderPrintedId")

	t.Run("FIND00", func(t *testing.T) {
		scenario := test.Scenario(testName(t))
		resp, err := c.WithNoAuth().Do("GET", test.UrlPattern, scenario.Vars(), nil)
		require.Error(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.StatusCode, 401)
	})

	t.Run("FIND01", func(t *testing.T) {
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("GET", test.UrlPattern, scenario.Vars(), nil)
		require.Error(t, err)
		require.NotNil(t, resp)
		require.Equal(t, resp.StatusCode, 401)
	})
}

func testName(t *testing.T) string {
	return strings.Split(t.Name(), "/")[1]
}
