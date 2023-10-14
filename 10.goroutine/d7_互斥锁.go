package main

import (
	"fmt"
	"sync"
)

/*
	互斥锁：用于保护并发修改数据的安全性
 */

func main7(){
	//noMutex1()  // 数据不等于20000
	useMutex1()  // 数据等于20000
}

func useMutex1(){
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var num int

	wg.Add(2)

	go func(pmutex *sync.Mutex, pwg *sync.WaitGroup) {
		for i:=0;i<10000;i++{
			pmutex.Lock()  // 上锁
			num += 1
			pmutex.Unlock()  // 解锁
		}
		pwg.Done()
	}(&mutex, &wg)

	go func(pmutex *sync.Mutex, pwg *sync.WaitGroup) {
		for i:=0;i<10000;i++{
			pmutex.Lock()
			num += 1
			pmutex.Unlock()
		}
		pwg.Done()
	}(&mutex, &wg)

	wg.Wait()
	fmt.Println(num)
}

func noMutex1() {
	var num int
	var wg sync.WaitGroup
	wg.Add(2)

	go func(pwg *sync.WaitGroup) {
		for i:=0;i<10000;i++{
			num += 1
		}
		wg.Done()
	}(&wg)

	go func(pwg *sync.WaitGroup) {
		for i:=0;i<10000;i++{
			num += 1
		}
		wg.Done()
	}(&wg)

	wg.Wait()
	fmt.Println(num)

}
