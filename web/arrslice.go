package main

import "fmt"

func foo(nums [3]int) int {
	acc := 0
	for i:=0; i<3; i++ {
		acc += nums[i]
	}
	return acc
}

func main() {
	arr := [3]int {1, 2, 3}
	sl := []int {4, 5, 6}

	// 配列もスライスと同様にlen, capに対応してしまうので区別がつかない。
	// => 最初、配列arrはlen, capに応じられないというコンパイルエラーになることを期待してこれを書いたが、
	//    なぜか期待通りにならなかった。
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Println(len(sl))

	// 配列もスライスもrangeに渡せてしまうので両者の区別には使えなかった。
	for _, x := range arr{
		fmt.Println(x)
	}

	for _, x := range sl{
		fmt.Println(x)
	}

	a1, a2 := [...]int{1, 2, 3}, [...]int{7, 8, 9}  // こっちは配列
	a3 := a2[:]  // スライス

	a1[1] += 10
	a2[1] += 10
	a3[1] += 10  // a3はa2のスライスなので、基底配列を共有している

	fmt.Println(a1)
	fmt.Println(a2)  // このため、a2の第２要素はa3の操作によって余分に加算される。
	fmt.Println(a3)

	// a3はスライスなので配列を引数に宣言する関数fooに渡せない
	// cannot use a3 (type []int) as type [3]int in argument to foo
	// fmt.Println(foo(a3))

	// しかし、a2はintの配列なので、fooに渡せる。
	// => よって、a2とa3が異なるオブジェクトであることがわかる。
	fmt.Println(foo(a2))
}
