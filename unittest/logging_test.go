package unittest

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// ログを出力するプログラムのテスト: logモジュールで出力先を潰してしまう(!)
func TestDoSomethingWithLog(t *testing.T) {
	setUp := func() {
		log.SetOutput(ioutil.Discard) // /dev/nullに出力
	}
	tearDown := func() {
		log.SetOutput(os.Stdout) // 出力先を戻す
	}

	func() {
		setUp()
		defer tearDown()

		var ok bool

		ok = DoSomethingWithLog(3)
		if !ok {
			t.Errorf("結果が期待通りではありませんでした: %v != true", ok)
		}

		ok = DoSomethingWithLog(2)
		if ok {
			t.Errorf("結果が期待通りではありませんでした: %v != false", ok)
		}
	}()

	// DoSomethingWithLog(2) // 当然、無名関数の外で呼び出したら標準出力にログが出力される
}
