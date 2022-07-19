package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Loop through ids and print each value.
	for i, id := range ids {
		fmt.Printf("%d -ID: %d\n", i, id)

	}
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id

	}
	fmt.Println("Sum:", sum)

	// Range with map
	emails := map[string]string{"Farshad": "Bolouri@gmail.com", "Sara": "Sara@gmail.com", "John": "John@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name:", k)
	}
}
