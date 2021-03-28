package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// server 端 随机出现慢反应

func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second*10) // 耗时10秒的慢响应
		fmt.Println(w, "slow response")
		return
	}
	fmt.Fprint(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8689", nil)
	if err != nil {
		panic(err)
	}
}