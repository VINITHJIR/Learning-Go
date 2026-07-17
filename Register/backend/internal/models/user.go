package models

import (
	"time"
)

// User represents the users table schema in MySQL database.
type User struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement;column:id"`
	Username    string    `gorm:"type:varchar(100);unique;not null;column:username"`
	Email       string    `gorm:"type:varchar(150);unique;not null;column:email"`
	PhoneNumber string    `gorm:"type:varchar(20);unique;not null;column:phone_number"`
	Address     string    `gorm:"type:text;not null;column:address"`
	Password    string    `gorm:"type:varchar(255);not null;column:password"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

// TableName overrides the table name used by GORM to "users".
func (User) TableName() string {
	return "users"
}
