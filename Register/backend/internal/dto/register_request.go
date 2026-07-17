package dto

// RegisterRequest holds validation rules for user registration payload.
type RegisterRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Password    string `json:"password" binding:"required,min=8"`
}
