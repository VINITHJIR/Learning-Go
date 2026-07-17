package service

import (
	"register/internal/dto"
)

// AuthService defines business logic workflows for authentication.
type AuthService interface {
	Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error)
}
