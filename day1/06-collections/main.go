package main

import "fmt"

func main() {

	// Task 1 - Array

	ratings := [5]int{4, 5, 3, 4, 5}

	fmt.Println("First Rating :", ratings[0])
	fmt.Println("Last Rating  :", ratings[4])

	// Task 2 - Slice

	students := []string{
		"Vinith",
		"Arun",
		"Jayanthi",
	}

	students = append(students, "Kumar")

	fmt.Println("Student Count  :", len(students))
	fmt.Println("Slice Capacity :", cap(students))

	fmt.Println("All Students:")

	for _, student := range students {
		fmt.Println(student)
	}

	// Task 3 - JWT Claims Map

	claims := map[string]string{
		"user_id": "101",
		"role":    "FACULTY",
	}

	role, exists := claims["role"]

	if !exists {
		fmt.Println("Role Claim Missing")
		return
	}

	switch role {
	case "ADMIN":
		fmt.Println("Admin Access")

	case "FACULTY":
		fmt.Println("Faculty Review Access")

	case "STUDENT":
		fmt.Println("Student Access")

	default:
		fmt.Println("Invalid Role")
	}
}
