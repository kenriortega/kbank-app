package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	domain "github.org/kbank/customer/domain"
	handlers "github.org/kbank/customer/handlers"
	service "github.org/kbank/customer/service"
	"github.org/kbank/internal/logger"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	ch := handlers.CustomerHandler{
		Service: service.NewCustomerService(domain.NewCustomerRepositoryDb()),
	}
	// define routes
	r.HandleFunc("/customers/{customerID}/status", ch.UpdateStatusCustomer).Methods(http.MethodPatch)
	r.HandleFunc("/customers/{customerID}", ch.DeleteCustomer).Methods(http.MethodDelete)
	r.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers", ch.CreateCustomer).Methods(http.MethodPost)

	// starting server
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	logger.Error(err.Error())
}
