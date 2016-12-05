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

func TestSum(t *testing.T) {
	var xs []int
	var result int
	var err error

	xs = []int{1, 2, 3, 4, 5}
	result, err = Sum(xs...)
	if err != nil {
		t.Error(err)
	}
	if result != 15 {
		t.Error("Sum(%v) != 15, but %v", xs, result)
	}

	xs = []int{10, 11, -5, 8, -2, 16, 0, -9}
	result, err = Sum(xs...)
	if err == nil {
		t.Errorf("Sum(%v) should be failed, but succeeded.", xs)
	}
	if result != 0 {
		t.Error("Sum(%v) != 0, but %v", xs, result)
	}

}
