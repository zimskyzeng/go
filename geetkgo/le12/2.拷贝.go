package main

import "fmt"

// 这里传入参数是数组切片，对元素的修改会生效
// 数组是值传递，但是数组内的值是切片的地址，所以传递过程中数组地址发生变化，但是切片地址不变
func modify(s [3][]string) {
	fmt.Printf("s:%p\n", &s)
	s[0][0] = "99"
}

func main2() {
	s := [3][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	fmt.Printf("s:%p\n", &s)
	modify(s)
	fmt.Println(s)
}
