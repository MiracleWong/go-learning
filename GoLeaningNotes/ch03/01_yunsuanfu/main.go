package main

import "fmt"

const (
	read byte = 1 << iota
	write
	exec
	freeze
)

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

	// read  = 0001
	// write = 0010
	// exec  = 0100
	// freeze = 1000

	g := read | write | freeze
	// 0001 | 0010 | 1000 = 1011
	h := read | freeze | exec
	// 0001 | 1000 | 0100 = 1101
	i := g &^ h
	// 1011 &^ 1101 = 0010
	fmt.Printf("%04b &^ %04b = %04b\n", g, h, i)

	// 自增自减不能作为运算符，只能作为独立语句。
}
