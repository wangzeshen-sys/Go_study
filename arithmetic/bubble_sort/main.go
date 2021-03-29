package main

import "fmt"

// 冒泡排序
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++{
		for j := i + 1; j < len(arr); j++{
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	} 
	return arr
}

// 冒泡排序获取最大值
func GetMax(arr []int) int {
	return arr[len(arr)-1]
}
func main() {
	initial_arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	change_arr := BubbleSort(initial_arr)
	arr_max_data := GetMax(change_arr)
	fmt.Println("冒泡排序后:", change_arr)
	fmt.Println("initual_arr is max:", arr_max_data)
}