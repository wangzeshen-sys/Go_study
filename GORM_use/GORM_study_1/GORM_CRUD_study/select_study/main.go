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
		fmt.Println("open mysql failed, err:", err)
		return
	}
	defer Db.Close()
	
	var u1 UserInfo
	// 👆 <=> u1 := new(UserInfo)
	Db.First(&u1)
	// 查询第一条记录  主键为int类型才能使用
	// select * from user_infos order by id  limit 1;
	fmt.Println("第一条记录")
	fmt.Printf("%v\n", u1)

	var u2 UserInfo
	Db.Last(&u2)
	// 查询最后一条记录
	// select * from user_infos order by id  desc limit 1;
	fmt.Println("最后一条记录")
	fmt.Printf("%v\n", u2)

	var u3 []UserInfo // 定义一个切片类型
	// 查询所有记录
	// select * from user_infos;
	Db.Debug().Find(&u3)
	fmt.Println("表中的所有记录")
	fmt.Println(u3)

	// where 查询

	// 获取第一个匹配的记录
	var u4 UserInfo
	Db.Where("name=?", "小丽").First(&u4)
	// select * from user_infos where name='小丽' limit 1;
	fmt.Println(u4)

	// 获取所有匹配记录
	Db.Where("name=?", "小刚").Find(&u3)
	// select * from user_infos where name='小刚';
	fmt.Println(u3)

	Db.Where("name<>?", "小红").Find(&u3)
	// select * from user_infos where name<>'小红'; 
	fmt.Println(u3)

	// IN
	fmt.Println("IN TEST")
	Db.Where("name IN (?)", []string{"小刚", "小丽"}).Find(&u3)
	fmt.Println(u3)

	// LIKE
	fmt.Println("LIKE TEST")
	Db.Where("name LIKE ?", "%丽%").Find(&u3)
	fmt.Println(u3)
	
	// AND
	fmt.Println("AND TEST")
	Db.Where("name=? AND gender=?", "小丽", "女").Find(&u3)
	fmt.Println(u3)

	//Time
	Db.Where("UpdatedAt > ?", 3).Find(&u3)
	Db.Where("CreatedAt BETWEEN ? AND ?", 3, 8).Find(&u3)

}
