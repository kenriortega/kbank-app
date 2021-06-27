package account

import (
	"encoding/json"
	"net/http"

	dto "github.org/kbank/account/dto"
	service "github.org/kbank/account/service"
	"github.org/kbank/internal/errs"
)

type AccountHandler struct {
	Service service.AccountService
}

func (ah *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newAccount dto.AccountRequest
	_ = json.NewDecoder(r.Body).Decode(&newAccount)
	result, err := ah.Service.CreateAccount(newAccount)

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
