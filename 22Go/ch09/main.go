package main

import (
	"fmt"
	"sync"
	"time"
)

//共享的资源

var (
	sum   int
	mutex sync.RWMutex
)

func main() {
	//开启100个协程让sum+10
	for i := 0; i < 100; i++ {
		go add(10)
	}
	for i := 0; i < 10; i++ {
		go fmt.Println("和为:", readSum())
	}
	// 防止提前退出
	time.Sleep(2 * time.Second)
	//fmt.Println("和为:", sum)
}

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i // 临界区
}

//增加了一个读取sum的函数，便于演示并发
func readSum() int {
	mutex.Lock()
	defer mutex.Unlock()
	b := sum
	return b
}
