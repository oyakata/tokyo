package main

import (
	"fmt"
	"strings"
	"strconv"
)

func Recorder(fn func(a ...string)) func(a...string){
	return func(a...string){
		fmt.Println("*** start ***")
		fn(a...)
		fmt.Println("*** end ***")
	}
}

type DoSomething (func(numbers...int) string)

func IntJoin(numbers...int) string {
	// 配列の宣言時に可変値を長さに指定するとエラーになる
	// non-constant array bound len(numbers)
	// var texts [len(numbers)]string

	// makeでスライスを作ると大丈夫
	texts := make([]string, len(numbers))

	for i, x := range numbers {
		texts[i] = strconv.Itoa(x)
	}
	return strings.Join(texts, ", ")
}

func RecordSomething(fn DoSomething) DoSomething {
	return func(numbers...int) string {
		fmt.Println("===== START =====")
		defer fmt.Println("===== END =====")
		return fn(numbers...)
	}
}

func main(){
	// 最初にやった方法。同じシグネチャの記述が何度も続いてくどい。
	fn := Recorder(func(a...string){
		fmt.Println(strings.Join(a, ", "))
	})
	fn([]string{"a", "b"}...)

	// 次に試した方法
	// funcにもtypeを付与できることが大発見
	fmt.Println(IntJoin(1, 2, 5, 7, 0))
}
