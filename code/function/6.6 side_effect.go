package function

import (
	"fmt"
	"time"
)

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	Start := time.Now()
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	end := time.Now()
	delta := end.Sub(Start)
	fmt.Println("The time:", delta)
}
