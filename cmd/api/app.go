package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	accountDomain "github.org/kbank/account/domain"
	accountHandler "github.org/kbank/account/handlers"
	accountService "github.org/kbank/account/service"
	authDomain "github.org/kbank/auth/domain"
	authHandler "github.org/kbank/auth/handlers"
	authService "github.org/kbank/auth/service"
	customerDomain "github.org/kbank/customer/domain"
	customerHandler "github.org/kbank/customer/handlers"
	customerService "github.org/kbank/customer/service"
	"github.org/kbank/internal/logger"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	mongoDbClient := GetMongoDbClient()

	// Customers Services
	customerRepository := customerDomain.NewCustomerRepositoryDb(mongoDbClient)
	ch := customerHandler.CustomerHandler{
		Service: customerService.NewCustomerService(customerRepository),
	}
	// define routes for customers
	customersRoutes := r.PathPrefix("/api/v1/customers").Subrouter()
	customersRoutes.HandleFunc("/{customerID}/status", ch.UpdateStatusCustomer).Methods(http.MethodPatch)
	customersRoutes.HandleFunc("/{customerID}", ch.DeleteCustomer).Methods(http.MethodDelete)
	customersRoutes.HandleFunc("/{customerID}", ch.GetCustomer).Methods(http.MethodGet)
	customersRoutes.HandleFunc("/", Chain(ch.GetAllCustomers, VerifyJWT())).Methods(http.MethodGet)
	customersRoutes.HandleFunc("/", ch.CreateCustomer).Methods(http.MethodPost)
	// End Customer service

	// define routes for accounts
	accountRepository := accountDomain.NewAccountRepositoryDb(mongoDbClient)
	ah := accountHandler.AccountHandler{
		Service: accountService.NewAccountService(accountRepository),
	}
	accoutsRoutes := r.PathPrefix("/api/v1/accounts").Subrouter()
	accoutsRoutes.HandleFunc("/", ah.GetAllAccount).Methods(http.MethodGet)
	accoutsRoutes.HandleFunc("/", ah.CreateAccount).Methods(http.MethodPost)
	// End Accounts service

	// define routes for accounts
	authRepository := authDomain.NewAuthRepositoryDb(mongoDbClient)
	auh := authHandler.AuthHandler{
		Service: authService.NewAuthService(authRepository),
	}
	authRoutes := r.PathPrefix("/api/v1/auth").Subrouter()
	authRoutes.HandleFunc("/register", auh.Register).Methods(http.MethodPost)
	authRoutes.HandleFunc("/login", auh.Login).Methods(http.MethodPost)
	// End Accounts service

	// swagger
	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// middleware
	r.Use(LogginMiddleware)
	r.Use(CORSMiddleware)
	// starting server
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()
	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	logger.Info(fmt.Sprintf("Got signal: %s", sig))

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}
