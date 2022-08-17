package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	//i, err := strconv.Atoi("a")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(i)
	//}
	sum, err := add(-1, 2)
	if cm, ok := err.(*commonError); ok {
		fmt.Println("错误代码为：", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}

	// 错误嵌套 Error Wrapping

	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap 了一个错误: %w", e)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))

	fmt.Println(errors.Is(w, e))

	var cm *commonError
	if errors.As(err, &cm) {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}

	//defer 关键字 + 匿名函数 + recover 函数从 panic 异常中恢复的方式。
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	connectMySQL("", "root", "123456")
}

func connectMySQL(ip, username, password string) {
	if ip == "" {
		panic("ip不能为空")
	}
	//省略其他代码
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

func ReadFile(filename string) ([]byte, error) {

	f, err := os.Open(filename)

	if err != nil {

		return nil, err

	}

	defer f.Close()

	//省略无关代码

	//return readAll(f, n)
	return nil, nil
}
