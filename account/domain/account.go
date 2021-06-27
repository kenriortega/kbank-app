package account

import (
	"time"

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

type AccountRepository interface {
	CreateOne(Account) (*mongo.InsertOneResult, *errs.AppError)
}
