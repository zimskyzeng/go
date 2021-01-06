package main

import (
	"fmt"
)

/*
	常量
	1、常量在执行过程中不可改变；
	2、定义方式
		显示定义： const 常量名 数据类型 = 值
		隐式定义： const 常量名 = 值
	3、常量一般大写命名；
*/

func main() {
	// 定义方式
	const PATH string = "/usr/local/bin/"
	const PI = 3.14

	// 同时定义多个常量
	

	fmt.Println(PATH, PI)
}
