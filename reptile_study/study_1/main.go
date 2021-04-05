package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
	爬虫步骤
	1 明确目标（目标网站）
	2 爬取内容
	3 筛选重要信息
	4 处理数据
*/

// 爬取邮箱
var (
	reQQEmail = `(\d+)@qq.com`
)

// 爬取QQ和QQ邮箱
func GetEmail() {
	// 1 获取网站数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()

	// 读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")

	// 字节转字符串
	pageStr := string(pageBytes)
	// fmt.Println(pageStr)

	// 过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQEmail)
	// -1 代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)

	// fmt.Println(results)
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}
}

// 异常处理
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func main() {
	GetEmail()
}