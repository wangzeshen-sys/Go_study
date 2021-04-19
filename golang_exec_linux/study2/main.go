package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

// 创建一个缓冲读取器

func main() {
	cmd := exec.Command("/bin/bash", "-c", `tail -3 main.go`)

	// 创建一个获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe:", err)
		return
	}

	// 执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Start", err)
		return
	}

	// 使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)

	for {
		// 判断是否到了文件的结尾否则出错
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				fmt.Printf("Error: %s\n", err)
			}
			return
		}
		fmt.Printf("%s\n", string(output))
	}

	// wait 方法会一直阻塞到其所有的命令完全运行结束为止
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println("wait:", err.Error())
	// 	return
	// }
}