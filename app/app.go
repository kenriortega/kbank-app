package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	domain "github.org/kbank/customer/domain"
	service "github.org/kbank/customer/service"
	"github.org/kbank/logger"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	ch := CustomerHandler{
		service: service.NewCustomerService(domain.NewCustomerRepositoryDb()),
	}
	// define routes
	r.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	logger.Error(err.Error())
}
