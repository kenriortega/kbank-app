package account

import (
	domain "github.org/kbank/account/domain"
	dto "github.org/kbank/account/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.org/kbank/internal/errs"
)

type AccountService interface {
	CreateAccount(dto.AccountRequest) (dto.ResultResponse, *errs.AppError)
	GetAllAccount() ([]dto.AccountResponse, *errs.AppError)
}
type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) GetAllAccount() (response []dto.AccountResponse, err *errs.AppError) {
	accounts, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	for _, a := range accounts {
		response = append(response, a.ToDto())
	}
	return response, nil
}

func (s DefaultAccountService) CreateAccount(newAccount dto.AccountRequest) (result dto.ResultResponse, err *errs.AppError) {
	customerObjectID, _ := primitive.ObjectIDFromHex(newAccount.CustomerID)
	account := domain.Account{
		CustomerID:  customerObjectID,
		AccountType: newAccount.AccountType,
		Amount:      newAccount.Amount,
	}

	_, err = s.repo.CreateOne(account)
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

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}