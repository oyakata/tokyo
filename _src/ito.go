package main

import "fmt"

type Samurai int

// Goは変数名に日本語を使っても大丈夫
const (
	島田勘兵衛 Samurai = 1 + iota // constが登場したところから0からカウントアップ
	菊千代
	岡本勝四郎
	片山五郎兵衛
	七郎次
	林田平八
	久蔵
)

// 関数の名前にも日本語が使える
func 戦(s Samurai, enemies []int) {
	for _, e := range enemies {
		up, down := e-int(s), Samurai(e)+s
		fmt.Printf("%v, %v\n", up, down)
		fmt.Printf("%s, %s\n", up, down)
	}
}

func main() {
	fmt.Println(七郎次 + 久蔵) // => 12
	戦(片山五郎兵衛, []int{10, 9, 8, 5, 3, 11})
}
