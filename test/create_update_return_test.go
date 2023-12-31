package test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	server "github.com/itsrever/integration/server/go"
)

const RETURN_ID = "return-1"

func Test_Create_Return(t *testing.T) {
	cfg, err := configFromEnv()
	require.NoError(t, err)
	c := clientFromConfig(cfg)
	test := cfg.Test("CreateReturn")
	if test == nil {
		t.Skip("Test CreateReturn not found. Skiping...")
	}
	returnRequest := server.ReturnRequest{
		Returns: []server.ReturnRequestItem{
			{
				LineItemId: "testing1",
				Quantity:   1,
				Status:     server.Status{Status: "APPROVED"},
			},
		},
	}

	t.Run("CREATERETURN001", func(t *testing.T) {
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("POST", test.UrlPattern, scenario.Vars(), returnRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
		body := getBodyFromResponse(t, resp)
		returnId, _ := getReturnIdfromResponseBody(body)
		assert.Equal(t, RETURN_ID, returnId)
	})

	t.Run("CREATERETURN002", func(t *testing.T) {
		resp, err := c.WithNoAuth().Do("POST", test.UrlPattern, nonExistingOrder(), returnRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})

	t.Run("CREATERETURN003", func(t *testing.T) {
		resp, err := c.WithAuth(&ApiKeyAuthInfo{
			HeaderName: cfg.ApiKeyAuth.HeaderName,
			ApiKey:     "invalid-api-key",
		}).Do("POST", test.UrlPattern, nonExistingOrderVars(), returnRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})

	t.Run("CREATERETURN004", func(t *testing.T) {
		resp, err := c.Do("POST", test.UrlPattern, nonExistingOrderVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("CREATERETURN005", func(t *testing.T) {
		resp, err := c.Do("POST", test.UrlPattern, nonExistingOrderVars(), returnRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 404, resp.StatusCode)
	})
}

func getBodyFromResponse(t *testing.T, resp *http.Response) []byte {
	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return data
}

// getReturnIdfromResponseBody decodes the response body into an order
func getReturnIdfromResponseBody(body []byte) (string, error) {
	var responseMap map[string]string
	err := json.Unmarshal(body, &responseMap)
	return responseMap["return_id"], err
}

func nonExistingOrder() map[string]string {
	return map[string]string{
		"customer_printed_order_id": "non-existing-order",
	}
}
