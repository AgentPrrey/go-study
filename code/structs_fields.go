package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "hello"
	fmt.Printf("%d\n", ms.i1)
	fmt.Printf("%f\n", ms.f1)
	fmt.Printf("%s\n", ms.str)
	fmt.Println(ms)
}
