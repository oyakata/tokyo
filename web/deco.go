package main

import (
	"fmt"
	"strings"
)


func Recorder(fn func(a ...string)) func(a...string){
	return func(a...string){
		fmt.Println("*** start ***")
		fn(a...)
		fmt.Println("*** end ***")
	}
}


func main(){
	fn := Recorder(func(a...string){
		fmt.Println(strings.Join(a, ", "))
	})
	fn([]string{"a", "b"}...)
}
