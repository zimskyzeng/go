package main

import "fmt"

func main1(){
	i := 10

	// if语句写法格式
	// 注意：在条件判断前可以使用前置语句
	if i-=1;i < 10 {
		fmt.Println("yes")
	} else if i==10 {
		fmt.Println("no")
	} else {
		fmt.Println("haha")
	}
}
