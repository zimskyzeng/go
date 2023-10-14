package main

import "fmt"

/*
	结构体嵌套
		即使用另一个结构体对象作为属性
 */

type School struct {
	name string
	addr string
}

type Student1 struct {
	name string
	school School
}

type Student2 struct {
	name string
	school *School
}

func main4() {
	school := School{
		name: "学校A",
		addr: "南京",
	}
	p_school := new(School)

	// student1 使用school值作为属性
	student1 := Student1{
		name:   "student1",
		school: school,
	}
	// student2 使用指针作为属性值
	student2 := Student2{
		name:   "student2",
		school: p_school,
	}
	fmt.Printf("student1 值：%+v\n", student1)
	fmt.Printf("student2 值：%+v\n", student2)

	// 修改值
	student1.school.name = "大学A"
	student2.school.addr = "深圳"
	fmt.Printf("student1 值：%+v\n", student1)
	fmt.Printf("student2 值：%+v\n", student2)

}
