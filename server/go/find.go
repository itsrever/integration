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
		Identification:    getIdentificationDetails(),
		Customer:          getCustomerDetails(),
		BillingAddress:    getBillingAddress(),
		ShippingAddress:   getShippingAddress(),
		LineItems:         getLineItems(),
		TotalAmount:       getTotalAmountInUSD(),
		TotalTaxes:        getTotalTaxesInEUR(),
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
			Currency: "USD",
		},
		AmountCustomer: IntegrationMultiMoneyAmountCustomer{
			Amount:   45,
			Currency: "USD",
		},
	}
}

func getTotalTaxesInEUR() IntegrationOrderTotalTaxes {
	return IntegrationOrderTotalTaxes{
		AmountShop: IntegrationMultiMoneyAmountShop{
			Amount:   5,
			Currency: "EUR",
		},
		AmountCustomer: IntegrationMultiMoneyAmountCustomer{
			Amount:   5,
			Currency: "EUR",
		},
	}
}

func getShippingDetails() IntegrationShipping {
	return IntegrationShipping{
		Amount: IntegrationShippingAmount{
			AmountShop: IntegrationMultiMoneyAmountShop{
				Amount:   5,
				Currency: "USD",
			},
			AmountCustomer: IntegrationMultiMoneyAmountCustomer{
				Amount:   5,
				Currency: "USD",
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
					Currency: "USD",
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
					Currency: "USD",
				},
				AmountCustomer: IntegrationMultiMoneyAmountCustomer{
					Amount:   50,
					Currency: "USD",
				},
			},
			Total: IntegrationLineItemTotal{
				AmountShop: IntegrationMultiMoneyAmountShop{
					Amount:   45,
					Currency: "USD",
				},
				AmountCustomer: IntegrationMultiMoneyAmountCustomer{
					Amount:   45,
					Currency: "USD",
				},
			},
			Id:        "lineitem1",
			Quantity:  1,
			Name:      "Product 1",
			VariantId: "variant1",
			TotalDiscounts: IntegrationLineItemTotalDiscounts{
				AmountShop: IntegrationMultiMoneyAmountShop{
					Amount:   10,
					Currency: "USD",
				},
				AmountCustomer: IntegrationMultiMoneyAmountCustomer{
					Amount:   10,
					Currency: "USD",
				},
			},
			TotalTaxes: IntegrationLineItemTotalTaxes{
				AmountShop: IntegrationMultiMoneyAmountShop{
					Amount:   5,
					Currency: "USD",
				},
				AmountCustomer: IntegrationMultiMoneyAmountCustomer{
					Amount:   5,
					Currency: "USD",
				},
			},
			UnitPrice: IntegrationLineItemUnitPrice{
				AmountShop: IntegrationMultiMoneyAmountShop{
					Amount:   50,
					Currency: "USD",
				},
				AmountCustomer: IntegrationMultiMoneyAmountCustomer{
					Amount:   50,
					Currency: "USD",
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
			Currency: "USD",
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
