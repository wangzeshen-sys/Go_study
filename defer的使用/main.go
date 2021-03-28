package main

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5 
	}()
	return t
}

func f3() (r int) {
	defer func() {
		r = r + 5
	}()
	return 1
}

func main() {
	r1 := f1()
	r2 := f2()
	r3 := f3()
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	/*
	 defer 语句并不会马上执行，而是会进入一个栈，函数return前，会按先进后出的顺序执行，也就是说最先被定义的defer语句最后执行，先进后出的原因是后面定义的函数可能会依赖前面的资源，自然要先执行，否则，如果前面执行，那后面函数的依赖就没有了

	 defer 语句定义时，在对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。作为函数参数，则在defer定义时就把值传递给defer，并被缓存起来，作为闭包引用的话，则会在defer函数真正调用时根据整个上下文确定当前的值
	 
	 避免掉坑的关键时要理解这条语句
		 return xxx
	 这条语句并不是一个原子指令，经过编译之后，变成了三条指令
	 1 返回值 = xxx
	 2 调用 defer 函数
	 3 空的 return
	 1,3步才是return语句真正的命令, 第二步时defer定义的语句,这里就有可能会操作返回值 

	*/
}