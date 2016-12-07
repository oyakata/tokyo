package unittest

import (
	"log"
)

// ログを出力するプログラム
func DoSomethingWithLog(x int) (ok bool) {
	if x%2 == 0 {
		log.Printf("偶数は受け付けません: %vが入力されました", x)
		return false
	}
	return true
}
