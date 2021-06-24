package service

import (
	"github.org/kbank/domain"
	"github.org/kbank/errs"
)

type CustomerService interface {
	GetAllConstumer() ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllConstumer() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
