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

	iota: 计数器
	当使用const定义常量组时，会初始化一个iota，值为0；
	每当在常量组中定义一个常量，计数器加1，直到下一个const出现清零
*/

func main3() {
	// 定义方式
	const PATH string = "/usr/local/bin/"
	const PI = 3.14

	// 同时定义多个常量
	const c1, c2, c3 = 1, 3, 4
	const (
		FEMALE = 0
		MALE   = 1
	)

	// 若某个常量没有设置值，则与上一个相同
	const (
		a int = 100
		b  // 此时与a相同
		c
		d string = "zim"
		e
	)

	// iota举例
	const (
		i = iota
		j
		k
	)

	fmt.Println(PATH, PI)
	fmt.Println(a, b, c, d, e)
	fmt.Println(i, j, k)
}
