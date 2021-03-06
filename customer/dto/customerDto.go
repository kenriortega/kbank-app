package customer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerResponse struct {
	ID          primitive.ObjectID `json:"_id" `
	Name        string             `json:"name"`
	City        string             `json:"city"`
	Zipcode     string             `json:"zip_code"`
	DateofBirth string             `json:"date_of_birth"`
	Status      string             `json:"status"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type CustomerRequest struct {
	ID          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	City        string             `json:"city"`
	Zipcode     string             `json:"zip_code"`
	DateofBirth string             `json:"date_of_birth"`
	Status      string             `json:"status"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type UpdateCustomerRequest struct {
	ID        primitive.ObjectID `json:"_id"`
	Status    string             `json:"status"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type ResultResponse struct {
	Message string `json:"message"`
}
