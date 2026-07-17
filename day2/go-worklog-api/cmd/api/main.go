package main

import (
	"log"
	"net/http"

	"go-worklog-api/internal/database"
	"go-worklog-api/internal/handler"
	"go-worklog-api/internal/repository"
	"go-worklog-api/internal/router"
	"go-worklog-api/internal/service"
)

func main() {

	dsn := "root:VIni@2003@tcp(localhost:3306)/go_worklog_db"

	db, err := database.NewMySQLConnection(dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	worklogRepository := repository.NewMySQLWorklogRepository(db)

	worklogService := service.NewWorklogService(worklogRepository)

	worklogHandler := handler.NewWorklogHandler(worklogService)

	router.RegisterRoutes(worklogHandler)

	log.Println("Server Running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
