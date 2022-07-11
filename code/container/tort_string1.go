package main

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	n, h := len(runes), len(runes)/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
func main() {
	str := "ABCDF"
	str1 := reverse(str)
	fmt.Println(str1)
}
