package main

import (
	"fmt"
	"time"
)

/*
	定向通道
		管道默认都是双向的，定向通道即只读或者只写

	使用场景
		函数中使用，用来限定管道的读或者写操作
 */

func main(){
	ch1 := make(chan int)  // 双向channel
	go writeOnly(ch1)
	go readOnly(ch1)
	time.Sleep(time.Second)

}

// 含只读通道的函数
func readOnly(ch <- chan int){
	res := <- ch
	fmt.Println("readOnly:", res)
}

// 含只写通道的函数
func writeOnly(ch chan <- int){
	ch <- 1
	fmt.Println("writeOnly:", 1)
}
