package main

import "fmt"

type Employee interface {
	Pay() int
}

func PrintWage(e *Employee) {
	fmt.Println(e.Pay())
}

type WageSlave struct {
	id int
	name string
	wage int
}

func (self WageSlave) Pay() int {
	return self.wage / 12
}

func main() {
	emp := WageSlave{1, "foo", 60}

	// これはコンパイルエラー
	// e.Pay undefined (type *Employee is pointer to interface, not interface)
	// fmt.Println(emp.Pay())

	// これもコンパイルエラー
	// e.Pay undefined (type *Employee is pointer to interface, not interface)
	// cannot use &emp (type *WageSlave) as type *Employee in argument to PrintWage:
	// 	*Employee is pointer to interface, not interface
	// PrintWage(&emp)

	/*
		何がダメなのか？

		どうも、PrintWageが引数にEmployeeのポインタを取るのが良くないらしい。

		逆に、Employee.Pay()メソッドをWageSlaveで定義するとき、ポインタにしないのが良くない。
	*/
}
