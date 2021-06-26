package customer

import (
	domain "github.org/kbank/customer/domain"
	dto "github.org/kbank/customer/dto"
	"github.org/kbank/internal/errs"
)

type CustomerService interface {
	GetAllConstumer(string) ([]dto.CustomerResponse, *errs.AppError)
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

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
