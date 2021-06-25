package service

import (
	"github.org/kbank/domain"
	"github.org/kbank/errs"
)

type CustomerService interface {
	GetAllConstumer(string) ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllConstumer(status string) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll(status)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
