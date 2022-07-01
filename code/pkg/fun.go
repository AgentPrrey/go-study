package pkg

import (
	"fmt"
	"strconv"
)

func TestErr() {
	var orig string = "ABC"
	var newS string
	fmt.Printf("The size of ints is:%d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %S is not an integar - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integar is %d/", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is:%s\n", newS)
}
