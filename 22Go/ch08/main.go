package main

import (
	"fmt"
)

// 无缓冲的channel，同步channel
// 定义好 chan 后就可以使用了，一个 chan 的操作只有两种：发送和接收。
// 1. 接收：获取 chan 中的值，操作符为 <- chan。
// 2. 发送：向 chan 发送值，把值放在 chan 中，操作符为 chan <-。

func main() {
	ch := make(chan string)
	go func() {
		fmt.Println("飞雪无情")
		ch <- "goroutine 完成"
	}()

	fmt.Println("我是 main goroutine")

	v := <-ch
	fmt.Println("chan 为: ", v)
}
