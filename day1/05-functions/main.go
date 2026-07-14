package main

import (
	"errors"
	"fmt"
)

// Function 1 - Calculate Total
func calculateTotal(
	python int,
	java int,
	database int,
	network int,
	operatingSystem int,
) int {

	total := python + java + database + network + operatingSystem

	return total
}

// Function 2 - Calculate Percentage
func calculatePercentage(total int) float64 {

	percentage := float64(total) / 5

	return percentage
}

// Function 3 - Get Result
func getResult(percentage float64) (bool, string) {

	if percentage >= 50 {
		return true, "PASS"
	}

	return false, "FAIL"
}

// Function 4 - Find Student
func findStudent(studentID int64) (string, error) {

	if studentID == 101 {
		return "Vinith", nil
	}

	return "", errors.New("student not found")
}

func main() {

	studentName, err := findStudent(101)

	if err != nil {
		fmt.Println(err)
		return
	}

	total := calculateTotal(85, 80, 90, 85, 85)

	percentage := calculatePercentage(total)

	passed, result := getResult(percentage)

	fmt.Println("Student Name :", studentName)
	fmt.Println("Total        :", total)
	fmt.Printf("Percentage   : %.2f\n", percentage)
	fmt.Println("Passed       :", passed)
	fmt.Println("Result       :", result)
}
