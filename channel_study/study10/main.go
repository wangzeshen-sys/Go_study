package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3) + 1
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

func main() {
	rand.Seed(time.Now().UnixNano())
	startTime := time.Now()
	c := make(chan int32, 5)
	for i := 0;i < cap(c); i++ {
		go source(c)
	}
	rnd := <-c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}