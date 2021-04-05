package main

import (
	"fmt"

	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}


var Db *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "root:asdasd123@tcp(127.0.0.1:3306)/db2")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
	// defer db.Close() // 注意这行代码要写在上面err判断的下面
}

func main() {

	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Printf("select succ:%#v\n", person)
}
