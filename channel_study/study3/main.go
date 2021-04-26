package main

import (
	"fmt"
	"time"
)

func main() {
	var ball = make(chan string)

	kicBall := func(playerName string) {
		for {
			fmt.Print(<-ball, "传球", "\n")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kicBall("张三")
	go kicBall("李四")
	go kicBall("王二麻子")
	go kicBall("刘大")
	ball <- "裁判"
	
	c := make(chan int, 1)
	c <- 1
	<- c

}