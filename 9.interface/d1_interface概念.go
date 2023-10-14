package main

import "fmt"

/*
	接口
	1.概念：方法描述的集合，仅是方法描述
		实现类：实现了该接口中所有方法的类，称为实现类
	2.注意
		a.使用接口类型变量的地方，都可以用实现类对象来代替，和java有点像
		b.接口对象不能访问实现类的属性
		c.接口中描述的方法可以有参数，可以有返回值，均可以省略参数仅写值
	3.语法
		定义接口
			type 接口名 interface{ 方法名集合 }
 */

type Action interface {
	run() string
	eat()
}

type Person1 struct {
	name string
}

type User1 struct {
	name string
	age int
}

// 实现eat方法
func (p *Person1) eat() {
	fmt.Println(fmt.Sprintln(p.name, "is eating -- Person method"))
}

// 实现run方法，此时Person1实现了接口Action
func (p *Person1) run() string {
	fmt.Println(p.name, "is running. -- Person method")
	return "Person"
}

func (u *User1) eat() {
	fmt.Println(fmt.Sprintln(u.name, "is eating -- User method"))
}

func (u *User1) run() string {
	fmt.Println(u.name, "is running. -- User method")
	return "User"
}

// 使用接口作为参数，传递类实例作为实参，那么会调用各自实现的接口方法
// 即多态的概念
func exec(act Action) {
	act.run()
}

func main1() {
	p1 := Person1{name: "zimsky"}
	u1 := User1{
		name: "zeng",
		age: 20,
	}

	exec(&p1)  // 调用Peron中定义的run方法
	exec(&u1)  // 调用User中定义的run方法


	var a Action
	fmt.Printf("类型:%T, 值:%v, 地址:%p\n",a, a, &a)
}

