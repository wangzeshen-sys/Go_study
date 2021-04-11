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

	// _, err = conn.Do("MSet", "a", 1, "b", 2, "c", 3)
	// if err != nil {
	// 	fmt.Println("MSet failed, err:", err)
	// 	return
	// }

	r, err := redis.Ints(conn.Do("MGet", "a", "b", "c"))
	if err!= nil {
		fmt.Println("get failed, err:", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}
}