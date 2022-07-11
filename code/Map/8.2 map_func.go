package main

import "fmt"

func main() {
	mf := map[string]func() int{
		"one":   func() int { return 100 },
		"two":   func() int { return 200 },
		"three": func() int { return 300 },
	}
	fmt.Println(mf)
}

//map[1:0x1089a00 2:0x1089a20 3:0x1089a40]
