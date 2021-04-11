package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn failed, err:", err)
		return
	}

	defer conn.Close()

	_, err = conn.Do("expire", "a", 5)
	if err != nil {
		fmt.Println("expire failed, err:", err)
		return
	}

	
}