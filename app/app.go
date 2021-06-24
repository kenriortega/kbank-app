package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.org/kbank/domain"
	"github.org/kbank/service"
)

func Start() {

	r := mux.NewRouter()

	// wiring
	ch := CustomerHandler{
		service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	}
	// define routes
	r.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe(":8000", r))
}
