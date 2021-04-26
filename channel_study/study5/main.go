package main

import (
	"fmt"
	"time"
)

// for-range
func main() {
	ch := make(chan int, 5)
	for a := 1; a < 6; a++ {
		ch <- a
	}
	// close(ch)
	// for x := range ch {
	// 	time.Sleep(time.Second)
	// 	fmt.Println(x)
	// }
	for {
		v, ok := <-ch
		if !ok {
			break
		} else {
			time.Sleep(time.Second)
			fmt.Println(v)
		}
	}
}