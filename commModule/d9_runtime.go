package main

import (
	"fmt"
	"runtime"
)

func main(){
	// 打印CPU核数
	fmt.Println(runtime.NumCPU())
}
