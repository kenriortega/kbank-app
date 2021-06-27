package customer

import (
	"time"

	dto "github.org/kbank/customer/dto"
	"github.org/kbank/internal/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Customer struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	City        string             `bson:"city"`
	Zipcode     string             `bson:"zip_code"`
	DateofBirth string             `bson:"date_of_birth"`
	Status      string             `bson:"status"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.Status,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindOne(primitive.ObjectID) (Customer, *errs.AppError)
	CreateOne(Customer) (*mongo.InsertOneResult, *errs.AppError)
	CreateMany(list []Customer) (*mongo.InsertManyResult, *errs.AppError)
	DeleteOne(primitive.ObjectID) (*mongo.DeleteResult, *errs.AppError)
	DeleteAll() (*mongo.DeleteResult, *errs.AppError)
	UpdateStatusByCustomerID(customerID primitive.ObjectID, status string, updatedAt time.Time) (*mongo.UpdateResult, *errs.AppError)
}
