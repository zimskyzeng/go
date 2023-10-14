package main

import "fmt"

/*
	匿名结构体 与匿名函数类似
	没有名称的结构体定义，一般仅能在定义时实例化，用作临时使用

	1.语法
		实例变量 := struct { 字段定义 } { 字段赋值 }
*/

func main3() {
	s1 := struct {
		name string
		age  int
	}{name: "zimsky", age: 20}
	fmt.Printf("s1 值：%+v, 类型：%T", s1, s1)
}
