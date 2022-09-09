package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		watchDog("[监控狗1]")
	}()
	wg.Wait()
}

func watchDog(name string) {
	for {
		select {
		default:
			fmt.Println(name, "正在监控……")
		}
		time.Sleep(1 * time.Second)
	}
}
