package main

import "fmt"

/*
	1.结构体作为函数参数
	2.结构体作为函数返回值
*/

type Person2 struct {
	name   string
	salary int
}

func modifySalary1(p *Person2) {
	p.salary += 10000
	fmt.Printf("modifySalary1_p 值：%+v, 地址：%p\n", p, p)
}

func modifySalary2(p Person2) {
	p.salary += 10000
	fmt.Printf("modifySalary2_p 值：%+v, 地址：%p\n", p, &p)
}

func getPerson2() Person2 {
	p := Person2{}
	return p
}

func getPerson1() *Person2 {
	return new(Person2)
}

func main2() {
	// 结构体的值或指针作为参数传递的区别
	p1 := Person2{}
	fmt.Printf("p1 值：%+v, 地址：%p\n", p1, &p1)
	modifySalary2(p1)
	fmt.Printf("p1 值：%+v, 地址：%p\n", p1, &p1)
	modifySalary1(&p1)
	fmt.Printf("p1 值：%+v, 地址：%p\n", p1, &p1)

	// 结构体作为返回值
	p2 := getPerson1()  // 返回的是指针对象
	fmt.Printf("p2 值：%+v, 地址：%p\n", p2, p2)
	p3 := getPerson2()  // 返回的是值对象
	fmt.Printf("p3 值：%+v, 地址：%p\n", p3, &p3)
}
