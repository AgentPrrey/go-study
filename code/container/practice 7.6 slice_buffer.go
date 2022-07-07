package main

func main() {
	buf := [10]int{1, 2, 3, 4, 5}
	s, b := buf[:5], buf[5:]
	println(s)
	println(b)
}
