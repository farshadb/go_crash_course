package main

import "fmt"

func main() {
	
	// In go arrays are fixed size.
	// The size of the array is defined by the first argument of the array declaration.
	// The second argument is the type of the array elements.
	// The third argument is the initial value of the array.
	// The fourth argument is the length of the array.
	// The fifth argument is the capacity of the array.
	// The sixth argument is the pointer to the array.
	// The seventh argument is the pointer to the first element of the array.
	// The eighth argument is the pointer to the last element of the array.
	// The ninth argument is the pointer to the next element of the array.
	// The tenth argument is the pointer to the previous element of the array.
	// The eleventh argument is the pointer to the first element of the array.
	// The twelfth argument is the pointer to the last element of the array.

	// var a [3]int 
	// var b [3]int64 = [3]int64{1, 2, 3}
	// var c [3]float64 = [3]float64{1.1, 2.2, 3.3}
	// var d [3]string = [3]string{"a", "b", "c"}
	// var e [3]bool = [3]bool{true, false, true}
	// var f [3]byte = [3]byte{1, 2, 3}
	// var g [3]rune = [3]rune{'a', 'b', 'c'}



	// fmt.Println(a)
	// fmt.Printf("%T\n", a)
	// fmt.Println(b)
	// fmt.Printf("%T\n", b)
	// fmt.Println(c)
	// fmt.Printf("%T\n", c)
	// fmt.Println(d)
	// fmt.Printf("%T\n", d)
	// fmt.Println(e)
	// fmt.Printf("%T\n", e)
	// fmt.Println(f)
	// fmt.Printf("%T\n", f)
	// fmt.Println(g)
	// fmt.Printf("%T\n", g)

	// Assignment of arrays.
	// create slices with make and specefic length and capacity.
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)





	var arrfruit [3]string = [3]string{"apple", "banana", "orange"}
	var arrfruit2 [3]string
	arrfruit2[0] = "apple"
	arrfruit2[1] = "banana"
	arrfruit2[2] = "orange"

	fmt.Println(arrfruit2)
	fmt.Println(arrfruit)

	//Slices are a reference to the array.
	//The size of the slice is defined by the first argument of the slice declaration.

	var fruitSlice []string = []string{"bluebery", "Ananas", "Aaple"}
	fruitSlice = append(fruitSlice, "Orange")
	fruitSlice = append(fruitSlice, "Andange")
	fruitSlice = append(fruitSlice, "Mango")
	fruitSlice2 := []string{"blueberry", "Cucamber", "Mango"}
	fruitSlice3 := []string{}


	fruitSlice3 = append(fruitSlice3, "blueberry")

	fmt.Println(fruitSlice)
	fmt.Printf("%T\n", fruitSlice)
	fmt.Println(fruitSlice2)
	fmt.Println(fruitSlice3)

	fmt.Println(len(fruitSlice2))
	//fmt.Println(fruitSlice2[2:4]) This line will cause an error.(panic: runtime error: slice bounds out of range)
 }

 func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
