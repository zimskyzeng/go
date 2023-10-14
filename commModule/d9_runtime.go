package main

import (
	"fmt"
	"runtime"
)

func main9(){
	// 打印CPU核数
	fmt.Println(runtime.NumCPU())
}
