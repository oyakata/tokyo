package main

import "fmt"

// func incr(i *int) int{
// 	i++
// 	return *i
// }

// func incr2(i int) int{
// 	i++
// 	// 引数iはポインタとして宣言されていないので、関数の中で間接参照を試みるとコンパイルエラー
// 	// invalid indirect of i (type int)
// 	return *i
// }

func incr3(i int) int {
	i++
	return i
}

func incr4(i *int) int {
	*i++
	return *i
}

func main() {
	x := 1

	// incrの引数iはintのポインタで宣言しているので変数のまま渡そうとするとコンパイルエラー
	// cannot use x (type int) as type *int in argument to incr
	// y := incr(x)

	// incr2の方がコンパイルエラー
	// y := incr2(x)

	y := incr3(x)
	fmt.Println(x, y) // => 1 2

	z := incr4(&x)
	fmt.Println(x, y, z) // => 2 2 2

	// これもダメ
	// 変数xはもはや間接参照先が無いのでコンパイルエラー
	// invalid indirect of x (type int)
	// fmt.Println(*x)

	// なので、一度&を通してから*を利用するとコンパイルが通る
	fmt.Println(*(&x))  // => 2
}
