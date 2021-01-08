package main

import (
	"fmt"
	"math"
)

func main(){
	// Abs 求绝对值
	fmt.Println(math.Abs(-10))

	// 向下取整
	fmt.Println(math.Floor(3.5))

	// 向上取整
	fmt.Println(math.Ceil(3.5))

	// 最大值最小值
	fmt.Println(math.Max(1.1, 2.2))
	fmt.Println(math.Min(1.1, 2.2))
}
