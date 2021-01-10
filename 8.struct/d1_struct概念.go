package main

import "fmt"

/*
	Go 语言中使用结构体struct来表示类
	1.语法
		定义结构体
			type 结构体名 struct{ 属性名 类型; 属性名 类型}  》 注意写成1行的时候是;分割
		创建实例
			var 变量名 结构体名
			变量名 := 结构体名{属性:值, 属性:值...}
			指针变量名 := new(结构体名)  》 注意返回的是指向结构体对象的指针
		修改属性值
			变量名.属性 = 新值
	2.应用场景
		实现面向对象的操作
		传递的是值实例，则是深拷贝
		传递的是引用实例，则是浅拷贝
	3.注意事项
		若结构体的属性名省略，则默认使用类型作为属性名，因此不能使用重复的类型作为属性
 */

type Person1 struct { name string; age int}

func main(){
	// 创建方法12
	p1 := Person1{name:"zimsky",age:20}
	p2 := Person1{}
	fmt.Printf("p1变量 类型：%T, 值：%+v, 地址：%p\n", p1, p1, &p1)
	fmt.Printf("p2变量 类型：%T, 值：%+v, 地址：%p\n", p2, p2, &p2)

	// 修改属性
	p1.name = "zeng"
	fmt.Printf("p1变量 类型：%T, 值：%+v, 地址：%p\n", p1, p1, &p1)

	// 使用new创建实例
	p3 := new(Person1)
	// 使用如下两种方法都可以修改属性，Go语言提供的语法糖
	(*p3).name = "zzm"
	p3.age = 22
	fmt.Printf("p3变量 类型：%T, 值：%+v, 实例地址：%p, 指针自身地址：%p\n", p3, p3, p3, &p3)

	// 通过下面代码可以更好的看出struct指针的地址关系
	p3 = &p2
	fmt.Printf("p3变量 类型：%T, 值：%+v, 实例地址：%p, 指针自身地址：%p\n", p3, p3, p3, &p3)
}