package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	authDto "github.org/kbank/auth/dto"

	"github.com/gbrlsnchs/jwt/v3"
	"github.org/kbank/internal/logger"
)

var (
	hs = jwt.NewHS256([]byte("secret"))
)

type ResponseMiddleware struct {
	Message string `json:"message"`
}
type Middleware func(http.HandlerFunc) http.HandlerFunc

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				trc := fmt.Sprintf("[%s]-[%s]", r.Method, r.RequestURI)
				logger.Info(trc)
			}()

			next.ServeHTTP(w, r)
		},
	)
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}

//
func VerifyJWT() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			var response ResponseMiddleware
			var header = r.Header.Get("Authorization")
			now := time.Now()
			if !strings.HasPrefix(header, "Bearer ") {
				response.Message = "Format is Authorization: Bearer [token]"
				fmt.Println(response)
				writeResponse(w, http.StatusBadRequest, response)
				return
			}
			token := strings.Split(header, " ")[1]
			pl := authDto.JWTPayload{}
			expValidator := jwt.ExpirationTimeValidator(now)
			validatePayload := jwt.ValidatePayload(&pl.Payload, expValidator)
			_, err := jwt.Verify([]byte(token), hs, &pl, validatePayload)

			if errors.Is(err, jwt.ErrExpValidation) {
				logger.Error(err.Error())
				response.Message = err.Error()
				writeResponse(w, http.StatusForbidden, response)
				return
			}
			if errors.Is(err, jwt.ErrHMACVerification) {
				logger.Error(err.Error())
				response.Message = err.Error()
				writeResponse(w, http.StatusForbidden, response)
				return
			}

			// ACL for ROLES
			switch pl.Role {
			case "ADMIN":
				f(w, r)
				return

			case "BASIC":
				logger.Error("You don`t have permissions")
				response.Message = "You don`t have permissions"
				writeResponse(w, http.StatusUnauthorized, response)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
