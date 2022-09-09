package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	stopCh := make(chan bool)
	go func() {
		defer wg.Done()
		watchDog(stopCh, "[监控狗1]")
	}()
	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stopCh <- true              //发停止指令
	wg.Wait()
}

func watchDog(stopCh chan bool, name string) {
	// 开启for select循环，一直后台监控
	for {
		select {
		case <-stopCh:
			fmt.Println(name, "停止指令已经收到，马上停止……")
		default:
			fmt.Println(name, "正在监控……")
		}
		time.Sleep(1 * time.Second)
	}
}
