//定义结构体 employee，它有一个 salary 字段，给这个结构体定义一个方法 giveRaise 来按照指定的百分比增加薪水。
package main

import "fmt"

type employee struct {
	salary float32
}

func (s *employee) giveRaise(up float32) float32 {
	return s.salary * (1 + up)
}
func main() {
	m := new(employee)
	m.salary = 1000
	fmt.Printf("The salary is %f\n", m.giveRaise(0.1))
}
