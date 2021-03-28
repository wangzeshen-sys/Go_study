package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker2() {
	defer wg.Done()
	for {
		fmt.Println("worker2...")
		time.Sleep(time.Second)
	}
}
func main() {
	wg.Add(1)
	go worker2()
	wg.Wait()
	fmt.Println("main over")
}