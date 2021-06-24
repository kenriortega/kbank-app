package domain

import "github.org/kbank/errs"

type Customer struct {
	ID          string `json:"_id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	City        string `json:"city" bson:"city"`
	Zipcode     string `json:"zip_code" bson:"zip_code"`
	DateofBirth string `json:"date_of_birth" bson:"date_of_birth"`
	Status      string `json:"status" bson:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
}
