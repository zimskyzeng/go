package main

/*
使用快速排序法对slice排序
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func mySort(s []int64) []int64 {
	if len(s) <= 1 {
		return s
	} else {
		first := s[0]
		s1 := make([]int64, 0)
		s2 := make([]int64, 0)

		for _, v := range s[1:] {
			if v < first {
				s1 = append(s1, v)
			} else {
				s2 = append(s2, v)
			}
		}
		tmp := append(mySort(s1), first)
		return append(tmp, mySort(s2)...)
	}
}

func main4() {
	s := make([]int64, 0, 10)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		s = append(s, rand.Int63n(100))
	}
	fmt.Println(s)
	ret := mySort(s)
	fmt.Println(ret)
}
