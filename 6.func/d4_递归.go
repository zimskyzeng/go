package main

import "fmt"

/*
	递归函数，即函数自己调用自己
 */

func fibe(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		return fibe(n-1) + fibe(n-2)
	}
}

func main4(){
	fmt.Println(fibe(5))
}
