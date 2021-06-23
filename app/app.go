package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	r := mux.NewRouter()

	// define routes
	r.HandleFunc("/greet", greet).Methods(http.MethodGet)
	r.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	r.HandleFunc("/customers/{customerID:[0-9]+}", getCustomerByID).Methods(http.MethodGet)

	// starting smuxer
	log.Fatal(http.ListenAndServe(":8000", r))
}
