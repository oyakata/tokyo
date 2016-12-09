package gdbsample

import (
	"fmt"
	"strconv"
)

func Sum(xs ...int) (acc int) {
	for i := 0; i < len(xs); i++ {
		acc += xs[i]
		if xs[i]%7 == 0 {
			panic("cannot accept multiples of 7.")
		}
	}
	return
}

func Sample() {
	var xs []int
	base := "1234567890"

	xs = make([]int, len(base))
	for i, x := range base {
		xs[i], _ = strconv.Atoi(string(x))
	}
	fmt.Println(Sum(xs...))
}
