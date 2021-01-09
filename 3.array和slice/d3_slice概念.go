package main

import "fmt"

/*
切片
	1.概念：变长数组，是一个引用类型的容器，底层维护了一个数组。
	2.语法
		创建slice
			var s []类型  》 声明，注意这个变量的值为nil，并未维护一个容器
			s := []int64{}  》 定义，此时s作为容器，已经存放了切片的地址
			s := make([]类型, len, cap)  》 使用make来创建引用对象
			s := arr[start, end]  》 通过数组来创建切片
		注意：从已有的数组来创建切片，维护的底层数据就是该数组
	3.方法
		len(slice) 切片的长度
		cap(slice) 切片的容量 容量大于等于长度
		append(slice, element)  追加元素
		append(slice, slice...)  追加slice
		copy(old, new)  将new中的元素拷贝到old中
	4.特点
		变长容器，当超过容量时会自动扩容(成倍增长)
*/

func main() {
	var s1 []int64
	fmt.Printf("类型: %T, 值：%v, 地址：%p\n", s1, s1, s1)  // 此时s1指向的地址为nil

	s2 := []int64{}
	fmt.Printf("类型: %T, 值：%v, 地址：%p\n", s2, s2, s2)  // 此时s2指向的地址有值

	s3 := make([]int64, 5, 10)
	fmt.Println(len(s3), cap(s3))
}
