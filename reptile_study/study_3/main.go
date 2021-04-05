package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 并发爬取图片
// 并发爬虫思路
// 1 初始化数据管道
// 2 爬虫写出: 26个协程向管道中添加图片链接
// 3 任务统计协程， 检查26个任务是否都完成，完成则关闭数据管道
// 下载协程：从管道里读取链接并下载

var (
	// 存放图片链接的数据管道
	chanImageUrls chan string
	wg            sync.WaitGroup
	//用于监控协程
	chanTask chan string
	reImg1   = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func HandleError1(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}

func GetPageStr1(url string) (pageStr11 string) {
	resp, err := http.Get(url)
	HandleError1(err, "http.Get failed")
	defer resp.Body.Close()
	// 2 读取页面内容
	pagebytes, err := ioutil.ReadAll(resp.Body)
	HandleError1(err, "ioutil.ReadAll failed")
	pageStr11 = string(pagebytes)
	return
}

// 获取当前页面图片链接
func GetImg1(url string) (urls []string) {
	pageStr := GetPageStr1(url)
	re := regexp.MustCompile(reImg1)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

// 将爬取到的链接发到送到管道
func GetImgUrls(url string) {
	urls := GetImg1(url)
	// 遍历切片里所有链接,存入数据管道
	for _, url := range urls {
		chanImageUrls <- url
	}
	// 标识当前协程完成
	// 每完成一个任务， 写入一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- url
	wg.Done()
}

// 任务统计协程
func CheaOk() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	wg.Done()
}

// 截取url名字
func GetFileNameFromUrl(url string) (fileName string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	fileName = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	fileName = timePrefix + "_" + fileName
	return
}

// 下载图片，传入的图片叫什么
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError1(err, "http.Get url")
	defer resp.Body.Close()
	pagebytes, err := ioutil.ReadAll(resp.Body)
	HandleError1(err, "ioutil.ReadAll failed")
	filename = "as" + filename

	// 写出数据
	err = ioutil.WriteFile(filename, pagebytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

// 是否下载成功
func DownloadImg() {
	for url := range chanImageUrls {
		filename := GetFileNameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
}
func main() {
	// 1 初始化管道
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, 26)

	// 2 爬虫协程
	for i := 1; i < 27; i++ {
		wg.Add(1)
		go  GetImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}

	// 3 任务统计协程，统计26个任务是否都完成，完成则关闭管道 
	wg.Add(1)
	go CheaOk()

	// 4 下载协程：从管道中读取链接并下载
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go DownloadImg()
	}
	wg.Wait()
}
