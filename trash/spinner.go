package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 配列やスライスはconstで定義できない。
// const numbers = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func spinner(message string, delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c%v", r, message)
			time.Sleep(delay)
		}
	}
}

func main() {
	var delay time.Duration
	var acc time.Duration

	go spinner("コンピュータ思考中...", 140*time.Millisecond)

	// 乱数のシードをナノ秒にしておかないと、デフォルトは秒なので短時間の処理では乱数が分布しない。

	// http://qiita.com/cubicdaiya/items/819886c57e9d17e4b019
	// ↑ここのコメント欄から

	// ちなみに↓ここでも同じことをして、Pythonのshuffleを模倣している。
	// http://marcelom.github.io/2013/06/07/goshuffle.html
	rand.Seed(time.Now().UnixNano())

	for _, x := range rand.Perm(9)[:5] {
		delay = time.Duration(100*x) * time.Millisecond
		acc += delay
		time.Sleep(delay)

		// fmt.Println(x, 100*x*time.Millisecond) // これはコンパイルエラー

		// ?? どういうこと?
		// invalid operation: 100 * x * time.Millisecond (mismatched types int and time.Duration)
		// time.Durationはint64の別名型なのでダメらしい
		// 掛け算とわり算でやり方が違う点がまたわかりにくい

		// https://golang.org/pkg/time/#Duration
		// fmt.Print(int64(second/time.Millisecond)) // prints 1000
		// fmt.Print(time.Duration(seconds)*time.Second) // prints 10s

		// [追記 11:45]

		// > 掛け算とわり算でやり方が違う

		// これはそういうことじゃなくて、割り算の例は、Duration -> int64に、
		// 掛け算の例はint64 -> Durationに変換している例だと思う。

		// time.SleepはDurationを引数に取るので注意
	}

	// ANSIエスケープコードを使ってコンソール制御
	// http://www.mm2d.net/main/prog/c/console-02.html

	// fmt.Printf("\033[H\033[2J") // これだと画面を消してしまう
	fmt.Printf("\033[2K\033[G") // 行を消し、その後カーソルを行頭に移動
	// [ok] だけ色を緑に変える
	fmt.Printf("\033[32m[ok]\033[39m %v かかりました\n", acc)
	// fmt.Printf("\r%v かかりました\n", acc)
}
