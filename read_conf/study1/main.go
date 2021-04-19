package main

import (
	"fmt"
	"log"


	"github.com/Unknwon/goconfig"
)

func main() {
	// 加载配置文件
	cfg, err := goconfig.LoadConfigFile("./myconf.conf")
	if err != nil {
		log.Println("LoadConfigFile failed, err: ", err)
		return
	}
	// 读取单个值
	mysql_passwd, err := cfg.GetValue("mysql", "password")
	if err != nil {
		fmt.Println("get mysql passwd failed, err:", err)
		return
	}
	redis_addr, err := cfg.GetValue("redis", "address")
	if err != nil {
		fmt.Println("get redis address failed, err:", err)
		return
	}
	fmt.Println(mysql_passwd)
	fmt.Println(redis_addr)

	// 读取整个mysql区的内容
	mysql_data, err := cfg.GetSection("mysql")
	if err != nil {
		fmt.Println("GetSection failed, err:", err)
		return
	}
	fmt.Println(mysql_data["url"])
}