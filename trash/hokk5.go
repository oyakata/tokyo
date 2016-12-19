package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte{228, 184, 150} // 世

	r, size := utf8.DecodeRune(b[:2])

	// 読みだしたRuneが不正だった場合、utf8.RuneErrorが返却される。

	// https://golang.org/pkg/unicode/utf8/#pkg-constants
	// RuneError = '\uFFFD'

	fmt.Printf("%c\t%v\t%v\n", r, size, r == utf8.RuneError) // �	1	true
}
