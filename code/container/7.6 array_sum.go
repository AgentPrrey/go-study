package container

import "fmt"

func main() {
	array := [4]float64{7.0, 8.5, 9.1, 6.0}
	x := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[4]float64) (sum float64) {
	for _, v := range *a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
