package main

import "fmt"

// select-case 控制代码
func main() {
	// var c chan struct{}
	
	// select {
	// 	case <-c:
	// 		fmt.Println("<-c 执行了")
	// 	case c <- struct{}{}:
	// 		fmt.Println("c <- struct{}{} 执行了")
	// 	default:
	// 		fmt.Println("Go here")
	// }
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}:
	case <- c:
		fmt.Println("<- c 执行")
	}
}