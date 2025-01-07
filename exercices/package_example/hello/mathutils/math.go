package mathutils

import "errors"

func Add(x int, y int) int {
	return x + y
}

func Multiply(x int, y int) int {
	return x * y
}

func Devide(x int, y int) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot devide by zero")
	}
	return float64(x) / float64(y), nil
}
