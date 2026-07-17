package repository

import "go-worklog-api/internal/domain"

type WorklogRepository interface {
	Save(worklog *domain.Worklog) error
}
