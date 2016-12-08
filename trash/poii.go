package main

import "fmt"

type Point struct { x int; y int }

func main() {
	var p Point
	var q *Point

	p = Point{10, 20}
	q = &p
	r := *q

	q.x, q.y = 7, 8
	fmt.Println(p, q, r) // => {7 8} &{7 8} {10 20}

	point := func(p Point) Point { return p }
	point2 := func(p *Point) *Point { return p }
	point3 := func(p *Point) Point { return *p }

	L := Point{1, 2}
	M := point(L)
	N := point2(&M)
	O := point3(N)

	M.x, M.y = -5, -6
	fmt.Println(L, M, N, O) // {1 2} {-5 -6} &{-5 -6} {1 2}

	fmt.Printf("L> %#v(%T) %p\n", L, L, &L)
	fmt.Printf("M> %#v(%T) %p\n", M, M, &M)
	fmt.Printf("N> %#v(%T) %p\n", N, N, &N)
	fmt.Printf("O> %#v(%T) %p\n", O, O, &O)

	// 結果

	// L> main.Point{x:1, y:2}(main.Point) 0xc82000a370
	// M> main.Point{x:-5, y:-6}(main.Point) 0xc82000a380
	// N> &main.Point{x:-5, y:-6}(*main.Point) 0xc820026028
	// O> main.Point{x:1, y:2}(main.Point) 0xc82000a390
}
