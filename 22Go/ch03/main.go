package main

import "fmt"

func main() {

	i := 6
	if i > 10 {
		fmt.Println("i>10")
	} else if i > 5 && i <= 10 {
		fmt.Println("5<i<=10")
	} else {
		fmt.Println("i<=5")
	}

	//sum := 0
	//
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	//fmt.Println("the sum is", sum)

	sum := 0
	i = 1
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println("the sum is", sum)

	switch j := 1; j {
	case 1:
		fallthrough
	case 2:
		fmt.Println("1")
	default:
		fmt.Println("没有匹配")

	}

}
