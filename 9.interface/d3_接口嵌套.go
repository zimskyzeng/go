package main

import "fmt"

/*
接口嵌套
	集成多个已有的接口
*/

type Action3 interface {
	eat(string)
}

type Sport3 interface {
	run(addr string)
}

// 此时 All3 接口集成了Action3 和Sport3接口
type All3 interface {
	Action3
	Sport3
}

type Person3 struct {
	name string
}

func (p *Person3) eat(food string) {
	fmt.Println(p.name, "is eating", food)
}

func (p *Person3) run(addr string){
	fmt.Println(p.name, "is running in", addr)
}

func ipEat(a All3, s string){
	a.eat(s)
}

func ipRun(a All3, s string){
	a.run(s)
}

func main3(){
	p3 := new(Person3)
	p3.name = "zimsky"

	ipEat(p3, "food")
	ipRun(p3, "shenzhen")
}