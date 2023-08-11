package server

import (
	"github.com/itsrever/integration/server/notes"
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
