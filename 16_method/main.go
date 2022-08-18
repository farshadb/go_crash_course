package main

import (
	"fmt"
)

type Engineer struct {
	Name	  string
	Age		  int
	Project	  Project
}

type Project struct {
	Name		 string
	Priority	 string
	Technologies []string
}

type Empty struct {}

func (e Engineer ) Print() {
	fmt.Println("Engineer Information: ")
	fmt.Printf("Name:           %s\n", e.Name)
	fmt.Printf("Age:            %d\n", e.Age)
	fmt.Printf("Current Project %s\n", e.Project.Name)
}

// Define a method Print() for the Engineer struct.
// Print() should print the Engineer's name, age and current project.
// Print() should be a method of the Engineer struct.
// Print() should be a method of the Engineer pointer.

func (e *Engineer ) UpdateAge() {
	e.Age += 1
}

func (e *Engineer ) GetProjectPriority() string {
	return e.Project.Priority
}





func main() {
	fmt.Println("Methods Tutorial")
	engieer := Engineer {
		Name:    "Farshad",
		Age:     32,
		Project: Project {
			Name:	      "Beginners Guide to Go",
			Priority:     "High",
			Technologies: []string{"Go"},
		},
	}

	//fmt.Println("engineer: ", engieer)

	engieer.Print()
	fmt.Println("  ")
	engieer.UpdateAge()
	engieer.Print()

	fmt.Println("\nGet Project Priority: ")
	fmt.Println(engieer.GetProjectPriority())
	fmt.Print("\n Empty Struct Print Test: \n")
	empty := Empty{}
	fmt.Println(empty)


}
