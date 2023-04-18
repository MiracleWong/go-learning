package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func test(x int) func() { // 返回函数类型
	return func() { // 匿名函数
		println(x) // 闭包
	}
}

func main() {
	fmt.Println("Hello World")

	// var 定义变量，支持类型推断
	// 变量被初始化为0值
	var x int32
	var s = "Hello World!"

	println(x, s)

	x1 := 100
	fmt.Println(x1)

	// if
	x2 := 100
	if x2 > 0 {
		println("x2")
	} else if x2 < 0 {
		println("-x2")
	} else {
		println("0")
	}

	// switch
	x3 := 100
	switch {
	case x3 > 0:
		println("x3")
	case x3 < 0:
		println("-x")
	default:
		println("0")
	}

	// for
	for i := 0; i < 5; i++ {
		println(i)
	}
	for i := 4; i >= 0; i-- {
		println(i)
	}

	i1 := 0
	for i1 < 5 {
		println(i1)
		i1++
	}

	i2 := 4
	for {
		println(i2)
		i2--

		if i2 < 0 {
			break
		}
	}

	// for range
	i4 := []int{100, 101, 102}
	for i, n := range i4 {
		println(i, ":", n)
	}

	// 函数 div
	a, b := 10, 2
	c, err := div(a, b)
	fmt.Println(c, err)

	d := 100
	f := test(d)
	f() // 函数调用

	// defer 定义延迟调用，无论函数是否出错，他都确保结束前被调用。

}
