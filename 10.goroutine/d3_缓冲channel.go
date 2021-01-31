package main

import (
	"fmt"
	"time"
)

/*
	缓冲channel
		默认创建的channel没有缓冲区，读或写操作会阻塞，而缓冲channel则可以写入多次而不阻塞

	创建方式
		make(chan 类型, 缓冲区大小)
 */

func main(){
	ch1 := make(chan int, 3)
	go func(ch chan int) {
		for i:=0;i<10;i++{
			fmt.Println("子goroutine写入:",i)
			ch1 <- i
			time.Sleep(time.Second)
		}
		close(ch1)
	}(ch1)

	for data := range ch1 {
		fmt.Println("主goroutine读取:", data)
		time.Sleep(time.Second*2)
	}

	fmt.Println("ch1关闭，读取完毕")
}
