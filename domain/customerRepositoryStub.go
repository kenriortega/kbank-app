package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{ID: "1001", Name: "Cust1", City: "Hav", Zipcode: "120202", DateofBirth: "2021-06-23", Status: "1"},
		{ID: "1002", Name: "Cust2", City: "Hav", Zipcode: "120202", DateofBirth: "2021-06-23", Status: "1"},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
