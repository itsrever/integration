package server

// tShirtVariants returns the variants for the t-shirt product
func tShirtVariants() []IntegrationVariant {
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
				Currency: CurrencyUSD,
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
