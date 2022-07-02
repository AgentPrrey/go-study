package main

import "fmt"

func main() {
	str := "Jump to the sky turn to a rider kick!"
	fmt.Printf("The length of str is:%d\n ", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on the position %d is: %c \n", ix, str[ix])
	}

	str2 := "日本语"
	fmt.Printf("The length of str2 is:%d\n ", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on the position %d is: %c \n", ix, str2[ix])
	}
}
