package main

import "fmt"

func main() {

	// 无显式类型声明的常量，自动转型
	const v = 20
	var a byte = 10
	b := v + a
	fmt.Printf("%T, %v\n", b, b)
	const c float32 = 1.2
	d := c + v
	fmt.Printf("%T, %v\n", d, d)

	e := 3
	x := 1 << e
	println(x)

}
