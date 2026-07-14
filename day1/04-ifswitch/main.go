package main

import "fmt"

func main() {
	role := "FACULTYg"
	isActive := true
	currentWeek := 3
	scheduleWeek := 3

	if !isActive {
		fmt.Println("Account Inactive")
		return
	}

	switch role {

	default:
		fmt.Println("Invalid Role")
		return

	case "ADMIN":
		fmt.Println("Admin Access Granted")

	case "FACULTY":
		fmt.Println("Faculty Access Granted")

	case "STUDENT":
		fmt.Println("Student Access Granted")

	}

	if scheduleWeek == currentWeek {
		fmt.Println("Weekly Log is Active")
	} else if scheduleWeek > currentWeek {
		fmt.Println("Weekly Log is Upcoming")
	} else {
		fmt.Println("Weekly Log is Closed")
	}
}
