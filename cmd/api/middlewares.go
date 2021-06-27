package api

import (
	"fmt"
	"net/http"
	"strings"

	authDto "github.org/kbank/auth/dto"

	"github.com/gbrlsnchs/jwt/v3"
	"github.org/kbank/internal/logger"
)

var hs = jwt.NewHS256([]byte("secret"))

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
			var header = r.Header.Get("Authorization")
			token := strings.Split(header, " ")[1]
			pl := authDto.JWTPayload{}
			hd, err := jwt.Verify([]byte(token), hs, &pl)
			if err != nil {
				logger.Error(err.Error())
			}
			fmt.Println(hd)
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
