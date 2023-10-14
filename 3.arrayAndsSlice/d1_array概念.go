package main

import "fmt"

/*
数组
	1.概念：存储一组相同类型的数据
	2.特点：
		数组定长、有序、长度和容量始终保持一致
	3.语法
		创建数组
			var 数组名 [长度]数据类型
			var 数组名 = [长度]数据类型{元素a, 元素b...}
			var 数组名 = [长度]数据类型{index:元素a, index:元素b...}
			数组名 := [长度]数据类型
			数组名 := [...]数据类型{元素a, 元素b...}  这里的长度用...代替，根据元素自动推断长度
	4.常用方法
		len(arr) 求数组长度
		cap(arr) 求数组容量
	5.遍历数组
		使用for循环
	6.二维数组
		数组里面再嵌套一层数组
 */

func main1(){
	var arr1 = [4]int64{}
	fmt.Printf("类型：%T, 值：%v, 地址：%p\n", arr1, arr1, &arr1)

	var arr2 = [4]string{"zimsky", "zeng"}
	arr2[0] = "shenzhen"  // 根据下标修改元素
	fmt.Printf("类型：%T, 值：%v, 地址：%p\n", arr2, arr2, &arr2)

	for index, value := range arr2 {
		fmt.Println(index, value)
	}
}
