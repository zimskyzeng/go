package main

import "fmt"

/*
	函数的可变参数：可以为函数传递不定长参数
	可变参数变量的本质为切片

	注意：
	1.若函数同时存在可变参数和普通参数，可变参数要写在最后
	2.一个函数的参数列表中，至多1个可变参数
 */

func main3(){
	myAdd(1,2,3,4,5)
}

// 使用可变参数
func myAdd(i... int){
	fmt.Printf("类型：%T, 值：%v\n", i, i)

	var ret int
	for j:=1;j<=len(i);j++ {
		ret += j
	}
	fmt.Println(ret)
}
