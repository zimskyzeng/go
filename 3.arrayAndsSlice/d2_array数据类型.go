package main

import "fmt"

/*
数组是值类型数据
	在下面例子可以发现，数组值传递过程中地址发生改变
*/

func sum(arr [4]int64) {
	fmt.Printf("sum内 arr地址: %p\n", &arr)
}

func main2() {
	arr1 := [4]int64{1, 2, 3, 4}
	fmt.Printf("main函数 arr地址: %p\n", &arr1)
	sum(arr1)
}
