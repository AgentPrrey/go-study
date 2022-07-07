//写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上
package main

import (
	"fmt"
)


func main() {
	var number [16]int
	for i := 0; i < 16; i++ {
		number[i] = i
	}
	fmt.Println(number)
}
