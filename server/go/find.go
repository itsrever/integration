package server

func FindResponses() map[string]*IntegrationOrder {
	return map[string]*IntegrationOrder{
		"simple_order_1": OrderWithProduct(),
	}
}

func FindResponseFor(id string) *IntegrationOrder {
	return FindResponses()[id]
}

func OrderWithProduct() *IntegrationOrder {
	return &IntegrationOrder{
		BillingAddress: IntegrationOrderBillingAddress{
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
		},
		ShippingAddress: IntegrationOrderShippingAddress{
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
		},

		Customer: IntegrationCustomer{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@test.com",
		},

		Identification: IntegrationIdentification{
			Id:                     "123456",
			CustomerPrintedOrderId: "#123456",
		},

		TaxesIncluded: true,
		TotalAmount: IntegrationOrderTotalAmount{
			AmountShop: IntegrationMultiMoneyAmountShop{
				Amount:   1250,
				Currency: "USD",
			},
			AmountCustomer: IntegrationMultiMoneyAmountCustomer{
				Amount:   1250,
				Currency: "USD",
			},
		},
		LineItems: []IntegrationLineItem{
			{
				Subtotal: IntegrationLineItemSubtotal{
					AmountShop: IntegrationMultiMoneyAmountShop{
						Amount:   1250,
						Currency: "USD",
					},
					AmountCustomer: IntegrationMultiMoneyAmountCustomer{
						Amount:   1250,
						Currency: "USD",
					},
				},
				Total: IntegrationLineItemTotal{
					AmountShop: IntegrationMultiMoneyAmountShop{
						Amount:   1250,
						Currency: "USD",
					},
					AmountCustomer: IntegrationMultiMoneyAmountCustomer{
						Amount:   1250,
						Currency: "USD",
					},
				},
				Id:        "lineitem1",
				Quantity:  1,
				Name:      "Product 1",
				VariantId: "variant1",
				TotalDiscounts: IntegrationLineItemTotalDiscounts{
					AmountShop: IntegrationMultiMoneyAmountShop{
						Amount:   0,
						Currency: "USD",
					},
					AmountCustomer: IntegrationMultiMoneyAmountCustomer{
						Amount:   0,
						Currency: "USD",
					},
				},
				TotalTaxes: IntegrationLineItemTotalTaxes{
					AmountShop: IntegrationMultiMoneyAmountShop{
						Amount:   0,
						Currency: "USD",
					},
					AmountCustomer: IntegrationMultiMoneyAmountCustomer{
						Amount:   0,
						Currency: "USD",
					},
				},
				UnitPrice: IntegrationLineItemUnitPrice{
					AmountShop: IntegrationMultiMoneyAmountShop{
						Amount:   1250,
						Currency: "USD",
					},
					AmountCustomer: IntegrationMultiMoneyAmountCustomer{
						Amount:   1250,
						Currency: "USD",
					},
				},
				Product: IntegrationProduct{
					Id:                "product1",
					Name:              "Product 1",
					Description:       "Product 1 description",
					Sku:               "skutestproduct",
					InventoryQuantity: 14,
					Price: IntegrationProductPrice{
						Amount:   1250,
						Currency: "USD",
					},
					Tags: []IntegrationTag{
						{
							Name: "Tag 1",
						},
						{
							Name: "Tag 2",
						},
					},
					Variants: []IntegrationVariant{
						{
							Id:               "variant1",
							Name:             "Variant 1",
							Description:      "Variant 1 description",
							ShortDescription: "Variant 1 short description",
							Enabled:          true,
							Sku:              "skutest",
							Weight:           1000,
							Price: IntegrationVariantPrice{
								Amount:   1250,
								Currency: "USD",
							},
							Images: []IntegrationImage{
								{
									Src:  "https://clientes.oxfamintermon.org/643-medium_default/camiseta-hombre-lisa-algorg-blanca-s.jpg",
									Name: "Camiseta hombre lisa algorg blanca S",
									Alt:  "Camiseta hombre lisa algorg blanca S",
								},
							},
							InventoryQuantity: 10,
							Options: []IntegrationOption{
								{
									Name:  "Option 1",
									Value: "Option 1 value",
								},
							},
						},
					},
				},
			},
		},
	}
}
