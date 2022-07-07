package main

import "fmt"

func main() {
	var ages [5]int = [5]int{1, 2, 1, 2, 5}
	var b = ages
	b[4] = 1
	fmt.Println(ages)
	fmt.Println(b)
}
