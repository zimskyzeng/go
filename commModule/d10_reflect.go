package main

import (
	"fmt"
	"reflect"
)

func main10(){
	u10 := User10{
		Name: "zimsky",
		Age:  10,
	}

	v := reflect.ValueOf(u10)
	user10 := v.Interface().(User10)
	fmt.Printf("type:%T, value:%+v\n", user10, user10)

	fmt.Println(v.Type())
}

type User10 struct{
	Name string
	Age int
}
