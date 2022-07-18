package main 

import "fmt"

func main() {
	
	// Long method
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Short method
	for i := 0; i < 10; i++ {
		fmt.Println("number", i)
	}

	// Range method
	for i, v := range "Hello" {
		fmt.Println(i, v)
	}

	// Range method with arrays
	a := [3]int32{12, 13, 14}
	for i, v := range a {
		fmt.Println(i, v)
	}

	// Range method with maps
	m := map[string]int32{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Range method with slices
	s := []int32{12, 13, 14}
	for i, v := range s {
		fmt.Println(i, v)
	} 

	// Fizzbuzz
	for i := 1; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}


