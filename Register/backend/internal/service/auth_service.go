package service

import (
	"errors"
	"fmt"
	"register/internal/dto"
	"register/internal/models"
	"register/internal/repository"
	"register/internal/utils"
)

// Sentinel errors for authentication business logic
var (
	ErrUsernameExists    = errors.New("username already exists")
	ErrEmailExists       = errors.New("email already exists")
	ErrPhoneNumberExists = errors.New("phone number already exists")
)

type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService is the constructor for authService.
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

// Register contains user registration business workflow.
func (s *authService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// 1. Check if username is already taken
	existingUser, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed checking username uniqueness: %w", err)
	}
	if existingUser != nil {
		return nil, ErrUsernameExists
	}

	// 2. Check if email is already taken
	existingUser, err = s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed checking email uniqueness: %w", err)
	}
	if existingUser != nil {
		return nil, ErrEmailExists
	}

	// 3. Check if phone number is already taken
	existingUser, err = s.userRepo.FindByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed checking phone number uniqueness: %w", err)
	}
	if existingUser != nil {
		return nil, ErrPhoneNumberExists
	}

	// 4. Hash the password before storing
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 5. Map DTO to GORM Model
	userModel := &models.User{
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Password:    hashedPassword,
	}

	// 6. Persist user to database
	if err := s.userRepo.Create(userModel); err != nil {
		return nil, fmt.Errorf("failed to create user record: %w", err)
	}

	// 7. Map GORM Model to Response DTO (excluding password/metadata)
	response := &dto.RegisterResponse{
		ID:          userModel.ID,
		Username:    userModel.Username,
		Email:       userModel.Email,
		PhoneNumber: userModel.PhoneNumber,
		Address:     userModel.Address,
	}

	return response, nil
}
