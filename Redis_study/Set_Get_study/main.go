package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis conn failed, err:", err)
		return
	}

	defer conn.Close()
	fmt.Println("redis conn success")
	// Set
	_, err =conn.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println("set failed, err:", err)
		return
	}
	fmt.Println("set success")
	// get
	r, err := redis.Int(conn.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed, err:", err)
		return
	}
	fmt.Println("abc:", r)
}