package server

const RETURN_ID = "return-1"

func getReturnResponse(order *Order, returnItemDetails []ReturnRequestItem) map[string]string {
	for _, returnItem := range returnItemDetails {
		//append returnItem to returns
		order.Returns = append(order.Returns, ReturnOrder{
			ReturnId: RETURN_ID,
			Returns : []Return{
				{
				LineItemId: returnItem.LineItemId,
				Quantity: returnItem.Quantity,
				Status: Status{Status: returnItem.Status.Status},
				},
			},
		})
	} 
	return map[string]string{
		"return_id": RETURN_ID,
	}
}