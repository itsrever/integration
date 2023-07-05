package server

import "time"

func FindResponses() map[string]*IntegrationOrder {
	return map[string]*IntegrationOrder{
		"simple_order_1": OrderWithSingleProductEUR(),
	}
}

func FindResponseFor(id string) *IntegrationOrder {
	return FindResponses()[id]
}

// CurrencyShop is the currency used by the shop, always EUR in all cases
const CurrencyShop = "EUR"

func OrderWithSingleProductEUR() *IntegrationOrder {
	return &IntegrationOrder{
		Date:          LastWeek(),
		TaxesIncluded: true,
		TotalAmount: IntegrationOrderTotalAmount{
			AmountShop: IntegrationMultiMoneyAmountShop{
				Amount:   15.13,
				Currency: CurrencyShop,
			},
			AmountCustomer: IntegrationMultiMoneyAmountCustomer{
				Amount:   15.13,
				Currency: ProductSunglasses().UnitPrice.Currency,
			},
		},
		TotalTaxes: IntegrationOrderTotalTaxes{},
		LineItems: []IntegrationLineItem{
			{
				Product: ProductSunglasses(),
			},
		},
		Shipping: OrderShipping(),
	}
}

func LastWeek() time.Time {
	return time.Now().Add(-time.Hour * 24 * 7)
}

// ProductSunglasses is a product with 21% VAT in EUR, with no variants
func ProductSunglasses() IntegrationProduct {
	return IntegrationProduct{
		Id:          "sunglasses_product_id",
		Name:        "Sunglasses",
		Description: "Sunglasses long description",
		UnitPrice: IntegrationProductUnitPrice{
			Amount:   12.50,
			Currency: CurrencyShop,
		},
	}
}

func OrderShipping() IntegrationShipping {
	return IntegrationShipping{
		Amount: IntegrationShippingAmount{
			AmountShop: IntegrationMultiMoneyAmountShop{
				Amount:   12.40,
				Currency: "EUR",
			},
			AmountCustomer: IntegrationMultiMoneyAmountCustomer{
				Amount:   12.40,
				Currency: "EUR",
			},
		},
	}
}
