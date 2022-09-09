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
	//run()
	doOnce()
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	//用于等待协程执行完毕
	done := make(chan bool)
	//启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			//把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
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
