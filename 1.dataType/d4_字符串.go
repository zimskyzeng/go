package main

import "fmt"

/*
	字符串
	定义：多个byte的集合，使用双引号包裹，值类型为string

	字符
	定义：单个字符，使用单引号包裹，值类型为int32
*/

func main() {
	// 默认零值为空字符串
	var s1 string
	fmt.Println("s1:", s1)

	c := 'a'
	// %T打印变量的类型 %v使用默认方式输出值 %p打印地址
	fmt.Printf("%T, %v, %p", c, c, &c)
}
