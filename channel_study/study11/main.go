package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

// 从一个通道接收一个值来实现单对单通知
func main() {
	values := make([]byte, 32* 1024 * 1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	done := make(chan struct{})

	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{}
	}()
	
	<- done
	fmt.Println(values[0], values[len(values)-1])
}