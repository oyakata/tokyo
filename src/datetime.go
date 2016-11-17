package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())                      // 2016-11-17 11:46:44.630877067 +0900 JST
	fmt.Println(time.Now().Format(time.RFC3339)) // 2016-11-17T11:46:44+09:00

	// わかる。
	fmt.Println(time.Now().Format("2006/01/02")) // 2016/11/17

	// なぜなのか？
	fmt.Println(time.Now().Format("2016/12/31")) // 17116/1117/1111
	fmt.Println(time.Now().Format("2005/01/02")) // 17044/11/17
}
