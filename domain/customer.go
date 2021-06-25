package domain

import (
	"github.org/kbank/dto"
	"github.org/kbank/errs"
)

type Customer struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	City        string `bson:"city"`
	Zipcode     string `bson:"zip_code"`
	DateofBirth string `bson:"date_of_birth"`
	Status      string `bson:"status"`
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.Status,
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
}
