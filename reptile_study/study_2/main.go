package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reEmail = `\w+@\w+\.\w+`
	//w 代表大小写字母+数字+下划线
	// +代表出现1次或多次
	// \s\S各种字符
	// +? 代表贪婪模式
	reLinke  = `href="(https?://[\s\S]+?)"`
    rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
    reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
    reImg    = `https?://\w\s\S+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
) 

func HandlerError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 抽取页面信息
func GetPageStr(url string) (pageStr1 string) {
	resp, err := http.Get(url)
	HandlerError(err, "http.Get failed")
	defer resp.Body.Close()
	
	// 读取页面信息
	pageBytes1, err := ioutil.ReadAll(resp.Body)
	HandlerError(err, "ioutil.ReadAll failed")

	// 字节转字符串
	pageStr1 = string(pageBytes1)
	return 
}

// 获取Email
func GetEmail2(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("邮箱:", result)
	}
}

// 获取身份证号
func GetIdCard(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reIdcard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results{
		fmt.Println("身份证号:", result[0])
	} 
}

// 爬取连接
func GetLink(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("链接:", result[1])
	}
}

// 爬取手机号
func GetPhone(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("手机号", result)
	}
}

func GetImg(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("图片", result[0])
	}
}

func main() {
	// 2.抽取的爬邮箱
    // GetEmail2("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
    // 3.爬链接
    // GetLink("http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")
    // 4.爬手机号
    // GetPhone("https://www.zhaohaowang.com/")
    // 5.爬身份证号
    // GetIdCard("https://henan.qq.com/a/20171107/069413.htm")
    // 6.爬图片
    GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
}