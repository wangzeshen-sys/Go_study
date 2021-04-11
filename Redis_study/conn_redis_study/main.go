package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "192.168.240.1:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("redis conn success")

}