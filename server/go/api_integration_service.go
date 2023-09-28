/*
 This file can be edited.
*/

package server

import (
	"context"

	"github.com/itsrever/integration/server/notes"
	"github.com/itsrever/integration/server/refund"
)

// IntegrationApiService is a service that implements the logic for the IntegrationApiServicer
// This service should implement the business logic for every endpoint for the IntegrationApi API.
// Include any external packages or services that will be required by this service.
type IntegrationApiService struct {
	refundManager refund.RefundManager
}

// NewIntegrationApiService creates a default api service
func NewIntegrationApiService(refundManager refund.RefundManager) IntegrationApiServicer {
	return &IntegrationApiService{
		refundManager: refundManager,
	}
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

func (s *IntegrationApiService) CreateReturn(ctx context.Context, orderID string, req ReturnRequest) (ImplResponse, error) {
	if orderID == "" || len(req.Returns) == 0 {
		return Response(400, nil), nil
	}
	order := FindOrderFor(orderID)
	if order == nil {
		return Response(404, nil), nil
	}
	payload := getReturnResponse(order, req.Returns)
	return Response(200, payload), nil
}

func (s *IntegrationApiService) CreateRefund(ctx context.Context, orderID string, req RefundRequest) (ImplResponse, error) {
	if orderID == "" || req.Items == nil {
		return Response(400, nil), nil
	}
	refunds := mapRefundRequest(req)
	order := FindOrderWithRefunds(orderID, refunds)
	if order == nil {
		return Response(404, nil), nil
	}

	s.refundManager.CreateRefund(orderID, refunds)

	return Response(200, nil), nil
}

func mapRefundRequest(req RefundRequest) refund.Refund {
	items := make([]refund.RefundRequestItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = mapRefundRequestItems(item)
	}
	return refund.Refund{
		Items: items,
	}
}

func mapRefundRequestItems(items RefundRequestItem) refund.RefundRequestItem {
	return refund.RefundRequestItem{
		LineItemId: items.LineItemId,
		Quantity:   items.Quantity,
		Amount: refund.RefundRequestItemAmount{
			Amount:   items.Amount.Amount,
			Currency: items.Amount.Currency,
		},
	}
}
