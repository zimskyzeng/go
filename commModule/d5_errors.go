package main

import (
	"errors"
	"fmt"
)

/*
	errors包
	1.提供了创建error对象的方法，用来抛出异常。传入的参数是string
	2.另一个提供error对象的方式为 fmt.Errorf() 传入的参数是格式化文本
*/

/*
	自定义error
	在Go语言中内置了一个error接口，其中描述Error() string这个方法
	因此自定义error即自己创建struct去实现这个interface
*/

type myError struct {
	msg string
}

func (m *myError) Error() string {
	return fmt.Sprintln("Error:", m.msg)
}

// 这里的返回值是error接口
func checkNum(n int) error {
	if n < 0 {
		err := new(myError)
		err.msg = "Number less then zero."
		return err
	}
	return nil
}

func main5() {
	// 方式1
	err := errors.New("an error")
	fmt.Println(err.Error())

	// 方式2
	err1 := fmt.Errorf("error: %s", "my error")
	fmt.Println(err1.Error())

	fmt.Printf("类型：%T, 值：%v\n", err, err)
	fmt.Printf("类型：%T, 值：%v\n", err1, err1)

	if err:=checkNum(-10);err!=nil {
		fmt.Println(err.Error())
	}
}
