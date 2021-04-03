package main

import "fmt"
// 快速排序
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0]          // 第一个数据
	small := make([]int, 0, 0)   // 比我小的数据
	big := make([]int, 0, 0)     // 比我大的数据
	mid := make([]int, 0, 0)     // 与我一样大的数据
	mid = append(mid, splitdata) // 加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			small = append(small, arr[i])
		} else if arr[i] > splitdata {
			big = append(big, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	small, big = QuickSort(small), QuickSort(big)
	myarr := append(append(small, mid...), big...)
	return myarr
}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}
