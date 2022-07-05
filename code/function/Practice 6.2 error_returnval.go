package main

import (
	"fmt"
	"math"
)

var m float64 = -3
var n float64

func main() {
	if m < 0 {
		fmt.Println("You need to make sure m >= 0")
	} else {
		n = MySqrt(m)
		see()
		n = MySqrt_2(m)
		see()
	}
}
func see() {
	fmt.Printf("The sqrt is %f\n ", n)
}
func MySqrt(p float64) float64 {
	return math.Sqrt(p)
}
func MySqrt_2(p float64) (q float64) {
	q = math.Sqrt(p)
	return
}
