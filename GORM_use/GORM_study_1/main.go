package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	db, err := gorm.Open("mysql", "root:asdasd123@tcp(127.0.0.1:3306)/db1")
	if err != nil {
		fmt.Println("open mysql failed, err:", err)
		return
	}
	fmt.Println("连接成功")
	db.Close()
}