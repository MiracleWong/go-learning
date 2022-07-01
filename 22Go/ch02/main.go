package main

import "fmt"

const (
	one = iota + 1
	two
	three
	four
)

func main() {
	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	fmt.Println("f32 is", f32, ",f64 is", f64)

	var bf bool = false
	var bt bool = true
	fmt.Println("bf is", bf, ",bt is", bt)

	var s1 string = "Hello"
	var s2 string = "世界"
	fmt.Println("s1 is", s1, ",s2 is", s2)

	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi, zf, zb, zs)

	i := 10
	//bf := false
	//s1 := "Hello"

	pi := &i
	fmt.Println(&i)
	fmt.Println(*pi)

	i = 20
	fmt.Println("i的新值是", i)

	fmt.Println(one, two, three, four)
}
