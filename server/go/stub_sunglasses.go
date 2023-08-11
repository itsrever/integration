package server

import (
	"github.com/shopspring/decimal"
)

type Decimal = decimal.Decimal

const sunglassesDescription = "Wonderful sunglasses to get around the city, looking for victims. One size."
const sunglassesProductName = "Wonderful Vampire (WV) Sunglasses"
const sunglassesUnitPrice = 21.55
const sunglassesTaxRate = 0.21
const sunglassesDiscountRate = 0.12
const sunglassesQuantity = 1
const sunglassesCurrency = CurrencyEUR

func sunglassesLineItem() IntegrationLineItem {
	decUnitPrice := decimal.NewFromFloat(sunglassesUnitPrice)
	decTaxRate := decimal.NewFromFloat(sunglassesTaxRate)
	decDiscount := decimal.NewFromFloat(sunglassesDiscountRate)
	decQuantity := decimal.NewFromInt(sunglassesQuantity)

	unitPrice := decUnitPrice.InexactFloat64()
	subtotal := decUnitPrice.Mul(decQuantity).InexactFloat64()
	discounts := decimal.NewFromFloat(subtotal).Mul(decDiscount).RoundBank(2).InexactFloat64()
	taxes := decimal.NewFromFloat(subtotal - discounts).Mul(decTaxRate).RoundBank(2).InexactFloat64()
	total := decimal.NewFromFloat(subtotal - discounts + taxes).RoundBank(2).InexactFloat64()

	return IntegrationLineItem{
		Quantity: int32(decQuantity.IntPart()),
		UnitPrice: IntegrationLineItemUnitPrice{
			AmountShop: IntegrationMultiMoneyAmountShop{
				Amount:   unitPrice,
				Currency: sunglassesCurrency,
			},
			AmountCustomer: IntegrationMultiMoneyAmountCustomer{
				Amount:   unitPrice,
				Currency: sunglassesCurrency,
			},
		},
		Subtotal: IntegrationLineItemSubtotal{
			AmountShop: IntegrationMultiMoneyAmountShop{
				Amount:   subtotal,
				Currency: sunglassesCurrency,
			},
			AmountCustomer: IntegrationMultiMoneyAmountCustomer{
				Amount:   subtotal,
				Currency: sunglassesCurrency,
			},
		},
		TotalDiscounts: IntegrationLineItemTotalDiscounts{
			AmountShop: IntegrationMultiMoneyAmountShop{
				Amount:   discounts,
				Currency: sunglassesCurrency,
			},
			AmountCustomer: IntegrationMultiMoneyAmountCustomer{
				Amount:   discounts,
				Currency: sunglassesCurrency,
			},
		},
		TotalTaxes: IntegrationLineItemTotalTaxes{
			AmountShop: IntegrationMultiMoneyAmountShop{
				Amount:   taxes,
				Currency: sunglassesCurrency,
			},
			AmountCustomer: IntegrationMultiMoneyAmountCustomer{
				Amount:   taxes,
				Currency: sunglassesCurrency,
			},
		},
		Total: IntegrationLineItemTotal{
			AmountShop: IntegrationMultiMoneyAmountShop{
				Amount:   total,
				Currency: sunglassesCurrency,
			},
			AmountCustomer: IntegrationMultiMoneyAmountCustomer{
				Amount:   total,
				Currency: sunglassesCurrency,
			},
		},
		Id:      "lineitem1",
		Name:    sunglassesProductName,
		Product: sunglassesProduct(),
	}
}

func sunglassesProduct() IntegrationProduct {
	return IntegrationProduct{
		Id:                "sun_wv",
		Name:              sunglassesProductName,
		Description:       sunglassesDescription,
		Sku:               "sku_sun_vw",
		InventoryQuantity: 14,
		UnitPrice: IntegrationProductUnitPrice{
			Amount:   sunglassesUnitPrice,
			Currency: sunglassesCurrency,
		},
		Images: []IntegrationImage{
			{
				Name: "Red sunglasses front",
				Src:  "https://mypartyshirt.com/media/catalog/product/cache/1/image/1000x1231/9df78eab33525d08d6e5fb8d27136e95/r/e/red-vampire-sunglasses1.jpg",
				Alt:  "Image 1",
			},
		},
		Tags: []IntegrationTag{
			{
				Name: "One size",
			},
			{
				Name: "Sale",
			},
		},
	}
}
