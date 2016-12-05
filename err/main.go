package err

import "fmt"

func PositiveInt(x int) error {
	if x < 0 {
		return fmt.Errorf("PositiveInt() argument must be a positive number, not negative")
	}
	return nil
}

func Even(x int) error {
	if x%2 == 0 {
		return fmt.Errorf("Even() argument must not be a odd number.")
	}
	return nil
}
