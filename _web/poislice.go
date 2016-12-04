package main

import "fmt"

// インターフェイスを定義するときは関数の外でやる。
// 関数の中でやるとコンパイルエラーになる。
// non-declaration statement outside function body
type B interface {
	Print()
}
type Bar struct {
	id int
}
func (b *Bar) Print() {
	fmt.Println(b, b.id)
}

func main(){
	type Foo struct{
		id int
	}

	x := []Foo{
		{1},
		{2},
		{3},
	}

	fmt.Println(x)  // [{1} {2} {3}]
	x[1].id += 10
	(&x[1]).id += 10
	fmt.Println(x[0].id, x[1].id, x[2].id)

	y := []*Foo{
		{4},
		{5},
		{6},
	}

	fmt.Println(y)  // [0xc82000a358 0xc82000a360 0xc82000a368]
	y[1].id += 10
	// これはコンパイルエラー
	// (&y[1]).id undefined (type **Foo has no field or method id)
	// (&y[1]).id += 10

	// こっちは大丈夫
	(*y[1]).id += 10
	fmt.Println(y[0].id, y[1].id, y[2].id)

	/*
		結論

		[]*Fooと書くと、Fooのポインタのスライスを宣言したことになる。
			([]Fooと書けば当然Fooのスライス)
		こうしておくと、各要素に都度&を付けなくて良い。

		*Fooをレシーバとするメソッドを定義した場合、Fooのスライスから取り出す要素からは呼べない。
			(*Fooから呼べる)
	*/

	// インターフェイスを定義するときは関数の外でやる。
	// 関数の中でやるとコンパイルエラーになる。
	// non-declaration statement outside function body
	// type B interface {
	// 	Print()
	// }
	// type Bar struct {
	// 	id int
	// }
	// func (b *Bar) Print() {
	// 	fmt.Println(b, b.id)
	// }

	b := Bar{5678}
	fmt.Println(b)
	b.Print()
}
