package main

import (
	"fmt"
	"strconv"
)

// Define a struct type named Person with two fields, name and age.
// Go has no Clases  and structs are using instance of classes in Golang.

//type Person struct {
//	fistName string
//	lastName string
//	city     string
//	gender   string
//	age      int
//}

// Better Define structs are

type Person struct {
	fistName, lastName, city, gender string
	age                              int
}

// Greetting method (value reciver)
func (p Person) greet() string {
	return "Hello, My name is" + p.fistName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old" + " and I live in " + p.city

}

// birthday method (pointer receiver)
// Because its point receiver we use *Person and Because it changes something Function should return anything.
func (p *Person) hasBirthday() {
	p.age++
}

// Getmarried method (pointer receiver)

func (p *Person) getMarried(spouseLastName string) {
	if p.lastName == spouseLastName {
		return
	}
	p.lastName = spouseLastName
}

func main() {
	// Init Person using struct type
	person := Person{fistName: "Farshad", lastName: " Bolouri", city: " NYC", gender: "Male", age: 32}
	// Alternative initial Person struct without  fields.
	person2 := Person{"Brad", "Travesi", "Palo Alto", "Male", 42}

	//fmt.Println(person)
	//fmt.Println(person2)
	//fmt.Println(person.fistName)
	//person.age++
	//fmt.Println(person.age)
	fmt.Println(person.greet())
	person.hasBirthday()

	fmt.Println(person.age)

	person2.getMarried("Williams")
	fmt.Println(person2)

}
