package err

import (
	"errors"
	"fmt"
)

func PositiveInt(x int) error {
	if x < 0 {
		return fmt.Errorf("PositiveInt() argument must be a positive number, not '%v'", x)
	}
	return nil
}

func Even(x int) error {
	if x%2 == 0 {
		return fmt.Errorf("Even() argument must be a odd number, not '%v'.", x)
	}
	return nil
}

// 合計値を返す。但し、引数に負の数が含まれたら即座に処理を打ち切ってnilを返す。
func Sum(xs ...int) (result int, e error) {
	var err error
	for _, x := range xs {
		err = PositiveInt(x)
		if err != nil {
			if e == nil {
				e = errors.New("Sum")
			}
			e = fmt.Errorf("%v: %v", e, err)
			return 0, e
		}
		result += x
	}
	return
}

// 合計を求める。但し、引数に負の数が含まれても計算を続行する。
func SumEasily(xs ...int) (result int, e error) {
	var err error
	for _, x := range xs {
		err = PositiveInt(x)
		if err != nil {
			if e == nil {
				e = errors.New("SumEasily")
			}
			e = fmt.Errorf("%v: %v", e, err)
		}
		result += x
	}
	return
}
