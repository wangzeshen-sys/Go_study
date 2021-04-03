package main

import "fmt"

// 二分查找
func bin_search(arr []int, finddata int) int {
	low := 0
	hight := len(arr) - 1
	for low <= hight {
		mid := (low + hight) / 2
		fmt.Println(mid)
		if arr[mid] > finddata {
			hight = mid - 1
		} else if arr[mid] < finddata {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := make([]int, 1024*1024, 1024*1024)
	for i := 0; i < 1024*1024; i++ {
		arr[i] = i + 1
	}
	id := bin_search(arr, 124)
	if id != -1 {
		fmt.Println(id, arr[id])
	} else {
		fmt.Println("没有找到数据")
	}
}