package main

import "fmt"

func foo(x int) int {
	return x + 10
}

// func foo(x float) float {
// 	return x * 10
// }

type Number interface {
	Show() int
}

type Foo int
type Bar int

func (x Foo) Show() int {
	return int(x) + 10
}

func (x Bar) Show() int {
	return int(x) * 10
}

type NumberWrapper Number

func (x NumberWrapper) Show int {
	fmt.Println("Hello")
	Number(x).Show()
	fmt.Println("world.")
}

func main(){
	fmt.Println(foo(8))
	fmt.Println(Foo(8).Show())
	fmt.Println(Bar(8).Show())

	var num Number
	num = Foo(17)
	fmt.Println(num.Show())
}
