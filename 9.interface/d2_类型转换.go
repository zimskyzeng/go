package main

import "fmt"

/*
	类型转换：
		接口类型 --> struct 反向转换，需要手动实现
			实现方式有两种：
			a. instance，ok:=接口对象.(实际类型)
			b. 使用switch case语句
		struct --> 接口类型 正向转换，不需要手动实现
			在接口中只能使用实例方法，不能使用实例属性
 */

type Do interface {
	eat()
}

type Person2 struct {
	name string
	age int
}

func (p *Person2) eat(){
	fmt.Println(p.name, "is eating.")
}

type Animal2 struct {
	name string
}

func (a *Animal2) eat(){
	fmt.Println(a.name, "is eating.")
}

// 类型转换的方式1
func checkInstance1(d Do) {
	// 此时实现类是以指针形式实现的，那么在这里的也要写成指针形式
	if instance, ok := d.(*Person2); ok {
		fmt.Println(instance.name, "is a Person")
	} else if instance, ok := d.(*Animal2); ok {
		fmt.Println(instance.name, "is an Animal")
	}
}

// 类型转换的方式2
func checkInstance2(d Do){
	switch instance:=d.(type) {
	// 此时实现类是以指针形式实现的，那么在这里的也要写成指针形式
	case *Person2:
		fmt.Println(instance.name, "is a Person")
	case *Animal2:
		fmt.Println(instance.name, "is an Animal")
	}
}


func main2(){
	p2 := new(Person2)
	p2.name = "zimsky"

	a2 := new(Animal2)
	a2.name = "哈士奇"

	checkInstance1(p2)
	checkInstance1(a2)

	checkInstance2(p2)
	checkInstance2(a2)
}