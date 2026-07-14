package main

import "fmt"

func main() {

	Application_Name := "Bridge.ai Training API"
	Port := 8080
	Environment := "Development"
	Batch_Capacity := 100
	API_Status := true

	fmt.Printf("Application Name: %s\n", Application_Name)
	fmt.Printf("Port: %d\n", Port)
	fmt.Printf("Environment: %s\n", Environment)
	fmt.Printf("Batch Capacity: %d\n", Batch_Capacity)
	fmt.Printf("API Status: %t\n", API_Status)
	fmt.Println("Server Started Successfully")

}
