package main

import "fmt"

/*
	深拷贝：拷贝的是值 值类型数据
	浅拷贝：拷贝的是地址 引用类型数据
*/

// 传递的参数是指针，能够直接通过地址来操作值
func modify(p *int) {
	*p += 10
}

func main2() {
	i := 1
	modify(&i)
	fmt.Println("i:", i)
}
