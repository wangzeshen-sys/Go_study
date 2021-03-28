package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool
/*

*/
func worker3() {
	defer wg.Done()
	for {
		fmt.Println("worker3...")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go worker3()
	time.Sleep(time.Second*3)
	exit = true
	wg.Wait()
	fmt.Println("main over")
}