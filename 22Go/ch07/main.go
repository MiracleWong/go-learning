package main

import (
	"fmt"
)

func main() {

	//i, err := strconv.Atoi("a")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(i)
	//}
	sum, err := add(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &commonError{
			errorCode: 1,
			errorMsg:  "a或者b不能为负数"}
	} else {
		return a + b, nil
	}
}

// 自定义的error

type commonError struct {
	errorCode int    //错误码
	errorMsg  string //错误信息

}

func (ce *commonError) Error() string {
	return ce.errorMsg
}
