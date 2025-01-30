package main

import (
	"fmt"
	"math"

	"github.com/farshadb/go_crash_course.git/03_packages/strutil"
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
