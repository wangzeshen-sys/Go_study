package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle: 16,
		// MaxActive: 100000,
		MaxActive: 0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get() // 从连接池取出一个连接
	defer conn.Close() // 函数运行结束，把连接放回连接池

	_, err := conn.Do("set", "api", 200)
	if err != nil {
		fmt.Println("set failed, err:", err)
		return
	}

	r, err := redis.Int(conn.Do("Get", "api"))
	if err != nil {
		fmt.Println("get api failed, err:", err)
		return
	}

	fmt.Println(r)
}
