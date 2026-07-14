package main

import (
	"errors"
	"fmt"
)

// Worklog represents one weekly worklog
type Worklog struct {
	ID       int64
	Week     int
	Status   string
	Rating   int
	Reviewed bool
}

// Student represents one student with multiple worklogs
type Student struct {
	ID       int64
	Name     string
	Email    string
	Worklogs []Worklog
}

// Find the first submitted worklog
func findSubmittedWorklog(worklogs []Worklog) (Worklog, error) {

	for _, worklog := range worklogs {

		if worklog.Status == "SUBMITTED" {
			return worklog, nil
		}
	}

	return Worklog{}, errors.New("submitted worklog not found")
}

func main() {

	// Create one student with multiple worklogs
	student := Student{
		ID:    101,
		Name:  "Vinith",
		Email: "vinith@example.com",

		Worklogs: []Worklog{
			{
				ID:       1001,
				Week:     1,
				Status:   "REVIEWED",
				Rating:   4,
				Reviewed: true,
			},
			{
				ID:       1002,
				Week:     2,
				Status:   "SUBMITTED",
				Rating:   0,
				Reviewed: false,
			},
			{
				ID:       1003,
				Week:     3,
				Status:   "DRAFT",
				Rating:   0,
				Reviewed: false,
			},
		},
	}

	// Print student details
	fmt.Println("Student ID    :", student.ID)
	fmt.Println("Student Name  :", student.Name)
	fmt.Println("Student Email :", student.Email)

	// Print total worklogs
	fmt.Println("Total Worklogs:", len(student.Worklogs))

	fmt.Println()

	// Print every worklog
	fmt.Println("All Worklogs:")

	for _, worklog := range student.Worklogs {
		fmt.Printf(
			"Week: %d | Status: %s | Rating: %d | Reviewed: %t\n",
			worklog.Week,
			worklog.Status,
			worklog.Rating,
			worklog.Reviewed,
		)
	}

	fmt.Println()

	// Find submitted worklog
	submittedWorklog, err := findSubmittedWorklog(student.Worklogs)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Submitted Worklog Found")
	fmt.Printf("Worklog: %+v\n", submittedWorklog)

	fmt.Println()

	// Create Week 4 worklog
	newWorklog := Worklog{
		ID:       1004,
		Week:     4,
		Status:   "DRAFT",
		Rating:   0,
		Reviewed: false,
	}

	// Add new worklog to student's worklogs
	student.Worklogs = append(
		student.Worklogs,
		newWorklog,
	)

	// Print new total count
	fmt.Println("Week 4 Worklog Added")
	fmt.Println("New Total Worklogs:", len(student.Worklogs))
}
