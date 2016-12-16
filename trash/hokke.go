package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	exp := "𩸽のひらきを居酒屋で注文して、1時間経つがまだ来ない。𠈻な客が店員を引き止めてなじるからだ。" // 46文字

	fmt.Println(exp)
	fmt.Println(
		len(exp),                    // 138
		utf8.RuneCountInString(exp), // 46
		len([]rune(exp)))            // 46

	for _, x := range exp {
		fmt.Printf("%c", x)
	}
	fmt.Println() // ここで改行を挟まないと何故か1回めのループの「𩸽」の文字が化けて表示される。

	var xs []byte

	xs = []byte(exp)
	for len(xs) > 0 {
		r, size := utf8.DecodeRune(xs)
		fmt.Printf("%c", r)
		xs = xs[size:]
	}
	fmt.Println()

	// DecodeLastRuneだと逆から反復
	xs = []byte(exp)
	for len(xs) > 0 {
		r, size := utf8.DecodeLastRune(xs)
		fmt.Printf("%c", r)
		xs = xs[:len(xs)-size]
	}
}
