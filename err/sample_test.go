package err

import "testing"

func TestFoo(t *testing.T){
	if 1 != 0 {
		t.Error("結果が不正です: 1 != 0") // ここで処理は止まらない
	}

	if 2 != 1 {
		t.Error("結果が不正です: 2 != 1") // こっちもテストしてくれる
	}
}
