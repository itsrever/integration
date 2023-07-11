package server

import "time"

const currencyUSD = "USD"

func FindResponses() map[string]*IntegrationOrder {
	return map[string]*IntegrationOrder{
		"simple_order_1": OrderWithSingleProductUSD(),
		"simple_order_2": GetSimpleOrderWithVartiantsInUSD(),
	}
}

func FindResponseFor(id string) *IntegrationOrder {
	return FindResponses()[id]
}

func GetSimpleOrderWithVartiantsInUSD() *IntegrationOrder {
	integrationOrder := OrderWithSingleProductUSD()
	integrationOrder.Identification.CustomerPrintedOrderId = "simple_order_2"
	integrationOrder.LineItems[0].Product.Variants = getVariants()
	return integrationOrder
}

func getVariants() []IntegrationVariant {
	return []IntegrationVariant{
		{
			Id:                "gopher-black-s",
			Enabled:           true,
			InventoryQuantity: 10,
			Name:              "Gopher Black S",
			Sku:               "gopher-black-s",
			Weight:            100,
			ShortDescription:  "Gopher Black S",
			Description:       "Gopher Black S",
			UnitPrice: IntegrationVariantUnitPrice{
				Amount:   10,
				Currency: currencyUSD,
			},

			Images: []IntegrationImage{
				{
					Src:  "https://magento.byrever.com/media/catalog/product/cache/584aced4a1dec0308dc2dca447b4d064/t/e/teegolang.jpg",
					Name: "Tee Golang",
					Alt:  "Tee Golang",
				},
			},
			Options: []IntegrationOption{
				{
					Name:  "Color",
					Value: "Black",
				},
				{
					Name:  "Size",
					Value: "S",
				},
			},
		},
	}
}

// CurrencyShop is the currency used by the shop, always EUR in all cases
const CurrencyShop = "EUR"

func OrderWithSingleProductUSD() *IntegrationOrder {
	return &IntegrationOrder{
		Identification:    getIdentificationDetails(),
		Customer:          getCustomerDetails(),
		BillingAddress:    getBillingAddress(),
		ShippingAddress:   getShippingAddress(),
		LineItems:         getLineItems(),
		TotalAmount:       getTotalAmountInUSD(),
		TotalTaxes:        getTotalTaxesInUSD(),
		Shipping:          getShippingDetails(),
		Payment:           getPaymentDetails(),
		FulfillmentOrders: getFulfillmentOrders(),
		Returns:           getReturns(),
		TaxesIncluded:     true,
		Date:              getOrderDate(),
	}
}

func getOrderDate() time.Time {
	return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
}

func getReturns() []IntegrationReturnOrder {
	return []IntegrationReturnOrder{
		{
			ReturnId:    "return1",
			Description: "Test return",
			Date:        getOrderDate(),
			Returns: []IntegrationReturn{
				{
					LineItemId: "lineitem1",
					Quantity:   1,
				},
			},
		},
	}
}

func getTotalAmountInUSD() IntegrationOrderTotalAmount {
	return IntegrationOrderTotalAmount{
		AmountShop: IntegrationMultiMoneyAmountShop{
			Amount:   45,
			Currency: currencyUSD,
		},
		AmountCustomer: IntegrationMultiMoneyAmountCustomer{
			Amount:   45,
			Currency: currencyUSD,
		},
	}
}

func getTotalTaxesInUSD() IntegrationOrderTotalTaxes {
	return IntegrationOrderTotalTaxes{
		AmountShop: IntegrationMultiMoneyAmountShop{
			Amount:   5,
			Currency: currencyUSD,
		},
		AmountCustomer: IntegrationMultiMoneyAmountCustomer{
			Amount:   5,
			Currency: currencyUSD,
		},
	}
}

func getShippingDetails() IntegrationShipping {
	return IntegrationShipping{
		Amount: IntegrationShippingAmount{
			AmountShop: IntegrationMultiMoneyAmountShop{
				Amount:   5,
				Currency: currencyUSD,
			},
			AmountCustomer: IntegrationMultiMoneyAmountCustomer{
				Amount:   5,
				Currency: currencyUSD,
			},
		},
	}
}

func getFulfillmentOrders() []IntegrationFulfillmentOrder {
	return []IntegrationFulfillmentOrder{
		{
			LocationId: "123456",
			Date:       getOrderDate(),
			Fulfillments: []IntegrationFulfillment{
				{
					LineItemId: "lineitem1",
					Quantity:   1,
				},
			},
		},
	}
}

func getPaymentDetails() IntegrationPayment {
	return IntegrationPayment{
		Date: getOrderDate(),
		Transactions: []IntegrationTransaction{
			{
				PaymentMethodType: "non-cash",
				TransactionId:     "123456",
				Amount: IntegrationTransactionAmount{
					Amount:   45,
					Currency: currencyUSD,
				},
				Date: getOrderDate(),
			},
		},
	}
}

func getBillingAddress() IntegrationOrderBillingAddress {
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
		Email:         "test@tets.com",
		Company:       "Test Company",
	}
}

func getShippingAddress() IntegrationOrderShippingAddress {
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
		Email:         "test@tets.com",
		Company:       "Test Company",
	}
}

func getCustomerDetails() IntegrationCustomer {
	return IntegrationCustomer{
		FirstName:     "John",
		LastName:      "Doe",
		Email:         "test@test.com",
		PreferredLang: "ES",
	}
}

func getIdentificationDetails() IntegrationIdentification {
	return IntegrationIdentification{
		Id:                     "123456",
		CustomerPrintedOrderId: "simple_order_1",
	}
}

func getLineItems() []IntegrationLineItem {
	return []IntegrationLineItem{
		{
			Subtotal: IntegrationLineItemSubtotal{
				AmountShop: IntegrationMultiMoneyAmountShop{
					Amount:   50,
					Currency: currencyUSD,
				},
				AmountCustomer: IntegrationMultiMoneyAmountCustomer{
					Amount:   50,
					Currency: currencyUSD,
				},
			},
			Total: IntegrationLineItemTotal{
				AmountShop: IntegrationMultiMoneyAmountShop{
					Amount:   45,
					Currency: currencyUSD,
				},
				AmountCustomer: IntegrationMultiMoneyAmountCustomer{
					Amount:   45,
					Currency: currencyUSD,
				},
			},
			Id:        "lineitem1",
			Quantity:  1,
			Name:      "Product 1",
			VariantId: "variant1",
			TotalDiscounts: IntegrationLineItemTotalDiscounts{
				AmountShop: IntegrationMultiMoneyAmountShop{
					Amount:   10,
					Currency: currencyUSD,
				},
				AmountCustomer: IntegrationMultiMoneyAmountCustomer{
					Amount:   10,
					Currency: currencyUSD,
				},
			},
			TotalTaxes: IntegrationLineItemTotalTaxes{
				AmountShop: IntegrationMultiMoneyAmountShop{
					Amount:   5,
					Currency: currencyUSD,
				},
				AmountCustomer: IntegrationMultiMoneyAmountCustomer{
					Amount:   5,
					Currency: currencyUSD,
				},
			},
			UnitPrice: IntegrationLineItemUnitPrice{
				AmountShop: IntegrationMultiMoneyAmountShop{
					Amount:   50,
					Currency: currencyUSD,
				},
				AmountCustomer: IntegrationMultiMoneyAmountCustomer{
					Amount:   50,
					Currency: currencyUSD,
				},
			},
			Product: getIntegrationProduct(),
		},
	}
}

func getIntegrationProduct() IntegrationProduct {
	return IntegrationProduct{
		Id:                "product1",
		Name:              "Product 1",
		Description:       "Product 1 description",
		Sku:               "skutestproduct",
		InventoryQuantity: 14,
		UnitPrice: IntegrationProductUnitPrice{
			Amount:   50,
			Currency: currencyUSD,
		},
		Images: []IntegrationImage{
			{
				Name: "Image 1",
				Src:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
				Alt:  "Image 1",
			},
		},
		Tags: []IntegrationTag{
			{
				Name: "Tag 1",
			},
			{
				Name: "Tag 2",
			},
		},
	}
}
