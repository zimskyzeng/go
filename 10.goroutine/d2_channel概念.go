package main

import (
	"fmt"
	"time"
)

/*
	channel	用于在goroutine中传递数据 是引用数据类型

	创建channel
		make(chan 数据类型)

	从channel中读取数据
		data := <- ch
		也可以不用变量来接收数据  `<- ch`
	向channel中写入数据
		ch <- data
	关闭管道，关闭后的管道不能写入数据，关闭后管道已写入的数据是能读完的
		close(ch)

	读取数据的高级方式，判断管道是否关闭
		方式1  if data, ok := <- ch ; ok {}
		方式2  for data := range ch {}  --推荐写法

	注意事项：
		1.默认情况下，创建的管道中不带buffer，读或写即阻塞，直到写和读来解除阻塞
		2.通道是安全的，同一时间只能允许一个goroutine读或者写

*/

func main() {
	ch1 := make(chan int)

	fmt.Println("================ 示例1 ===============")
	go func(ch chan int) {
		time.Sleep(time.Second * 2)
		fmt.Println("子goroutine写入数据: 1")
		ch <- 1
	}(ch1)

	fmt.Println("主goroutine等待数据")
	res := <-ch1
	fmt.Println("主goroutine读取到数据:", res)

	fmt.Println("================ 示例2 ===============")

	ch2 := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("子goroutine写入数据:", i)
			time.Sleep(time.Second)
		}
		close(ch2)
	}(ch2)

	//for {
	//	// 判断管道是否关闭
	//	if data, ok := <-ch2; ok {
	//		fmt.Println("主goroutine输出ch2数据:", data)
	//	} else {
	//		fmt.Println("主goroutine发现ch2管道关闭，结束")
	//		break
	//	}
	//}

	// 推荐写法
	for data := range ch2 {
		fmt.Println("主goroutine输出ch2数据:", data)
	}
	fmt.Println("主goroutine发现ch2管道关闭，结束")

}
