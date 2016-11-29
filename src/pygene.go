package main

import (
	"fmt"
	"strconv"
	"strings"
)

type DoSomething func(numbers ...int) string

func SayStart(fn DoSomething) DoSomething {
	return func(numbers ...int) string {
		fmt.Println("=== START ===")
		return fn(numbers...)
	}
}

// intの可変長引数を受け取ってカンマ区切りの文字列に変換して返す
func IntJoin(numbers ...int) string {
	texts := make([]string, len(numbers))

	for i := 0; i < len(numbers); i++ {
		texts[i] = strconv.Itoa(numbers[i])
	}

	return strings.Join(texts, ", ")
}

func main() {
	fn := SayStart(IntJoin)
	fmt.Println(fn(1, 2, 3, 9, 8, 7))
}
