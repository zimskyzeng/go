package _1_写法分析

type IdCard struct {
	name string
	age  int
	addr string
}

// 将属性配置放在结构体内
type Person struct {
	IdCard
}

// 操作属性的函数，操作的属性是结构体，因此传入结构体的引用。同时这也是闭包返回的函数
type Option func(*IdCard)

// 通过闭包的形式给属性设置值，因为操作的是引用，可以修改对象的属性
func ChangeName(name string) Option {
	return func(card *IdCard) {
		card.name = name
	}
}

func ChangeAge (age int) Option {
	return func(card *IdCard) {
		card.age = age
	}
}

func ChangeAddr (addr string) Option {
	return func(card *IdCard) {
		card.addr = addr
	}
}

// 设置属性的调用
func (p *Person) InitPerson (opts ...Option) {
	for _, o := range opts {
		o(&p.IdCard)
	}
}
