package account

import (
	"time"

	dto "github.org/kbank/account/dto"
	"github.org/kbank/internal/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	ID          primitive.ObjectID `bson:"_id"`
	CustomerID  primitive.ObjectID `bson:"customer_id"`
	OpeningDate time.Time          `bson:"opening_date"`
	AccountType string             `bson:"account_type"`
	Amount      float64            `bson:"amount"`
	Status      string             `bson:"status"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (a Account) ToDto() dto.AccountResponse {
	return dto.AccountResponse{
		ID:          a.ID,
		CustomerID:  a.CustomerID,
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.Status,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

type AccountRepository interface {
	FindAll() ([]Account, *errs.AppError)
	CreateOne(Account) (*mongo.InsertOneResult, *errs.AppError)
}
