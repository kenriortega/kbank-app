package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

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
	// define routes for customers
	customersRoutes := r.PathPrefix("/customers").Subrouter()
	customersRoutes.HandleFunc("/{customerID}/status", ch.UpdateStatusCustomer).Methods(http.MethodPatch)
	customersRoutes.HandleFunc("/{customerID}", ch.DeleteCustomer).Methods(http.MethodDelete)
	customersRoutes.HandleFunc("/{customerID}", ch.GetCustomer).Methods(http.MethodGet)
	customersRoutes.HandleFunc("/", ch.GetAllCustomers).Methods(http.MethodGet)
	customersRoutes.HandleFunc("/", ch.CreateCustomer).Methods(http.MethodPost)

	// middleware
	r.Use(LogginMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	// starting server
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	err = srv.ListenAndServe()
	logger.Error(err.Error())
}
