/*
 This file can be edited.
*/

package server

import (
	"context"

	"github.com/itsrever/integration/server/notes"
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
	payload := FindOrderFor(customerOrderPrintedId)
	if payload == nil {
		return Response(404, nil), nil
	}
	return Response(200, payload), nil
}

func (s *IntegrationApiService) AddNoteToOrder(ctx context.Context, orderID string, req AddNoteToOrderRequest) (ImplResponse, error) {
	if orderID == "" || req.Note == "" {
		return Response(400, nil), nil
	}
	order := FindOrderFor(orderID)
	if order == nil {
		return Response(404, nil), nil
	}

	notes.New().AddNoteToOrder(orderID, req.Note)

	return Response(200, nil), nil
}

func (s *IntegrationApiService) CreateOrUpdateReturn(ctx context.Context, orderID string, req ReturnRequest) (ImplResponse, error) {
	if orderID == "" || len(req.Returns) == 0 {
		return Response(400, nil), nil
	}
	order := FindOrderFor(orderID)
	if order == nil {
		return Response(404, nil), nil
	}
	payload := getReturnResponse()
	return Response(200, payload), nil
}

func (s *IntegrationApiService) CreateReturn(context.Context, string, ReturnRequest) (ImplResponse, error) {
	return Response(200, nil), nil
}
func (s *IntegrationApiService) CreateRefund(context.Context, string, RefundRequest) (ImplResponse, error) {
	return Response(200, nil), nil
}
