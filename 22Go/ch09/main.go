package main

import (
	"fmt"
	"sync"
)

//共享的资源

var (
	sum   int
	mutex sync.RWMutex
)

func main() {
	run()
}

func run() {
	var wg sync.WaitGroup
	wg.Add(110)
	//开启100个协程让sum+10
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			add(10)
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("和为:", readSum())
		}()
	}
	wg.Wait()
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
