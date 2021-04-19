package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", `df -lh`)

	// 创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe", err)
		return
	}

	// 执行命令
	err = cmd.Start()
	if err != nil {
		fmt.Println("Start:", err)
		return
	}

	// 读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ioutil.ReadAll:", err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Wait:", err)
		return
	}
	fmt.Printf("stdout:\n\n %s", bytes)
}