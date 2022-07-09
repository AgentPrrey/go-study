package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3, 4}
	b := append(a, 5)
	c := append(a[:1], b...)

	fmt.Println(len(a), ",", cap(a))
	fmt.Println(len(b), ",", cap(b))
	fmt.Println(len(c), ",", cap(c))
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(a)

	vB := reflect.ValueOf(&b)
	vC := reflect.ValueOf(&c)
	vA := reflect.ValueOf(&a)

	fmt.Println(vB.Elem().Pointer())
	fmt.Println(vC.Elem().Pointer())
	fmt.Println(vA.Elem().Pointer())

}
