package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//测试 for range
	nameAgeMap := make(map[string]int)
	nameAgeMap["飞雪无情0"] = 20
	nameAgeMap["飞雪无情1"] = 21
	nameAgeMap["飞雪无情2"] = 22

	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)
	}
	fmt.Println(len(nameAgeMap))

	s := "Hello飞雪无情"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(bs[0], bs[1])
	fmt.Println(utf8.RuneCountInString(s))

	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])
	for i, v := range array {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}
	for _, v := range array {
		fmt.Printf("对应值:%s\n", v)
	}

	// 从数组 array 的索引 2 开始到索引 5 结束
	// [2,5)
	slice := array[2:5]
	fmt.Println(slice)

	slice1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(slice1), cap(slice1))

	slice2 := append(slice1, "f")
	fmt.Println(slice2)

	slice3 := append(slice1, "f", "g")
	fmt.Println(slice3)

	slice4 := append(slice1, slice2...)
	fmt.Println(slice4)
}
