package main

import (
	"fmt"
)

func main() {
	exp := "𩸽のひらきを居酒屋で注文して、1時間経つがまだ来ない。𠈻な客が店員を引き止めてなじるからだ。" // 46文字

	for _, x := range []byte(exp) {
		fmt.Printf("%c", x)
	}
}
