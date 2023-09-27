/*
 * Integration stubs
 *
 * Stubs for implementing a REVER integration
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"context"
	"net/http"
)



// IntegrationApiRouter defines the required methods for binding the api requests to a responses for the IntegrationApi
// The IntegrationApiRouter implementation should parse necessary information from the http request,
// pass the data to a IntegrationApiServicer to perform the required actions, then write the service results to the http response.
type IntegrationApiRouter interface { 
	AddNoteToOrder(http.ResponseWriter, *http.Request)
	CreateRefund(http.ResponseWriter, *http.Request)
	CreateReturn(http.ResponseWriter, *http.Request)
	FindOrderByCustomerPrintedOrderId(http.ResponseWriter, *http.Request)
}


// IntegrationApiServicer defines the api actions for the IntegrationApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type IntegrationApiServicer interface { 
	AddNoteToOrder(context.Context, string, AddNoteToOrderRequest) (ImplResponse, error)
	CreateRefund(context.Context, string, RefundRequest) (ImplResponse, error)
	CreateReturn(context.Context, string, ReturnRequest) (ImplResponse, error)
	FindOrderByCustomerPrintedOrderId(context.Context, string) (ImplResponse, error)
}
