package refund

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Refund(t *testing.T) {
	manager := New()
	assert.NotNil(t, manager)

	manager.CreateRefund("order1", createRefund())
	refunds := manager.GetRefund("order1")
	assert.NotEmpty(t, refunds)

	assert.Equal(t, 1, len(refunds.Items))
	assert.Equal(t, "line_item_1", refunds.Items[0].LineItemId)
}

func createRefund() Refund {
	return Refund{
		Items: refundItems(),
	}
}

func refundItems() []RefundRequestItem {
	return []RefundRequestItem{
		{
			LineItemId: "line_item_1",
			Quantity:   1,
			Amount: RefundRequestItemAmount{
				Amount:   10.0,
				Currency: "EUR",
			},
		},
	}
}
