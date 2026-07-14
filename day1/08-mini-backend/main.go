package main

import (
	"errors"
	"fmt"
)

const searchStudentID int64 = 102
const facultyRating int = 4

// Function 1 - Validate Role
func validateRole(role string) (bool, string) {

	switch role {

	case "ADMIN":
		return true, "Admin Access Granted"

	case "FACULTY":
		return true, "Faculty Review Access Granted"

	case "STUDENT":
		return false, "403 Forbidden"

	default:
		return false, "Invalid Role"
	}
}

// Function 2 - Find Student
func findStudent(
	studentIDs []int64,
	searchID int64,
) (int64, error) {

	for _, studentID := range studentIDs {

		if studentID == searchID {
			return studentID, nil
		}
	}

	return 0, errors.New("student not found")
}

// Function 3 - Find Submitted Worklog
func findSubmittedWorklog(
	worklogs []string,
) (string, error) {

	for _, status := range worklogs {

		if status == "DRAFT" {
			fmt.Println("Draft Worklog - Skipping")
			continue
		}

		if status == "INVALID" {
			fmt.Println("Invalid Worklog - Skipping")
			continue
		}

		if status == "REVIEWED" {
			fmt.Println("Already Reviewed - Skipping")
			continue
		}

		if status == "SUBMITTED" {
			return status, nil
		}
	}

	return "", errors.New("submitted worklog not found")
}

// Function 4 - Validate Rating
func validateRating(rating int) (bool, string) {

	if rating < 1 || rating > 5 {
		return false, "Invalid Rating"
	}

	return true, "Rating Accepted"
}

// Function 5 - Review Worklog
func reviewWorklog(
	studentID int64,
	status string,
	rating int,
) string {

	fmt.Println("Student ID     :", studentID)
	fmt.Println("Worklog Status :", status)
	fmt.Println("Rating         :", rating)

	return "Worklog Reviewed Successfully"
}

func main() {

	fmt.Println("Starting IECC Student Worklog Service")
	fmt.Println()

	// Database Students
	studentIDs := []int64{
		101,
		102,
		103,
	}

	// JWT Claims
	claims := map[string]string{
		"user_id": "501",
		"role":    "FACULTY",
	}

	// Database Worklogs
	worklogs := []string{
		"DRAFT",
		"INVALID",
		"SUBMITTED",
		"REVIEWED",
	}

	// Step 1 - Get JWT Role

	role, exists := claims["role"]

	if !exists {
		fmt.Println("401 Unauthorized")
		return
	}

	// Step 2 - Validate Role

	allowed, message := validateRole(role)

	fmt.Println(message)
	fmt.Println()

	if !allowed {
		return
	}

	// Step 3 - Find Student

	studentID, err := findStudent(
		studentIDs,
		searchStudentID,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Student Found:", studentID)
	fmt.Println()

	// Step 4 - Find Submitted Worklog

	status, err := findSubmittedWorklog(worklogs)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Submitted Worklog Found:", status)
	fmt.Println()

	// Step 5 - Validate Rating

	validRating, ratingMessage := validateRating(facultyRating)

	fmt.Println(ratingMessage)
	fmt.Println()

	if !validRating {
		return
	}

	// Step 6 - Review Worklog

	resultMessage := reviewWorklog(
		studentID,
		status,
		facultyRating,
	)

	fmt.Println(resultMessage)
}
