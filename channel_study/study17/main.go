package main

import (
	"log"
	"math/rand"
	"time"
)

// 通过发送操作来获取所有权的实现相对简单一些

type Customer struct{ id int }
type Bar chan Customer

func (bar Bar) serveCustomer(c Customer) {
	log.Print("++顾客#", c.id, "开始饮酒")
	time.Sleep(time.Second * time.Duration(3+rand.Intn(16)))
	log.Print("--顾客#", c.id, "离开酒吧")
	<-bar
}

func main() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10)
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second * 2)
		customer := Customer{customerId}
		bar24x7 <- customer
		go bar24x7.serveCustomer(customer)
	}
}
