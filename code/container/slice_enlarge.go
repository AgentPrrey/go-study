package main

import "fmt"

var s []int

func main() {
	s = []int{1, 2, 3}
	fmt.Printf("The length of s before enlarging is %d", len(s))
	fmt.Println(s)
	s = enlarge(s, 5)
	fmt.Printf("The length of s after enlarging is %d", len(s))
	fmt.Println(s)
}
func enlarge(s []int, factor int) []int {
	l := len(s)
	nl := l * factor
	ns := make([]int, nl)
	copy(ns, s)
	s = ns
	return s
}
