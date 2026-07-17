package handler

import (
	"errors"
	"net/http"
	"register/internal/dto"
	"register/internal/service"
	"register/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AuthHandler handles HTTP requests for user authentication.
type AuthHandler struct {
	authService service.AuthService
}

// NewAuthHandler is the constructor for AuthHandler.
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register processes user registration requests.
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	// Bind and validate JSON body automatically using the validator tags in DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			// Format validation errors to be readable
			errMsgs := make(map[string]string)
			for _, fe := range ve {
				errMsgs[fe.Field()] = fe.Tag()
			}
			utils.JSONResponse(c, http.StatusBadRequest, false, "Validation failed", errMsgs)
			return
		}
		utils.JSONResponse(c, http.StatusBadRequest, false, "Invalid request payload", err.Error())
		return
	}

	// Invoke business logic layer
	res, err := h.authService.Register(&req)
	if err != nil {
		// Handle duplication/conflict errors
		if errors.Is(err, service.ErrUsernameExists) ||
			errors.Is(err, service.ErrEmailExists) ||
			errors.Is(err, service.ErrPhoneNumberExists) {
			utils.JSONResponse(c, http.StatusConflict, false, err.Error(), nil)
			return
		}

		// Handle general server errors
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Internal server error occurred", nil)
		return
	}

	// Return successful response
	utils.JSONResponse(c, http.StatusCreated, true, "User registered successfully", res)
}
