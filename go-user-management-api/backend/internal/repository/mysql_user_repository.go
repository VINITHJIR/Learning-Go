package repository

import (
	"user-management-api/internal/domain"

	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &MySQLUserRepository{
		db: db,
	}
}

func (r *MySQLUserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *MySQLUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
