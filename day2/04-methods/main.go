package main

import "fmt"

// Custom type
type WorklogStatus string

// Constants
const (
	StatusDraft     WorklogStatus = "DRAFT"
	StatusSubmitted WorklogStatus = "SUBMITTED"
	StatusReviewed  WorklogStatus = "REVIEWED"
)

// Worklog struct
type Worklog struct {
	ID     int64
	Week   int
	Status WorklogStatus
	Rating int
}

// Normal Function
func canReviewWorklog(
	worklog Worklog,
) bool {

	return worklog.Status == StatusSubmitted
}

// Method
func (w Worklog) CanReview() bool {

	return w.Status == StatusSubmitted
}

// Method with Value Receiver
func (w Worklog) Review(
	rating int,
) string {

	fmt.Printf("Inside Review - Before: %+v\n", w)

	w.Status = StatusReviewed
	w.Rating = rating

	fmt.Printf("Inside Review - After : %+v\n", w)

	return "Worklog Reviewed Successfully"
}

func main() {

	worklog := Worklog{
		ID:     1001,
		Week:   2,
		Status: StatusSubmitted,
		Rating: 0,
	}

	// Normal function
	functionResult := canReviewWorklog(worklog)

	fmt.Println(
		"Normal Function Result:",
		functionResult,
	)

	// Method
	methodResult := worklog.CanReview()

	fmt.Println(
		"Method Result:",
		methodResult,
	)

	fmt.Println()

	// Before Review
	fmt.Println("Before Review")
	fmt.Printf("Original Worklog: %+v\n", worklog)

	fmt.Println()

	// Call Review method
	message := worklog.Review(4)

	fmt.Println()
	fmt.Println(message)

	// After Review
	fmt.Println()
	fmt.Println("After Review")
	fmt.Printf("Original Worklog: %+v\n", worklog)
}
