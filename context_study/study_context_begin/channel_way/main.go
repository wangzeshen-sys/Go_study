package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 管道方式存在的问题
// 1.使用全局变量在跨包调用时不容易实现规范和统一

func worker4(exitChan chan struct{}) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("worker4...")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}
}
func main() {
	var exitChan = make(chan struct{})
	wg.Add(1)
	go worker4(exitChan)
	time.Sleep(time.Second*3)
	exitChan <- struct{}{}
	close(exitChan)
	wg.Wait()
	fmt.Println("main over")
}