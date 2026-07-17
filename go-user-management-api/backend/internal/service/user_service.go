package service

import (
	"errors"

	"user-management-api/internal/config"
	"user-management-api/internal/domain"
	"user-management-api/internal/repository"
	"user-management-api/internal/utils"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Register(user *domain.User) error {

	_, err := s.repository.FindByEmail(user.Email)

	if err == nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return s.repository.Create(user)
}
func (s *UserService) Login(email, password string) (string, error) {

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return "", errors.New("Email Not Found")
	}

	err = utils.CheckPassword(user.Password, password)
	if err != nil {
		println("Password Compare Failed")
		return "", errors.New("Password Incorrect")
	}

	token, err := utils.GenerateJWT(
		user.ID,
		user.Email,
		config.AppConfig.JWTSecret,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
