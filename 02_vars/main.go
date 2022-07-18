package main

import "fmt"

func main() {
	
	//techicaly, is not necessary to use the string  or var keyword.
	var name string = "Farshad"
	var age int64 = 30
	var size float64 = 1.8

	// Shorthand Definition for the above variables.
	lastName := "Bolouri"
	size2 := 1.8
	name2, email := "Farshad", "Bolouri.Farshad@gmai.com"



	fmt.Println("Hello, ", name, lastName,  "! You are ", age, " years old.")
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Println(size, size2)
	fmt.Println(name2, email)

}