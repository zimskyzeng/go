package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	生成随机数
*/

func main() {
	// 需要先设置种子数，而且该种子数是动态的
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 20; i++ {
		// 使用rand.Intn获取随机数
		fmt.Println(rand.Intn(20))
	}
}
