package main

import "fmt"

type Point struct {
	x int
	y int
}

func NewPointer(x int, y int) (p *Point) {
	fmt.Printf("??> %#v(%T) %p\n", p, p, p)
	p = &Point{x, y}
	return p
}

func NewIndirect(p *Point) (P Point) {
	return *p
}

func NewDirect(p *Point) (P *Point) {
	return p
}

func main() {
	p, q := NewPointer(1, 2), NewPointer(3, 4)
	fmt.Printf("%#v(%T) %p\n", p, p, p)
	fmt.Printf("%#v(%T) %p\n", q, q, q)
	fmt.Println()

	// 出力内容

	// ??> (*main.Point)(nil)(*main.Point) 0x0
	// ??> (*main.Point)(nil)(*main.Point) 0x0
	// &main.Point{x:1, y:2}(*main.Point) 0xc82000a340
	// &main.Point{x:3, y:4}(*main.Point) 0xc82000a360

	// NewPointerの戻り値pはnilで初期化されていることが0x0の出力から判る。
	// このため、名前付き結果パラメータから提供される変数を書き換えて返却してもすべて別管理なので心配いらない。

	r, s := NewIndirect(p), NewIndirect(p)
	fmt.Printf("%#v(%T) %p\n", r, r, &r)
	fmt.Printf("%#v(%T) %p\n", s, s, &s)
	fmt.Println()

	r.x, s.y = 77, 99
	fmt.Printf("p> %#v(%T) %p\n", p, p, p)
	fmt.Printf("q> %#v(%T) %p\n", q, q, q)
	fmt.Printf("r> %#v(%T) %p\n", r, r, &r)
	fmt.Printf("s> %#v(%T) %p\n", s, s, &s)
	fmt.Println()

	// 出力内容

	// main.Point{x:1, y:2}(main.Point) 0xc82000a3d0
	// main.Point{x:1, y:2}(main.Point) 0xc82000a3e0

	// p> &main.Point{x:1, y:2}(*main.Point) 0xc82000a340
	// q> &main.Point{x:3, y:4}(*main.Point) 0xc82000a360
	// r> main.Point{x:77, y:2}(main.Point) 0xc82000a3d0
	// s> main.Point{x:1, y:99}(main.Point) 0xc82000a3e0

	// rのフィールドを書き換えてもpに影響しない。(ｺﾚｶﾞﾜｶﾗﾅｲ)
	// => たぶん、ポインタのエイリアスを返す(=戻り値がPointであって*Pointではない)ので、
	//    通常のルールに則りPointの複製を返していると推測される。

	t := NewDirect(&s)
	t.x, t.y = 30, 45
	fmt.Printf("p> %#v(%T) %p\n", p, p, p)
	fmt.Printf("q> %#v(%T) %p\n", q, q, q)
	fmt.Printf("r> %#v(%T) %p\n", r, r, &r)
	fmt.Printf("s> %#v(%T) %p\n", s, s, &s)
	fmt.Printf("t> %#v(%T) %p\n", t, t, &t)

	// 出力内容

	// p> &main.Point{x:1, y:2}(*main.Point) 0xc82000a340
	// q> &main.Point{x:3, y:4}(*main.Point) 0xc82000a360
	// r> main.Point{x:77, y:2}(main.Point) 0xc82000a3d0
	// s> main.Point{x:30, y:45}(main.Point) 0xc82000a3e0
	// t> &main.Point{x:30, y:45}(*main.Point) 0xc820026028

	// 今度はsがtの変更から影響を受けている。
	// => 戻り値が*Pointなので、複製せず返却されている。
}
