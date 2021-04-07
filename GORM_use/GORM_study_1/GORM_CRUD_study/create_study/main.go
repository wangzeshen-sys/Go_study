package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}
var (
	Db  *gorm.DB
	err error
)

func main() {
	Db, err = gorm.Open("mysql", "root:asdasd123@tcp(127.0.0.1:3306)/db1")
	if err != nil {
		fmt.Println("open mysql failed, err:", err)
		return
	}
	defer Db.Close()
	
	// 自动迁移（创建表）
	Db.AutoMigrate(&UserInfo{})

	// 插入数据
	u1 := UserInfo{1, "小明", "男", "篮球"}
	u2 := UserInfo{2, "小刚", "女", "排球"}
	u3 := UserInfo{3, "小丽", "女", "羽毛球"}
	u4 := UserInfo{4, "小红", "女", "跑步"}
	u5 := UserInfo{5, "小智", "男", "跳绳"}
	Db.Create(&u1)
	Db.Create(&u2)
	Db.Create(&u3)
	Db.Create(&u4)
	Db.Create(&u5)

}
