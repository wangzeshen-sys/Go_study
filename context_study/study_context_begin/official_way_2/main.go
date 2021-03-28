package main

import (
    "context"
    "fmt"
    "sync"

    "time"
)

var wg sync.WaitGroup

func worker6(ctx context.Context) {
    go worker7(ctx)
LOOP:
    for {
        fmt.Println("worker6...")
        time.Sleep(time.Second)
        select {
        case <-ctx.Done(): // 等待上级通知
            break LOOP
        default:
        }
    }
    wg.Done()
}

func worker7(ctx context.Context) {
LOOP:
    for {
        fmt.Println("worker7...")
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
    go worker6(ctx)
    time.Sleep(time.Second * 3)
    cancel() // 通知子goroutine结束
    wg.Wait()
    fmt.Println("over")
}