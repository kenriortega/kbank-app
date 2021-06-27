package account

import (
	"encoding/json"
	"net/http"

	dto "github.org/kbank/auth/dto"
	service "github.org/kbank/auth/service"
	"github.org/kbank/internal/errs"
)

type AuthHandler struct {
	Service service.AuthService
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser dto.RegisterRequest
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	result, err := ah.Service.Register(newUser)

	if err != nil {
		if err.Message == errs.InsertOneError.Message {
			writeResponse(w, err.Code, err)
		}
	} else {
		writeResponse(w, http.StatusCreated, result)
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var authReq dto.LoginRequest
	_ = json.NewDecoder(r.Body).Decode(&authReq)
	result, err := ah.Service.Login(authReq)

	if err != nil {
		if err.Message == errs.InsertOneError.Message {
			writeResponse(w, err.Code, err)
		}
	} else {
		writeResponse(w, http.StatusCreated, result)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
