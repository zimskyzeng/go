package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//var s1 string
	//// 方式1
	//// 使用Scanln来读取键盘输入，不推荐，因为没法读取空格后面的内容
	//n, err := fmt.Scanln(&s1)
	//fmt.Println(n, err)
	//fmt.Println(s1)
	//
	//// 方式2
	//// 使用bufio来读取键盘输入，推荐
	//r := bufio.NewReader(os.Stdin)
	//s2, err := r.ReadString('\n')  // 传参为分隔符
	//fmt.Println(s2)
	//
	//// 方式3
	//// Reader.ReadLine，推荐
	//arr, _, _ :=r.ReadLine()
	//fmt.Println(string(arr))


	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
