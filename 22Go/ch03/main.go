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

	// switch 相关
	switch i := 6; {
	case i > 10:
		fmt.Println("i>10")
	case i > 5 && i <= 10:
		fmt.Println("5<i<=10")
	default:
		fmt.Println("i<=5")
	}

	switch j := 1; j {
	case 1:
		// 加入fallthrough 后，输出1，顺延
		fallthrough
	case 2:
		fmt.Println("1")
	default:
		fmt.Println("没有匹配")
	}

	switch 2 > 1 {
	case true:
		fmt.Println("2>1")
	case false:
		fmt.Println("2<=1")
	}

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("the sum is", sum)

	// for 模拟 while循环
	sum1 := 0
	i = 1
	for i <= 100 {
		sum1 += i
		i++
	}
	fmt.Println("the sum is", sum1)

	//continue 可以跳出本次循环，继续执行下一个循环。
	//break 可以跳出整个 for 循环，哪怕 for 循环没有执行完，也会强制终止。
	sum2 := 0
	k := 1
	for k <= 120 {
		sum2 += k
		k++
		if k > 100 {
			break
		}
	}
	fmt.Println("the sum is", sum2)

	sum3 := 0
	l := 1
	for l <= 110 {
		sum3 += l
		l++
		if l > 100 {
			continue
		}
	}
	fmt.Println("the sum is", sum3)
}
