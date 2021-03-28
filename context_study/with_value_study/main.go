package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

func workerer(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("workerer, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
			case <-ctx.Done():
				break LOOP
			default:
		}
	}
	fmt.Println("worlerer done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在斯通的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg.Add(1)
	go workerer(ctx)
	time.Sleep(time.Second*5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
