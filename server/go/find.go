package server

import (
	"time"

	"github.com/itsrever/integration/server/notes"
	"github.com/itsrever/integration/server/refund"
)

const CurrencyUSD = "USD"
const CurrencyEUR = "EUR"

// CurrencyShop is the currency used by the customer
const CurrencyShop = CurrencyEUR

// CurrencyShop is the currency used by the customer
const CurrencyCustomer = CurrencyEUR

// SingleOrderWithoutVariantsID is the ID of the order that returns a
// simple order without variants, the sunglasses
const SingleOrderWithoutVariantsID = "simple_order_1"

// SingleOrderWithVariantsID is the ID of the order that returns a
// simple order with variants, the t-shirt
const SingleOrderWithVariantsID = "simple_order_2"

// responsesByID returns a map of orders, indexed by order_id
func responsesByID() map[string]*Order {
	return map[string]*Order{
		SingleOrderWithoutVariantsID: OrderWithSingleProduct(SingleOrderWithoutVariantsID),
		SingleOrderWithVariantsID:    SimpleOrderWithVariants(SingleOrderWithVariantsID),
	}
}

// FindOrderFor returns the order corresponding to the given id
func FindOrderFor(id string) *Order {
	order := responsesByID()[id]
	if order != nil {
		order.Notes = toNotes(notes.New().GetNotesFromOrder(id))
	}
	return order
}

func FindOrderWithRefunds(id string, refunds refund.Refund) *Order {
	order := responsesByID()[id]
	if order != nil {
		order.Refunds = toRefundOrder(refunds.Items)
	}
	return order
}

func toNotes(notes []notes.Note) []Note {
	var Notes []Note
	for _, note := range notes {
		Notes = append(Notes, Note{
			Text: note.Text,
			Date: note.Date,
			User: "unknown",
		})
	}
	return Notes
}

func toRefundOrder(items []refund.RefundRequestItem) []RefundOrder {
	return []RefundOrder{
		{
			Refunds:     toRefund(items),
			Description: "Refund for order",
			Amount: RefundOrderAmount{
				Amount:   calculateItemsTotalAmount(items),
				Currency: CurrencyShop,
			},
			Date:          time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			RefundId:      "refund_id",
			TransactionId: "transaction_id",
		},
	}
}

func toRefund(items []refund.RefundRequestItem) []Refund {
	var Refunds []Refund
	for _, item := range items {
		Refunds = append(Refunds, Refund{
			LineItemId: item.LineItemId,
			Quantity:   item.Quantity,
		})
	}
	return Refunds
}

func calculateItemsTotalAmount(items []refund.RefundRequestItem) float64 {
	var totalAmount float64
	for _, item := range items {
		totalAmount += item.Amount.Amount
	}
	return totalAmount
}
