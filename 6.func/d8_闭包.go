package main

import "fmt"

/*
	闭包，函数内部有函数，实际调用的是内部的函数
 */

// 这里用到了闭包结构，返回 y=kx+b 表达式，通过外层函数给内层函数设置初始值
func outer(k int, b int) func(int) int {
	line := func(x int) int {
		return k*x +b
	}
	return line
}

func main8(){
	fun1 := outer(1, 1)
	fmt.Println(fun1(10))  // 11
	fmt.Println(fun1(20))  // 21
}
