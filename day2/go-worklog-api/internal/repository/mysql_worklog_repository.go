package repository

import (
	"database/sql"

	"go-worklog-api/internal/domain"
)

type MySQLWorklogRepository struct {
	db *sql.DB
}

func NewMySQLWorklogRepository(
	db *sql.DB,
) *MySQLWorklogRepository {

	return &MySQLWorklogRepository{
		db: db,
	}
}

func (r *MySQLWorklogRepository) Save(
	worklog *domain.Worklog,
) error {

	query := `
	INSERT INTO worklogs(
		week,
		status,
		rating
	)
	VALUES (?, ?, ?)
	`

	result, err := r.db.Exec(
		query,
		worklog.Week,
		worklog.Status,
		worklog.Rating,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	worklog.ID = id

	return nil
}
