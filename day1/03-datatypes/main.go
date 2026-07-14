package main

import "fmt"

func main() {
	studentName := "Vinith"
	studentAge := 23
	studentID := int64(100001)
	percentage := 85.75
	active := true

	fmt.Printf("Name       : %v | Type: %T\n", studentName, studentName)
	fmt.Printf("Age        : %v | Type: %T\n", studentAge, studentAge)
	fmt.Printf("Student ID : %v | Type: %T\n", studentID, studentID)
	fmt.Printf("Percentage : %v | Type: %T\n", percentage, percentage)
	fmt.Printf("Active     : %v | Type: %T\n", active, active)
}
