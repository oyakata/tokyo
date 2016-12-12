package main

import (
	"fmt"
)


var naisho = make([]*int, 0)

func NewFoo(value int) *int {
	naisho = append(naisho, &value)
	return &value
}

func Add(x int) {
	for _, y := range naisho {
		*y += x  // こういうことをするときに*が必要なのか
	}
}

func main() {
	x, y, z := NewFoo(70), NewFoo(65), NewFoo(-8)
	Add(10)

	*x = 99
	fmt.Println(*x, *y, *z)
	for _, n := range naisho {
		fmt.Println(*n)
	}
}
