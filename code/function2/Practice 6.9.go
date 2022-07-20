package function2

import "fmt"

const ai = 40

func main() {
	var a [ai]int
	fi := fibonacci3()
	for j := 0; j < ai; j++ {
		a[j] = fi()
	}
	fmt.Println(a)
}
func fibonacci3() func() int {

	back1, back2 := 0, 1
	return func() int {
		back1, back2 = back2, back1+back2
		return back1
	}

}
