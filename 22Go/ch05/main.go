package main

import (
	"errors"
	"fmt"
)

// 函数的声明格式

/*	func funcName(params) result {
		body
	}
*/

// 在 Go 语言中，函数也是一种类型，它也可以被用来声明函数类型的变量、参数或者作为另一个函数的返回值类型。

func sum(a int, b int) (sum int, err error) {

	if a < 0 || b < 0 {
		// error 是 Go 语言内置的一个接口，用于表示程序的错误信息
		return 0, errors.New("a或者b不能是负数")
	}
	// 命名返回参数
	sum = a + b
	err = nil
	//return a + b, nil
	return
}

// 可变参数
func sum1(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

// 闭包
func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/* 方法 */

type Age uint

func (age Age) String() {
	fmt.Println("the age is", age)
}

func (age *Age) Modify() {
	*age = Age(30)
}

func main() {

	// result,_ := sum(1, 2)
	result, err := sum(11, -2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// 可变参数
	fmt.Println()
	fmt.Println("飞雪")
	fmt.Println("飞雪", "无情")

	fmt.Println(sum1(1, 2))
	fmt.Println(sum1(1, 2, 3))
	fmt.Println(sum1(1, 2, 3, 4))

	// 函数名称首字母小写代表私有函数，只有在同一个包中才可以被调用；
	// 函数名称首字母大写代表公有函数，不同的包也可以调用；
	// 任何一个函数都会从属于一个包。

	// 匿名函数
	sum2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum2(1, 2))

	// 闭包
	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	age := Age(25)
	age.String()
	age.Modify()
	age.String()
	(&age).String()

	// 作业

	age2 := Age(35)
	sm := age2.String
	sm()
	//sm(age2)
}
