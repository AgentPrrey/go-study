package function

import "fmt"

var i, j int = 8, 2
var x, y, z int

func main() {
	x, y, z = xyz(i, j)
	look()
	x, y, z = xyz_2(i, j)
	look()
}
func look() {
	fmt.Printf("sum = %d,sub = %d, dif= %d\n", x, y, z)
}
func xyz(t1, t2 int) (int, int, int) {
	return t1 + t2, t1 * t2, t1 - t2
}
func xyz_2(t1, t2 int) (x1 int, y1 int, z1 int) {
	x1 = t1 + t2
	y1 = t1 * t2
	z1 = t1 - t2
	return
}
