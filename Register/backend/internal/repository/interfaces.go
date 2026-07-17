package repository

import (
	"register/internal/models"
)

// UserRepository defines the contract for user-related database operations.
type UserRepository interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByPhoneNumber(phoneNumber string) (*models.User, error)
}
