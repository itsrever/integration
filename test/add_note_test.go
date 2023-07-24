package test

import (
	"math/rand"
	"testing"
	"time"

	server "github.com/itsrever/integration/server/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Add_Note_Into_Order(t *testing.T) {
	cfg, err := configFromEnv()
	require.NoError(t, err)
	c := NewClient(cfg.BaseURL).WithAuth(cfg.Auth)
	require.NoError(t, err)
	test := cfg.Test("AddNoteToOrder")
	testFindOrder := cfg.Test("FindOrderByCustomerOrderPrintedId")
	val, err := NewJsonValidator(schemaLocation)
	require.NoError(t, err)
	noteBody := server.AddNoteToOrderRequest{
		Note: "Note" + generateRandomString(10),
	}

	t.Run("ADDNOTE00", func(t *testing.T) {
		resp, err := c.WithNoAuth().Do("POST", test.UrlPattern, nonExistingOrderVars(), noteBody)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})

	t.Run("ADDNOTE01", func(t *testing.T) {
		resp, err := c.WithAuth(&AuthenticationInfo{
			HeaderName: cfg.Auth.HeaderName,
			ApiKey:     "invalid-api-key",
		}).Do("POST", test.UrlPattern, nonExistingOrderVars(), noteBody)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 401, resp.StatusCode)
	})
	t.Run("ADDNOTE02", func(t *testing.T) {
		resp, err := c.Do("POST", test.UrlPattern, nonExistingOrderVars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("ADDNOTE03", func(t *testing.T) {
		resp, err := c.Do("POST", test.UrlPattern, nonExistingOrderVars(), noteBody)
		println(resp.Request.URL.String())
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 404, resp.StatusCode)
	})

	t.Run("ADDNOTE04", func(t *testing.T) {
		test.FailTestIfScenarioNotPresent(t, testName(t))
		scenario := test.Scenario(testName(t))
		resp, err := c.Do("POST", test.UrlPattern, scenario.Vars(), noteBody)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)

		// recover the order and see that the note has been added
		resp, err = c.Do("GET", testFindOrder.UrlPattern, scenario.Vars(), nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.StatusCode)
		body := requireBodyFromResponse(t, resp)
		val.RequireModel(t, "integration.Order", body)
		order, err := orderFromBody(body)
		require.NoError(t, err)
		assertSanity(t, order)
		assertOrderHasNote(t, order, noteBody.Note)
	})
}

func generateRandomString(length int) string {
	var letters = []rune("bcdfghijklmnpqrstwxyz")
	rand.Seed(time.Now().UTC().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func assertOrderHasNote(t *testing.T, order *server.IntegrationOrder, text string) {
	assert.NotEmpty(t, order.Notes)
	for _, note := range order.Notes {
		if note.Text == text {
			assert.NotEmpty(t, note.Date)
			return
		}
	}
	t.Errorf("Note with text \"%s\" not found", text)
}
