package main

import (
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (b Bar) ServeCustomer(c int) {
	log.Print("顾客#", c, "进入酒吧")
	seat := <- b // 需要一个位置来喝酒
	log.Print("++ customer#", c, " drinks at seat#", seat)
	log.Print("++顾客", c, "在第", seat, "个座位开始饮酒")
	time.Sleep(time.Second * time.Duration(2 * rand.Intn(6)))
	log.Print("--顾客#", c, "离开了第", seat, "个座位")
	b <- seat // 释放座位， 离开酒吧
}

func main() {
	rand.Seed(time.Now().UnixNano())
	bar24x7 := make(Bar, 10) // 此酒吧有10个位置
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		go bar24x7.ServeCustomer(customerId)
	}
	// for {time.Sleep(time.Second)}

}