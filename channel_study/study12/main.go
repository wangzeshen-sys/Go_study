package main

import (
	"log"
	"time"
)

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<- ready  // 阻塞在此，等待通知
	log.Print("worker#", id, "开始工作") // 模拟一个工作负载
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("worker#", id, "工作完成")
	done <- T{} // 通知主协程
}

// 多对单和单对多的通知

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// 模拟一个初始化过程
	time.Sleep(time.Second* 3 / 2)
	// 单对多通知
	ready <- T{}; ready <- T{}; ready <- T{}
	// 等待被多对单通知
	<-done; <-done; <-done
}