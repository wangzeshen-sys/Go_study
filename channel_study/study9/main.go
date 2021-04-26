package main

import (
	"fmt"
	"math/rand"
	"time"
)
// 将单项发送通道类型用作函数实参
func longTimeRequest(r chan<- int32) {
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ra, rb := make(chan int32), make(chan int32)
	go longTimeRequest(ra)
	go longTimeRequest(rb)

	fmt.Println(sumSquares(<-ra, <-rb))
}