// Schemes: http, https
// Host: 127.0.0.1
// BasePath: /api/v1
// Version: 0.0.1
// License: MIT http://opensource.org/licenses/MIT
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
//
// swagger:meta
package api

import (
	customerDto "github.org/kbank/customer/dto"
)

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response ErrorResponse
type ErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body customerDto.ResultResponse
}

// A list of customers
// swagger:response CustomersResponse
type CustomersResponseWrapper struct {
	// All current customers
	// in: body
	Body []customerDto.CustomerResponse
}

// Data structure representing a single product
// swagger:response CustomerResponse
type CustomerResponseWrapper struct {
	// Newly created customer
	// in: body
	Body customerDto.CustomerResponse
}

// swagger:parameters UpdateCustomerRequest
type UpdateCustomerRequestWrapper struct {
	// The status of the customers for which the operation relates
	// in: body
	UpdateRequest customerDto.UpdateCustomerRequest
}

// No content is returned by this API endpoint
// swagger:response NoContentResponse
type NoContentResponseWrapper struct {
}
