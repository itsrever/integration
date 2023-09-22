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
	c := NewClient(cfg.BaseURL).WithAuth(cfg.Auth)
	if cfg.Debug {
		c = c.Debug()
	}
	test := cfg.Test("CreateOrUpdateReturn")
	returnRequest := server.ReturnRequest{
		Returns: []server.ReturnRequestItem{
			{
				LineItemId: "testing1",
				Quantity : 1,
				
			},
		},
	}

	t.Run("CREATERETURN001", func(t *testing.T) {
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("PUT", test.UrlPattern, scenario.Vars(), returnRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
		body := getBodyFromResponse(t, resp)
		returnId,err := getReturnIdfromResponseBody(body)
		assert.Equal(t, RETURN_ID, returnId)
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
