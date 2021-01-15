package main

import "fmt"

/*
	接口
	1.概念：方法描述的集合，仅是方法描述
		实现类：实现了该接口中所有方法的类，称为实现类
	2.注意
		使用接口类型变量的地方，都可以用实现类对象来代替，和java有点像
		接口对象不能访问实现类的属性
	3.语法
		定义接口
			type 接口名 interface{ 方法名集合 }
 */

type Action interface {
	run()
	eat()
}

type Person1 struct {
	name string
}

func (p *Person1) eat(food string) {
	fmt.Println(fmt.Sprintln(p.name, "is eating", food))
}

func (p *Person1) run() {
	fmt.Println(p.name, "is running.")
}

func main() {
	p1 := Person1{name: "zimsky"}

}

