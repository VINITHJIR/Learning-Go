package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"register/internal/config"
	"register/internal/handler"
	"register/internal/models"
	"register/internal/repository"
	"register/internal/routes"
	"register/internal/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting Authentication System Backend...")

	// 1. Load Configurations
	cfg := config.LoadConfig()

	// Set Gin mode
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 2. Connect to MySQL Database
	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// 3. Auto-Migrate schema
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	log.Println("Database migration completed successfully")

	// 4. Initialize Clean Architecture Layers (Constructor-based Dependency Injection)
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// 5. Initialize Router and Setup Routes
	router := gin.Default()
	routes.SetupRoutes(router, authHandler)

	// 6. Graceful Shutdown Configuration
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server is running on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server gracefully...")

	// Context with 5-second timeout for cleanup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited cleanly")
}
