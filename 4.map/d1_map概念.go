package main

import "fmt"

/*
map
	1.概念：专门用于存储键值对的集合，属于引用类型，即变量存储的是一个地址
	2.语法
		创建map
			var 变量 map[key类型]value类型  》 此时变量存储的地址为nil，不能使用
			var 变量 = map[key类型]value类型{key:value, key:value}  》 存储已定义，可以使用
			var 变量 = make(map[key类型]value类型)  》 存储已定义，可以使用
		修改map
			m[key] = value  》 若key不存在，则添加，若存在则修改
		获取值
			value, ok := m[key]  》 若键存在则ok为true，value为其值。若键不存在则ok为false
		遍历
			for key, value := range m
	3.特点
		引用类型、键不能重复，若重复会覆盖
	4.方法
		delete(map, key) 删除key
 */

func main1(){
	// 创建方式1
	var m1 map[string]string
	fmt.Printf("类型：%T, 值: %v, 地址：%p\n", m1, m1, m1)  // 可以看到存储的地址为nil

	// 创建方式2
	var m2 =  map[string]string{"name": "zimsky", "addr": "shenzhen"}

	// 获取值
	if value , ok := m2["name"] ; ok {
		fmt.Println("name值=", value)
	}

	// 遍历map
	fmt.Println("开始遍历")
	for key, value := range m2{
		fmt.Println("key=", key, "value=", value)
	}

	delete(m2, "salay")  // 删除不存在的键，程序不会报错
	delete(m2, "addr")
	fmt.Printf("类型：%T, 值: %v, 地址：%p\n", m2, m2, m2)  // 可以看到存储的地址为nil
}