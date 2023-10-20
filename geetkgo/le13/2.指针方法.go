package main

import "fmt"

type job interface {
	programe() string
	teach() string
}

type father struct {
	name string
}

func (f *father) programe() string {
	return fmt.Sprintf("%s can programe", f.name)
}

func (f father) teach() string {
	return fmt.Sprintf("%s can teach", f.name)
}

func main() {
	// 由于引用对象默认会包含值方法和指针方法，所有使用&father可以作为job接口类型的值
	// 但是，father值对象缺少实现programe方法，所以使用father不能作为job接口类型的值
	// 核心点: 值类型仅会包含其值方法，引用类型会同时包含值方法和指针方法
	f := father{
		name: "f",
	}
	// 相当于 var j job = father  报错
	// 相当于 var j job = &father 通过
	if _, ok := interface{}(f).(job); ok {
		fmt.Println("f inplement job interface")
	} else {
		fmt.Println("f do not inplement job interface")
	}

}
