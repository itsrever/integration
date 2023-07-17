/*
 This file can be edited.
*/

package server

import (
	"context"
)

// IntegrationApiService is a service that implements the logic for the IntegrationApiServicer
// This service should implement the business logic for every endpoint for the IntegrationApi API.
// Include any external packages or services that will be required by this service.
type IntegrationApiService struct {
}

// NewIntegrationApiService creates a default api service
func NewIntegrationApiService() IntegrationApiServicer {
	return &IntegrationApiService{}
}

// FindOrderByCustomerPrintedOrderId - Find Order by customer_order_id
func (s *IntegrationApiService) FindOrderByCustomerPrintedOrderId(ctx context.Context, customerOrderPrintedId string) (ImplResponse, error) {
	if customerOrderPrintedId == "" {
		return Response(400, nil), nil
	}
	payload := FindResponseFor(customerOrderPrintedId)
	if payload == nil {
		return Response(404, nil), nil
	}
	return Response(200, payload), nil
}

func (s *IntegrationApiService) AddNoteToOrder(ctx context.Context, orderID string, note AddNoteToOrderRequest) (ImplResponse, error) {
	if orderID == "" {
		return Response(400, nil), nil
	}

	payload := FindResponseFor(orderID)
	if payload == nil {
		return Response(404, nil), nil
	}

	return Response(200, nil), nil
}
