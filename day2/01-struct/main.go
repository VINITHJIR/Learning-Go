package main

import "fmt"

// Student struct
type Student struct {
	ID     int64
	Name   string
	Email  string
	Active bool
}

// Worklog struct
type Worklog struct {
	ID        int64
	StudentID int64
	Week      int
	Status    string
	Rating    int
	Reviewed  bool
}

func main() {

	// Create one Student
	student := Student{
		ID:     101,
		Name:   "Vinith",
		Email:  "vinith@example.com",
		Active: true,
	}

	// Create one Worklog
	worklog := Worklog{
		ID:        1001,
		StudentID: 101,
		Week:      3,
		Status:    "SUBMITTED",
		Rating:    0,
		Reviewed:  false,
	}

	// Print individual Student fields
	fmt.Println("Student ID     :", student.ID)
	fmt.Println("Student Name   :", student.Name)
	fmt.Println("Student Email  :", student.Email)
	fmt.Println("Student Active :", student.Active)

	// Print complete Student struct
	fmt.Printf("Complete Student: %+v\n", student)

	fmt.Println()

	// Print individual Worklog fields
	fmt.Println("Worklog ID        :", worklog.ID)
	fmt.Println("Student ID        :", worklog.StudentID)
	fmt.Println("Week              :", worklog.Week)
	fmt.Println("Status            :", worklog.Status)
	fmt.Println("Rating            :", worklog.Rating)
	fmt.Println("Reviewed          :", worklog.Reviewed)

	// Print Worklog before update
	fmt.Printf("Worklog Before Update: %+v\n", worklog)

	// Update Worklog
	worklog.Status = "REVIEWED"
	worklog.Rating = 4
	worklog.Reviewed = true

	// Print Worklog after update
	fmt.Printf("Worklog After Update : %+v\n", worklog)
}
