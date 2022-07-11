package main

import "fmt"

func main() {
	s := "123456789"
	s1, s2 := split(s, 5)
	fmt.Printf("s1=%s, s2=%s", s1, s2)
}

func split(s string, idx int) (s1 string, s2 string) {
	cs := []byte(s)
	s1 = string(cs[0:idx])
	s2 = string(cs[idx:])
	return
}
