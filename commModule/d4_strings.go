package main

import (
	"fmt"
	"strings"
)

/*
	strings包中有大量对于字符串操作的工具
 */

func main4(){
	// 转换大小写
	s1 := strings.ToUpper("zimsky")
	s2 := strings.ToLower("ZIMSKY")
	fmt.Println(s1, s2)

	// 切割与拼接
	s3 := strings.Split("zimsky,zeng,shenzhen", ",")  // 直接切割
	s4 := strings.SplitN("zimsky,zeng,shenzhen", ",", 2)  // n 表示根据分隔符，至多返回的子串数量
	s5 := strings.Join(s3, ",")  // 将字符串切片拼接字符串
	fmt.Println(s3, s4, s5)

	// 查找前缀后缀
	b1 := strings.HasPrefix("zimskyzeng", "zim")  // 查找前缀
	b2 := strings.HasSuffix("zimskyzeng", "zeng")  // 查找后缀
	fmt.Println(b1, b2)

	// 包含
	b3 := strings.Contains("zimskyzeng", "sky")  // 是否包含对应的子串
	b4 := strings.ContainsAny("zimskyzeng", "yks")  // 是否包含给定的子串中的任意1个字符
	fmt.Println(b3, b4)

	// 索引
	i1 := strings.Index("zimskyzeng", "ze")  // 查找子串位置的索引
	i2 := strings.IndexAny("zimskyzeng", "yks")  // 查找含有子串任意1字符的索引
	fmt.Println(i1, i2)

	// 统计次数
	i3 := strings.Count("zimskyzeng,zeng", "ze")  // 返回的是总次数
	fmt.Println(i3)

	// 替换字符串
	s6 := strings.Replace("zimskyzeng, zeng", "zeng", "shenzhen", 1)  // n 指替换次数
	s7 := strings.ReplaceAll("zimskyzeng, zeng", "zeng", "shenzhen")  // 替换出现的所有子串
	fmt.Println(s6, s7)

	// 创建含有重复子串的字符串
	s8 := strings.Repeat("zimsky", 3)
	fmt.Println(s8)

	//
	s9 := strings.Trim(".,zimsky,zeng,", ",|.")  // 去除字符串左右两端的 给定字符
	s10 := strings.TrimSpace("  zimsky ")  // 去除给定字符串左右两端的 空格
	fmt.Println(s9, s10)
}
