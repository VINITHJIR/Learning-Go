package main

import "fmt"

func main() {

	const applicationName = "Student Management API"

	studentName := "Vinith"
	studentAge := 23
	studentEmail := "vinith@gmail.com"
	percentage := 85.5
	active := true

	fmt.Printf("Application: %s\n", applicationName)

	fmt.Printf("Student Name: %s | Type: %T\n", studentName, studentName)
	fmt.Printf("Student Age: %d | Type: %T\n", studentAge, studentAge)
	fmt.Printf("Student Email: %s | Type: %T\n", studentEmail, studentEmail)
	fmt.Printf("Percentage: %.2f | Type: %T\n", percentage, percentage)
	fmt.Printf("Active: %t | Type: %T\n", active, active)

}
