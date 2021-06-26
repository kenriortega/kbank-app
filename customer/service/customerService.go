package customer

import (
	domain "github.org/kbank/customer/domain"
	dto "github.org/kbank/customer/dto"
	"github.org/kbank/internal/errs"
)

type CustomerService interface {
	GetAllConstumer(string) ([]dto.CustomerResponse, *errs.AppError)
	CreateCustomer(dto.CustomerRequest) (dto.ResultResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllConstumer(status string) (response []dto.CustomerResponse, err *errs.AppError) {
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) CreateCustomer(newCustomer dto.CustomerRequest) (result dto.ResultResponse, err *errs.AppError) {

	customer := domain.Customer{
		Name:        newCustomer.Name,
		City:        newCustomer.City,
		Zipcode:     newCustomer.Zipcode,
		DateofBirth: newCustomer.DateofBirth,
	}
	s.repo.CreateOne(customer)
	result = dto.ResultResponse{
		Message: "Customer created",
	}
	return result, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
