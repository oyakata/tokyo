package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	p := Point{0, 0}
	q := &p
	r := *q // rにはqの大元の変数pを保持しておく。

	// ここで、qの大元の変数に新しいPointのオブジェクトを入れる。
	// -> pに代入しているのと同じで、なおかつ、rには影響しない。
	*q = Point{3, 3} // 左辺に間接参照演算子「*」を持ってこられる点に驚き。

	p.y += 7

	fmt.Println(p, q, r) // {3 10} &{3 10} {0 0}
}
