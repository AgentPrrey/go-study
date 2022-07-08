package main

import "fmt"

func main() {
	fmt.Printf("min slice:%d\n", Minslice([]int{1, 2, 3, 4}))
	fmt.Printf("max slice:%d\n", Maxslice([]int{1, 2, 3, 4}))
}
func Minslice(mins []int) int {
	var min int = mins[0]
	for _, v := range mins {
		if v < min {
			min = v
		}
	}
	return min
}
func Maxslice(maxs []int) int {
	var max int = maxs[0]
	for _, v := range maxs {
		if v > max {
			max = v
		}
	}
	return max
}
