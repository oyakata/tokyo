package main

import "fmt"


func main(){
	var foo = func(x int) int {
		return x + 5
	}

	var deco = func(fn func(int) int) func(int) int {
		return func(x int) int {
			return fn(x)
		}
	}

	var fn = deco(foo)
	fmt.Println(fn(100))
}
