package main

import (
	"fmt"
	"sync"
	"time"
)

//共享的资源

var (
	sum   int
	mutex sync.Mutex
)

func main() {
	//开启100个协程让sum+10
	for i := 0; i < 100; i++ {
		go add(10)
	}

	// 防止提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("和为:", sum)
}

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i // 临界区
}
