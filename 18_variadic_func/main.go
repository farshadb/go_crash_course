package main

import "fmt"

// Variadic function to calculate sum
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func print(values ...string) {
	mainPrintFunc(values...)
}

func mainPrintFunc(values ...string) {
	for i, value := range values {
		fmt.Printf("%d, %s\n", i, value)
	}
}

func printAny(values ...any) {
	mainPrintAny(values...)
}

func mainPrintAny(values ...any) {
	for i, value := range values {
		fmt.Printf("%d, %s\n", i, value)
	}
}

func main() {
	print()
	print("one", "two", "three")
	print("eeeeeeee", "2", "3", "4", "100")
	printAny("eeeeeeee", "2", "3", "4", "100", "one", "two", "three", 1, 10, false)

	fmt.Println("Sum of 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println("Sum of 4, 5:", sum(4, 5))
	fmt.Println("Sum of no numbers:", sum())
}
