package main

import "fmt"

/*
slice是引用类型变量
*/

func main() {
	// 方式1
	// 通过make创建新的slice然后赋值
	s1 := []int64{1, 2, 3, 4, 5, 6}
	s2 := make([]int64, 0, 0)
	for _, value := range s1 {
		s2 = append(s2, value)
	}
	fmt.Printf("s1值：%v, s1地址：%p\n", s1, s1)
	fmt.Printf("s2值：%v, s2地址：%p\n", s2, s2)

	// 方式2
	// 使用copy()方法
	s3 := []int64{2, 3, 1, 4, 12, 1}
	s4 := make([]int64, 0, len(s3))
	copy(s4, s3)
	fmt.Printf("s3值：%v, s3地址：%p\n", s3, s3)
	fmt.Printf("s4值：%v, s4地址：%p\n", s4, s4)

	/*
		注意：
		使用copy方式时，接收的slice一定要有足够的len来接收，不然无法拷贝
	*/
}
