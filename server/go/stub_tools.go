package server

import (
	"fmt"
	"time"

	"github.com/itsrever/integration/tools"
	"github.com/shopspring/decimal"
)

const customerEmail = "test@itsrever.com"

// NewMultiMoney is an auxiliary function to help with the stubs
// and create a MultiMoney with the given amount and currency
func NewMultiMoney(amountShop float64, currencyShop string,
	amountCustomer float64, currencyCustomer string) IntegrationMultiMoney {
	return IntegrationMultiMoney{
		AmountShop: IntegrationMultiMoneyAmountShop{
			Amount:   amountShop,
			Currency: currencyShop,
		},
		AmountCustomer: IntegrationMultiMoneyAmountCustomer{
			Amount:   amountCustomer,
			Currency: currencyCustomer,
		},
	}
}

// shippingDetails returns a generic shipping of the given amount in the shop currency
// TODO: exchange if the Currency customer is different and
// more generic approach to pass currencies
func shippingDetails(costWithTax float64, taxRate float64) IntegrationShipping {
	decimalTaxRate := decimal.NewFromFloat(taxRate)
	decimalOppositeTaxRate := decimal.NewFromFloat(1 - taxRate)
	cost := decimal.NewFromFloat(costWithTax).Mul(decimalOppositeTaxRate).RoundBank(2).InexactFloat64()
	tax := decimal.NewFromFloat(costWithTax).Mul(decimalTaxRate).RoundBank(2).InexactFloat64()

	return IntegrationShipping{
		Amount: NewMultiMoney(cost, CurrencyShop, cost, CurrencyCustomer),
		Taxes:  NewMultiMoney(tax, CurrencyShop, tax, CurrencyCustomer),
	}
}

// payOrder returns the order with a payment for the total due
func payOrder(order *IntegrationOrder) *IntegrationOrder {
	amount := order.TotalAmount.AmountCustomer.Amount
	currency := order.TotalAmount.AmountCustomer.Currency
	order.Payment = IntegrationPayment{
		Date: orderDate(),
		Transactions: []IntegrationTransaction{
			{
				PaymentMethodType: "non-cash",
				TransactionId:     tools.RandomString(10),
				Amount: IntegrationTransactionAmount{
					Amount:   amount,
					Currency: currency,
				},
				Date: orderDate(),
			},
		},
	}
	return order
}

// billingAddress returns an arbitrary billing address
func billingAddress() IntegrationOrderBillingAddress {
	return IntegrationOrderBillingAddress{
		FirstName:     "John",
		LastName:      "Doe",
		AddressLine1:  "1234 Main Street",
		City:          "Anytown",
		Postcode:      "123456",
		Phone:         "555-123-4567",
		StateProvince: "California",
		Country:       "United States",
		CountryCode:   "US",
		Email:         customerEmail,
		Company:       "Test Company",
	}
}

// shippingAddress returns an arbitrary shipping address
func shippingAddress() IntegrationOrderShippingAddress {
	return IntegrationOrderShippingAddress{
		FirstName:     "John",
		LastName:      "Doe",
		AddressLine1:  "1234 Main Street",
		City:          "Anytown",
		Postcode:      "123456",
		Phone:         "555-123-4567",
		StateProvince: "California",
		Country:       "United States",
		CountryCode:   "US",
		Email:         customerEmail,
		Company:       "Test Company",
	}
}

// customerDetails returns an arbitrary customer
func customerDetails() IntegrationCustomer {
	return IntegrationCustomer{
		FirstName:     "John",
		LastName:      "Doe",
		Email:         customerEmail,
		PreferredLang: "es",
	}
}

// identification returns the IntegrationIdentification section
// corresponding to the given order id
func identification(orderID string) IntegrationIdentification {
	return IntegrationIdentification{
		// long format, internal ID
		Id:                     fmt.Sprintf("order-%s", orderID),
		CustomerPrintedOrderId: orderID,
	}
}

func orderDate() time.Time {
	return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
}

func emptyOrderReturns() []IntegrationReturnOrder {
	return []IntegrationReturnOrder{}
}

// calculateTotals, given the order line items and shipping,
// calculates the totals (tax, total) and returns the order
func calculateTotals(order *IntegrationOrder) *IntegrationOrder {
	totalCustomer := float64(0)
	totalShop := float64(0)
	totalTaxCustomer := float64(0)
	totalTaxShop := float64(0)
	for _, li := range order.LineItems {
		totalCustomer += li.Total.AmountCustomer.Amount
		totalShop += li.Total.AmountShop.Amount
		totalTaxCustomer += li.TotalTaxes.AmountCustomer.Amount
		totalTaxShop += li.TotalTaxes.AmountShop.Amount
	}
	totalCustomer += order.Shipping.Amount.AmountCustomer.Amount
	totalShop += order.Shipping.Amount.AmountShop.Amount
	totalTaxCustomer += order.Shipping.Taxes.AmountCustomer.Amount
	totalTaxShop += order.Shipping.Taxes.AmountShop.Amount

	order.TotalAmount = IntegrationOrderTotalAmount{
		AmountShop: IntegrationMultiMoneyAmountShop{
			Amount:   decimal.NewFromFloat(totalShop).RoundBank(2).InexactFloat64(),
			Currency: CurrencyShop,
		},
		AmountCustomer: IntegrationMultiMoneyAmountCustomer{
			Amount:   decimal.NewFromFloat(totalShop).RoundBank(2).InexactFloat64(),
			Currency: CurrencyCustomer,
		},
	}

	order.TotalTaxes = IntegrationOrderTotalTaxes{
		AmountShop: IntegrationMultiMoneyAmountShop{
			Amount:   decimal.NewFromFloat(totalTaxShop).RoundBank(2).InexactFloat64(),
			Currency: CurrencyShop,
		},
		AmountCustomer: IntegrationMultiMoneyAmountCustomer{
			Amount:   decimal.NewFromFloat(totalTaxCustomer).RoundBank(2).InexactFloat64(),
			Currency: CurrencyCustomer,
		},
	}
	return order
}

func fulfillOrder(order *IntegrationOrder) *IntegrationOrder {
	fOrder := IntegrationFulfillmentOrder{
		LocationId:   tools.RandomString(10),
		Date:         orderDate(),
		Fulfillments: []IntegrationFulfillment{},
	}
	for _, li := range order.LineItems {
		fOrder.Fulfillments = append(fOrder.Fulfillments,
			IntegrationFulfillment{
				LineItemId: li.Id,
				Quantity:   li.Quantity,
			},
		)
	}
	order.FulfillmentOrders = []IntegrationFulfillmentOrder{fOrder}
	return order
}
