package repository

import (
	"errors"
	"register/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository is the constructor for GORM-based user repository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create inserts a new user record in the database.
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByUsername retrieves a user by their username.
func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found is not treated as a system failure here
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail retrieves a user by their email address.
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByPhoneNumber retrieves a user by their phone number.
func (r *userRepository) FindByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
