package main

import "fmt"

func main() {

	// Task 1 - Weekly Schedule

	for week := 1; week <= 12; week++ {
		fmt.Println("Creating Weekly Schedule: Week", week)
	}

	// Task 2 - Worklog Batch Validation

	worklogs := []string{
		"SUBMITTED",
		"DRAFT",
		"SUBMITTED",
		"INVALID",
		"REVIEWED",
		"SUBMITTED",
	}

	for _, status := range worklogs {

		if status == "DRAFT" {
			fmt.Println("Draft - Skipping")
			continue
		}

		if status == "INVALID" {
			fmt.Println("Invalid - Skipping")
			continue
		}

		if status == "REVIEWED" {
			fmt.Println("Already Reviewed - Skipping")
			continue
		}

		if status == "SUBMITTED" {
			fmt.Println("Processing Worklog")
		}
	}

	// Task 3 - Student Search

	studentIDs := []int64{
		101,
		102,
		103,
		104,
		105,
	}

	searchID := int64(104)

	for _, studentID := range studentIDs {

		fmt.Println("Checking Student ID:", studentID)

		if studentID == searchID {
			fmt.Println("Student Found")
			break
		}
	}

	// Task 4 - Nested Loop

	for week := 1; week <= 3; week++ {

		fmt.Println("Week", week)

		for question := 1; question <= 5; question++ {
			fmt.Println("  Question", question)
		}
	}
}
