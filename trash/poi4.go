package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	p := Point{0, 0}
	q := &p
	r := *q // rにはqの大元の変数p(の値)を保持しておく。rとpは別のメモリが割り当てられる。

	fmt.Printf("p: %p q: %p  r:%p\n", &p, q, &r) // p: 0xc82000a320 q: 0xc82000a320  r:0xc82000a330

	p.x = 8
	fmt.Println(p, q, r) // {3 10} &{3 10} {0 0}

	// ここで、qの大元の変数に新しいPointのオブジェクトを入れる。
	// -> pに代入しているのと同じで、なおかつ、rには影響しない。
	*q = Point{3, 3} // 左辺に間接参照演算子「*」を持ってこられる点に驚き。

	p.y += 7

	fmt.Println(p, q, r) // {3 10} &{3 10} {0 0}
}
