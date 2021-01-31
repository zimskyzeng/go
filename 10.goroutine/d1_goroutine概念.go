package main

import (
	"fmt"
	"time"
)

/*
	Go 通过go 关键字创建并发goroutine

	程序执行过程
	1.创建主goroutine
	2.初始化操作
	3.执行main()
	4.main()结束，主goroutine结束，所有的子goroutine结束，程序结束

	注意：
	1.系统的主goroutine自动创建并执行main
	2.子goroutine由用户创建并启动，执行对应的函数
 */

func main(){
	fmt.Println("Program start...")
	go func() {
		for i:=0;i<5;i++{
			time.Sleep(time.Second)
			fmt.Println("11111")
		}
	}()
	for i:=0;i<5;i++{
		time.Sleep(time.Second)
		fmt.Println("22222")
	}
	// 注意：若主goroutine执行完，子goroutine未执行完，会中止子goroutine
	// 后续会使用进程等待组来实现优雅退出
}
