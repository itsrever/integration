package test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	server "github.com/itsrever/integration/server/go"
)

const schemaLocation = "./schema.json"

func TestValidateNOK(t *testing.T) {
	val, err := NewJsonValidator(schemaLocation)
	require.NoError(t, err)
	order := server.Order{
		Identification: server.Identification{
			CustomerPrintedOrderId: "123",
		},
	}
	bytes, err := json.Marshal(order)
	require.NoError(t, err)
	err = val.Validate("order", bytes)
	require.Error(t, err)
}

func TestValidateOK(t *testing.T) {
	val, err := NewJsonValidator(schemaLocation)
	require.NoError(t, err)
	order_id := []byte(`"123"`)
	err = val.Validate("customer_printed_order_id", order_id)
	require.NoError(t, err)
}
