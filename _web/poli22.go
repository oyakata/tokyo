package main

import "fmt"

type Employee interface {
	Pay() int
}

func PrintWage(e Employee) {
	fmt.Println(e.Pay())
}

type WageSlave struct {
	id int
	name string
	wage int
}

func (self *WageSlave) Pay() int {
	return self.wage / 12
}

func main() {
	emp := &WageSlave{1, "foo", 60}
	fmt.Printf("address? -> %p", emp)
	fmt.Println(emp.Pay())
	PrintWage(emp)
}
