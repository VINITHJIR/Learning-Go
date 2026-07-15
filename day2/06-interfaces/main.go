package main

import (
	"errors"
	"fmt"
)

// ========================================
// DOMAIN
// ========================================

type WorklogStatus string

const (
	StatusDraft     WorklogStatus = "DRAFT"
	StatusSubmitted WorklogStatus = "SUBMITTED"
	StatusReviewed  WorklogStatus = "REVIEWED"
)

type Worklog struct {
	ID     int64
	Week   int
	Status WorklogStatus
	Rating int
}

// Domain Method
func (w *Worklog) Review(
	rating int,
) error {

	if w.Status != StatusSubmitted {
		return errors.New(
			"only SUBMITTED worklog can be reviewed",
		)
	}

	if rating < 1 || rating > 5 {
		return errors.New(
			"rating must be between 1 and 5",
		)
	}

	w.Status = StatusReviewed
	w.Rating = rating

	return nil
}

// ========================================
// REPOSITORY INTERFACE
// ========================================

type WorklogRepository interface {
	FindByID(
		id int64,
	) (*Worklog, error)

	Update(
		worklog *Worklog,
	) error
}

// ========================================
// MEMORY REPOSITORY
// ========================================

type MemoryWorklogRepository struct {
	worklogs map[int64]*Worklog
}

// Constructor
func NewMemoryWorklogRepository() *MemoryWorklogRepository {

	return &MemoryWorklogRepository{
		worklogs: make(map[int64]*Worklog),
	}
}

// Save a new Worklog
func (r *MemoryWorklogRepository) Save(
	worklog *Worklog,
) error {

	r.worklogs[worklog.ID] = worklog

	return nil
}

// Find Worklog by ID
func (r *MemoryWorklogRepository) FindByID(
	id int64,
) (*Worklog, error) {

	worklog, exists := r.worklogs[id]

	if !exists {
		return nil, errors.New("worklog not found")
	}

	return worklog, nil
}

// Update existing Worklog
func (r *MemoryWorklogRepository) Update(
	worklog *Worklog,
) error {

	_, exists := r.worklogs[worklog.ID]

	if !exists {
		return errors.New("worklog not found")
	}

	r.worklogs[worklog.ID] = worklog

	return nil
}

// ========================================
// SERVICE
// ========================================

type WorklogService struct {
	repository WorklogRepository
}

// Service Constructor
func NewWorklogService(
	repository WorklogRepository,
) *WorklogService {

	return &WorklogService{
		repository: repository,
	}
}

// Review Worklog Use Case
func (s *WorklogService) ReviewWorklog(
	id int64,
	rating int,
) error {

	// Step 1 - Find Worklog
	worklog, err := s.repository.FindByID(id)

	if err != nil {
		return err
	}

	// Step 2 - Apply Domain Logic
	err = worklog.Review(rating)

	if err != nil {
		return err
	}

	// Step 3 - Update Worklog
	err = s.repository.Update(worklog)

	if err != nil {
		return err
	}

	return nil
}

// ========================================
// MAIN
// ========================================

func main() {

	// Step 1 - Create Memory Repository
	repository := NewMemoryWorklogRepository()

	// Step 2 - Create SUBMITTED Worklog
	worklog := &Worklog{
		ID:     1001,
		Week:   2,
		Status: StatusSubmitted,
		Rating: 0,
	}

	// Step 3 - Save Worklog
	err := repository.Save(worklog)

	if err != nil {
		fmt.Println("Save Error:", err)
		return
	}

	fmt.Println("Before Review:")
	fmt.Printf("%+v\n", worklog)

	// Step 4 - Inject Repository into Service
	service := NewWorklogService(repository)

	// Step 5 - Review Worklog
	err = service.ReviewWorklog(1001, 4)

	if err != nil {
		fmt.Println("Review Error:", err)
		return
	}

	// Step 6 - Find Updated Worklog
	updatedWorklog, err := repository.FindByID(1001)

	if err != nil {
		fmt.Println("Find Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("After Review:")
	fmt.Printf("%+v\n", updatedWorklog)
}
