package function2

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now()
	for i := 0; i <= 10; i++ {
		result = fibonacci1(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci1(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci1(n-1) + fibonacci1(n-2)
	}
	return

}
