package main

import (
	"fmt"
	"github.com/F4r5h4d/go_crash_course/03_packages/strutil"
	"math"
)

func main() {

	fmt.Println(math.Floor(2.78))
	fmt.Println(math.Ceil(2.78))
	fmt.Println(math.Sqrt(2.78))
	fmt.Println(math.Pi)
	fmt.Println(math.Sin(2.78))
	fmt.Println(math.Atan2(2.78, 3.78))
	fmt.Println(math.Pow(2, 10))
	fmt.Println(strutil.Reverse("Farshad"))

}
