package main

import "fmt"

/*
	继承：面向对象的特性，继承父类的特性
	1.特点
		子类可以直接访问父类的方法和属性
		子类可以新增方法和属性
		子类可以重写父类已有的方法
*/

type Person5 struct {
	name string
	age  int
}

type Student5 struct {
	Person5
	school string
}

func main5() {
	// 子类的实例化1
	s1 := Student5{
		Person5: Person5{
			name: "zimsky",
			age:  20,
		},
		school: "小学",
	}

	// 子类的实例化2
	s2 := new(Student5)
	s2.Person5.name = "zeng"
	s2.Person5.age = 10
	s2.school = "幼儿园"
	fmt.Printf("s1 类型：%T, 值：%+v\n", s1, s1)
	fmt.Printf("s2 类型：%T, 值：%+v\n", s2, s2)
}
