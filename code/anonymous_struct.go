package main

import "fmt"

//创建一个结构体，它有一个具名的 float 字段，2 个匿名字段，类型分别是 int 和 string。通过结构体字面量新建一个结构体实例并打印它的内容
type struct2 struct {
	f1 float32
	int
	string
}

func main() {
	a := new(struct2)
	a.f1 = 1.2
	a.int = 2
	a.string = "hello"
	fmt.Printf("a:%v", a)
}
