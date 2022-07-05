package main

import "fmt"

func main() {
	fibonacci2(10)
}
func fibonacci2(n int) (index int, value int) {
	index = n
	if n <= 1 {
		value = 1
	} else {
		_, value0 := fibonacci2(n - 1)
		_, value1 := fibonacci2(n - 2)
		value = value0 + value1
	}
	fmt.Printf("index:%d,value:%d\n", n, value)
	return
}
