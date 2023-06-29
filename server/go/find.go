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
		LineItems: []IntegrationLineItem{
			{
				Product: IntegrationProduct{
					Id:          "product1",
					Name:        "Product 1",
					Description: "Product 1 description",
					Price: IntegrationProductPrice{
						Amount:   1250,
						Currency: "USD",
					},
				},
			},
		},
	}
}
