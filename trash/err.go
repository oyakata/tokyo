package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error

	err = errors.New("処理に失敗しました") // エラーの原因などが何も無いので、あまり良くないメッセージ
	fmt.Println(err)

	err = errors.New(fmt.Sprintf("処理中に接続が切れたためリトライを開始します(%v回め)", 3))
	fmt.Println(err)

	// 前のエラー内容を引き継いで新規エラーメッセージを追記する場合
	err = fmt.Errorf("%v: リトライ回数が一定に達したため、処理を打ち切ります(%v回試しました)", err, 7)
	fmt.Println(err)
}
