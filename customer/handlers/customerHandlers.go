package customer

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	dto "github.org/kbank/customer/dto"
	service "github.org/kbank/customer/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")
	customers, err := ch.Service.GetAllConstumer(status)
	if err != nil {
		if strings.Contains(err.Message, "no documents") {
			writeResponse(w, err.Code, err)
		}
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer dto.CustomerRequest
	_ = json.NewDecoder(r.Body).Decode(&customer)
	customer.ID = primitive.NewObjectID()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	writeResponse(w, http.StatusCreated, customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
