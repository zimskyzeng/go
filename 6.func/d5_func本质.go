package main

import "fmt"

/*
	函数的本质：一种特殊的变量，引用类型变量
	函数名 指向 函数体代码的内存地址
*/

func fun1(name string) {
	fmt.Println("name:", name)
}

func main5() {
	// 函数有对应的类型
	fmt.Printf("类型：%T\n", fun1) // func(string)

	// 定义了fun2，指向fun1的函数体
	var fun2 func(string)
	fun2 = fun1
	fun2("zimsky")
}
