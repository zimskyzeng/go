package main

import "fmt"

/*
	defer函数
	1.执行时间
		若函数无返回值，即在函数调用结束之前，调用defer函数
		若函数有return返回值，那么defer在最后一步返回之前执行
		若引发了panic，会在panic返回前执行defer，因此defer可以用作panic处理
	2.defer函数的执行顺序
		当有多个defer函数时，先定义的后执行，后定义的先执行
*/

func fun2(x int) int {
	defer func() {
		fmt.Println("This is first define defer func")
	}()  // 先定义，后执行

	defer func() {
		fmt.Println("This is second define defer func")
	}()  // 后定义，先执行

	fmt.Println("开始返回")
	return x + 1
}

func fun3(){
	fmt.Println("fun3")
}

func fun4(){
	fmt.Println("fun4")
	defer func() {
		err := recover()  // 在这里捕获panic，那么程序不会报错中止
		fmt.Println("err:", err)
	}()

	panic("fun4 panic")  // 此处触发panic
	fmt.Println("fun4 end")  // 后续语句不再执行
}

func fun5(){
	fmt.Println("fun5")
}

func main9(){
	fmt.Println(fun2(20))

	fun3()
	fun4()
	fun5()
}


