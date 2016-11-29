package main

import (
	"fmt"
	"strings"
)

func foo(texts...string) string {
	return strings.Join(texts, ", ")
}

type Foo func(xs...string) string

func FooFactory(fn Foo) Foo {
	return func(xs...string) string {
		fmt.Println("*** START ***")
		defer fmt.Println("*** END ***")
		return fn(xs...)
	}
}

func main(){
	fn := FooFactory(foo)
	fmt.Println(fn("1", "2", "3"))
}
