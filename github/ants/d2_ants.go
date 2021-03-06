package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"time"
)


var sum2 int32
func main(){
	var wg sync.WaitGroup
	pool, _ := ants.NewPool(10, ants.WithPreAlloc(true))
	fmt.Println(time.Now().Format("15:4:5"))
	for i:=1;i<15;i++{
		wg.Add(1)
		pool.Submit(func() {
			fmt.Println("helloworld")
			time.Sleep(time.Second)
			wg.Done()
		})
	}

	//fmt.Println(sum2)
	wg.Wait()
	fmt.Println(time.Now().Format("15:4:5"))
}


func addNum(i interface{}){
	n := i.(int32)
	atomic.AddInt32(&sum2, n)
}