package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("𩸽のひらきを居酒屋で注文して、1時間経つがまだ来ない。𠈻な客が店員を引き止めてなじるからだ。") // 46文字

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c\t%v\n", r, size)
		b = b[size:]
	}
}

