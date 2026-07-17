package handler

import (
	"encoding/json"
	"net/http"

	"go-worklog-api/internal/domain"
	"go-worklog-api/internal/service"
)

type WorklogHandler struct {
	service *service.WorklogService
}

func NewWorklogHandler(
	service *service.WorklogService,
) *WorklogHandler {

	return &WorklogHandler{
		service: service,
	}
}

func (h *WorklogHandler) CreateWorklog(
	writer http.ResponseWriter,
	request *http.Request,
) {

	var worklog domain.Worklog

	err := json.NewDecoder(request.Body).Decode(&worklog)

	if err != nil {

		http.Error(
			writer,
			"Invalid JSON",
			http.StatusBadRequest,
		)

		return
	}

	err = h.service.SubmitWorklog(&worklog)

	if err != nil {

		http.Error(
			writer,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	writer.Header().Set(
		"Content-Type",
		"application/json",
	)

	writer.WriteHeader(http.StatusCreated)

	json.NewEncoder(writer).Encode(worklog)
}
