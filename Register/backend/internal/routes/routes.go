package routes

import (
	"register/internal/handler"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the route groups and endpoints.
func SetupRoutes(r *gin.Engine, authHandler *handler.AuthHandler) {
	// Root api/v1 group
	api := r.Group("/api/v1")
	{
		// Auth sub-group
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
		}
	}
}
