package main

import "fmt"

//方法的定义

// Person struct
type Person struct {
	name string
	age  uint8
}

// NewPerson 构造函数
func NewPerson(name string, age uint8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Say of Person
func (p Person) Say() {
	fmt.Printf("My name is %s, I am %d years old.", p.name, p.age)
}

func main() {
	p1 := NewPerson("张三", uint8(18))
	fmt.Printf("p1: type:%T value:%#v", p1, p1)
	p1.Say()
}
