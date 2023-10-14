package main

import "fmt"

func main1() {
	/*变量的定义
	a.Go是静态语言，变量的类型和赋值必须一致；
	b.同一个作用域内变量名不能冲突；
	c.简短定义方式，左边至少有1个变量是新的；
	d.最外层作用域必须使用声明语句，无法使用简短声明；
	e.变量存在零值：
	*/

	// 方式1： var 变量 数据类型
	// 变量 = 赋值
	var a int64
	a = 64

	// 方式2： var 变量 数据类型 = 值
	var b string = "zimsky"

	// 方式3： var 变量名 = 值
	var c = 32

	// 方式4：简短声明 变量名 := 值
	d := "zeng"

	fmt.Println(a, b, c, d)
}
