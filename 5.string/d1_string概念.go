package main

import "fmt"

/*
string
	1.概念：本质上是字节的切片，Go中的字符串是Unicode兼容的，并且是UTF-8编码。
	2.语法
		创建：直接使用双引号 或 反引号 创建
	3.特点
		字符串由字符组成
			byte字符 ==》 uint8
			rune字符 ==》 int32
	4.方法
		len(s)  长度
		fmt.Sprintf(s1, s2...) 拼接字符串
		string.Split()  切割字符串
		string.Join()  join操作
 */

func main1() {
	s := "zimskyzeng"
	fmt.Println(s, len(s))
}