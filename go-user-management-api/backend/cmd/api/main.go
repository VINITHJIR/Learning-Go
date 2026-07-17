package main

import (
	"log"
	"net/http"

	"user-management-api/internal/config"
	"user-management-api/internal/database"
	"user-management-api/internal/handler"
	"user-management-api/internal/repository"
	"user-management-api/internal/router"
	"user-management-api/internal/service"
)

func main() {

	// Load Environment Variables
	config.LoadConfig()

	// Connect Database
	database.ConnectDatabase()

	// Dependency Injection
	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Register Routes
	router.RegisterRoutes(userHandler)

	log.Println("🚀 Server Running on Port :", config.AppConfig.AppPort)

	err := http.ListenAndServe(":"+config.AppConfig.AppPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
