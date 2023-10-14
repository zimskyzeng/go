package main

import "fmt"

/*
	switch语句注意点
	1、switch后面的变量需要与case值相匹配
	2、case值唯一，值可以为多个数值，以逗号分隔
	3、case无序
	4、default可选
	5、若switch后面的变量省略，相当于true
 */

func main2(){
	a := 1
	switch a {
	case 3:
		fmt.Println(3)
	case 2:
		fmt.Println(2)
	case 1:
		fmt.Println(1)
		fallthrough  // 使用fallthrough可以继续往下执行，而不会case直接终止
	default:
		fmt.Println("default")
	}
}
