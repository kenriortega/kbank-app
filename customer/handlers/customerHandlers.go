package customer

import (
	"encoding/json"
	"net/http"

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

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
