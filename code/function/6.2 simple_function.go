package function

import "fmt"

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", Multipiy3Nums(2, 5, 6))
}
func Multipiy3Nums(a int, b int, c int) int {
	return a * b * c
}
