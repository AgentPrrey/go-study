package main

import "fmt"

func main() {
	test(10)
}
func test(n int) {
	if n == 0 {
		return
	}
	fmt.Printf("%d\n", n)
	test(n - 1)
}
