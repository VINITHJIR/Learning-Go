package utils

import (
	"github.com/gin-gonic/gin"
)

// Response defines the standardized structure for API responses.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// JSONResponse sends a standardized JSON response.
func JSONResponse(c *gin.Context, statusCode int, success bool, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: success,
		Message: message,
		Data:    data,
	})
}
