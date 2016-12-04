package main

import (
	"fmt"
	"time"
	"reflect"
)

func main() {
	fmt.Println(time.Now())                      // 2016-11-17 11:46:44.630877067 +0900 JST
	fmt.Println(time.Now().Format(time.RFC3339)) // 2016-11-17T11:46:44+09:00

	// わかる。
	fmt.Println(time.Now().Format("2006/01/02")) // 2016/11/17

	// なぜなのか？
	fmt.Println(time.Now().Format("2016/12/31")) // 17116/1117/1111
	fmt.Println(time.Now().Format("2005/01/02")) // 17044/11/17

	// !! 答え合わせ
	// https://golang.org/src/time/format.go?s=14349:14407#L76

	// 2:0:05 [stdDay "2", nil, stdZeroSecond "05"]
	// 2016-11-17 11:46:44.630877067 +0900 JST
	//         ^^       ^^
	//         "17" + "044" => "17044"

	// 2:0:16 [stdDay "2", nil, stdNumMonth "1", nil]
	// 2016-11-17 11:46:44.630877067 +0900 JST
	//      ^^ ^^
	//         "17" + "11" + "6" => "17116"

	// 1:2 [stdNumMonth "1", stdDay "2"] --> "11" + "17" => "1117"
	// 3:1 [stdHour12 "3", stdNumMonth "1"] --> "11" + "11" => "1111"
}
