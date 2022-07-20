package main

import (
	"fmt"
	"go-study/person"
)

func main() {
	p := new(person.Person)
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName())
}
