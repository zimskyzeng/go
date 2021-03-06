package main

import "fmt"

type Animal struct {}

func main(){
	var a *Animal
	a = new(Animal)
	if a == nil {
		fmt.Println(111)
	}
}
