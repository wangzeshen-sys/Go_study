package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

var (
	Db  *gorm.DB
	err error
)

func main() {
	Db, err = gorm.Open("mysql", "root:asdasd123@tcp(127.0.0.1:3306)/db1?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	var uu1 UserInfo
	Db.First(&uu1)
	// uu1.Name = "小兰"
	// uu1.Gender = "女"
	uu1.Name = "小明"
	uu1.Gender = "男"
	Db.Debug().Save(&uu1) // 默认会修改所有字段
	// UPDATE `user_infos` SET `name` = '小明', `gender` = '男', `hobby` = '篮球'  WHERE `user_infos`.`id` = 1 
	fmt.Println(uu1)

	Db.Debug().Model(&uu1).Update("name", "小明")
	// UPDATE `user_infos` SET `name` = '小明'  WHERE `user_infos`.`id` = 1 
	fmt.Println(uu1)

}
