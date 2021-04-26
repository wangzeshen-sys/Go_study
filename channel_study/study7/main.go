package main

import "fmt"

func main() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default:
		}
	}
	tryReceive := func() string {
		select {
			case v := <- c: return v
			default: return "-"
		}
	}
	trySend("hello")
	trySend("hi!")
	trySend("bye")

	fmt.Println(tryReceive())
	fmt.Println(tryReceive())
	fmt.Println(tryReceive())
}