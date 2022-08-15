package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	/*
		Array（数组）
	*/

	// 数组存放的是固定长度、相同类型的数据，存放的元素是连续的
	// array := [...]string{"a", "b", "c","d","e"}
	// array:=[5]string{1:"b",3:"d"}
	// array:=[...]string{1:"b",3:"d"}	错误 invalid argument: array index 5 out of bounds [0:5]
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])
	for i, v := range array {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}
	for _, v := range array {
		fmt.Printf("对应值:%s\n", v)
	}

	/*
		Slice 切片
	*/
	// 从数组 array 的索引 2 开始到索引 5 结束
	// [2,5) 左闭右开: 2, 3, 4
	slice := array[2:5]
	fmt.Println("slice: ", slice)
	// 切片元素循环
	for i, v := range slice {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}

	// 使用 make 函数。slice1:=make([]string,4,8) 长度为4，容量为8
	// 切片的容量不能比切片的长度小。
	slice1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(slice1), cap(slice1))

	slice2 := append(slice1, "f")
	fmt.Println(slice2)

	slice3 := append(slice1, "f", "g")
	fmt.Println(slice3)

	slice4 := append(slice1, slice2...)
	fmt.Println(slice4)

	/*
		Map（映射）
	*/
	// 在 Go 语言中，map 是一个无序的 K-V 键值对集合，结构为 map[K]V
	// map 中所有的 Key 必须具有相同的类型，Value 也同样
	// Key 的类型必须支持 == 比较运算符，这样才可以判断它是否存在，并保证 Key 的唯一。

	//测试 for range
	nameAgeMap := make(map[string]int)
	// nameAgeMap:=map[string]int{"飞雪无情":20}
	nameAgeMap["飞雪无情0"] = 20
	nameAgeMap["飞雪无情1"] = 21
	nameAgeMap["飞雪无情2"] = 22

	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)
	}
	fmt.Println("nameAgeMap的长度len为：", len(nameAgeMap))

	/* map 的 [] 操作符可以返回两个值：
	1. 第一个值是对应的 Value；
	2. 第二个值标记该 Key 是否存在，如果存在，它的值为 true。
	*/
	nameAgeMap1 := map[string]int{"飞雪无情": 20}
	age, ok := nameAgeMap1["飞雪无情"]
	if ok {
		fmt.Println(age)
	}
	delete(nameAgeMap1, "飞雪无情")
	fmt.Println("nameAgeMap1: ", nameAgeMap1)

	/*
		String 和 []byte
	*/
	s := "Hello飞雪无情"
	// 字符串是字节序列，每一个索引对应的是一个字节，而在 UTF8 编码下，一个汉字对应三个字节，所以字符串 s 的长度其实是 17。
	fmt.Println(len(s))

	// 把一个汉字当做按一个长度计算
	fmt.Println(utf8.RuneCountInString(s))

	// 字符串s 转换为 字节切片 []byte
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(bs[0], bs[1])
	fmt.Println(s[0], s[1], s[15])

	// for range 循环在处理字符串的时候，会自动地隐式解码 unicode 字符串。
	for i, r := range s {
		fmt.Println(i, r)
	}
}
