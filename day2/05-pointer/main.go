package main

import "fmt"

type WorklogStatus string

const (
	StatusSubmitted WorklogStatus = "SUBMITTED"
	StatusReviewed  WorklogStatus = "REVIEWED"
)

type Worklog struct {
	ID     int64
	Week   int
	Status WorklogStatus
	Rating int
}

// Value receiver - w is a copy
func (w Worklog) Review(rating int) {

	w.Status = StatusReviewed
	w.Rating = rating

	// Use the modified copy
	fmt.Println("Inside Review:")
	fmt.Printf("%+v\n", w)
}

func main() {

	worklog := Worklog{
		ID:     1001,
		Week:   2,
		Status: StatusSubmitted,
		Rating: 0,
	}

	fmt.Println("Before Review:")
	fmt.Printf("%+v\n", worklog)

	worklog.Review(4)

	fmt.Println("After Review:")
	fmt.Printf("%+v\n", worklog)
}
