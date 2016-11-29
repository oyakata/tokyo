package main

import "fmt"

type Employee interface {
	Labor()
}

func Pay(emp Employee) {
	emp.Labor()
}

func PayAll(emps ...*Employee) {
	for _, x := range emps {
		Pay(*x)
	}
}

type WageSlave struct {
	wage   int
	wallet int
}

func (w *WageSlave) Labor() {
	w.wallet += w.wage
}

func mainx() {
	var emp Employee
	var L []*Employee
	emp = &WageSlave{1, 0}
	fmt.Println(emp)
	fmt.Println(&emp)

	L = append(L, &emp)
	fmt.Println(L)
	PayAll(L...)

	// emps := []*Employee([]*WageSlave{
	// 	{wage: 5, wallet: 0},
	// 	{wage: 60, wallet: 12},
	// 	{wage: 23, wallet: 23},
	// })
	// PayAll(emps...)
	// fmt.Println(emps)

	var LL []*Employee
	var p, q, r Employee
	p = &WageSlave{wage: 5, wallet: 0}
	q = &WageSlave{wage: 60, wallet: 12}
	r = &WageSlave{wage: 23, wallet: 23}

	LL = append(LL, &p, &q, &r)
	fmt.Println(LL)
	PayAll(LL...)
}

type Point struct {
	x int
	y int
}

func (p *Point) Show() {
	fmt.Println("--------------------")
	fmt.Println(p.x, p.y)
	fmt.Println((*p).x, (*p).y)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", *p)
}

func main() {

	var x, y *Point
	x, y = &Point{0, 0}, &Point{1, 5}
	// fmt.Println(x, y)
	x.Show()
	y.Show()

	// var x, y *int
	// var z int
	// z = 10
	// x, y = &z, &z
	// fmt.Println(x, y)
}
