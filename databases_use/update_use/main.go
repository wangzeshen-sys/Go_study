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
	if  err !=nil {
		fmt.Println("open mysql failed, err:", err)
		return
	}
	
	// defer Db.Close()  
	// 通常要关闭，这里步关闭
}
func main() {
	res, err :=Db.Exec("update person set username=? where user_id=?", "stu00022", 2)
	if err != nil {
		fmt.Println("update failed, err:", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, err:", err)
	}
	fmt.Println("update succ", row)
}