package router

import (
	"net/http"

	"user-management-api/internal/handler"
)

func RegisterRoutes(userHandler *handler.UserHandler) {

	http.HandleFunc(
		"/api/v1/auth/register",
		userHandler.Register,
	)

	http.HandleFunc(
		"/api/v1/auth/login",
		userHandler.Login,
	)
}
