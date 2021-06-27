package account

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountResponse struct {
	ID          primitive.ObjectID `json:"_id"`
	CustomerID  primitive.ObjectID `json:"customer_id"`
	OpeningDate time.Time          `json:"opening_date"`
	AccountType string             `json:"account_type"`
	Amount      float64            `json:"amount"`
	Status      string             `json:"status"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
type AccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}
type ResultResponse struct {
	Message string `json:"message"`
}
