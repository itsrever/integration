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
func responsesByID() map[string]*IntegrationOrder {
	return map[string]*IntegrationOrder{
		SingleOrderWithoutVariantsID: OrderWithSingleProduct(SingleOrderWithoutVariantsID),
		SingleOrderWithVariantsID:    SimpleOrderWithVariants(SingleOrderWithVariantsID),
	}
}

// FindOrderFor returns the order corresponding to the given id
func FindOrderFor(id string) *IntegrationOrder {
	order := responsesByID()[id]
	if order != nil {
		order.Notes = toIntegrationNotes(notes.New().GetNotesFromOrder(id))
	}
	return order
}

func toIntegrationNotes(notes []notes.Note) []IntegrationNote {
	var integrationNotes []IntegrationNote
	for _, note := range notes {
		integrationNotes = append(integrationNotes, IntegrationNote{
			Text: note.Text,
			Date: note.Date,
			User: "unknown",
		})
	}
	return integrationNotes
}
