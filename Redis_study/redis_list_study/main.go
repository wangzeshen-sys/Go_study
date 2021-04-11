package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("reddis conn failed, err:", err)
		return
	}

	defer conn.Close()

	_, err = conn.Do("lpush", "book_list", "qwe", "asd", "zxc")
	if err != nil {
		fmt.Println("lpush failed, err:", err)
		return
	}

	r, err := redis.String(conn.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get qwe failed, err:", err)
		return
	}
	r1, err := redis.String(conn.Do("rpop", "book_list"))
	if err != nil {
		fmt.Println("get qwe failed, err:", err)
		return
	}
	fmt.Println("r:", r)
	fmt.Println("r1", r1)
}