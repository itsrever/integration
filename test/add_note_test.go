package test

import (
	"math/rand"
	"testing"
	"time"

	server "github.com/itsrever/integration/server/go"
	"github.com/stretchr/testify/require"
)

func Test_Add_Note_Into_Order(t *testing.T) {
	cfg, err := configFromEnv()
	require.NoError(t, err)
	c := NewClient(cfg.BaseURL).WithAuth(cfg.Auth)
	require.NoError(t, err)
	test := cfg.Test("AddNoteToOrder")
	noteBody := server.AddNoteToOrderRequest{
		Note: "Note" + GenerateRandomString(10),
	}

	t.Run("ADDNOTE00", func(t *testing.T) {
		resp, err := c.WithNoAuth().Do("POST", test.UrlPattern, nil, noteBody)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})

	t.Run("ADDNOTE01", func(t *testing.T) {
		resp, err := c.WithAuth(&AuthenticationInfo{
			HeaderName: cfg.Auth.HeaderName,
			ApiKey:     "invalid-api-key",
		}).Do("POST", test.UrlPattern, nil, noteBody)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})
	t.Run("ADDNOTE02", func(t *testing.T) {
		resp, err := c.Do("POST", test.UrlPattern, emptyVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("ADDNOTE03", func(t *testing.T) {
		resp, err := c.Do("POST", test.UrlPattern, map[string]string{"order_id": "non-existing-order"}, noteBody)
		println(resp.Request.URL.String())
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("ADDNOTE04", func(t *testing.T) {
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("POST", test.UrlPattern, scenario.Vars(), noteBody)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
	})
}

func GenerateRandomString(length int) string {
	var letters = []rune("testuserndjdismaluehsmsldjd8a6egenidhnalsoduenkdoshsyabdlieee")
	rand.Seed(time.Now().UTC().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
