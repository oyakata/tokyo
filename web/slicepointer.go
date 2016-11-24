package main

import "fmt"

type Foo struct {
	id int
	name string
}

// ポインタを渡してアドレスを含む文字列を返す。
func id(obj *Foo) string {
	obj.id += 100
	return fmt.Sprintf("%v> %p -> %p", obj.id, obj, &obj)
}

// こちらはポインタではなく、オブジェクトを渡す。戻り値はアドレスを含む文字列。
func id2(obj Foo) string {
	obj.id += 100
	return fmt.Sprintf("%v> %p -> %p", obj.id, obj, &obj)
}

func main() {
	foo := Foo{1, "bob"}

	fmt.Println(id(&foo))
	fmt.Println(id(&foo))

	fmt.Println(id2(foo))
	fmt.Println(id2(foo))

	bar := &foo

	fmt.Println(id(bar))
	fmt.Println(id(bar))

	// ポインタのポインタを渡そうとするとコンパイルエラー
	// cannot use bar (type *Foo) as type Foo in argument to id2
	// fmt.Println(id2(bar))
	// fmt.Println(id2(bar))

	/*	結果

		p1. 引数objのポインタが左側に来るので、左側の値は１行目と２行目で同じ。
		但し、ふたつのポインタ自体は別物なので、右側に表示される「ポインタそのもののアドレス」は違う。		
		また、ポインタを引数に受けてobj.idを書き換えるので、プロンプト部分の数値は呼び出しごとに101 -> 201と、繰り上がる。

		101> 0xc820010420 -> 0xc82002e020
		201> 0xc820010420 -> 0xc82002e030

		p2. 引数objはポインタではなく別個作成された仮引数なので、右側に表示される値は３行目と４行目で違う。
		左側がなぜこういう表記になってしまうのかは謎。
		仮引数と呼び出し側の変数fooは別物なので、プロンプト部分の数値は301から決して増えない。

		301> %!p(main.Foo={301 bob}) -> 0xc820010460
		301> %!p(main.Foo={301 bob}) -> 0xc8200104a0

		p3. ここはp1と同じ動きを示す。

		301> 0xc820010420 -> 0xc82002e038
		401> 0xc820010420 -> 0xc82002e040
	*/
}
