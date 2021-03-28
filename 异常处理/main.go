package main

import "fmt"

/*
1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。

捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。

*/

func test() {
	defer func() {
		fmt.Println(recover()) // 有效
	}()

	defer recover()              // 无效
	defer fmt.Println(recover()) // 无效
	defer func() {
		func() {
			println("defer inner")
			recover() // 无效
		}()
	}()
	panic("test panic")
}

func except() {
	fmt.Println(recover())
}
func test1() {
	defer except()
	panic("test1 panic")
}
func main() {
	// test()
	test1()
}
