package main

import "fmt"

/*
	结构体中的方法
	1.语法
		func (接受者) 方法名(参数列表) (返回值) { 方法实现 }
		这里的接受者表示，哪个对象调用即哪个对象为接受者
*/

type Person6 struct {
	name string
	age  string
}

// 分别以值传递和引用传递来定义方法
func (p Person6) showInfo() {
	fmt.Println(fmt.Sprint("Name:", p.name, " Age:", p.age))
}

func (p *Person6) changeAge(age string) {
	p.age = age
}

func main() {
	// 以值实例的方式调用方法
	p1 := new(Person6)
	p1.name = "zim"
	p1.age = "20"
	p1.showInfo()

	// 以引用实例的方式调用方法
	p2 := Person6{
		name: "zeng",
		age:  "22",
	}
	p2.showInfo()

	// 调用修改值的方法，并打印结果
	p1.changeAge("26")
	p2.changeAge("26")
	p1.showInfo()
	p2.showInfo()
}
