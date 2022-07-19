package main

import "fmt"

func main() {

	a := 5
	b := &a
	fmt.Println(a, b)
	fmt.Println(&a, b)

	fmt.Printf("%T\n", b)

	// Use * to read value from address
	fmt.Println(*b)
	fmt.Println(*&a)
	// Change value with pointer
	// Pass the address instead of data itself it can increase the performance
	*b = 20
	fmt.Println(a)
}
