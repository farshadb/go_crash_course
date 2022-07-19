package main

import "fmt"

func main() {

	// Define a map.
	// First Data Type: Key and second Data Type: Value
	emails := make(map[string]string)

	// Assign key values.
	emails["Farshad"] = "Bolouri.Farshad@gmail.com"
	emails["Sara"] = "Sara@gmail.com"
	emails["John"] = "John@gmail.com"

	// Declare and assign key values at the same time.
	emails2 := map[string]string{"Farshad": "Bolouri@gmail.com", "Sara": "Sara@gmail.com", "John": "John@gmail.com"}
	fmt.Println(emails)
	fmt.Println(len(emails))
	delete(emails, "John")
	fmt.Println(emails["Farshad"])
	fmt.Println(emails)
	fmt.Println(emails2)

}
