/*
Integration stubs

Stubs for implementing a REVER integration

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// IntegrationApiService IntegrationApi service
type IntegrationApiService service

type ApiFindOrderByCustomerOrderPrintedIdRequest struct {
	ctx context.Context
	ApiService *IntegrationApiService
	customerOrderPrintedId *string
}

// This is the &#x60;order_id&#x60; as seen by the customer
func (r ApiFindOrderByCustomerOrderPrintedIdRequest) CustomerOrderPrintedId(customerOrderPrintedId string) ApiFindOrderByCustomerOrderPrintedIdRequest {
	r.customerOrderPrintedId = &customerOrderPrintedId
	return r
}

func (r ApiFindOrderByCustomerOrderPrintedIdRequest) Execute() (*IntegrationOrder, *http.Response, error) {
	return r.ApiService.FindOrderByCustomerOrderPrintedIdExecute(r)
}

/*
FindOrderByCustomerOrderPrintedId Find Order by customer_order_id

Finds a single order by its `customer_order_printed_id`.
In case of multiple matches, the most recent one is returned.

Please note that `customer_order_printed_id` is the `order_id` as seen by the customer, typically in the order mail.
It's not necessarily the same as the `order_id` as seen by the platform.

The path provided is the suggested one for this methods, you can choose your own in your integration.
In case of receiving a different response code than the ones listed below, REVER might contact the integration for further investigation.

(WIP) The implementation of this method must support the following scenarios:

- **(FIND00)** The requests does not contain the authentication header. The response must be `401`.
- **(FIND01)** The requests contains an invalid authentication header. The response must be `401`.
- **(FIND02)** The `customer_order_printed_id` is empty. The response must be `400`.
- **(FIND03)** The `customer_order_printed_id` is not empty but not found in the integration. The response code must be `404`.
- **(FIND04)** The simplest returning scenario. The `customer_order_printed_id` is found and it returns a valid order:  
  - a single product without variants
  - quantity 1
  - with a positive amount in EUR
  - with taxes
  - with shipping costs
  - with a non-cash, non-cash on delivery, non-BNPL payment method 
  - without discounts applied
  - with a valid customer name and email
  - fulfilled, with a valid shipping address
  - with a valid purchase date
  In this case, the response code must be `200`.
- **(FIND05)** A return scenario with multiple products. 
  - More than one product without variants
  - quantity > 0
  - with a positive amount in EUR
  - with taxes
  - with shipping costs
  - with a non-cash, non-cash on delivery, non-BNPL payment method 
  - without discounts applied
  - with a valid customer name and email
  - fulfilled, with a valid shipping address
  - with a valid purchase date
  In this case, the response code must be `200`.
- **more cases to be added soon**


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFindOrderByCustomerOrderPrintedIdRequest
*/
func (a *IntegrationApiService) FindOrderByCustomerOrderPrintedId(ctx context.Context) ApiFindOrderByCustomerOrderPrintedIdRequest {
	return ApiFindOrderByCustomerOrderPrintedIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IntegrationOrder
func (a *IntegrationApiService) FindOrderByCustomerOrderPrintedIdExecute(r ApiFindOrderByCustomerOrderPrintedIdRequest) (*IntegrationOrder, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntegrationOrder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationApiService.FindOrderByCustomerOrderPrintedId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/integration/orders/find"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerOrderPrintedId == nil {
		return localVarReturnValue, nil, reportError("customerOrderPrintedId is required and must be specified")
	}

	localVarQueryParams.Add("customer_order_printed_id", parameterToString(*r.customerOrderPrintedId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
