package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	WaitGroup 同步等待组，用于等待goroutine执行完成。当counter值不为零时阻塞。
	本质上是结构体对象

	常用方法
	1.Add()
		设置counter值
	2.Done()
		将counter值减1
	3.Wait()
		WaitGroup对象counter不为零时阻塞，直到counter为零时放行。

	注意：counter不能小于零
 */

func main() {
	var wg sync.WaitGroup
	wg.Add(2)


	go func(wg *sync.WaitGroup) {
		time.Sleep(time.Second * 2)
		fmt.Println("group1 done...")
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		time.Sleep(time.Second * 3)
		fmt.Println("group2 done...")
		wg.Done()
	}(&wg)

	wg.Wait()
	fmt.Println("程序结束")
}