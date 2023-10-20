package main

import "fmt"

type Dog struct {
	action string
	Animal
}

// 若这里没有显示写明该方法，那么在打印Dog对象时，会使用Animal父类对象的方法
func (d *Dog) String() string {
	return fmt.Sprintf("Dog action: %s, Animal name: %s", d.action, &d.Animal)
}

type Animal struct {
	name string
}

func (a *Animal) String() string {
	return a.name
}

func main1() {
	a := &Animal{
		name: "Dog",
	}
	fmt.Printf("%s\n", a)

	d := &Dog{
		action: "wangwang",
		Animal: *a,
	}
	fmt.Printf("%s\n", d)
}
