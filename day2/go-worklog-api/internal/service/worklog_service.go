package service

import (
	"errors"

	"go-worklog-api/internal/domain"
	"go-worklog-api/internal/repository"
)

type WorklogService struct {
	repository repository.WorklogRepository
}

func NewWorklogService(
	repository repository.WorklogRepository,
) *WorklogService {

	return &WorklogService{
		repository: repository,
	}
}

func (s *WorklogService) SubmitWorklog(
	worklog *domain.Worklog,
) error {

	if worklog.Week <= 0 {
		return errors.New("invalid week")
	}

	return s.repository.Save(worklog)
}
