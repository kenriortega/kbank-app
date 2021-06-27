package customer

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	dto "github.org/kbank/customer/dto"
	service "github.org/kbank/customer/service"
	"github.org/kbank/internal/errs"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")
	customers, err := ch.Service.GetAllConstumer(status)
	if err != nil {
		if err.Message == errs.NoDocumentsError.Message {
			writeResponse(w, err.Code, err)
		}
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

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

// TODO: Check work flow
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

// TODO: Check work flow
func (ch *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := ch.Service.DeleteCustomer(params["customerID"])
	if err != nil {
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusAccepted, result)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
