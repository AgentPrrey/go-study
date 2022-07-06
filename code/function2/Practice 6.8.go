package main

import "fmt"

func main() {
	lv := func() { fmt.Println("Hello World") }
	lv()
	fmt.Printf("%T", lv)
}
