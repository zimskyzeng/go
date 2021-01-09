package main

import "fmt"

// 定义全局变量
var a = 100

func main(){
	// 定义局部变量
	b := 200

	// 可以使用_来舍弃额外的变量
	x, y, _ := 1, 2, 3

	fmt.Println(a, b, x, y)
}
