package main

import (
	"errors"
	"fmt"
)

// Custom Worklog Status Type
type WorklogStatus string

// Worklog Status Constants
const (
	StatusDraft     WorklogStatus = "DRAFT"
	StatusSubmitted WorklogStatus = "SUBMITTED"
	StatusReviewed  WorklogStatus = "REVIEWED"
)

// Worklog Struct
type Worklog struct {
	ID     int64
	Week   int
	Status WorklogStatus
	Rating int
}

// Constructor-style Function
// Returns a pointer to a new Worklog
func NewWorklog(
	id int64,
	week int,
) *Worklog {

	return &Worklog{
		ID:     id,
		Week:   week,
		Status: StatusDraft,
		Rating: 0,
	}
}

// Submit Worklog
// Pointer receiver modifies the original Worklog
func (w *Worklog) Submit() error {

	if w.Status != StatusDraft {
		return errors.New("only DRAFT worklog can be submitted")
	}

	w.Status = StatusSubmitted

	return nil
}

// Review Worklog
// Pointer receiver modifies the original Worklog
func (w *Worklog) Review(
	rating int,
) error {

	if w.Status != StatusSubmitted {
		return errors.New("only SUBMITTED worklog can be reviewed")
	}

	if rating < 1 || rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	w.Status = StatusReviewed
	w.Rating = rating

	return nil
}

// Find Worklog
// Returns *Worklog if found
// Returns nil if not found
func findWorklog(
	worklogs []*Worklog,
	searchID int64,
) *Worklog {

	for _, worklog := range worklogs {

		if worklog.ID == searchID {
			return worklog
		}
	}

	return nil
}

func main() {

	// Create a new Worklog
	worklog := NewWorklog(1001, 1)

	fmt.Println("New Worklog:")
	fmt.Println("ID     :", worklog.ID)
	fmt.Println("Week   :", worklog.Week)
	fmt.Println("Status :", worklog.Status)
	fmt.Println("Rating :", worklog.Rating)

	// Submit Worklog
	err := worklog.Submit()

	if err != nil {
		fmt.Println("Submit Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("After Submit:")
	fmt.Println("Status:", worklog.Status)

	// Review Worklog
	err = worklog.Review(4)

	if err != nil {
		fmt.Println("Review Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("After Review:")
	fmt.Println("Status:", worklog.Status)
	fmt.Println("Rating:", worklog.Rating)

	// Create a slice of Worklog pointers
	worklogs := []*Worklog{
		worklog,
	}

	// Search for a Worklog that does not exist
	missingWorklog := findWorklog(worklogs, 9999)

	fmt.Println()
	fmt.Println("Missing Worklog:")

	if missingWorklog == nil {
		fmt.Println("Worklog Not Found")
	}
}
