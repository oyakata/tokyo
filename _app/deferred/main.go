package main

import "fmt"

func Sum1(xs ...int) int {
	acc := 0
	for i := 0; i < len(xs); i++ {
		acc += xs[i]
	}
	return acc
}

// func Sum2(xs ...int) (result int) {
// 	acc := 0
// 	for i := 0; i < len(xs); i++ {
// 		acc += xs[i]
// 	}
// 	return acc
// }

func Sum2(xs ...int) (result int) {
	for i := 0; i < len(xs); i++ {
		result += xs[i]
	}
	return result
}

func Sum3(xs ...int) (result int) {
	for i := 0; i < len(xs); i++ {
		result += xs[i]
	}
	return // これはok. 暗黙的に戻り値に宣言した変数resultが使われる
	// return 0  // これもok. 戻り値として宣言した型が一致するため
	// return "NG"  // 型がintではないのでこれはNG.
	// 当然、returnを書かないのもNG. (戻り値voidとみなされるので)
}

func Fn() (x int, y string, z [2]int) {
	return
}

func Sum4(xs ...int) (result int) {
	defer fmt.Printf("Sum4> result = %s\n", result)
	defer func() { fmt.Printf("Sum4> result = %s\n", result) }()
	for _, x := range xs[:len(xs)-1] {
		result += x
	}
	return result + xs[len(xs)-1]
}

func Sum5(xs ...int) (result int) {
	defer func() {
		result -= 500
		fmt.Printf("Sum5> result = %s\n", result)
	}()
	for _, x := range xs[:len(xs)-1] {
		result += x
	}
	return result + xs[len(xs)-1]
}

func main() {
	fmt.Println(Sum1(1, 2, 3, 4))
	fmt.Println(Sum2(1, 2, 3, 4))
	fmt.Println(Sum3(1, 2, 3, 4))
	fmt.Println(Fn()) // 0  [0 0]
	fmt.Println(Sum4(1, 2, 3, 4))
	fmt.Println(Sum5(1, 2, 3, 4))
}
