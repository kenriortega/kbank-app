package auth

import "github.com/gbrlsnchs/jwt/v3"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Role     string `json:"role"`
	Username string `json:"username"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ResultResponse struct {
	Message string `json:"message"`
}

type JWTPayload struct {
	jwt.Payload
	Username string
	Role     string
}
