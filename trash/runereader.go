package main

import (
	"bufio"
	"bytes"
	"fmt"
	// "unicode/utf8"
)

func main() {
	bs := []byte{
		240, 169, 184, 189,
		227, 129, 174,
		// 227, 129, 178,
		227, 178,
		227, 130, 137,
		227, 129, 141,
	}

	r := bytes.NewReader(bs)
	br := bufio.NewReader(r)

	// length := utf8.UTFMax // 4
	// xs := make([]byte, length)
	// fmt.Println(xs)

	for {
		r, size, err := br.ReadRune()
		
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("%c\t%v\n", r, size)
	}
}
