package main

import "fmt"

func main() {
	//测试 for range
	nameAgeMap := make(map[string]int)
	nameAgeMap["飞雪无情"] = 20

	nameAgeMap["飞雪无情1"] = 21

	nameAgeMap["飞雪无情2"] = 22

	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)

	}

}
