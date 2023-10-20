package main

import "fmt"

type operator func(a, b int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// 高阶函数
// 在执行过程中，真正决定具体的操作方式
func calculate(a, b int, oper operator) int {
	return oper(a, b)
}

func main1() {
	ret := calculate(1, 2, add)
	fmt.Println(ret)
}
