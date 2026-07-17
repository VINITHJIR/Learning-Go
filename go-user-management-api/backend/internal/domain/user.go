package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(100);not null" json:"username"`
	Email     string    `gorm:"type:varchar(150);unique;not null" json:"email"`
	Phone     string    `gorm:"type:varchar(15);unique;not null" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
