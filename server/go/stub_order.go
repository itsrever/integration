package server

// SimpleOrderWithVariants returns the order corresponding to the given id
func SimpleOrderWithVariants(orderID string) *Order {
	Order := OrderWithSingleProduct(orderID)
	Order.LineItems[0].Product.Variants = tShirtVariants()
	return Order
}

// OrderWithSingleProduct returns a simple order without variants, the vampire sunglasses
func OrderWithSingleProduct(orderID string) *Order {
	order := &Order{
		Identification:  identification(orderID),
		Customer:        customerDetails(),
		BillingAddress:  billingAddress(),
		ShippingAddress: shippingAddress(),
		LineItems:       []LineItem{sunglassesLineItem()},
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
