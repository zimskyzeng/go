package main

import (
	"fmt"
)

/*
	string由两种字符组成，byte和rune，如下是遍历举例，得出结论
		1.使用range进行遍历，对于byte字符和rune字符均无问题，能够自动检测出rune字符
		2.使用原始for循环进行遍历，ASCII正常显示，但是无法自动检测rune字符，会有乱码
	结论：尽量使用range来遍历字符串
 */

func main2() {
	s1 := "zimskyzeng"
	s2 := "hello世界"

	for i:=0;i<len(s1);i++ {
		fmt.Printf("%c", s1[i])
	}
	fmt.Println("")

	for i:=0;i<len(s2);i++ {
		fmt.Printf("%c", s2[i])  // 这里显示会有问题
	}
	fmt.Println("")

	for _, value := range s1 {
		fmt.Printf("%c", value)
	}
	fmt.Println("")

	for _, value := range s2 {
		fmt.Printf("%c", value)
	}
}
