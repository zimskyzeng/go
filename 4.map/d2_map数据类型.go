package main

import "fmt"

/*
map的数据类型为引用类型
下例中打印的的地址相同
*/

func fun(m map[string]string) {
	fmt.Printf("值：%v, 地址：%p\n", m, m)
}

func main2() {
	m := make(map[string]string)
	m["name"] = "zimsky"
	m["addr"] = "shenzhen"
	fmt.Printf("值：%v, 地址：%p\n", m, m)
	fun(m)
}
