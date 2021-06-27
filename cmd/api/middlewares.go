package api

import (
	"fmt"
	"net/http"

	"github.org/kbank/internal/logger"
)

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			trc := fmt.Sprintf("[%s]-[%s]", r.Method, r.RequestURI)
			logger.Info(trc)
			next.ServeHTTP(w, r)
		},
	)
}
