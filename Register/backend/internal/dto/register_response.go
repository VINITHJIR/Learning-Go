package dto

// RegisterResponse represents the client-facing user data after registration.
type RegisterResponse struct {
	ID          uint64 `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}
