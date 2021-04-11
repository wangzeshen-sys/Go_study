package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis conn failed, err:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("HSet", "books", "qws", 100)
	if err != nil {
		fmt.Println("HSet failed, err:", err)
		return
	}

	r, err := redis.Int(conn.Do("Hget", "books", "qws"))
	if err != nil {
		fmt.Println("Hget failed, err:", err)
		return
	}
	fmt.Println(r)
}