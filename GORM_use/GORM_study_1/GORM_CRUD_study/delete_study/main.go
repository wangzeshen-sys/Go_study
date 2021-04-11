package main

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
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

	var uu2 UserInfo
	Db.Where("name=?", "小明").Delete(&uu2)

	Db.Unscoped().Where("name=?", "小刚").Delete(&uu2)
}