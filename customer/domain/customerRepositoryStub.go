package customer

import "go.mongodb.org/mongo-driver/bson/primitive"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{ID: primitive.NewObjectID(), Name: "Cust1", City: "Hav", Zipcode: "120202", DateofBirth: "2021-06-23", Status: "1"},
		{ID: primitive.NewObjectID(), Name: "Cust2", City: "Hav", Zipcode: "120202", DateofBirth: "2021-06-23", Status: "1"},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
