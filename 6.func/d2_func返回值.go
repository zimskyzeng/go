package main

import "fmt"

/*
	函数返回值基础
	1.在定义函数是可以写明返回哪个变量，return后面的变量可以省略
	2.函数定义时的返回值可以仅写类型
*/

func main() {
	fmt.Println(getSize1(1, 2))
	fmt.Println(getSize2(1, 2))
}

// 1.在定义函数是可以写明返回哪个变量，return后面的变量可以省略
func getSize1(a int, b int) (size int, round int) {
	size = a * b
	round = 2 * (a + b)
	return
}

// 2.函数定义时的返回值可以仅写类型
func getSize2(a int, b int) (int, int) {
	size := a * b
	round := 2 * (a + b)
	return size, round
}
