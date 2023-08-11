package server

// SimpleOrderWithVariants returns the order corresponding to the given id
func SimpleOrderWithVariants(orderID string) *IntegrationOrder {
	integrationOrder := OrderWithSingleProduct(orderID)
	integrationOrder.LineItems[0].Product.Variants = tShirtVariants()
	return integrationOrder
}

// OrderWithSingleProduct returns a simple order without variants, the vampire sunglasses
func OrderWithSingleProduct(orderID string) *IntegrationOrder {
	order := &IntegrationOrder{
		Identification:  identification(orderID),
		Customer:        customerDetails(),
		BillingAddress:  billingAddress(),
		ShippingAddress: shippingAddress(),
		LineItems:       []IntegrationLineItem{sunglassesLineItem()},
		Shipping:        shippingDetails(2.50, 0.21),
		Returns:         emptyOrderReturns(),
		TaxesIncluded:   true,
		Date:            orderDate(),
	}
	order = fulfillOrder(order)
	order = calculateTotals(order)
	order = payOrder(order)
	return order
}
