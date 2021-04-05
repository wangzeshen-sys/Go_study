package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	Db *sqlx.DB
	err error
)
func init() {
	Db, err = sqlx.Open("mysql", "root:asdasd123@tcp(127.0.0.1:3306)/db2")
	if err != nil {
		fmt.Println("open mysql failed, err:", err)
		return
	}
}

func main() {
	res, err := Db.Exec("delete from person where user_id=?", 2)
	if err != nil {
		fmt.Println("delete failed, err:", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, err:", err)
	}
	fmt.Println("delete succ", row)
}