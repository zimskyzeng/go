package main

import "fmt"

/*
	参考d2，这里提到的更新字符串依然会出现 byte和rune的问题
 */

func main(){
	s1 := "zimskyzeng"
	s2 := "hello世界"

	// 由于string不可变性，需要先转换成变长数组slice
	newS1 := []byte(s1)  // 对于ASCII字符，可以使用byte切片
	newS1[1] = 'a'
	fmt.Println(string(newS1))

	newS2 := []rune(s2)  // 因为包含了中文字符，这里只能使用rune
	newS2[5] = '中'
	fmt.Println(string(newS2))
}
