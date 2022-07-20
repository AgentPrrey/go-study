package container

import "fmt"

func main() {
	sum := [...]int{1, 2, 3}
	for i := 3; i < 5; i++ {
		sum[i] = i
	}
	fmt.Println(sum)
}
