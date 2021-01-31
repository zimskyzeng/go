package main

import "fmt"

/*
	fmt包 与输入输出有关系

 */

func main(){
	// 输出相关方法
	fmt.Println("输出普通字符串（带换行）")
	fmt.Print("输出普通字符串（不带换行）")

	fmt.Println()
	fmt.Printf("格式化字符串: %s", "zimskyzeng")
	fmt.Println()

	// 读取相关方法
	// 读取键盘建议使用bufio方法来读取
	// 注意：若使用Scan相关方法时，参数是指针类型 Scan(&s) 类似这样

	// 字符串拼接方法
	// 1.Sprintf() 格式化拼接字符串
	fmt.Println(fmt.Sprintf("%s say: %d", "zimsky", 1))

}