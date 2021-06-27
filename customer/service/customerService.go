package customer

import (
	"fmt"
	"time"

	domain "github.org/kbank/customer/domain"
	dto "github.org/kbank/customer/dto"
	"github.org/kbank/internal/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerService interface {
	GetAllConstumer(string) ([]dto.CustomerResponse, *errs.AppError)
	CreateCustomer(dto.CustomerRequest) (dto.ResultResponse, *errs.AppError)
	DeleteCustomer(string) (dto.ResultResponse, *errs.AppError)
	UpdateCustomerByStatus(string, dto.UpdateCustomerRequest) (dto.ResultResponse, *errs.AppError)
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

	_, err = s.repo.CreateOne(customer)
	result = dto.ResultResponse{
		Message: "1",
	}
	if err != nil {
		result = dto.ResultResponse{
			Message: "0",
		}
		return result, err
	}

	return result, nil
}

func (s DefaultCustomerService) DeleteCustomer(customerID string) (result dto.ResultResponse, err *errs.AppError) {

	customerObjectID, _ := primitive.ObjectIDFromHex(customerID)
	fmt.Println(customerObjectID)
	result = dto.ResultResponse{
		Message: "Customer deleted",
	}
	return result, nil
}
func (s DefaultCustomerService) UpdateCustomerByStatus(customerID string, updateRequest dto.UpdateCustomerRequest) (result dto.ResultResponse, err *errs.AppError) {

	customerObjectID, _ := primitive.ObjectIDFromHex(customerID)
	updateRequest.ID = customerObjectID
	updateRequest.UpdatedAt = time.Now()

	fmt.Println(updateRequest)
	result = dto.ResultResponse{
		Message: "Customer Updated",
	}
	return result, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
