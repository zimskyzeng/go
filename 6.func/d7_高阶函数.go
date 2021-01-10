package main

import "fmt"

/*
	高阶函数：含有函数作为参数的 函数称为高阶函数
	回调函数：作为参数，传递给另一个函数的函数称为回调函数

	本质在于，函数是一种特殊的变量，能够作为参数传递
*/

// 此时，exec函数称为高级函数，具备函数作为其参数
// oper函数作为回调函数
func exec(a int, b int, oper func(int, int) int) int {
	return oper(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println(exec(20, 10, add))
	fmt.Println(exec(20, 10, subtract))
}
