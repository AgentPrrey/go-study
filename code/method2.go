package main

import "fmt"

type Intvector []int

func (v Intvector) sum() (s int) {
	for _, v := range v {
		s += v
	}
	return
}
func main() {
	fmt.Println(Intvector{1, 2, 3}.sum())
}