package main

import "fmt"

/*
	匿名函数：函数没有名字
	特点：通常只能调用1次，定义时直接调用
*/

func main() {
	// 定义了一个匿名函数并自动执行
	func(name string) {
		fmt.Println("Name:", name)
	}("zimsky")

	// 定义一个匿名函数，并用一个变量来指向匿名函数，作用上与全局定义类似
	fun1 := func(age int) {
		fmt.Println(age + 10)
	}
	fun1(20)
}
