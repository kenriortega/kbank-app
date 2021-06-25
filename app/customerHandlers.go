package app

import (
	"encoding/json"
	"net/http"
	"strings"

	service "github.org/kbank/customer/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllConstumer(status)
	if err != nil {
		if strings.Contains(err.Message, "no documents") {
			writeResponse(w, err.Code, err)
		}
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
