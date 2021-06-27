package account

import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	domain "github.org/kbank/auth/domain"
	dto "github.org/kbank/auth/dto"

	"github.org/kbank/internal/errs"
)

var hs = jwt.NewHS256([]byte("secret"))

type AuthService interface {
	Register(dto.RegisterRequest) (dto.ResultResponse, *errs.AppError)
	Login(dto.LoginRequest) (dto.LoginResponse, *errs.AppError)
}
type DefaultAuthService struct {
	repo domain.AuthRepository
}

func (s DefaultAuthService) Register(newUser dto.RegisterRequest) (result dto.ResultResponse, err *errs.AppError) {

	user := domain.User{
		Username: newUser.Username,
		Password: newUser.Password,
		Role:     newUser.Role,
	}

	_, err = s.repo.CreateOne(user)
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

func (s DefaultAuthService) Login(authReq dto.LoginRequest) (response dto.LoginResponse, err *errs.AppError) {

	user, err := s.repo.Login(authReq.Username, authReq.Password)
	response.Username = user.Username
	response.Role = user.Role
	now := time.Now()
	pl := dto.JWTPayload{
		Payload: jwt.Payload{
			Issuer:         "Bank",
			Subject:        "SystemApp",
			Audience:       jwt.Audience{"http://localhost:8000"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "auth-server-1",
		},
		Username: user.Username,
		Role:     user.Role,
	}

	token, errToken := jwt.Sign(pl, hs)
	if errToken != nil {
		// ...
	}

	if err != nil {
		return response, err
	}
	response.Token = string(token)
	return response, nil
}

func NewAuthService(repository domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{repo: repository}
}
