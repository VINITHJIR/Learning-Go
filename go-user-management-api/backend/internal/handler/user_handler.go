package handler

import (
	"encoding/json"
	"net/http"

	"user-management-api/internal/domain"
	"user-management-api/internal/dto"
	"user-management-api/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	// Allow only POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read Request Body
	var request dto.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(dto.APIResponse{
			Success: false,
			Message: "Invalid Request Body",
		})
		return
	}

	// Convert DTO -> Domain Model
	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Phone:    request.Phone,
		Address:  request.Address,
		Password: request.Password,
	}

	// Call Service
	err = h.service.Register(&user)
	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(dto.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Prepare Response DTO
	response := dto.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Address:  user.Address,
	}

	// Send Success Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(dto.APIResponse{
		Success: true,
		Message: "User Registered Successfully",
		Data:    response,
	})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var request dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(dto.APIResponse{
			Success: false,
			Message: "Invalid Request",
		})

		return
	}

	token, err := h.service.Login(
		request.Email,
		request.Password,
	)

	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		json.NewEncoder(w).Encode(dto.APIResponse{
			Success: false,
			Message: err.Error(),
		})

		return
	}

	response := dto.LoginResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dto.APIResponse{
		Success: true,
		Message: "Login Successful",
		Data:    response,
	})
}
