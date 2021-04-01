package main

import (
	"flag"
	"fmt"
)

// 格式: 属性名  默认值  使用提示信息
var cliName = flag.String("name", "zimsky", "Name ?")
var cliAge = flag.Int("age", 20, "Age ?")

// 这里是手动进行赋值，实际使用中用上面的方式
var cliFlag int
func Init(){
	flag.IntVar(&cliFlag, "flagName", 20, "demo")
}

func main(){
	Init()

	// 根据命令行的参数进行解析，若解析异常会返回help提示
	flag.Parse()
	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
}
