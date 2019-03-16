package error

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should be less than 2")
var LargerThanHundredError = errors.New("n should be larger than 100")

// The Fibonacci Func
func GerFibonacciSeries(n int) ([]int, error) {
	// if n < 0 || n > 100 {
	// 	return nil, errors.New("n should be in [2, 100]")
	// }
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n < 0 || n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestFibonacci(t *testing.T) {
	if v, error := GerFibonacciSeries(1); error != nil {
		if error == LessThanTwoError {
			fmt.Println("It is less")
		}
		if error == LargerThanHundredError {
			fmt.Println("It is larger")
		}
		t.Error(error)
	} else {
		t.Log(v)
	}
}
