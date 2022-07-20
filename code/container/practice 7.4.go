package container

import "fmt"

func main() {
	fmt.Println(fibslice(50))
}
func fibarray1(term int) []int {
	farr := make([]int, term)
	farr[0], farr[1] = 1, 1
	for i := 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}
	return farr
}
func fibslice(term int) []int {
	slice1 := fibarray1(term)[:term]
	return slice1
}
