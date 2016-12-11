package main

import (
	"fmt"
	"time"
)

func main() {
	// 正式なN秒スリープの仕方
	time.Sleep(1 * time.Second)
	// time.Sleep(0.5 * time.Second)  // これはエラー(floatと掛け算できない)
	time.Sleep(500 * time.Millisecond) // 0.5秒待つときはMillisecondと乗算する

	// こちらは冗長
	time.Sleep(time.Duration(1 * time.Second))

	var s time.Duration // ここをint64で宣言すると掛け算の式がコンパイルエラーになる
	// invalid operation: 100 * s * time.Millisecond (mismatched types int64 and time.Duration)
	s = 5
	time.Sleep(100 * s * time.Millisecond)
	fmt.Println("Hello, world.")

	var t int
	t = 5
	time.Sleep(time.Duration(t*100) * time.Millisecond)
	fmt.Println("Hello, world.")
}
