package errs

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var NoDocumentsError = &AppError{
	Message: "No documents",
	Code:    http.StatusNoContent,
}
var NotFoundError = &AppError{
	Message: "Document Not found",
	Code:    http.StatusNoContent,
}

var InsertOneError = &AppError{
	Message: "Unexpected error on insert document",
	Code:    http.StatusBadRequest,
}

var UpdateError = &AppError{
	Message: "Unexpected error on update document",
	Code:    http.StatusBadRequest,
}

var DeleteOneError = &AppError{
	Message: "Unexpected error on delete one document",
	Code:    http.StatusBadRequest,
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
