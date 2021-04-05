package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:asdasd123@tcp(127.0.0.1:3306)/db2")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	// defer Db.Close() // 注意这行代码要写在上面err判断的下面
	// 这里不要关闭，否则无法插入数据
}

func main() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)

}
