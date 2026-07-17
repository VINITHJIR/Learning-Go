package router

import (
	"net/http"

	"go-worklog-api/internal/handler"
)

func RegisterRoutes(
	worklogHandler *handler.WorklogHandler,
) {

	http.HandleFunc(
		"/worklogs",
		worklogHandler.CreateWorklog,
	)
}
