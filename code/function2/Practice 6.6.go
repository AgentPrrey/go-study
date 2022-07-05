package main

import "fmt"

func main() {
	var n int
	//fmt.Scanln(&n)
	m := factorial(n)
	fmt.Printf("n=%d,m=%d\n", n, m)
}
func factorial(n int) (m int) {

	if n <= 1 {
		m = 1
	} else {
		m = n * factorial(n-1)
	}
	return
}
