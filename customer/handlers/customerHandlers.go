package customer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	dto "github.org/kbank/customer/dto"
	service "github.org/kbank/customer/service"
	"github.org/kbank/internal/errs"
)

type CustomerHandler struct {
	Service service.CustomerService
}

// swagger:route GET /customers/ customers GetAllCustomers
// Return a list of customers from the database
// responses:
//	200: CustomersResponseWrapper
//  204: NoContentResponseWrapper
func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status := r.URL.Query().Get("status")
	customers, err := ch.Service.GetAllConstumer(status)
	name := context.Get(r, "username")
	fmt.Println(name)
	if err != nil {
		if err.Message == errs.NoDocumentsError.Message {
			writeResponse(w, err.Code, err)
		}
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

// swagger:route GET /customers/{customerID} customers GetCustomer
// Return a customer from the database
// responses:
//	200: CustomerResponse
//	204: NoContentResponseWrapper

// ListSingle handles GET requests
func (ch *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := ch.Service.GetCustomer(params["customerID"])
	if err != nil {
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusOK, result)
	}

}

// swagger:route POST /customers customers CreateCustomer
// Create a new customer
//
// responses:
//	201: ResultResponse

// Create handles POST requests to add new customer
func (ch *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCustomer dto.CustomerRequest
	_ = json.NewDecoder(r.Body).Decode(&newCustomer)
	result, err := ch.Service.CreateCustomer(newCustomer)

	if err != nil {
		if err.Message == errs.InsertOneError.Message {
			writeResponse(w, err.Code, err)
		}
	} else {
		writeResponse(w, http.StatusCreated, result)
	}
}

// swagger:route PATCH /customers/{customerID} customers UpdateStatusCustomer
// Update a customer
//
// responses:
//	202: ResultResponse

// Create handles POST requests to add new customer
func (ch *CustomerHandler) UpdateStatusCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updateCustomerRequest dto.UpdateCustomerRequest
	_ = json.NewDecoder(r.Body).Decode(&updateCustomerRequest)
	result, err := ch.Service.UpdateCustomerByStatus(params["customerID"], updateCustomerRequest)
	if err != nil {
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusAccepted, result)
	}
}

// swagger:route DELETE /customers/{customerID} customers DeleteCustomer
// Delete a customer
//
// responses:
//	202: ResultResponse

// Create handles DELETE requests to add new customer
func (ch *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	result, err := ch.Service.DeleteCustomer(params["customerID"])
	if err != nil {
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusAccepted, result)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
