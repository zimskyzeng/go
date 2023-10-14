package main

import (
	"fmt"
	"time"
)

/*
	select 分支语句 专门用于处理通道读写操作

	语法：
		详见示例
	执行顺序：
		1.在进入select语句时，若有多个分支可以执行，则随机挑选一个。
		2.若没有分支可以执行，则执行default或者阻塞等待。
		注意：case分支在阻塞时也是属不可执行状态。

*/

func main5() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	select {
	case data := <-ch1:  // 此时还不能执行
		fmt.Println("ch1 data:", data)
	case <-ch2:  // 此时还不能执行
		fmt.Println("ch2 data:", "data")
	default:  // 默认走default
		fmt.Println("default")
	}
}
