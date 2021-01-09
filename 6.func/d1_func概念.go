package main

import "fmt"

/*
	函数
	1.概念：一段具有独立功能的代码
	2.函数定义
		func 函数名(参数 参数类型, ...) (返回值 返回值类型) { 函数体}
	3.注意事项
		函数必须先定义，再使用
		函数名在同一个包下不能冲突
		main()函数是系统的执行入口，自动调用
		如果函数没有返回值，也可以使用return中止函数执行
 */

func main(){
	fmt.Println(getSum(100))
}

func getSum(n int64) int64 {
	var ret int64
	for i:=1;int64(i)<=n;i++ {
		ret += int64(i)
	}
	return ret
}
