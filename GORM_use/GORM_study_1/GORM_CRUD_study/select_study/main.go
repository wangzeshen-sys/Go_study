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
	// Db.Debug().Find(&u3)
	Db.Find(&u3)
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
	// select * from user_infos where name in ("小刚", "小丽");
	fmt.Println(u3)

	// LIKE
	fmt.Println("LIKE TEST")
	Db.Where("name LIKE ?", "%丽%").Find(&u3)
	// select * from user_infos where name like "%丽%";
	fmt.Println(u3)
	
	// AND
	fmt.Println("AND TEST")
	Db.Where("name=? AND gender=?", "小丽", "女").Find(&u3)
	// select * from user_infos where name="小丽" and gender="女";
	fmt.Println(u3)

	//Time  目前没有测试效果
	// Db.Where("UpdatedAt > ?", 3).Find(&u3)
	// Db.Where("CreatedAt BETWEEN ? AND ?", 3, 8).Find(&u3)


	// where 查询（struct map）
	// struct
	var u5 UserInfo
	Db.Where(&UserInfo{Name: "小刚", Gender: "女"}).First(&u5)
	// select * from user_infos where name="小刚" and gender="女" limit 1;
	fmt.Println(u5)
	// map
	Db.Where(map[string]interface{}{"name":"小刚", "gender":"女"}).Find(&u3)
	// select * from user_infos where name="小刚" and gender="女";
	fmt.Println(u3)

	// not 
	fmt.Println("Not")
	Db.Not("name", "小丽").Find(&u3)
	// select * from user_infos where name <> "小丽";
	fmt.Println(u3)
	
	Db.Not("name", []string{"小明", "小明"}).Find(&u3)
	// select * from user_infos where name not in ("小明","小明");
	fmt.Println(u3)

	var u6 UserInfo
	Db.Not([]int64{1,2,3}).First(&u6)
	// select * from user_infos where id not  in (1,2,3);
	fmt.Println(u6)

	var u7 []UserInfo
	Db.Not([]int64{}).Find(&u7)
	// select * from user_infos;
	fmt.Println("u7:", u7)

	var u8 []UserInfo
	// 下面的SQL 不知道为什么没有查到数据
	// Db.Not("name=?","小刚").Find(&u8)
	// fmt.Println(u8)
	Db.Not(UserInfo{Name: "小丽"}).Find(&u8)
	// select * from user_infos where name <> '小丽';
	fmt.Println("u8", u8)

	// 带内联条件的查询
	fmt.Println("=============================================")
	fmt.Println("带内联条件的查询")
	var u9 UserInfo
	Db.First(&u9, 2)
	// select * from user_infos where id=2 limit 1;
	fmt.Println("u9", u9)

	var u10 UserInfo
	Db.Find(&u10, "name=?", "小丽")
	// select * from user_infos where name='小丽';
	fmt.Println("u10", u10)
	Db.Find(&u3, "name <> ? and gender <> ?", "小刚", "女")
	// select * from user_infos where name <>'小刚 and gender<>'女';
	fmt.Println("u3", u3)

	// struct 
	Db.Find(&u3, UserInfo{Gender: "女"})
	// select * from user_infos where gender='女';
	fmt.Println("u3",u3)
	// map
	Db.Find(&u3, map[string]interface{}{"gender":"女"})
	// selct * from user_infos where gender='女';
	fmt.Println("u3", u3)

	// or条件查询
	fmt.Println("or 条件查询")
	Db.Where("name=?", "小丽").Or("name=?", "小刚").Find(&u3)
	// select * from user _infos where name='小丽' or name='小刚';
	fmt.Println("u3", u3)
	// struct
	Db.Where("name=?", "小智").Or(UserInfo{Name: "小明"}).Find(&u3)
	// selct * from user_infos where name='小明' or name='小明';
	fmt.Println("u3", u3)

	// 查询链
	fmt.Println("查询链")
	Db.Debug().Where("name <> ?", "小刚").Where("gender=?", "女").Find(&u3)
	// SELECT * FROM `user_infos`  WHERE (name <> '小刚') AND (gender='女')
	fmt.Println("u3", u3)

	










}
