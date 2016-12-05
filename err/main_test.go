package err

import "testing"

func TestPositiveInt(t *testing.T) {
	// $ go test ./err // なぜ ./ をつけないとダメ？
	var err error
	var param int

	param = -1
	err = PositiveInt(param)
	if err == nil {
		t.Error("負の数にもかかわらずエラーになりませんでした")
	}

	param = 0
	err = PositiveInt(param)
	if err != nil {
		t.Errorf("%vにもかかわらずエラーになりました", param)
	}

	param = 1
	err = PositiveInt(param)
	if err != nil {
		t.Errorf("%vにもかかわらずエラーになりました", param)
	}
}
