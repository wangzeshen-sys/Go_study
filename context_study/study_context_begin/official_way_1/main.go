package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


var wg sync.WaitGroup

func worker5(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("worker5...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker5(ctx)
	time.Sleep(time.Second*3)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("main over")
}