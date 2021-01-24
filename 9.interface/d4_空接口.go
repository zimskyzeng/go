package main

import (
	"fmt"
)

/*
	空接口 接口中没有任何的方法
	空接口也可以看成是任意类的实现

	定义: type N interface {}
 */

type N interface {}

type Person4 struct {}
type Cat4 struct {}
type Dog4 struct {}

func checkInstance4 (n N) {
	switch n.(type) {
	case *Person4:
		fmt.Println("Person")
	case *Cat4:
		fmt.Println("Cat")
	case *Dog4:
		fmt.Println("Dog")
	}
}

func main(){
	p4 := new(Person4)
	c4 := new(Cat4)
	d4 := new(Dog4)
	checkInstance4(p4)
	checkInstance4(c4)
	checkInstance4(d4)
}